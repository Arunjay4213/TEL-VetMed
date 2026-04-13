package asr

import (
	"bytes"
	"net/http"
)

type Deepgram struct {
	apiKey string
}

func NewDeepgram(apiKey string) *Deepgram {
	return &Deepgram{
		apiKey: apiKey,
	}
}

func (d *Deepgram) Transcribe(audio []byte) (Transcript, error) {
	req, err := http.NewRequest("POST", "https://api.deepgram.com/v1/listen?model=nova-2&language=en", bytes.NewReader(audio))
	if err != nil {
		return Transcript{}, err
	}
	req.Header.Set("Authorization", "Token "+d.apiKey)
	req.Header.Set("Content-Type", "audio/wav")

	return Transcript{}, nil
}
