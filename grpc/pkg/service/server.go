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
	contact3,
}
var (
	contact1 = contact{1, "Liam", "Neeson", "65454354", "liamneeson@gmail.com", "director"}
	contact2 = contact{2, "Kurt", "Russel", "6873465457", "kurtruss@gmail.com", "programmer"}
	contact3 = contact{3, "Tim", "Allen", "787546846", "allentim@gmail.com", "team-lead"}
)

type GRPCServer struct{}

func (s *GRPCServer) Read(ctx context.Context, req *contacts.ReadReq) (con *contacts.ReadResp, err error) {
	testId := req.GetId()
	log.Println("hello")
	for i := range cntcts {
		if testId == cntcts[i].id {
			con = &contacts.ReadResp{}
			con.Id = cntcts[i].id
			con.Firstname = cntcts[i].firstName
			con.Lastname = cntcts[i].lastName
			con.Phone = cntcts[i].phone
			con.Email = cntcts[i].email
			con.Position = cntcts[i].position
			return con, nil
		}
	}
	if con == nil {
		err = errors.New("no contact with this id")
		return con, err
	}
	return con, err
}

// Create
func (s *GRPCServer) Create(ctx context.Context, req *contacts.CreateReq) (con *contacts.CreateResp, err error) {
	c := contact{}
	c.id = req.GetId()

	for i := range cntcts {
		if c.id == cntcts[i].id {
			err = errors.New("already exsists contact with this id")
		} else {
			c.firstName = req.GetFirstname()
			c.lastName = req.GetLastname()
			c.phone = req.GetPhone()
			c.email = req.GetEmail()
			c.position = req.GetPosition()
			cntcts = append(cntcts, c)
		}
	}
	return &contacts.CreateResp{Id: c.id}, err
}

func (s *GRPCServer) Update(ctx context.Context, req *contacts.UpdateReq) (con *contacts.UpdateResp, err error) {
	con = &contacts.UpdateResp{}
	con.Id = req.GetId()
	var coun int
	for i := range cntcts {
		if con.Id == cntcts[i].id {
			log.Println(cntcts[i])
			cntcts[i].firstName = req.GetFirstname()
			cntcts[i].lastName = req.GetLastname()
			cntcts[i].phone = req.GetPhone()
			cntcts[i].email = req.GetEmail()
			cntcts[i].position = req.GetPosition()

			con.Id = cntcts[i].id
			con.Firstname = cntcts[i].firstName
			con.Lastname = cntcts[i].lastName
			con.Phone = cntcts[i].phone
			con.Email = cntcts[i].email
			con.Position = cntcts[i].position
		} else {
			coun++
			if coun == len(cntcts) {
				err = errors.New("no contact with this id")
			}
		}
	}

	return &contacts.UpdateResp{}, err
}

func (s *GRPCServer) Delete(ctx context.Context, req *contacts.DeleteReq) (con *contacts.DeleteResp, err error) {
	con = &contacts.DeleteResp{}
	testId := req.GetId()
	var coun int
	for i := range cntcts {
		if testId != cntcts[len(cntcts)-1].id {
			if testId == cntcts[i].id {
				copy(cntcts[i:], cntcts[i+1:])
				cntcts = cntcts[:len(cntcts)-1]
				break
			} else {
				coun++
				if coun == len(cntcts) {
					log.Fatal("no contact with this id")
				}
			}
		} else {
			if testId == cntcts[i].id {
				cntcts = cntcts[:len(cntcts)-1]
				break
			} else {
				coun++
				if coun == len(cntcts) {
					err = errors.New("no contact with this id")
				}
			}

		}

	}
	for j := range cntcts {
		if testId == cntcts[j].id {
			err = errors.New("not deleted")
		} else {
			coun++
			if coun == len(cntcts) {
				con.Result = "deleted succesfully"
			}
		}
	}
	return con, err
}
