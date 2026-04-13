package tts

type Stub struct{}

func NewStub() *Stub {
	return &Stub{}
}

func (s *Stub) Synthesize(text string, language string) (AudioOutput, error) {
	return AudioOutput{
		Audio:      []byte(text),
		Format:     "wav",
		SampleRate: 16000,
	}, nil
}
