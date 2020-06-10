package client

import (
	"fmt"
	"gotest.tools/assert"
	"testing"
)

func TestGetDefaultApiExtensionClientSet(t *testing.T) {
	_, err := GetDefaultApiExtensionClientSet()
	assert.Equal(t, nil, err)
	var bs []byte
	_, err = GetApiExtensionClientSet(bs)
	assert.Equal(t, nil, err)
}

func TestGetDefaultK8sClientSet(t *testing.T) {
	for i := 0; i < 3; i++ {
		fmt.Println("i: ", i)
		for j := 0; j < 2; j++ {
			if j == 1 {

				break
			}
			fmt.Println("j: ", j)
		}
	}
}