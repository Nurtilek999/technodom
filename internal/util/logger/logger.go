package logger

type Fields map[string]interface{}

type Logger interface {
	Info(message string, args map[string]interface{})
	Error(message string, args map[string]interface{})
}
