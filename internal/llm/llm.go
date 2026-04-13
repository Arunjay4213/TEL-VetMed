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
