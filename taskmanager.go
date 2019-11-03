package main

import (
	"fmt"
	"log"
	"mongopkg"
	"os"
	"strconv"
	"strings"
)

func main() {

	command, input := validate(os.Args)

	switch command {
	case "do":
		i, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		newTask := new(mongopkg.Task)
		do(newTask, i)
		list(newTask)
	case "list":
		list(new(mongopkg.Task))
	case "add":
		newTask := new(mongopkg.Task)
		newTask.Name = input
		add(newTask)
		list(newTask)
	default:
		log.Fatal("Wrong command")
	}
}

func do(task mongopkg.MongoOps, index int) {
	tasks := task.Get()
	for i, v := range tasks {
		if index == (i + 1) {
			v.Update()
		}
	}
}

func add(task mongopkg.MongoOps) {
	task.Insert()
}

func list(task mongopkg.MongoOps) {
	tasks := task.Get()
	fmt.Println("You have the following tasks:")
	for i, v := range tasks {
		tempTask := v.(mongopkg.Task) //Type Assertion : Converting Interface to Type
		fmt.Printf("%v. %v\n", i+1, tempTask.Name)
	}
}

func validate(args []string) (string, string) {
	var command string
	var input string

	command = args[1]
	for i := 2; i < len(args); i++ {
		input += args[i] + " "
	}

	return command, input
}
