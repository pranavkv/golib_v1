package libError

type ClientError interface {

	Error() string
	
	ResponseBody() ([]byte, error)

	ResponseHeaders() (int, map[string]string)
}
