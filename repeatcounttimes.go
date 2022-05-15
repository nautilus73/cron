package cron

import "time"

type RepeatCountTimesSchedule struct {
	Start time.Time
	Delay time.Duration
	Count int
	valid bool
}

func RepeatCountTimes(delay time.Duration, count int) RepeatCountTimesSchedule {
	return RepeatCountTimesFrom(time.Now(), delay, count)
}

func RepeatCountTimesFrom(start time.Time, delay time.Duration, count int) RepeatCountTimesSchedule {
	if start.IsZero() {
		start = time.Now()
	}
	if delay < time.Second {
		delay = time.Second
	}
	if count < 1 {
		count = 1
	}
	return RepeatCountTimesSchedule{
		Start: start.Truncate(time.Second),
		Delay: delay - time.Duration(delay.Nanoseconds())%time.Second,
		Count: count,
		valid: true,
	}
}

func (schedule RepeatCountTimesSchedule) Next(t time.Time) time.Time {
	start := schedule.Start.In(t.Location())
	passedPeriods := t.Sub(start)
	if passedPeriods < schedule.Delay {
		passedPeriods = 0
	} else {
		passedPeriods = passedPeriods.Truncate(schedule.Delay)
	}
	lastPeriod := schedule.Delay * time.Duration(schedule.Count)
	schedule.valid = passedPeriods < lastPeriod
	if !schedule.valid {
		return time.Time{}
	}
	return start.Add(passedPeriods + schedule.Delay)
}

func (schedule RepeatCountTimesSchedule) Valid() bool {
	return schedule.valid
}
