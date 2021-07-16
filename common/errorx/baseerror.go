package errorx

const defaultCode = 400

type CodeError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

type CodeErrorResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func NewCodeError(code int, message string, data interface{}) error {
	return &CodeError{Code: code, Message: message, Result: data}
}

func NewDefaultError(message string, data interface{}) error {
	return NewCodeError(defaultCode, message, data)
}

func (e *CodeError) Error() string {
	return e.Message
}

func (e *CodeError) DataInfo() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code:    e.Code,
		Message: e.Message,
		Result:  e.Result,
	}
}
