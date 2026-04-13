package pipeline

import (
	"fmt"

	"github.com/Arunjay4213/vetmed/internal/asr"
	"github.com/Arunjay4213/vetmed/internal/langdetect"
	"github.com/Arunjay4213/vetmed/internal/llm"
	"github.com/Arunjay4213/vetmed/internal/session"
	"github.com/Arunjay4213/vetmed/internal/translation"
	"github.com/Arunjay4213/vetmed/internal/tts"
)

const ConfidenceThreshold = 0.6

type Pipeline struct {
	ASR         asr.Service
	LangDetect  langdetect.Service
	LLM         llm.Service
	Translation translation.Service
	TTS         tts.Service
	Sessions    *session.Manager
}

type Result struct {
	Audio          tts.AudioOutput
	ResponseText   string
	SourceLanguage string
	TargetLanguage string
}

func (p *Pipeline) ProcessInteraction(audio []byte, sessionID string) (Result, error) {
	transcript, err := p.ASR.Transcribe(audio)
	if err != nil {
		return Result{}, fmt.Errorf("asr failed: %w", err)
	}

	if transcript.Confidence < ConfidenceThreshold {
		return p.requestRepeat(sessionID)
	}

	langResult, err := p.LangDetect.Detect(transcript.Text)
	if err != nil {
		return Result{}, fmt.Errorf("language detection failed: %w", err)
	}
	sourceLang := langResult.PrimaryLanguage

	sess, exists := p.Sessions.Get(sessionID)
	if !exists {
		sess = p.Sessions.Create(sessionID)
	}

	p.Sessions.AddTurn(sessionID, llm.ConversationTurn{
		Role:     "user",
		Text:     transcript.Text,
		Language: sourceLang,
	})

	response, err := p.LLM.GenerateResponse(transcript.Text, sess.History, sourceLang)
	if err != nil {
		return Result{}, fmt.Errorf("llm failed: %w", err)
	}

	responseText := response.Text

	targetLang := determineTargetLanguage(sourceLang)
	if response.Language != targetLang {
		transResult, err := p.Translation.Translate(responseText, response.Language, targetLang)
		if err != nil {
			return Result{}, fmt.Errorf("translation failed: %w", err)
		}
		responseText = transResult.Text
	}

	audioOut, err := p.TTS.Synthesize(responseText, targetLang)
	if err != nil {
		return Result{}, fmt.Errorf("tts failed: %w", err)
	}

	p.Sessions.AddTurn(sessionID, llm.ConversationTurn{
		Role:     "assistant",
		Text:     responseText,
		Language: targetLang,
	})

	return Result{
		Audio:          audioOut,
		ResponseText:   responseText,
		SourceLanguage: sourceLang,
		TargetLanguage: targetLang,
	}, nil
}

func (p *Pipeline) requestRepeat(sessionID string) (Result, error) {
	repeatText := "Could you please repeat that?"
	lang := "en"

	audioOut, err := p.TTS.Synthesize(repeatText, lang)
	if err != nil {
		return Result{}, fmt.Errorf("tts failed for repeat request: %w", err)
	}

	return Result{
		Audio:          audioOut,
		ResponseText:   repeatText,
		SourceLanguage: lang,
		TargetLanguage: lang,
	}, nil
}

func determineTargetLanguage(sourceLang string) string {
	switch sourceLang {
	case "en":
		return "es"
	case "es":
		return "en"
	default:
		return "en"
	}
}
