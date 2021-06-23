package errors

type ErrorCode int

var codeStringMap map[ErrorCode]string

const (
	CodePivotSwitch ErrorCode = iota
)

func init() {
	codeStringMap = make(map[ErrorCode]string)
	codeStringMap[CodePivotSwitch] = "CodePivotSwitch"
}

func (e ErrorCode) String() string {
	return codeStringMap[e]
}
