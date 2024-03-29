# Cucumber Clock steps for Golang

[![GitHub Releases](https://img.shields.io/github/v/release/godogx/clocksteps)](https://github.com/godogx/clocksteps/releases/latest)
[![Build Status](https://github.com/godogx/clocksteps/actions/workflows/test.yaml/badge.svg)](https://github.com/godogx/clocksteps/actions/workflows/test.yaml)
[![codecov](https://codecov.io/gh/godogx/clocksteps/branch/master/graph/badge.svg?token=eTdAgDE2vR)](https://codecov.io/gh/godogx/clocksteps)
[![Go Report Card](https://goreportcard.com/badge/github.com/godogx/clocksteps)](https://goreportcard.com/report/github.com/godogx/clocksteps)
[![GoDevDoc](https://img.shields.io/badge/dev-doc-00ADD8?logo=go)](https://pkg.go.dev/github.com/godogx/clocksteps)

`clocksteps` uses [`go.nhat.io/clock`](https://go.nhat.io/clock) to provide steps for `cucumber/godog` and
makes it easy to run tests with `time`.

## Prerequisites

- `Go >= 1.16`

## Usage

Initiate the clock and register it to the scenario.

```go
package mypackage

import (
	"testing"

	"github.com/cucumber/godog"
	"github.com/godogx/clocksteps"
)

func TestIntegration(t *testing.T) {
	clock := clocksteps.New()
	suite := godog.TestSuite{
		Name:                 "Integration",
		TestSuiteInitializer: nil,
		ScenarioInitializer: func(ctx *godog.ScenarioContext) {
			clock.RegisterContext(ctx)
		},
		Options: &godog.Options{
			Strict:    true,
			Output:    out,
			Randomize: rand.Int63(),
		},
	}
	
	// Inject the clock to your application then run the suite.
	status := suite.Run()
}
```

Read more about [`go.nhat.io/clock`](https://go.nhat.io/clock)

## Steps

### Set the time

By default, the clock always returns `time.Now()` unless you freeze or set it. For setting, you can use one of
these:
- `(?:the )?clock is at "([^"]*)"`
- `(?:the )?clock is set to "([^"]*)"`
- `sets? (?:the )?clock to "([^"]*)"`
- `now is "([^"]*)"`

They have the same effect, the clock will be set at a specific `time.Time`. The given can be in `RFC3339` 
(`2006-01-02T15:04:05Z07:00`) or `YMD` (`2006-01-02`)

For example:

```gherkin
    Scenario: Set time
        Given the clock is at "2020-01-02T03:04:05Z"
        Then the time is "2020-01-02T03:04:05Z"

        Given the clock is set to "2020-02-03T04:05:06Z"
        Then the time is "2020-02-03T04:05:06Z"

        Given Someone sets the clock to "2020-03-04T05:06:07Z"
        Then the time is "2020-03-04T05:06:07Z"

        Given now is "2020-04-05T06:07:08Z"
        Then the time is "2020-04-05T06:07:08Z"
```

### Adjust the time

After setting the clock, you can adjust the time by adding a `time.Duration`, some days, months or years with these:
- `adds? ([^\s]*) to (?:the )?clock`
- `adds? ([0-9]+) days? to (?:the )?clock`
- `adds? ([0-9]+) months? to (?:the )?clock`
- `adds? ([0-9]+) years? to (?:the )?clock`
- `adds? ([0-9]+) months?,? ([0-9]+) days? to (?:the )?clock`
- `adds? ([0-9]+) years?,? ([0-9]+) days? to (?:the )?clock`
- `adds? ([0-9]+) years?,? ([0-9]+) months? to (?:the )?clock`
- `adds? ([0-9]+) years?,? ([0-9]+) months?,? ([0-9]+) days? to (?:the )?clock`

**Important**: You have to set the clock before adjusting it. Otherwise you will get `clocksteps.ErrClockIsNotSet` 

For example:

```gherkin
    Scenario: Add time
        Given the clock is at "2020-01-02T03:04:05Z"
        And someone adds 1h5s to the clock
        Then the time is "2020-01-02T04:04:10Z"

        Given someone adds 2 days to the clock
        Then the time is "2020-01-04T04:04:10Z"
```

### Freeze and release the clock

```gherkin
    Scenario: Freeze and Release
        Given the time is now

        When I freeze the clock
        And I wait for 50ms
        Then the time is not now

        When I wait for 50ms
        And I release the clock
        Then the time is now
```
