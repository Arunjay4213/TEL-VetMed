package asr

type Transcript struct {
	Text       string
	Confidence float64
	Language   string
}

type Service interface {
	Transcribe(audio []byte) (Transcript, error)
}
