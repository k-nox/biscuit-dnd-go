package app_test

import (
	"bytes"
	"testing"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/k-nox/biscuit-dnd-go/app"
	"github.com/k-nox/biscuit-dnd-go/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// dummy test just to get it up & running.
func TestApp_Run(t *testing.T) {
	if !testing.Short() {
		var cfg config.Config
		err := cleanenv.ReadConfig("../config/config.yaml", &cfg)
		require.NoError(t, err)

		var output bytes.Buffer
		app, err := app.New(cfg, &output)
		require.NoError(t, app.Run())
		expected := "Barbarian\n"

		assert.Equal(t, expected, output.String())
	}
}
