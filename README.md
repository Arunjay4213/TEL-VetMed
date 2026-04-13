# Bilingual Veterinary Voice Agent

A bilingual English–Spanish voice agent for dairy farm practice, built at the University of Wisconsin–Madison School of Veterinary Medicine. It helps English-speaking veterinarians communicate with Spanish-speaking farm workers in real time during herd health visits, clinical procedures, and emergencies.

## What it does

A vet or farm worker speaks into a device. The system transcribes the audio, detects the language, generates an appropriate response via an LLM, translates if needed, and speaks the response back — all in under 3 seconds.

```
Audio → ASR → Language Detection → LLM → Translation → TTS → Audio
                      Session Manager tracks conversation history
```

## Project structure

```
cmd/
  server/
    main.go           Entry point. Starts the HTTP server and wires all services together.

internal/
  asr/
    asr.go            Shared types and Service interface for speech recognition.
    deepgram.go       Real ASR implementation using Deepgram Nova-3 (multilingual).
    keyterms.go       Veterinary keyterms grouped by farm environment for Deepgram boosting.
    stub.go           Fake ASR for development/testing without a microphone.

  langdetect/
    langdetect.go     Interface for language detection (English, Spanish, or mixed).
    stub.go           Fake detector that always returns English.

  llm/
    llm.go            Interface for the LLM reasoning core.
    stub.go           Fake LLM that echoes input back.

  translation/
    translation.go    Interface for English–Spanish translation.
    stub.go           Fake translator for development.

  tts/
    tts.go            Interface for Text-to-Speech synthesis.
    stub.go           Fake TTS for development.

  pipeline/
    pipeline.go       Orchestrates all services for a single conversation turn.

  session/
    session.go        Tracks conversation history across turns within a session.
```

## Getting started

### Prerequisites

- Go 1.26+
- A Deepgram API key

### Setup

1. Clone the repo:
   ```bash
   git clone https://github.com/Arunjay4213/vetmed.git
   cd vetmed
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Create a `.env` file in the project root (never commit this):
   ```
   DEEPGRAM_API_KEY=your_key_here
   ```

4. Run the server:
   ```bash
   go run ./cmd/server/
   ```

   You should see:
   ```
   Veterinary Voice Agent server starting on :8080
   ```

### Build

```bash
go build ./...
```

## API endpoints

### `GET /health`
Check the server is running.

```bash
curl http://localhost:8080/health
# {"status":"ok"}
```

### `POST /interact?session_id=<id>`
Send raw audio (WAV, 16-bit PCM, 16kHz+) and get back a JSON response.

```bash
curl -X POST "http://localhost:8080/interact?session_id=session1" \
  --data-binary @audio.wav
```

Response:
```json
{
  "response_text": "...",
  "source_language": "es",
  "target_language": "en",
  "audio_format": "wav",
  "audio_length": 12400
}
```

The `session_id` ties multiple turns together so the system remembers context across a conversation. Use a new ID to start a fresh session.

## Services and stubs

Every service (ASR, LLM, translation, TTS) has a real implementation and a stub. The server currently uses:

- **Deepgram Nova-3** for ASR (real)
- **Stubs** for everything else (LLM, translation, TTS)

To swap in a real implementation, replace the relevant `NewStub()` call in `cmd/server/main.go` with the real constructor. Nothing else needs to change.

## Keyterm environments

The ASR is configured with veterinary-specific keyterms grouped by farm environment to improve recognition accuracy. Pass the environment name when calling `Transcribe`:

| Environment | Description |
|---|---|
| `milking_parlor` | Mastitis, somatic cell counts, teat dipping |
| `general_barn` | Downer cows, hypocalcemia, displaced abomasum |
| `hoof_trimming` | Sole ulcers, digital dermatitis, lameness scoring |
| `calving` | Dystocia, OB chains, colostrum, calf care |
| `treatment_pen` | Medication administration, jugular injections |
| `breeding` | AI protocols, pregnancy checks, OvSynch |
| `biosecurity` | Isolation, disinfection, quarantine |

Leave the environment empty (`""`) to use only the global keyterms (drug names, withdrawal periods, common Spanish terms).

## Important constraints

This system is a **communication facilitator only**. It must never:
- Diagnose a condition
- Recommend a specific treatment
- Calculate drug dosages
- Provide a prognosis

All clinical decisions remain the responsibility of the attending licensed veterinarian.

## Environment variables

| Variable | Required | Description |
|---|---|---|
| `DEEPGRAM_API_KEY` | Yes | Deepgram API key for speech recognition |
| `PORT` | No | HTTP server port (default: `8080`) |
