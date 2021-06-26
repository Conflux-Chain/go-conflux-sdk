package errors

type ErrorCode int

var codeStringMap map[ErrorCode]string

const (
	CodePivotAssumption ErrorCode = 1
	CodeBlockNotFound
)

func init() {
	codeStringMap = make(map[ErrorCode]string)
	codeStringMap[CodePivotAssumption] = "CodePivotAssumption"
	codeStringMap[CodeBlockNotFound] = "CodeBlockNotFound"
}

func (e ErrorCode) String() string {
	return codeStringMap[e]
}
