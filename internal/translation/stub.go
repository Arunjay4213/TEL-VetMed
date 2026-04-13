// stub.go is a fake translator that labels text with the language direction
// but does not actually translate. Used during development.
package translation

import "fmt"

type Stub struct{}

func NewStub() *Stub {
	return &Stub{}
}

func (s *Stub) Translate(text string, sourceLang string, targetLang string) (Result, error) {
	return Result{
		Text:           fmt.Sprintf("[stub translation %s->%s] %s", sourceLang, targetLang, text),
		SourceLanguage: sourceLang,
		TargetLanguage: targetLang,
		Confidence:     0.99,
	}, nil
}
