package pretty

import (
	"context"
	"errors"
	"fmt"
	"reflect"
)

type PrettyInterface interface {
	New(o interface{}) (pretty *Pretty)
	Print() error
}

type Pretty struct {
	o      interface{}
	data   [][]interface{}
	length []int
	err    error
	real   bool
}

type TAG string

type SizeError string

const (
	DefaultTag  TAG    = "json"
	SizeInvalid string = "<invalid reflect.Value>"
	SizeNil     int    = 5
)

type data struct {
	length int
	data   []string
}

//Set 初始化
func New(o interface{}) (pretty *Pretty) {
	return &Pretty{o: o}
}

func (pretty *Pretty) IsReal() *Pretty {
	pretty.real = true
	return pretty
}

func (pretty *Pretty) Print() error {
	pretty.ToSlice()
	if pretty.err != nil {
		return pretty.err
	}
	pretty.max().Run()
	return nil
}

//Max
func (pretty *Pretty) max() *Pretty {
	length := []int{}
	// O = i*j
	for i := 0; i < len(pretty.data); i++ {
		for j := 0; j < len(pretty.data[i]); j++ {
			size := getSize(pretty.data[i][j], pretty.real)
			if i == 0 {
				length = append(length, size)
			} else if size > length[j] {
				length[j] = size
			}
		}
	}
	pretty.length = length
	return pretty
}

func (pretty *Pretty) ToSlice(tag ...TAG) *Pretty {
	dx, dy := 0, 0
	//set heard
	value := indirect(reflect.ValueOf(pretty.o))
	// i is size =0
	if value.Kind() == reflect.Slice {
		length := value.Len()
		if length == 0 {
			pretty.err = errors.New("slice length size 0")
			return pretty
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

	//set body
	for i := 0; i < value.Len(); i++ {
		types := value.Index(i).Type()
		values := value.Index(i)
		for j := 0; j < types.NumField(); j++ {
			data[i+1][j] = getValues(values.Field(j))
		}
	}
	pretty.data = data
	return pretty
}

func indirect(reflectValue reflect.Value) reflect.Value {
	for reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}
	return reflectValue
}

//GetValues 获取当前字段的值
func getValues(values reflect.Value) interface{} {
	defer func() {
		err := recover()
		if err != nil {
		}
	}()
	return values.Interface()
}

func getSize(o interface{}, real bool) int {
	if o == nil {
		return SizeNil
	}
	if real && reflect.ValueOf(o).Kind() == reflect.Ptr {
		return getSize(reflect.ValueOf(o).Elem(), real)
	}
	s := fmt.Sprintf("%v", o)
	if s == SizeInvalid {
		return SizeNil
	}
	b := []byte(s)
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

func (pretty *Pretty) Run() {
	p := ""
	for i := 0; i < len(pretty.data); i++ {
		for j := 0; j < len(pretty.data[i]); j++ {
			d := ""
			if !pretty.real {
				d = TrimSpace(pretty.data[i][j], pretty.length[j], pretty.real)
			} else {
				d = TrimSpacePre(pretty.data[i][j], pretty.length[j], pretty.real)
			}

			p = fmt.Sprintf("%v%v", p, d)
		}
		fmt.Println(p)
		p = ""
	}
}

func TrimSpace(d interface{}, length int, real bool) string {
	p := ""
	for i := getSize(d, real); i < length+2; i++ {
		p = fmt.Sprintf("%v%s", p, " ")
	}
	return fmt.Sprintf("%v%v", d, p)
}

func TrimSpacePre(d interface{}, length int, real bool) string {
	p := ""
	for i := getSize(d, real); i < length+2; i++ {
		p = fmt.Sprintf("%v%s", p, " ")
	}
	d = realVal(reflect.ValueOf(d))
	return fmt.Sprintf("%v%v", d, p)
}

func realVal(value reflect.Value) (o interface{}) {
	if value.Kind() == reflect.Ptr {
		return realVal(value.Elem())
	}
	o = getValues(value)
	return
}

func A(ch chan int, ctx context.Context) {
	for {
		select {
		case c := <-ch:
			go func() {
				fmt.Println("chan:", c)
			}()
		case <-ctx.Done():
			fmt.Println("down")
			return
		}
	}
}
