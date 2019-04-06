package session

import (
	"fmt"
	"math/rand"
	"time"
)

type Session struct {
	id            string
	valid         bool
	expiresAt     time.Time
	auth0Identity string

	// Google auth token is obtained from Auth0.
	googleAuthToken string
}

const (
	sessionIDLength = 36

	// SessionDefaultLifetime is how long a session will be valid for
	// if the request to create a new session specifies zero duration.
	SessionDefaultLifetime = time.Hour
)

func newSession(rand *rand.Rand, lifetime time.Duration, identity string) Session {
	if lifetime == 0 {
		lifetime = SessionDefaultLifetime
	}
	s := Session{
		id:            newSessionID(rand),
		valid:         true,
		expiresAt:     time.Now(),
		auth0Identity: identity,
	}
	return s
}

func newSessionID(rand *rand.Rand) string {
	return fmt.Sprintf("%016x%016x", rand.Uint64(), rand.Uint64())
}

func (s Session) ID() string {
	return s.id
}

func (s Session) Valid() bool {
	return s.valid
}

func (s Session) ExpiresAt() time.Time {
	return s.expiresAt
}

func (s Session) GoogleAuthToken() string {
	return s.googleAuthToken
}
