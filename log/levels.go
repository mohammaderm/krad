package log

type (
	Logger interface {
		Info(field *Field)
		Warning(field *Field)
		Error(field *Field)
		Panic(field *Field)
	}
	Field struct {
		Package  string
		Function string
		Params   string
		Message  string
	}
)
