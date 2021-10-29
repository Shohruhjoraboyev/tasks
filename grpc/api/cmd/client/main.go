package main

import (
	"GOPATH/grpc/pkg/contacts"
	"context"
	"flag"
	"log"
	"strconv"

	"google.golang.org/grpc"
)

var conn, err = grpc.Dial(":9000", grpc.WithInsecure())

var c = contacts.NewContactServiceClient(conn)

func main() {
	checkErr(err)

	//read()

	//create()

	//update()

	delete()
}
func read() {
	flag.Parse()

	if flag.NArg() != 1 {
		log.Fatal("too many or less arguments, enter only one argument")
	}

	id, err := strconv.Atoi(flag.Arg(0))
	checkErr(err)

	res, err := c.Read(context.Background(), &contacts.ReadReq{Id: uint32(id)})
	checkErr(err)

	log.Println(res.GetId(), res.GetFirstname(), res.GetLastname(), res.GetPhone(), res.GetEmail(), res.GetPosition())
}

func create() {
	flag.Parse()

	if flag.NArg() != 6 {
		log.Fatal("too many or less arguments, enter 6 arguments, first is int, others are string")
	}
	id, err := strconv.Atoi(flag.Arg(0))
	checkErr(err)
	firstname := flag.Arg(1)
	lastname := flag.Arg(2)
	phone := flag.Arg(3)
	email := flag.Arg(4)
	position := flag.Arg(5)

	res, err := c.Create(context.Background(), &contacts.CreateReq{Id: uint32(id), Firstname: firstname, Lastname: lastname, Phone: phone, Email: email, Position: position})
	checkErr(err)
	log.Println(res.GetId())

}

func update() {
	flag.Parse()

	if flag.NArg() != 6 {
		log.Fatal("too many or less arguments, enter 6 arguments, first is int, others are string")
	}
	id, err := strconv.Atoi(flag.Arg(0))
	checkErr(err)
	firstname := flag.Arg(1)
	lastname := flag.Arg(2)
	phone := flag.Arg(3)
	email := flag.Arg(4)
	position := flag.Arg(5)
	res, err := c.Update(context.Background(), &contacts.UpdateReq{Id: uint32(id), Firstname: firstname, Lastname: lastname, Phone: phone, Email: email, Position: position})
	checkErr(err)
	log.Println(res.GetId(), res.GetFirstname(), res.GetLastname(), res.GetPhone(), res.GetEmail(), res.GetPosition())

}

func delete() {
	flag.Parse()

	if flag.NArg() != 1 {
		log.Fatal("too many or less arguments, enter only one argument")
	}

	id, err := strconv.Atoi(flag.Arg(0))
	checkErr(err)
	res, err := c.Delete(context.Background(), &contacts.DeleteReq{Id: uint32(id)})
	checkErr(err)
	log.Println(res.GetResult())
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
