package alexa

import "fmt"

type Response struct {
	Speech     *Speech `json:"outputSpeech,omitempty"`
	EndSession bool    `json:"shouldEndSession"`
}

func Plaintextf(text string, args ...interface{}) *Speech {
	return &Speech{
		Type: "PlainText",
		Text: fmt.Sprintf(text, args...),
	}
}

func SSMLf(text string, args ...interface{}) *Speech {
	return &Speech{
		Type: "SSML",
		SSML: fmt.Sprintf(text, args...),
	}
}

type Speech struct {
	Type string `json:"type"`
	Text string `json:"text,omitempty"`
	SSML string `json:"ssml,omitempty"`
}
