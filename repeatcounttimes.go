package cron

import "time"

type RepeatCountTimesSchedule struct {
	Start time.Time
	Delay time.Duration
	Count int
	valid bool
}

func RepeatCountTimes(delay time.Duration, count int) *RepeatCountTimesSchedule {
	return RepeatCountTimesFrom(time.Now(), delay, count)
}

func RepeatCountTimesFrom(start time.Time, delay time.Duration, count int) *RepeatCountTimesSchedule {
	if start.IsZero() {
		start = time.Now()
	}
	if delay < time.Second {
		delay = time.Second
	}
	if count < 1 {
		count = 1
	}
	return &RepeatCountTimesSchedule{
		Start: start.Truncate(time.Second),
		Delay: delay - time.Duration(delay.Nanoseconds())%time.Second,
		Count: count,
		valid: true,
	}
}

func (s *RepeatCountTimesSchedule) Next(t time.Time) time.Time {
	start := s.Start.In(t.Location())
	passedPeriods := t.Sub(start)
	if passedPeriods < s.Delay {
		passedPeriods = 0
	} else {
		passedPeriods = passedPeriods.Truncate(s.Delay)
	}
	lastPeriod := s.Delay * time.Duration(s.Count)
	if passedPeriods >= lastPeriod {
		s.valid = false
		return time.Time{}
	}
	return start.Add(passedPeriods + s.Delay)
}

func (s *RepeatCountTimesSchedule) Valid() bool {
	return s.valid
}
