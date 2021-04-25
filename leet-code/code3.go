/*
Copyright 2020 The go-harbor Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
*/

package leet_code

// 获取0-n之间的所有偶数
func even(a int) (array []int) {
	for i := 0; i < a; i++ {
		// 同时为0 的时候就为0
		if i&1 == 0 { // 位操作符&与C语言中使用方式一样
			array = append(array, i)
		}
	}
	return array
}

// 获取0-n之间的所有基数
func even1(a int) (array []int) {
	for i := 0; i < a; i++ {
		// 同时为0 的时候就为0
		if i|1 == 1 { // 位操作符&与C语言中使用方式一样
			array = append(array, i)
		}
	}
	return array
}

func singleNumber(nums []int) []int {
	v := 0
	for _, n := range nums {
		v = v ^ n
	}
	r := 1
	for r&v == 0 {
		r = r << 1
	}
	a, b := 0, 0
	for _, n := range nums {
		if r&n != 0 {
			a = a ^ n
		} else {
			b = b ^ n
		}
	}
	return []int{a, b}
}
