package main

import (
	"github.com/fitchlol/yacop/cmd/yacop/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	if err := config.LoadConfig("../../config"); err != nil {
		assert.Error(t, err)
	}
	assert.NotNil(t, config.Config.App.Name)
}
