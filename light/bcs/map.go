package bcs

import (
	"bytes"
	"io"
	"reflect"
	"sort"
)

type entry struct {
	key   []byte
	value []byte
}

type entries []entry

func (list entries) Len() int { return len(list) }

func (list entries) Less(i, j int) bool { return bytes.Compare(list[i].key, list[j].key) < 0 }

func (list entries) Swap(i, j int) { list[i], list[j] = list[j], list[i] }

func writeMap(w io.Writer, val reflect.Value) (int, error) {
	count, err := WriteLen(w, val.Len())
	if err != nil {
		return 0, err
	}

	var entries entries
	iter := val.MapRange()

	for iter.Next() {
		key, err := encodeToBytes(iter.Key())
		if err != nil {
			return 0, err
		}

		value, err := encodeToBytes(iter.Value())
		if err != nil {
			return 0, err
		}

		entries = append(entries, entry{key, value})

		count += len(key) + len(value)
	}

	sort.Sort(entries)

	for _, v := range entries {
		if _, err = w.Write(v.key); err != nil {
			return 0, err
		}

		if _, err = w.Write(v.value); err != nil {
			return 0, err
		}
	}

	return count, nil
}
