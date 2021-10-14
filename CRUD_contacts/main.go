package main

import "fmt"

var contacts []contact

type contact struct {
	id        uint16
	firstName string
	lastName  string
	phone     string
	email     string
	position  string
}

func selectMethod() {
	fmt.Println("Enter a method name you want to use:")
	var method string
	fmt.Scanf("%s", &method)
	switch method {
	case "create", "Create":
		create()
	case "update", "Update":
		update()
	case "get", "Get":
		get()
	case "getAll", "getall", "get All", "GetAll", "Get All":
		getAll()
	case "delete", "Delete":
		delete()

	}
}

func create() {

	var (
		id                                          uint16
		firstName, lastName, phone, email, position string
	)
	fmt.Println("Enter properties in following order: id, firstName, lastName, phone, email, position ")
	fmt.Scanf("%d %s %s %s %s %s", &id, &firstName, &lastName, &phone, &email, &position)

	fmt.Println(id, firstName, lastName, phone, email, position)
	var contactMem contact
	contactMem.id = id
	contactMem.firstName = firstName
	contactMem.lastName = lastName
	contactMem.phone = phone
	contactMem.email = email
	contactMem.position = position
	contacts = append(contacts, contactMem)
	getAll()
}
func update() {
	var id uint16

	fmt.Println("Enter an id of contact you want to update:")
	fmt.Scanf("%d", &id)
	fmt.Println("Enter new properties in following order: firstName, lastName, phone, email, position ")
	fmt.Scanf("%s %s %s %s %s", &contacts[id-1].firstName, &contacts[id-1].lastName, &contacts[id-1].phone, &contacts[id-1].email, &contacts[id-1].position)
	getAll()
}
func get() {
	var id uint16
	fmt.Println("Enter an id of contact you want to see:")
	fmt.Scanf("%d", &id)
	fmt.Println(contacts[id-1])
}
func getAll() {
	for _, value := range contacts {
		fmt.Println(value)
	}

}
func delete() {
	var id uint16
	fmt.Println("Enter an id of contact you want to delete:")
	fmt.Scanf("%d", &id)
	copy(contacts[id-1:], contacts[id:])
	contacts = contacts[:len(contacts)-1]
	getAll()
}
func defaultV() {

	var contact1, contact2, contact3 contact
	contact1.id = 1
	contact1.firstName = "Liam"
	contact1.lastName = "Neeson"
	contact1.phone = "65454354"
	contact1.email = "liamneeson@gmail.com"
	contact1.position = "Director"

	contact2.id = 2
	contact2.firstName = "Kurt"
	contact2.lastName = "Russel"
	contact2.phone = "6873465457"
	contact2.email = "kurtruss@gmail.com"
	contact2.position = "project manager"

	contact3.id = 3
	contact3.firstName = "Tim"
	contact3.lastName = "Allen"
	contact3.phone = "787546846"
	contact3.email = "allentim@gmail.com"
	contact3.position = "architect"

	contacts = append(contacts, contact1, contact2, contact3)
}

func main() {

	defaultV()

	selectMethod()
}
