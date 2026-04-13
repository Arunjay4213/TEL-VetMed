package session

import (
	"sync"
	"time"

	"github.com/Arunjay4213/vetmed/internal/llm"
)

const MaxHistoryTurns = 10

type Session struct {
	ID        string
	History   []llm.ConversationTurn
	Language  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Manager struct {
	mu       sync.Mutex
	sessions map[string]*Session
}

func NewManager() *Manager {
	return &Manager{
		sessions: make(map[string]*Session),
	}
}

func (m *Manager) Create(id string) *Session {
	m.mu.Lock()
	defer m.mu.Unlock()

	s := &Session{
		ID:        id,
		History:   make([]llm.ConversationTurn, 0),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	m.sessions[id] = s
	return s
}

func (m *Manager) Get(id string) (*Session, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	s, exists := m.sessions[id]
	return s, exists
}

func (m *Manager) AddTurn(id string, turn llm.ConversationTurn) {
	m.mu.Lock()
	defer m.mu.Unlock()

	s, exists := m.sessions[id]
	if !exists {
		return
	}

	s.History = append(s.History, turn)

	if len(s.History) > MaxHistoryTurns {
		s.History = s.History[len(s.History)-MaxHistoryTurns:]
	}

	s.UpdatedAt = time.Now()
}

func (m *Manager) Delete(id string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.sessions, id)
}
