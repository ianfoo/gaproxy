package session_test

import (
	"testing"
	"time"

	"github.com/ianfoo/gaproxy/session"
	"github.com/matryer/is"
)

type testRandSource64 struct{}

func (testRandSource64) Seed(_ int64) {
}

func (testRandSource64) Int63() int64 {
	return 0
}

func (testRandSource64) Uint64() uint64 {
	return 0
}

func TestNewSession(t *testing.T) {
	is := is.New(t)

	m := session.NewManager(session.WithRandSource64(testRandSource64{}))
	now := time.Now()

	s, err := m.New(time.Hour, "session-test")
	is.NoErr(err)
	is.Equal(s.ID(), "00000000000000000000000000000000")
	is.True(s.ExpiresAt().After(now.Add(time.Minute)))
}
