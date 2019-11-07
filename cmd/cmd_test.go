package cmd

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/manishjagtap/taskmanager/mongopkg"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

var tempFile = "./deleteMe.txt"

func TestDoCommandConversionActualImplementation(t *testing.T) {
	expectedOutput := "Error while calling actual implementation of stringconv.Atoi()"

	file, old := setStdoutToFile()

	var tempCmd *cobra.Command
	args := []string{"fake-task-1"}
	DoCommand.Run(tempCmd, args)

	op := resetStdoutAndGetFileContent(file, old)

	assert.Truef(t, op != "", "Expected : %v", expectedOutput)
}

func TestAddCommandActualImplementation(t *testing.T) {
	expectedOutput := "Error while calling actual implementation of vault.Insert()"

	file, old := setStdoutToFile()

	var tempCmd *cobra.Command
	args := []string{"fake-task-1"}
	AddCommand.Run(tempCmd, args)

	op := resetStdoutAndGetFileContent(file, old)

	assert.Truef(t, op != "", "Expected : %v", expectedOutput)
}

func TestListCommandActualImplemenation(t *testing.T) {
	expectedOutput := "Error while calling actual implementation of vault.Insert()"

	file, old := setStdoutToFile()

	var tempCmd *cobra.Command
	var args []string
	ListCommand.Run(tempCmd, args)

	op := resetStdoutAndGetFileContent(file, old)

	assert.Truef(t, op != "", "Expected : %v", expectedOutput)
}

func TestDoCommandActualImplementation(t *testing.T) {
	expectedOutput := "Error while calling actual implementation of vault.Update()"

	mockGetTask()
	mockConvertStringToInt()

	file, old := setStdoutToFile()

	var tempCmd *cobra.Command
	args := []string{"fake-task-1"}
	DoCommand.Run(tempCmd, args)

	op := resetStdoutAndGetFileContent(file, old)

	resetConvertStringToInt()
	resetGetTask()

	assert.Truef(t, op != "", "Expected : %v", expectedOutput)
}

func TestAddCommandSuccess(t *testing.T) {
	expectedOutput := "Task added!"

	mockInsertTask()

	file, old := setStdoutToFile()

	var tempCmd *cobra.Command
	args := []string{"fake-task-1"}
	AddCommand.Run(tempCmd, args)

	op := resetStdoutAndGetFileContent(file, old)

	resetInsertTask()

	assert.Equalf(t, expectedOutput, op, "Expected : %v", expectedOutput)
}

func TestDoCommandSuccess(t *testing.T) {
	expectedOutput := "Task completed!"

	mockGetTask()
	mockUpdateTask()
	mockConvertStringToInt()

	file, old := setStdoutToFile()

	var tempCmd *cobra.Command
	args := []string{"fake-task-1"}
	DoCommand.Run(tempCmd, args)

	op := resetStdoutAndGetFileContent(file, old)

	resetConvertStringToInt()
	resetUpdateTask()
	resetGetTask()

	assert.Equalf(t, expectedOutput, op, "Expected : %v", expectedOutput)
}

func TestListCommandSuccess(t *testing.T) {
	expectedOutput := "You have the following tasks"

	mockGetTask()

	file, old := setStdoutToFile()

	var tempCmd *cobra.Command
	var args []string
	ListCommand.Run(tempCmd, args)

	op := resetStdoutAndGetFileContent(file, old)

	resetGetTask()

	assert.Truef(t, strings.Contains(op, expectedOutput), "Expected : %v", expectedOutput)
}

func TestAddCommandInsertError(t *testing.T) {
	expectedOutput := "Error while adding new task"

	errorInsertTask(expectedOutput)

	file, old := setStdoutToFile()

	var tempCmd *cobra.Command
	args := []string{"fake-task-1"}
	AddCommand.Run(tempCmd, args)

	op := resetStdoutAndGetFileContent(file, old)

	resetInsertTask()

	assert.Equalf(t, expectedOutput, op, "Expected : %v", expectedOutput)
}

func TestDoCommandGetError(t *testing.T) {
	expectedOutput := "Error while fetching the tasks"

	errorGetTask(expectedOutput)
	mockUpdateTask()
	mockConvertStringToInt()

	file, old := setStdoutToFile()

	var tempCmd *cobra.Command
	args := []string{"fake-task-1"}
	DoCommand.Run(tempCmd, args)

	op := resetStdoutAndGetFileContent(file, old)

	resetConvertStringToInt()
	resetUpdateTask()
	resetGetTask()

	assert.Equalf(t, expectedOutput, op, "Expected : %v", expectedOutput)
}

