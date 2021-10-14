package main

import "fmt"

func selectMethod() {
	var method string
	t := task{}
	fmt.Println("Enter a method name:")
	fmt.Scanf("%s", &method)
	switch method {
	case "create", "Create":
		t.createTask()
	case "update", "Update":
		t.update()
	case "get", "Get":
		t.get()
	case "getAll", "getall", "get All", "GetAll", "Get All":
		t.getAll()
	case "delete", "Delete":
		t.delete()

	}
}

type create interface {
	createTask()
	update()
	get()
	getAll()
	delete()
}

func (t1 *task) createTask() {

	t1.id = uint16(len(tasks))
	fmt.Println("Enter infos following order: name status priority createdAt createdBy dueDate")
	fmt.Scanf("%s, %s, %d, %s, %s, %s", &t1.name, &t1.status, &t1.priority, &t1.createdAt, &t1.createdBy, &t1.dueDate)
	tasks = append(tasks, *t1)
	t1.getAll()
}

func (t1 *task) update() {
	fmt.Println("Enter an id of task you want to update:")
	fmt.Scanf("%d", &t1.id)
	t1 = &tasks[t1.id-1]
	fmt.Println("Enter infos following order: name status priority createdAt createdBy dueDate")
	fmt.Scanf("%s %s %d %s %s %s", &t1.name, &t1.status, &t1.priority, &t1.createdAt, &t1.createdBy, &t1.dueDate)
	t1.getAll()
}

func (t1 *task) get() {
	fmt.Println("Enter an id of task you want to update:")
	fmt.Scanf("%d", &t1.id)
	t1 = &tasks[t1.id-1]
	fmt.Println(t1)
}

func (t1 *task) getAll() {
	for _, value := range tasks {
		fmt.Println(value)
	}
}
func (t1 *task) delete() {
	fmt.Println("Enter an id of task you want to delete:")
	fmt.Scanf("%d", &t1.id)
	copy(tasks[t1.id-1:], tasks[t1.id:])
	tasks = tasks[:len(tasks)-1]
	t1.getAll()
}

var tasks = []task{
	task1,
	task2,
	task3,
}
var (
	task1 = task{1, "fibonacci numbers", "done", 1, "10-10-2021", "director", "15-10-2021"}
	task2 = task{2, "fizzBuzz", "started", 3, "9-10-2021", "programmer", "20-10-2021"}
	task3 = task{3, "palindrome", "notStarted", 2, "14-10-2021", "team-lead", "18-10-2021"}
)

type task struct {
	id        uint16
	name      string
	status    string
	priority  uint8
	createdAt string
	createdBy string
	dueDate   string
}

func main() {

	selectMethod()

}
