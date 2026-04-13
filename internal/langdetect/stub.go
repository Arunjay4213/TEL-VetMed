// stub.go is a fake language detector that always returns English.
// Used during development so the pipeline runs end-to-end without a real detector.
package langdetect

type Stub struct{}

func NewStub() *Stub {
	return &Stub{}
}

func (s *Stub) Detect(text string) (Result, error) {
	return Result{
		PrimaryLanguage: "en",
		Segments: []Segment{
			{Text: text, Language: "en", Start: 0, End: len(text)},
		},
	}, nil
}
