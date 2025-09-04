package clocksteps

import (
	"errors"
	"sync"
	"time"

	"go.nhat.io/clock"
)

// ErrClockIsNotSet indicates that the clock must be set by either Clock.Set() or Clock.Freeze() before adding some
// time.Duration into it.
var ErrClockIsNotSet = errors.New("clock is not set")

var _ clock.Clock = (*Clock)(nil)

// Clock is a clock.Clock.
type Clock struct {
	timestamps []time.Time
	mu         sync.Mutex
}

// Now returns a fixed timestamp or time.Now().
func (c *Clock) Now() time.Time {
	c.mu.Lock()
	defer c.mu.Unlock()

	if len(c.timestamps) == 0 {
		return time.Now()
	}

	result := c.timestamps[0]

	if len(c.timestamps) > 1 {
		c.timestamps = c.timestamps[1:]
	}

	return result
}

// Set fixes the clock at a time.
func (c *Clock) Set(t time.Time) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.timestamps = []time.Time{t}
}

// Next sets the next timestamps to be returned by Now().
func (c *Clock) Next(t ...time.Time) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.timestamps = append(c.timestamps, t...)
}

// Add adds time to the clock.
func (c *Clock) Add(d time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if len(c.timestamps) == 0 {
		return ErrClockIsNotSet
	}

	c.timestamps[0] = c.timestamps[0].Add(d)

	return nil
}

// AddDate adds date to the clock.
func (c *Clock) AddDate(years, months, days int) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if len(c.timestamps) == 0 {
		return ErrClockIsNotSet
	}

	c.timestamps[0] = c.timestamps[0].AddDate(years, months, days)

	return nil
}

// Freeze freezes the clock.
func (c *Clock) Freeze() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.timestamps = []time.Time{time.Now()}
}

// Unfreeze unfreezes the clock.
func (c *Clock) Unfreeze() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.timestamps = nil
}

// Clock provides clock.Clock.
func (c *Clock) Clock() clock.Clock {
	return c
}

// New initiates a new Clock.
func New() *Clock {
	return &Clock{}
}
