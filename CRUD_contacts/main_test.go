package main

import (
	"testing"
)

func TestCreate(t *testing.T) {
	var (
		c  int
		id uint16 = 1
	)
	con := Create(id, "Liam", "Neeson", "65454354", "liamneeson@gmail.com", "Director")
	for _, v := range contacts {
		if v.id == id && con == v {
			c++

		}
		if c < 0 {
			t.Error("Creating error")
		}
	}
}
func TestUpdate(t *testing.T) {
	var id uint16 = 1
	con := Update(id, "Liam", "Neeson", "65454354", "liamneeson@gmail.com", "Director")
	if !(contacts[id-1].id == id && contacts[id-1] == con) {
		t.Error("Updating error")

	}
}

func TestGet(t *testing.T) {
	var id uint16 = 1
	con := Get(id)
	if contacts[id-1] != con {
		t.Error("Getting error")
	}
}

func TestGetAll(t *testing.T) {
	con := GetAll()
	for i, _ := range con {
		if con[i] != contacts[i] {
			t.Error("Getting error")
		}
	}
}

func TestDelete(t *testing.T) {
	var id uint16 = 1
	con := Delete(id)
	if len(contacts) == 0 {
	} else if con[0] != contacts[id] {
		t.Error("Deleting error")
	}
}
