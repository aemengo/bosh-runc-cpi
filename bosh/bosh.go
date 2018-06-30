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

func CPIError(prefix string, err error, logMessage ...string) Response {
	var log string

	if len(logMessage) > 0 {
		log = logMessage[0]
	}

	return Response{
		Error: &Error{
			Type:      "Bosh::Clouds::CPIError",
			Message:   prefix + ": " + err.Error(),
			OkToRetry: false,
		},
		Log: log,
	}
}