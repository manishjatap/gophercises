package main

import (
	"fmt"
	"mongopkg"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var taskList []mongopkg.Task

type taskMock struct {
	mock.Mock
}

func (t taskMock) Insert() {
	fmt.Println("-- Mock : mongopkg.Insert() --")
	taskList = []mongopkg.Task{
		{"Go to office-mock", "2019-10-30 19:39:16.439946488 +0530 IST m=+0.005931177", "incomplete", ""},
		{"Go to school-mock", "2019-10-30 19:39:16.439946488 +0530 IST m=+0.005931177", "incomplete", ""},
		{"Go to gym-mock", "2019-10-30 19:39:16.439946488 +0530 IST m=+0.005931177", "incomplete", ""},
	}
}

func (t taskMock) Get() []mongopkg.MongoOps {
	fmt.Println("-- Mock : mongopkg.Get() --")
	tsOp := make([]mongopkg.MongoOps, 1)
	tsOp[0] = taskMock{}
	return tsOp
}

func (t taskMock) Delete() {
	fmt.Println("-- Mock : mongopkg.Delete() --")
	taskList = taskList[:2]
}

func (t taskMock) Update() {
	fmt.Println("-- Mock : mongopkg.Update() --")
	taskList[0].Status = "completed"
	taskList[0].CompletionDate = "2019-10-30 19:39:16.439946488 +0530 IST m=+0.005931177"
}

func TestAdd(t *testing.T) {
	var myTask taskMock
	add(myTask)
	assert.NotEmpty(t, taskList)
}

func TestDo(t *testing.T) {
	var myTask taskMock
	do(myTask, 1)
	assert.Equal(t, taskList[0].Status, "completed", "Task status should be updated")
}

// func TestList(t *testing.T) {
// 	var myTask taskMock
// 	list(myTask)
// 	assert.True(t, true)
// }
