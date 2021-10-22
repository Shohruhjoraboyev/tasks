package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "myPass"
	dbname   = "myDb"
)

func main() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()

	// Add

	contact1 := `INSERT INTO "contacts"("firstname","lastname","phone","email","position") VALUES('Liam','Neeson','65454354','liam@gmail.com','Director')`
	contact2 := `INSERT INTO "contacts"("firstname","lastname","phone","email","position") VALUES('Kurt','Russel','6873465457','kurt@gmail.com','project manager')`
	contact3 := `INSERT INTO "contacts"("firstname","lastname","phone","email","position") VALUES('Tim','Allen','787546846','allen@gmail.com','architect')`
	_, e := db.Exec(contact1)
	CheckError(e)
	_, e = db.Exec(contact2)
	CheckError(e)
	_, e = db.Exec(contact3)
	CheckError(e)

	// Update

	updateCon1 := `UPDATE "contacts" SET "firstname"="liam","lastname"="neeson","phone"="123456","email"="nees@gmail.com","position"="manager" WHERE "id"=3`
	_, e = db.Exec(updateCon1)
	CheckError(e)

	// Delete

	delete := `DELETE FROM "contacts" WHERE "id"=3`
	_, e = db.Exec(delete)
	CheckError(e)

	// List

	rows, err := db.Query(`SELECT*FROM "contacts"`)
	CheckError(err)

	defer rows.Close()
	for rows.Next() {
		var id int
		var firstname, lastname, phone, email, position string

		err = rows.Scan(&id, &firstname, &lastname, &phone, &email, &position)
		CheckError(err)

		fmt.Println(id, firstname, lastname, phone, email, position)
	}

	CheckError(err)

	// GET
	rows, err = db.Query(`SELECT*FROM "contacts" WHERE "i"=2`)
	CheckError(err)

	defer rows.Close()
	for rows.Next() {
		var id int
		var firstname, lastname, phone, email, position string

		err = rows.Scan(&id, &firstname, &lastname, &phone, &email, &position)
		CheckError(err)

		fmt.Println(id, firstname, lastname, phone, email, position)
	}

	CheckError(err)

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
