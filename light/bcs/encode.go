package bcs

import (
	"bytes"
	"encoding/binary"
	"io"
	"reflect"

	"github.com/pkg/errors"
)

const (
	MAX_SEQUENCE_LENGTH = (1 << 31) - 1
)

func WriteUint8(w io.Writer, v uint8) (int, error) {
	return w.Write([]byte{v})
}

func WriteUint16(w io.Writer, v uint16) (int, error) {
	var buf [2]byte
	binary.LittleEndian.PutUint16(buf[:], v)
	return w.Write(buf[:])
}

func WriteUint32(w io.Writer, v uint32) (int, error) {
	var buf [4]byte
	binary.LittleEndian.PutUint32(buf[:], v)
	return w.Write(buf[:])
}

func WriteUint64(w io.Writer, v uint64) (int, error) {
	var buf [8]byte
	binary.LittleEndian.PutUint64(buf[:], v)
	return w.Write(buf[:])
}

func WriteBool(w io.Writer, v bool) (int, error) {
	if v {
		return WriteUint8(w, uint8(1))
	}

	return WriteUint8(w, uint8(0))
}

func WriteInt8(w io.Writer, v int8) (int, error) {
	return WriteUint8(w, uint8(v))
}

func WriteInt16(w io.Writer, v int16) (int, error) {
	return WriteUint16(w, uint16(v))
}

func WriteInt32(w io.Writer, v int32) (int, error) {
	return WriteUint32(w, uint32(v))
}

func WriteInt64(w io.Writer, v int64) (int, error) {
	return WriteUint64(w, uint64(v))
}

func writeULEB128(w io.Writer, value uint32) (int, error) {
	var len int

	for value >= 0x80 {
		data := byte(value&0x7F) | 0x80
		n, err := w.Write([]byte{data})
		if err != nil {
			return 0, err
		}

		len += n
		value >>= 7
	}

	n, err := w.Write([]byte{byte(value)})
	if err != nil {
		return 0, err
	}

	return len + n, nil
}

func WriteLen(w io.Writer, len int) (int, error) {
	if len > MAX_SEQUENCE_LENGTH {
		return 0, errors.Errorf("exceeded max sequence length: %v", len)
	}

	return writeULEB128(w, uint32(len))
}

func WriteBytes(w io.Writer, v []byte) (int, error) {
	len1, err := WriteLen(w, len(v))
	if err != nil {
		return 0, err
	}

	len2, err := w.Write(v)
	if err != nil {
		return 0, err
	}

	return len1 + len2, nil
}

func WriteString(w io.Writer, v string) (int, error) {
	return WriteBytes(w, []byte(v))
}

func WriteOption(w io.Writer, some bool) (int, error) {
	return WriteBool(w, some)
}

func WriteEnumIndex(w io.Writer, index int) (int, error) {
	return writeULEB128(w, uint32(index))
}

func encodeToBytes(val reflect.Value) ([]byte, error) {
	var buf bytes.Buffer

	if _, err := write(&buf, val); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func write(w io.Writer, val reflect.Value) (int, error) {
	switch val.Kind() {
	case reflect.Bool:
		return WriteBool(w, val.Bool())
	case reflect.Int8:
		return WriteInt8(w, int8(val.Int()))
	case reflect.Int16:
		return WriteInt16(w, int16(val.Int()))
	case reflect.Int32:
		return WriteInt32(w, int32(val.Int()))
	case reflect.Int64:
		return WriteInt64(w, val.Int())
	case reflect.Uint8:
		return WriteUint8(w, uint8(val.Uint()))
	case reflect.Uint16:
		return WriteUint16(w, uint16(val.Uint()))
	case reflect.Uint32:
		return WriteUint32(w, uint32(val.Uint()))
	case reflect.Uint64:
		return WriteUint64(w, val.Uint())
	case reflect.Array: // e.g. Hash, Address
		return writeArray(w, val)
	case reflect.Interface:
		return write(w, val.Elem())
	case reflect.Map:
		return writeMap(w, val)
	case reflect.Ptr: // for rust Option type
		return writeOption(w, val)
	case reflect.Slice:
		return writeSlice(w, val)
	case reflect.String:
		return WriteString(w, val.String())
	case reflect.Struct:
		return writeStruct(w, val)
	default:
		return 0, errors.Errorf("unsupported type kind: %v", val.Kind())
	}
}

func writeArray(w io.Writer, val reflect.Value) (int, error) {
	var count int

	for i, len := 0, val.Len(); i < len; i++ {
		n, err := write(w, val.Index(i))
		if err != nil {
			return 0, err
		}

		count += n
	}

	return count, nil
}

func writeOption(w io.Writer, val reflect.Value) (int, error) {
	if val.IsNil() {
		return WriteOption(w, false)
	}

	len1, err := WriteOption(w, true)
	if err != nil {
		return 0, err
	}

	len2, err := write(w, val.Elem())
	if err != nil {
		return 0, err
	}

	return len1 + len2, nil
}

func writeSlice(w io.Writer, val reflect.Value) (int, error) {
	if val.Type().Elem().Kind() == reflect.Uint8 {
		return WriteBytes(w, val.Bytes())
	}

	count, err := WriteLen(w, val.Len())
	if err != nil {
		return 0, err
	}

	for i, len := 0, val.Len(); i < len; i++ {
		n, err := write(w, val.Index(i))
		if err != nil {
			return 0, err
		}

		count += n
	}

	return count, nil
}

func writeStruct(w io.Writer, val reflect.Value) (int, error) {
	rtype := val.Type()

	var count int

	for i, len := 0, val.NumField(); i < len; i++ {
		// ignore private field
		if !rtype.Field(i).IsExported() {
			continue
		}

		n, err := write(w, val.Field(i))
		if err != nil {
			return 0, err
		}

		count += n
	}

	return count, nil
}
