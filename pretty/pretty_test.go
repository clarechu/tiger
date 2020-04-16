package pretty

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"reflect"
	"testing"
)

type Foo struct {
	Name     string `json:"NAME"`
	To       *int   `json:"to"`
	Ready    string `json:"READY"`
	Status   string `json:"STATUS"`
	Restarts int32  `json:"RESTARTS"`
	age      string `json:"AGE"`
	Bar      Bar    `json:"BAR"`
	baz      Baz    `json:"BAZ"`
}

type Bar struct {
	Age string `json:"age" json:"age"`
}

type Baz struct {
	Age string `json:"age" json:"age"`
}

func TestMax(t *testing.T) {
	//to := 1
	foos := &[]Foo{{
		Name:     "manager-v1-5955846b7c-29nm8",
		Ready:    "1/1",
		Status:   "Running",
		Restarts: 1,
		age:      "34d",
		//To:       &to,
	}, {
		Name:     "manager-v1-5955846b7c-29nm8",
		Ready:    "1/1",
		Status:   "Running",
		Restarts: 1,
		age:      "34d",
	},
	}
	pre := New(foos).ToSlice()
	fmt.Println(pre.data)
	assert.Equal(t, nil, pre.err)
}

func TestGetSize(t *testing.T) {
	size := getSize("manager-v1-5955846b7c-29nm8", true)
	assert.Equal(t, size, 27)
	size = getSize(1, true)
	assert.Equal(t, size, 1)
	size = getSize(nil, true)
	assert.Equal(t, size, 5)
	to := 1
	size = getSize(&to, true)
	assert.Equal(t, size, 1)
	size = getSize("to", true)
	a := Foo{}
	size = getSize(a.To, true)
	assert.Equal(t, size, 5)
}

func TestEnd(t *testing.T) {
	to := 1
	too := &to
	//tooo := &too
	//toooo := &tooo
	foos := &[]Foo{{
		Name:     "manager-v1-5955846b7c-29nm8",
		Ready:    "1/1",
		Status:   "Running",
		Restarts: 1,
		age:      "34d",
		To:       too,
	}, {
		Name:     "manager-v1-5955846b7c-29nm8",
		Ready:    "1/1",
		Status:   "Running",
		Restarts: 1,
		age:      "34d",
	},
	}
	err := New(foos).IsReal().Print()
	assert.Equal(t, nil, err)
}

func TestTrimSpacePre(t *testing.T) {
	to := 1
	too := &to
	tooo := &too
	toooo := &tooo
	a := realVal(reflect.ValueOf(toooo))
	fmt.Println(a)

	a = realVal(reflect.ValueOf(too))
	fmt.Println(a)

	a = realVal(reflect.ValueOf(to))
	fmt.Println(a)
}

func TestPrint(t *testing.T) {
	to := 1
	too := &to
	//tooo := &too
	//toooo := &tooo
	foos := &[]Foo{{
		Name:     "manager-v1-5955846b7c-29nm8",
		Ready:    "1/1",
		Status:   "Running",
		Restarts: 1,
		age:      "34d",
		To:       too,
	}, {
		Name:     "manager-v1-5955846b7c-29nm8",
		Ready:    "1/1",
		Status:   "Running",
		Restarts: 1,
		age:      "34d",
	},
	}
	err := New(foos).Print()
	assert.Equal(t, nil, err)
}
