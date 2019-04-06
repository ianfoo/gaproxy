package session

import (
	"errors"
	"math/rand"
	"sync"
	"time"
)

type Manager struct {
	rand     *rand.Rand
	mu       *sync.RWMutex
	sessions map[string]Session
}

func NewManager(options ...func(*Manager)) Manager {
	m := Manager{
		rand:     rand.New(rand.NewSource(time.Now().UnixNano())),
		mu:       &sync.RWMutex{},
		sessions: make(map[string]Session),
	}
	for _, o := range options {
		o(&m)
	}
	return m
}

func WithRandSource64(src rand.Source64) func(*Manager) {
	return func(m *Manager) {
		m.rand = rand.New(src)
	}
}

func (m Manager) New(lifetime time.Duration, auth0Identity string) (Session, error) {
	s, err := func() (Session, error) {
		m.mu.RLock()
		defer m.mu.RUnlock()
		for i := 0; i < 100; i++ {
			session := newSession(m.rand, lifetime, auth0Identity)
			if _, ok := m.sessions[session.id]; !ok {
				return session, nil
			}
		}
		return Session{}, errors.New("unable to create unique session")
	}()
	if err != nil {
		return Session{}, err
	}

	m.mu.Lock()
	m.sessions[s.id] = s
	m.mu.Unlock()

	go m.reaper()

	return s, nil
}

func (m Manager) reaper() {
	t := time.Tick(time.Minute)
	for range t {
		m.reapExpired()
	}
}

func (m Manager) reapExpired() {
	m.mu.RLock()
	expired := make([]string, len(m.sessions)/10)
	now := time.Now()
	for _, s := range m.sessions {
		if s.ExpiresAt().Before(now) {
			expired = append(expired, s.ID())
		}
	}
	m.mu.RUnlock()
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, id := range expired {
		if m.sessions[id].ExpiresAt().Before(now) {
			delete(m.sessions, id)
		}
	}
}

func (m Manager) Get(id string) (Session, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	s, ok := m.sessions[id]
	if !ok {
		return Session{}, errors.New("session does not exist")
	}
	return s, nil
}

func (m Manager) Invalidate(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	s, ok := m.sessions[id]
	if !ok {
		return errors.New("session does not exist")
	}
	s.valid = false
	s.expiresAt = time.Now()
	m.sessions[id] = s
	return nil
}
