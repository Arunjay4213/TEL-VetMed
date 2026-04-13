// stub.go is a fake ASR implementation used during development and testing.
// It always returns the same hardcoded transcript so you can test the rest
// of the pipeline without needing a real microphone or Deepgram API call.
package asr

type Stub struct{}

func NewStub() *Stub {
	return &Stub{}
}

func (s *Stub) Transcribe(audio []byte, env string) (Transcript, error) {
	return Transcript{
		Text:     "[stub] transcribed audio",
		Language: "en",
		Words:    []Word{},
	}, nil
}
