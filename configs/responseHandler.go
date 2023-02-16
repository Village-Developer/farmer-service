package configs

type Success struct {
	Success bool
	Message string
}

type Error struct {
	Success bool
	Message string
}

func (Configs) ResponseSuccess(message string) Success {
	return Success{
		Success: true,
		Message: "message",
	}
}

func (Configs) ResponseFailed(message string) Success {
	return Success{
		Success: false,
		Message: message,
	}
}
