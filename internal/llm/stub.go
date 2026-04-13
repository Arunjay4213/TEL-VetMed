// stub.go is a fake LLM that echoes back whatever the user said.
// Used during development so the pipeline runs without a real API call.
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
