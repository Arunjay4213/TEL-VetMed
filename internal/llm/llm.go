// llm.go defines the interface for the Large Language Model reasoning core.
// The LLM is the brain of the system -- it takes what the user said plus the
// conversation history and generates an appropriate response. It is strictly
// a communication facilitator and must never diagnose, prescribe, or make
// clinical decisions. Those always stay with the veterinarian.
package llm

type Response struct {
	Text           string
	Language       string
	ViolatesScope  bool
	DeferralReason string
}

type ConversationTurn struct {
	Role     string
	Text     string
	Language string
}

type Service interface {
	GenerateResponse(userText string, history []ConversationTurn, sourceLang string) (Response, error)
}
