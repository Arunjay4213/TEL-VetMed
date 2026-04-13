package translation

type Result struct {
	Text           string
	SourceLanguage string
	TargetLanguage string
	Confidence     float64
}

type Service interface {
	Translate(text string, sourceLang string, targetLang string) (Result, error)
}
