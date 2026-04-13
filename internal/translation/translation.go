// translation.go defines the interface for the translation layer. After the LLM
// generates a response, this converts it into the other language if needed --
// for example, turning an English response into Spanish for a farm worker.
// It is designed to preserve veterinary terminology rather than doing
// word-for-word literal translation.
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
