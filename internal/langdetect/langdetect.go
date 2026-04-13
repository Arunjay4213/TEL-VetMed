// langdetect.go defines the interface for language detection. After speech is
// transcribed, this determines whether the speaker used English, Spanish, or
// switched between both mid-sentence. The result tells the pipeline which
// language to respond in.
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
