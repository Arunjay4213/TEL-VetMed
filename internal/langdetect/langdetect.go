package langdetect

type Result struct {
	PrimaryLanguage string
	Segments        []Segment
}

type Segment struct {
	Text     string
	Language string
	Start    int
	End      int
}

type Service interface {
	Detect(text string) (Result, error)
}
