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

type Queue struct {
	tasks []int
}

func NewQueue() *Queue {
	return &Queue{
		tasks: make([]int, 0),
	}
}

func (q *Queue) push(a int) {
	q.tasks = append(q.tasks, a)
}

func (q *Queue) pop() int {
	if q.size() == 0 {
		return 0
	}
	var task int
	task, q.tasks = q.tasks[0], q.tasks[1:]
	return task
}

func (q *Queue) size() int {
	return len(q.tasks)
}