func TestDoCommandUpdateError(t *testing.T) {
	expectedOutput := "Error while updating task"

	mockGetTask()
	errorUpdateTask(expectedOutput)
	mockConvertStringToInt()

	file, old := setStdoutToFile()

	var tempCmd *cobra.Command
	args := []string{"fake-task-1"}
	DoCommand.Run(tempCmd, args)

	op := resetStdoutAndGetFileContent(file, old)

	resetConvertStringToInt()
	resetUpdateTask()
	resetGetTask()

	assert.Equalf(t, expectedOutput, op, "Expected : %v", expectedOutput)
}

func TestDoCommandConversionError(t *testing.T) {
	expectedOutput := "Invalid task number provided"

	mockGetTask()
	mockUpdateTask()
	errorConvertStringToInt(expectedOutput)

	file, old := setStdoutToFile()

	var tempCmd *cobra.Command
	args := []string{"fake-task-1"}
	DoCommand.Run(tempCmd, args)

	op := resetStdoutAndGetFileContent(file, old)

	resetConvertStringToInt()
	resetUpdateTask()
	resetGetTask()

	assert.Equalf(t, expectedOutput, op, "Expected : %v", expectedOutput)
}

func TestListCommandList(t *testing.T) {
	expectedOutput := "Error while fetching the task list"

	errorGetTask(expectedOutput)

	file, old := setStdoutToFile()

	var tempCmd *cobra.Command
	var args []string
	ListCommand.Run(tempCmd, args)

	op := resetStdoutAndGetFileContent(file, old)

	resetGetTask()

	assert.Truef(t, strings.Contains(op, expectedOutput), "Expected : %v", expectedOutput)
}

func mockUpdateTask() {
	updateTask = func(task mongopkg.Task) error {
		return nil
	}
}

func errorUpdateTask(errMsg string) {
	updateTask = func(task mongopkg.Task) error {
		return errors.New(errMsg)
	}
}

func resetUpdateTask() {
	updateTask = func(task mongopkg.Task) error {
		return task.Update()
	}
}

func mockConvertStringToInt() {
	convertStringToInt = func(num string) (int, error) {
		return 1, nil
	}
}

func errorConvertStringToInt(errMsg string) {
	convertStringToInt = func(num string) (int, error) {
		return 1, errors.New(errMsg)
	}
}

func resetConvertStringToInt() {
	convertStringToInt = func(num string) (int, error) {
		return strconv.Atoi(strings.TrimSpace(num))
	}
}

func mockInsertTask() {
	insertTask = func(newTask mongopkg.Task) error {
		return nil
	}
}

func errorInsertTask(errMsg string) {
	insertTask = func(newTask mongopkg.Task) error {
		return errors.New(errMsg)
	}
}

func resetInsertTask() {
	insertTask = func(newTask mongopkg.Task) error {
		return newTask.Insert()
	}
}

func mockGetTask() {
	getTask = func(task mongopkg.Task) ([]mongopkg.Task, error) {
		taskList := []mongopkg.Task{
			{"fake-task", "fake-date", "Incomplete", "fake-date"},
		}
		return taskList, nil
	}
}

func errorGetTask(errMsg string) {
	getTask = func(task mongopkg.Task) ([]mongopkg.Task, error) {
		taskList := []mongopkg.Task{
			{"fake-task", "fake-date", "Incomplete", "fake-date"},
		}
		return taskList, errors.New(errMsg)
	}
}

func resetGetTask() {
	getTask = func(task mongopkg.Task) ([]mongopkg.Task, error) {
		return task.Get()
	}
}

func setStdoutToFile() (*os.File, *os.File) {
	var file *os.File

	if _, err := os.Stat(tempFile); os.IsNotExist(err) {
		file, _ = os.Create(tempFile)
	} else {
		file, _ = os.OpenFile(tempFile, os.O_WRONLY, 0777)
	}

	file.Truncate(0)
	file.Seek(0, 0)

	old := os.Stdout
	os.Stdout = file

	return file, old
}

func resetStdoutAndGetFileContent(file *os.File, old *os.File) string {
	//close old file
	file.Close()
	os.Stdout = old

	//open the same file in read mode
	file, _ = os.OpenFile(tempFile, os.O_RDONLY, 0777)
	reader := bufio.NewReader(file)
	op, _, _ := reader.ReadLine()

	return string(op)
}
