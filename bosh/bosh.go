package bosh

import (
	"fmt"
)

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
	return errr("Bosh::Clouds::CPIError", prefix, err, logMessage...)
}

func CloudError(prefix string, err error, logMessage ...string) Response {
	return errr("Bosh::Clouds::CloudError", prefix, err, logMessage...)
}

func UnimplementedError(method string) Response {
	return errr("Bosh::Clouds::NotImplemented", "", fmt.Errorf("'%s' is not yet supported. Please call implemented method", method))
}

func errr(kind string, prefix string, err error, logMessage ...string) Response {
	var (
		log string
		message string
	)

	if len(logMessage) > 0 {
		log = logMessage[0]
	}

	if prefix == "" {
		message = err.Error()
	} else {
		message = prefix + ": " + err.Error()
	}

	return Response{
		Error: &Error{
			Type:      kind,
			Message:   message,
			OkToRetry: false,
		},
		Log: log,
	}
}