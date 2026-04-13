package asr

type Stub struct{}

func NewStub() *Stub {
	return &Stub{}
}

func (s *Stub) Transcribe(audio []byte) (Transcript, error) {
	return Transcript{
		Text:       "[stub] transcribed audio",
		Confidence: 0.95,
		Language:   "en",
	}, nil
}
