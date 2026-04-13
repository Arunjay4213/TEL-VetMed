// stub.go is a fake TTS service that returns the response text as raw bytes
// instead of real audio. Used during development.
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
