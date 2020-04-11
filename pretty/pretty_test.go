package pretty

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
)

type Foo struct {
	Name     string `json:"name" json:"name"`
	Ready    string `json:"ready" json:"ready"`
	Status   string `json:"status" json:"status"`
	Restarts int32  `json:"restarts" json:"restarts"`
	Age      string `json:"age" json:"age"`
}

type Bar struct {
}

func TestMax(t *testing.T) {
	foos := &[]Foo{{
		Name:     "manager-v1-5955846b7c-29nm8",
		Ready:    "1/1",
		Status:   "Running",
		Restarts: 1,
		Age:      "34d",
	}, {
		Name:     "manager-v1-5955846b7c-29nm8",
		Ready:    "1/1",
		Status:   "Running",
		Restarts: 1,
		Age:      "34d",
	},
	}
	data, err := ToSlice(foos)
	fmt.Println(data)
	assert.Equal(t, nil, err)
	//Max(foo)
}

func TestGetSize(t *testing.T) {
	size := GetSize( "manager-v1-5955846b7c-29nm8")
	assert.Equal(t, size, 27)
	size = GetSize( 1)
	assert.Equal(t, size, 1)
}