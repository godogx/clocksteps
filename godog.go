package clocksteps

import (
	"context"
	"strconv"
	"time"

	"github.com/cucumber/godog"
	"go.nhat.io/timeparser"
)

// RegisterContext registers clock to godog tests.
//
// Deprecated: use RegisterSteps instead.
func (c *Clock) RegisterContext(s *godog.ScenarioContext) {
	c.RegisterSteps(s)
}

// RegisterSteps registers clock to godog tests.
func (c *Clock) RegisterSteps(s *godog.ScenarioContext) {
	s.After(func(context.Context, *godog.Scenario, error) (context.Context, error) {
		// Unfreeze the clock.
		c.Unfreeze()

		return nil, nil
	})

	s.Step(`(?:the )?clock is at "([^"]*)"`, c.set)
	s.Step(`(?:the )?clock is set to "([^"]*)"`, c.set)
	s.Step(`sets? (?:the )?clock to "([^"]*)"`, c.set)
	s.Step(`now is "([^"]*)"`, c.set)

	s.Step(`adds? ([^\s]*) to (?:the )?clock`, c.add)
	s.Step(`adds? ([0-9]+) days? to (?:the )?clock`, c.addDay)
	s.Step(`adds? ([0-9]+) months? to (?:the )?clock`, c.addMonth)
	s.Step(`adds? ([0-9]+) years? to (?:the )?clock`, c.addYear)
	s.Step(`adds? ([0-9]+) months?,? ([0-9]+) days? to (?:the )?clock`, c.addMonthDay)
	s.Step(`adds? ([0-9]+) years?,? ([0-9]+) days? to (?:the )?clock`, c.addYearDay)
	s.Step(`adds? ([0-9]+) years?,? ([0-9]+) months? to (?:the )?clock`, c.addYearMonth)
	s.Step(`adds? ([0-9]+) years?,? ([0-9]+) months?,? ([0-9]+) days? to (?:the )?clock`, c.addDate)

	s.Step(`\s*freeze (?:the )?clock`, c.freeze)
	s.Step(`(?:(?:release)|(?:unset)|(?:reset)) (?:the )?clock$`, c.unfreeze)
}

func (c *Clock) set(t string) error {
	ts, err := timeparser.Parse(t)
	if err != nil {
		return err
	}

	c.Set(ts)

	return nil
}

func (c *Clock) add(s string) error {
	d, err := time.ParseDuration(s)
	if err != nil {
		return err
	}

	return c.Add(d)
}

func (c *Clock) addDay(days string) error {
	return c.addDate("0", "0", days)
}

func (c *Clock) addMonth(months string) error {
	return c.addDate("0", months, "0")
}

func (c *Clock) addYear(years string) error {
	return c.addDate(years, "0", "0")
}

func (c *Clock) addMonthDay(months, days string) error {
	return c.addDate("0", months, days)
}

func (c *Clock) addYearDay(years, days string) error {
	return c.addDate(years, "0", days)
}

func (c *Clock) addYearMonth(years, months string) error {
	return c.addDate(years, months, "0")
}

func (c *Clock) addDate(years, months, days string) error {
	y, err := strconv.Atoi(years)
	if err != nil {
		return err
	}

	m, err := strconv.Atoi(months)
	if err != nil {
		return err
	}

	d, err := strconv.Atoi(days)
	if err != nil {
		return err
	}

	return c.AddDate(y, m, d)
}

func (c *Clock) freeze() error {
	c.Freeze()

	return nil
}

func (c *Clock) unfreeze() error {
	c.Freeze()

	return nil
}
