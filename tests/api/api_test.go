package tests

import (
	"github.com/hirsch88/go-micro-framework/tests/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	app := testUtils.BootstrapApp()

	w := testUtils.Get(app, "/api/ping")

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
