package main

import (
	"Order-Service/legacy/status"
	"github.com/jmoiron/sqlx"
	"log"
)

func main() {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=Order sslmode=disable password=14062002")
	if err != nil {
		log.Fatalln(err)
	}

	statusRepo := status.NewRepository(db)

	// Create table statuses
	err = statusRepo.CreateTable()
	if err != nil {
		log.Fatalln(err)
		return
	}

	// Insert random Status in table
	/*statusVar := testing_module.GenerateMockStatus()
	fmt.Printf("Status id : %d, name : %s", statusVar.Id, statusVar.StatusName)

	id, err := statusRepo.Insert(statusVar)
	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Printf("Id of inserted status : %d", id)*/

	// Update chosen Status in table
	/*err = statusRepo.Update(3, &status.Status{
		StatusName: "Really New Name",
	})
	if err != nil {
		log.Fatalln(err)
		return
	}*/

	// Delete chosen Status in table
	/*err = statusRepo.Delete(2)
	if err != nil {
		log.Fatalln(err)
		return
	}*/

	//Select all entries in table
	/*statuses, err := statusRepo.Select()
	if err != nil {
		log.Fatalln(err)
		return
	}

	for _, s := range statuses {
		fmt.Printf("Id : %d, Status Name : %s\n", s.Id, s.StatusName)
	}*/

	// Delete all entries in table
	/*err = statusRepo.Clean()
	if err != nil {
		log.Fatalln(err)
		return
	}*/

	// Reset counter Id
	/*err = statusRepo.ResetIdCounter(1)
	if err != nil {
		log.Fatalln(err)
		return
	}*/

	// Delete table
	/*err = statusRepo.DeleteTable()
	if err != nil {
		log.Fatalln(err)
		return
	}*/

}
