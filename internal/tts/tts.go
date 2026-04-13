package tts

type AudioOutput struct {
	Audio      []byte
	Format     string
	SampleRate int
}

type Service interface {
	Synthesize(text string, language string) (AudioOutput, error)
}
