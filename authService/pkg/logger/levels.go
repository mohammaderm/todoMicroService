package logger

type (
	Logger interface {
		Info(msg string, extras ...map[string]interface{})
		Warning(msg string, extras ...map[string]interface{})
		Error(msg string, extras ...map[string]interface{})
		Panic(msg string, extras ...map[string]interface{})
	}
)
