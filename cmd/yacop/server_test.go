package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/xtradesoft/nexxchange-user-service/cmd/whs-user/config"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	if err := config.LoadConfig("../../config"); err != nil {
		assert.Error(t, err)
	}
	assert.NotNil(t, config.Config.App.Name)
}
