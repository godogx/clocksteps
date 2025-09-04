package clocksteps_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/godogx/clocksteps"
)

func TestClock(t *testing.T) {
	t.Parallel()

	c := clocksteps.New()

	now := time.Now()

	assert.True(t, now.Before(c.Now()))

	// Errors while adding time to a live clock.
	require.ErrorIs(t, c.Add(time.Hour), clocksteps.ErrClockIsNotSet)
	require.ErrorIs(t, c.AddDate(0, 0, 1), clocksteps.ErrClockIsNotSet)

	// Freeze the clock.
	c.Freeze()

	ts := c.Now()

	<-time.After(50 * time.Millisecond)

	assert.Equal(t, ts, c.Now())

	// Set to another time.
	ts = time.Date(2020, 1, 2, 3, 4, 5, 0, time.UTC)

	c.Set(ts)

	<-time.After(50 * time.Millisecond)

	assert.Equal(t, ts, c.Now())

	// Change the time.
	ts = ts.Add(2 * time.Hour)
	err := c.Add(2 * time.Hour)
	require.NoError(t, err)

	<-time.After(50 * time.Millisecond)

	assert.Equal(t, ts, c.Now())

	// Change the date.
	ts = ts.AddDate(2, 1, 3)
	err = c.AddDate(2, 1, 3)
	require.NoError(t, err)

	<-time.After(50 * time.Millisecond)

	assert.Equal(t, ts, c.Now())

	// Add more timestamps.
	ts2 := time.Date(2021, 2, 3, 4, 5, 6, 0, time.UTC)
	c.Next(ts2)

	oldTs := c.Now()

	assert.Equal(t, ts, oldTs)
	assert.NotEqual(t, ts2, oldTs)
	assert.Equal(t, ts2, c.Now())
	assert.Equal(t, ts2, c.Now())

	// Unfreeze the clock.
	c.Unfreeze()

	now = time.Now()

	assert.True(t, now.Before(c.Now()))
}

func TestClock_Clock(t *testing.T) {
	t.Parallel()

	c := clocksteps.New()

	assert.Equal(t, c, c.Clock())
}
