package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCronMessages(t *testing.T) {
	data := []string{"MSG_11_30__TESTS=ahoj"}
	cronMessages := EnvToCronMessages(data)

	assert.Equal(t, "30 11 * * MON-FRI *", cronMessages[0].CronLine)
}

func TestWronNameOfEnvVariableForCronMessages(t *testing.T) {
	data := []string{"MSG_11_30=ahoj"}
	cronMessages := EnvToCronMessages(data)

	assert.Equal(t, 0, len(cronMessages))
}
