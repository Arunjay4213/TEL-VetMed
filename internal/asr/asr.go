// asr.go defines the shared types and interface for Automatic Speech Recognition.
// Any ASR provider (Deepgram, Whisper, etc.) must implement the Service interface
// defined here. The rest of the codebase only talks to this interface, so you
// can swap providers without changing anything outside this package.
package asr

type Word struct {
	Word       string  `json:"word"`
	Start      float64 `json:"start"`
	End        float64 `json:"end"`
	Confidence float64 `json:"confidence"`
}

type Transcript struct {
	Text     string
	Language string
	Words    []Word
}

type Service interface {
	Transcribe(audio []byte, env string) (Transcript, error)
}
