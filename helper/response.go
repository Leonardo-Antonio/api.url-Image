package helper

type Response struct {
	MessageType string      `json:"message_type" xml:"message_type"`
	Message     string      `json:"message" xml:"message"`
	Error       bool        `json:"error" xml:"error"`
	Data        interface{} `json:"data" xml:"data"`
}

func NewResponseJSON(
	msgT, msg string,
	err bool, data interface{},
) *Response {
	return &Response{
		MessageType: msgT,
		Message:     msg,
		Error:       err,
		Data:        data,
	}
}
