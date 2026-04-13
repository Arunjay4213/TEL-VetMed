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
