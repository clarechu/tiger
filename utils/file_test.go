package utils

import (
	"github.com/bmizerany/assert"
	"testing"
)

const (
	old = "${SONARQUBE_HOST}"
	new = "10.10.13.888456"
)

func TestReplace(t *testing.T) {
	file := "./script.yml"
	newFile := "./new-script.yml"
	out, err := ReadFile(file, old, new)
	assert.Equal(t, nil, err)
	err = WriteToFile(newFile, out)
	assert.Equal(t, nil, err)
}
