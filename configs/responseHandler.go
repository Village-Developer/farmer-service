package configs

type Success struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type SuccessData struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Error struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type Response interface {
	ResponseSuccess() Success
}

func (s *Success) ResponseSuccess(message string) Success {
	return Success{
		Success: true,
		Message: message,
	}
}

func (s *SuccessData) ResponseSuccess(message string, data interface{}) SuccessData {
	return SuccessData{
		Success: true,
		Message: message,
		Data:    data,
	}
}

func (e *Error) ResponseFailed(message string) Error {
	return Error{
		Success: false,
		Message: message,
	}
}
