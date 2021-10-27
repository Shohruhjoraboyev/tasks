package service

import (
	"GOPATH/grpc/pkg/contacts"
	"context"
	"errors"
	"log"
)

type contact struct {
	id        uint32
	firstName string
	lastName  string
	phone     string
	email     string
	position  string
}

var cntcts = []contact{
	contact1,
	contact2,
}
var (
	contact1           = contact{1, "Liam", "Neeson", "65454354", "liamneeson@gmail.com", "director"}
	contact2           = contact{2, "Kurt", "Russel", "6873465457", "kurtruss@gmail.com", "programmer"}
	testContact        = contact{3, "Tim", "Allen", "787546846", "allentim@gmail.com", "team-lead"}
	testId      uint32 = 2
)

type GRPCServer struct{}

func (s *GRPCServer) Read(ctx context.Context, req *contacts.ReadReq) (con *contacts.ReadResp, err error) {
	c := &contacts.ReadResp{}
	var coun int
	testId = c.GetId()
	for i := range cntcts {
		if testId == cntcts[i].id {
			c.Id = cntcts[i].id
			c.Firstname = cntcts[i].firstName
			c.Lastname = cntcts[i].lastName
			c.Phone = cntcts[i].phone
			c.Email = cntcts[i].email
			c.Position = cntcts[i].position
			return c, nil

		} else {
			coun++
			if coun == len(cntcts) {
				err = errors.New("no contact with this id")
				return c, err
			}
		}
	}
	return &contacts.ReadResp{Id: c.GetId(), Firstname: c.GetFirstname(), Lastname: c.GetLastname(), Phone: c.GetPhone(), Email: c.GetEmail(), Position: c.GetPosition()}, nil
}

func (s *GRPCServer) Create(ctx context.Context, req *contacts.CreateReq) (*contacts.CreateResp, error) {
	c := &contacts.CreateResp{}
	for i := range cntcts {
		if testContact.id == cntcts[i].id {
			log.Fatal("already exsists contact with this id")
		} else {
			cntcts = append(cntcts, testContact)
			c.Id = testContact.id
		}
	}
	return &contacts.CreateResp{}, nil
}

func (s *GRPCServer) Update(ctx context.Context, req *contacts.UpdateReq) (*contacts.UpdateResp, error) {
	c := &contacts.UpdateResp{}
	var coun int
	for i := range cntcts {
		if testId == cntcts[i].id {
			log.Println(cntcts[i])
			cntcts[i].firstName = testContact.firstName
			cntcts[i].lastName = testContact.lastName
			cntcts[i].phone = testContact.phone
			cntcts[i].email = testContact.email
			cntcts[i].position = testContact.position

			c.Id = cntcts[i].id
			c.Firstname = cntcts[i].firstName
			c.Lastname = cntcts[i].lastName
			c.Phone = cntcts[i].phone
			c.Email = cntcts[i].email
			c.Position = cntcts[i].position
		} else {
			coun++
			if coun == len(cntcts) {
				log.Fatal("no contact with this id")
			}
		}
	}

	return &contacts.UpdateResp{}, nil
}

func (s *GRPCServer) Delete(ctx context.Context, req *contacts.DeleteReq) (*contacts.DeleteResp, error) {
	c := &contacts.DeleteResp{}
	var coun int
	for i := range cntcts {
		if testId == cntcts[i].id {
			copy(cntcts[testId-1:], cntcts[testId:])
			cntcts = cntcts[:len(cntcts)-1]
		} else {
			coun++
			if coun == len(cntcts) {
				log.Fatal("no contact with this id")
			}
		}
	}
	for i := range cntcts {
		if testId == cntcts[i].id {
			log.Fatal("not deleted")
		} else {
			coun++
			if coun == len(cntcts) {
				c.Result = "deleted succesfully"
			}
		}
	}
	return &contacts.DeleteResp{}, nil
}
