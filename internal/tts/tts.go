// tts.go defines the interface for Text-to-Speech synthesis. This is the last
// step in the pipeline -- it takes the final response text and converts it into
// spoken audio that plays back to the user. It supports both English and Spanish
// output.
package tts

type AudioOutput struct {
	Audio      []byte
	Format     string
	SampleRate int
}

type Service interface {
	Synthesize(text string, language string) (AudioOutput, error)
}
