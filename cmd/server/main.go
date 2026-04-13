// main.go is the entry point for the veterinary voice agent server. It wires
// all the services together (ASR, language detection, LLM, translation, TTS),
// loads environment variables like the Deepgram API key, and starts an HTTP
// server with two endpoints: /health to check the server is running, and
// /interact to process a voice interaction.
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/Arunjay4213/vetmed/internal/asr"
	"github.com/Arunjay4213/vetmed/internal/langdetect"
	"github.com/Arunjay4213/vetmed/internal/llm"
	"github.com/Arunjay4213/vetmed/internal/pipeline"
	"github.com/Arunjay4213/vetmed/internal/session"
	"github.com/Arunjay4213/vetmed/internal/translation"
	"github.com/Arunjay4213/vetmed/internal/tts"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	apiKey := os.Getenv("DEEPGRAM_API_KEY")

	p := &pipeline.Pipeline{
		ASR:         asr.NewDeepgram(apiKey),
		LangDetect:  langdetect.NewStub(),
		LLM:         llm.NewStub(),
		Translation: translation.NewStub(),
		TTS:         tts.NewStub(),
		Sessions:    session.NewManager(),
	}

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	http.HandleFunc("/interact", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		sessionID := r.URL.Query().Get("session_id")
		if sessionID == "" {
			http.Error(w, "session_id is required", http.StatusBadRequest)
			return
		}

		audio, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "failed to read audio", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		result, err := p.ProcessInteraction(audio, sessionID)
		if err != nil {
			log.Printf("pipeline error: %v", err)
			http.Error(w, "processing failed", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"response_text":   result.ResponseText,
			"source_language": result.SourceLanguage,
			"target_language": result.TargetLanguage,
			"audio_format":    result.Audio.Format,
			"audio_length":    len(result.Audio.Audio),
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Veterinary Voice Agent server starting on :%s\n", port)
	fmt.Println("Endpoints:")
	fmt.Println("  GET  /health   - Health check")
	fmt.Println("  POST /interact - Process voice interaction")

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
