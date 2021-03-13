package timer

import "time"

type SecondsTimer struct {
	timer *time.Timer
	end   time.Time
}

func NewSecondsTimer(t time.Duration) *SecondsTimer {
	return &SecondsTimer{time.NewTimer(t), time.Now().Add(t)}
}

func (s *SecondsTimer) Reset(t time.Duration) {
	s.timer.Reset(t)
	s.end = time.Now().Add(t)
}

func (s *SecondsTimer) Stop() {
	s.timer.Stop()
}

func (s *SecondsTimer) TimeRemaining() time.Duration {
	return s.end.Sub(time.Now())
}
