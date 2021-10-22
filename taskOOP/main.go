package main

import (
	"errors"
	"fmt"
	"time"
)

var tasks []task

type task struct {
	id        int
	name      string
	status    string
	deadline  time.Time
	createdAt time.Time
	updatedAt time.Time
}

var task1 = task{1, "newProject", "initial", time.Date(2021, time.Month(10), 30, 12, 0, 0, 0, time.UTC), time.Date(2021, time.Month(10), 19, 10, 0, 0, 0, time.UTC), time.Date(2021, time.Month(10), 21, 9, 14, 19, 54, time.UTC)}
var task2 = task{1, "Project2", "dev", time.Date(2021, time.Month(10), 25, 18, 0, 0, 0, time.UTC), time.Date(2021, time.Month(10), 15, 9, 0, 0, 0, time.UTC), time.Date(2021, time.Month(10), 21, 19, 25, 34, 15, time.UTC)}

type Director struct {
}
type teamLead struct {
}
type developer struct {
}

type methods interface {
	giveTask()
	change()
	develop()
}

func (dir *Director) giveTask(newTask task) (id int, err error) {
	for _, v := range tasks {
		if v.id == newTask.id {
			id = 0
			err = errors.New("you already have task with this ID")
			return 0, err
		} else if newTask.status == "initial" {
			tasks = append(tasks, newTask)
			err = nil
			return id, err
		} else {
			err = errors.New("new task must be on 'initial' status")
			id = 0
			return 0, err
		}
	}
	return id, err
}
func (l *teamLead) change(id int, status string) (st string, err error) {
	c := 0
	for _, v := range tasks {
		if v.id == id {
			if v.status == "test" {
				if status == "done" {
					v.status = status
					st = status
					err = nil
					return st, err
				} else {
					st = ""
					err = errors.New("this task is on 'test' status now, you can change status to 'done'")
					return st, err
				}
			} else {
				st = ""
				err = errors.New("this task isn't on 'test' status yet")
			}
		} else {
			c++
		}
	}
	if c == len(tasks) {
		st = ""
		err = errors.New("there is no task with this ID")
	}
	return st, err
}

func (d *developer) develop(id int, status string) (st string, err error) {
	c := 0
	for _, v := range tasks {
		if v.id == id {
			if v.status == "initial" {
				if status == "dev" {
					v.status = status
					st = status
					err = nil
					return st, err
				} else {
					st = ""
					err = errors.New("this task is on 'initial' status now, you can develop it to 'dev' status")
					return st, err
				}
			} else if v.status == "dev" {
				if status == "test" {
					v.status = status
					st = status
					err = nil
					return st, err
				} else {
					st = ""
					err = errors.New("this task is on 'dev' status now,you can develop it to 'test' status")
					return st, err
				}
			} else {
				st = ""
				err = errors.New("this task is on 'test' status now,you can't develop it")
				return st, err
			}
		} else {
			c++
		}
	}
	if c == len(tasks) {
		st = ""
		err = errors.New("there is no task with this ID")
	}
	return st, err
}

func main() {
	tasks = append(tasks, task1, task2)
	fmt.Println(tasks[0].deadline)
	director := Director{}
	newTask := task{1, "newProject", "initial", time.Date(2021, time.Month(10), 30, 12, 0, 0, 0, time.UTC), time.Date(2021, time.Month(10), 19, 10, 0, 0, 0, time.UTC), time.Date(2021, time.Month(10), 21, 9, 14, 19, 54, time.UTC)}
	fmt.Println(director.giveTask(newTask))
	fmt.Println("hello")
	teamLead := teamLead{}
	teamLead.change(1, "done")

	developer1 := new(developer)
	developer1.develop(1, "dev")

}
