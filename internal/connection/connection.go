package connection

type Connection interface {
	SendByParams(params map[string]interface{}) error
}
