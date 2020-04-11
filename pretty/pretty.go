package pretty

import (
	"errors"
	"reflect"
	"unicode/utf8"
)

type Pretty struct {
	datas []data
}

type TAG string

const (
	DefaultTag TAG = "json"
)

type data struct {
	length int
	data   []string
}

func New(o interface{}) (pretty *Pretty) {
	return &Pretty{}
}

func Max(i []interface{}) int {
	var max int
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	for _, a := range i {
		kind := reflect.TypeOf(a).Elem().Kind()
		if kind == reflect.Struct || kind == reflect.Ptr {
			continue
		}

		str := a.(string)
		count := utf8.RuneCountInString(str)
		if count > max {
			max = count
		}
	}
	return max
}

func Length(l string) int32 {

	return 0
}

func ToSlice(i interface{}, tag ...TAG) ([][]interface{}, error) {
	dx, dy := 0, 0
	//set heard
	value := indirect(reflect.ValueOf(i))
	// i is size =0
	if value.Kind() == reflect.Slice {
		length := value.Len()
		if length == 0 {
			return nil, errors.New("slice length size 0")
		}
		dx = value.Index(0).Type().NumField()
		dy = length + 1
	}
	data := InitSlice(dx, dy)
	//set heard
	for i := 0; i < value.Len(); {
		types := value.Index(i).Type()
		for j := 0; j < types.NumField(); j++ {
			key := types.Field(j).Tag.Get(IsTag(tag...))
			data[0][j] = key
		}
		break
	}
	//tp := reflect.TypeOf(i)

	//set body
	for i := 0; i < value.Len(); i++ {
		types := value.Index(i).Type()
		values := value.Index(i)
		for j := 0; j < types.NumField(); j++ {
			key := values.Field(j).Interface()
			data[i+1][j] = key
		}
	}
	return data, nil
}

func indirect(reflectValue reflect.Value) reflect.Value {
	for reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}
	return reflectValue
}

/*
func isPointer(p reflect.Type) reflect.Value {
	kind := p.Kind()
	if kind == reflect.Ptr {
		p.Elem()
	}

	return
}
*/

func GetSize(o interface{}) int {
	b := []byte(o.(string))
	return len(b)
}



func InitSlice(dx, dy int) [][]interface{} {
	a := make([][]interface{}, dy)
	for i := range a {
		a[i] = make([]interface{}, dx)
	}
	return a
}

func IsTag(tags ...TAG) string {

	for _, tag := range tags {
		return string(tag)
	}
	return string(DefaultTag)
}
