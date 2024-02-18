package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewErrorJson(t *testing.T) {

	r := jsonError("error")
	expected := []byte(`{"message":"error"}`)

	assert.Equal(t, expected, r)
}
