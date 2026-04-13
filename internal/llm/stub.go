package llm

import "fmt"

type Stub struct{}

func NewStub() *Stub {
	return &Stub{}
}

func (s *Stub) GenerateResponse(userText string, history []ConversationTurn, sourceLang string) (Response, error) {
	return Response{
		Text:          fmt.Sprintf("[stub llm] You said: %s", userText),
		Language:      "en",
		ViolatesScope: false,
	}, nil
}
