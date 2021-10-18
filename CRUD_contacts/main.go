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

var (
	id                                          uint16
	firstName, lastName, phone, email, position string
)

func selectMethod() {
	fmt.Println("Enter a method name you want to use:")
	var method string
	fmt.Scanf("%s", &method)
	switch method {
	case "create", "Create":
		{

			fmt.Println("Enter properties in following order: id, firstName, lastName, phone, email, position ")
			fmt.Scanf("%d %s %s %s %s %s", &id, &firstName, &lastName, &phone, &email, &position)

			Create(id, firstName, lastName, phone, email, position)
		}
	case "update", "Update":
		{
			fmt.Println("Enter an id of contact you want to update:")
			fmt.Scanf("%d", &id)
			fmt.Println("Enter new properties in following order: firstName, lastName, phone, email, position ")
			fmt.Scanf("%s %s %s %s %s", firstName, lastName, phone, email, position)

			Update(id, firstName, lastName, phone, email, position)
		}
	case "get", "Get":
		{
			fmt.Println("Enter an id of contact you want to see:")
			fmt.Scanf("%d", &id)
			fmt.Println(Get(id))
		}
	case "getAll", "getall", "get All", "GetAll", "Get All":
		{

			c := GetAll()
			for _, value := range c {
				fmt.Println(value)
			}
		}
	case "delete", "Delete":
		{
			fmt.Println("Enter an id of contact you want to delete:")
			fmt.Scanf("%d", &id)
			Delete(id)
		}

	}
}

func Create(id uint16, firstName string, lastName string, phone string, email string, position string) contact {

	var ContactMem contact
	ContactMem.id = id
	ContactMem.firstName = firstName
	ContactMem.lastName = lastName
	ContactMem.phone = phone
	ContactMem.email = email
	ContactMem.position = position
	contacts = append(contacts, ContactMem)
	GetAll()
	return ContactMem
}
func Update(id uint16, firstName string, lastName string, phone string, email string, position string) contact {
	contacts[id-1].id = id
	contacts[id-1].firstName = firstName
	contacts[id-1].lastName = lastName
	contacts[id-1].phone = phone
	contacts[id-1].email = email
	contacts[id-1].position = position
	GetAll()
	var updated = contact{id, firstName, lastName, phone, email, position}
	return updated
}
func Get(id uint16) contact {

	return contacts[id-1]
}
func GetAll() []contact {

	return contacts
}
func Delete(id uint16) []contact {
	con := contacts[id-1 : id]
	copy(contacts[id-1:], contacts[id:])
	contacts = contacts[:len(contacts)-1]
	GetAll()
	return con
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
