// deepgram.go is the real ASR implementation using Deepgram's Nova-3 model.
// It takes raw audio bytes, sends them to Deepgram's API over HTTP, and returns
// the transcribed text along with the detected language. Nova-3 supports
// mid-sentence code-switching between English and Spanish, which is important
// for bilingual farm conversations.
package asr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Deepgram struct {
	apiKey string
}

type deepgramResponse struct {
	Results struct {
		Channels []struct {
			Alternatives []struct {
				Transcript string  `json:"transcript"`
				Confidence float64 `json:"confidence"`
				Words      []Word  `json:"words"`
			} `json:"alternatives"`
			DetectedLanguage string `json:"detected_language"`
		} `json:"channels"`
	} `json:"results"`
}

func NewDeepgram(apiKey string) *Deepgram {
	return &Deepgram{
		apiKey: apiKey,
	}
}

func (d *Deepgram) Transcribe(audio []byte, env string) (Transcript, error) {
	endpoint := "https://api.deepgram.com/v1/listen"

	params := url.Values{}
	params.Set("model", "nova-3")
	params.Set("language", "multi")
	params.Set("smart_format", "true")
	params.Set("punctuate", "true")

	for _, term := range KeytermsForEnvironment(env) {
		params.Add("keyterm", term)
	}

	req, err := http.NewRequest("POST", endpoint+"?"+params.Encode(), bytes.NewReader(audio))
	if err != nil {
		return Transcript{}, fmt.Errorf("building request: %w", err)
	}
	req.Header.Set("Authorization", "Token "+d.apiKey)
	req.Header.Set("Content-Type", "audio/wav")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Transcript{}, fmt.Errorf("sending request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Transcript{}, fmt.Errorf("reading response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return Transcript{}, fmt.Errorf("deepgram returned %d: %s", resp.StatusCode, string(body))
	}

	var dgResp deepgramResponse
	if err := json.Unmarshal(body, &dgResp); err != nil {
		return Transcript{}, fmt.Errorf("parsing response: %w", err)
	}

	if len(dgResp.Results.Channels) == 0 ||
		len(dgResp.Results.Channels[0].Alternatives) == 0 {
		return Transcript{}, fmt.Errorf("no transcription results returned")
	}

	alt := dgResp.Results.Channels[0].Alternatives[0]
	lang := dgResp.Results.Channels[0].DetectedLanguage

	return Transcript{
		Text:     alt.Transcript,
		Language: lang,
		Words:    alt.Words,
	}, nil
}
