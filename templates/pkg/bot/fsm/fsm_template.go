package fsm

import "sync"

const (
	Idle              = "idle"
	WaitingForContent = "waiting_for_content"
)

type FSM struct {
	states map[int64]string
	mu     sync.Mutex
}

func NewFSM() *FSM {
	return &FSM{
		states: make(map[int64]string),
	}
}

func (f *FSM) SetState(chatID int64, state string) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.states[chatID] = state
}

func (f *FSM) GetState(chatID int64) string {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.states[chatID]
}

func (f *FSM) ClearState(chatID int64) {
	f.mu.Lock()
	defer f.mu.Unlock()
	delete(f.states, chatID)
}
