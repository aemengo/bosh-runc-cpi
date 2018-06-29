package bosh

type Response struct {
	Result interface{} `json:"result"`
	Error  *Error      `json:"error"`
	Log    string      `json:"log"`
}

type Error struct {
	Type      string `json:"type"`
	Message   string `json:"message"`
	OkToRetry bool   `json:"ok_to_retry"`
}

func CPIError(message string, logMessage ...string) Response {
	var log string

	if len(logMessage) > 0 {
		log = logMessage[0]
	}

	return Response{
		Error: &Error{
			Type:      "Bosh::Cpi::CPIError",
			Message:   message,
			OkToRetry: false,
		},
		Log: log,
	}
}
