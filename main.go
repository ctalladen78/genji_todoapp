package main

import (
	"golang-projects/genji_todoapp/db"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

var (
	db        store.Store
	storePath string
)

const (
	ALL      = "all"
	NEW      = "new"
	DELETE   = "delete"
	UPDATE   = "update"
	FIND     = "find"
	RUNTESTS = "test"
)

func main() {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	taskStorePath := filepath.Join(dir, "/genji.db")
	log.Println("DB Used ",taskStorePath)

	taskDb, err := store.NewDB(taskStorePath)
	if err != nil {
		log.Fatal("DB NEW error ", err)
	}

	// init with new bucket, initial user data
	err = taskDb.Init(&store.Task{})
	err = userDb.Init(&store.User{})

	u1 := &store.User{Status: []byte("busy"), Name: []byte("sara")}
	u2 := &store.User{Status: []byte("busy"), Name: []byte("chris")}
	err = userDb.CreateUser(u1)
	err = userDb.CreateUser(u2)

	if err != nil {
		log.Fatal("DB init error ", err)
	}

	input := os.Args[1]
	if input == "" {
		log.Fatal("pls enter a command")
	}
	log.Print("user input: ", string(input))

	switch input {
	case ALL:
		// Get All
		getAllResult, err := taskDb.GetAll(taskStorePath)
		if err != nil {
			log.Fatal("GET all error", err)
		}
		// first item should be updated
		log.Println("GET ALL")
		for i, s := range getAllResult {
			log.Printf("%d %s", i, s)
		}
	case NEW:
		todoval := os.Args[2]
		if input == "" || todoval == "" {
			log.Fatal("pls enter a command")
		}
		t := &store.Task{Value: []byte(todoval)}
		// Create
		err = taskDb.CreateTask(t)
		if err != nil {
			log.Fatal("Create error ", err)
		}
	case UPDATE:
		newval := os.Args[3]
		idx, err := strconv.Atoi(os.Args[2])
		if input == "" || newval == "" || err != nil {
			log.Fatal("pls enter a command")
		}
		getAllResult, err := taskDb.GetAll(taskStorePath)
		if len(getAllResult) > 1 {
			// r := rand.New(rand.NewSource(99)) // Update random item
			// rnum := r.Intn(len(getAllResult))
			// err = taskDb.UpdateTask(getAllResult[rnum].ID, newval)
			err = taskDb.UpdateTask(getAllResult[idx].ID, newval)
		}
		if err != nil {
			log.Fatal("Create error ", err)
		}
	case FIND: // find by list index
		idx, err := strconv.Atoi(os.Args[2])
		getAllResult, err := taskDb.GetAll(taskStorePath) // post transaction routine TODO in db routine
		// to := &store.Task{}
		// to, err := taskDb.FindOne(getAllResult[idx].ID)
		if err != nil {
			log.Fatal("GET all error", err)
		}
		log.Printf("FIND ONE: %s", getAllResult[idx])
	case RUNTESTS:
		// store.GetTasksWhereCreatorStatus("open") // relational
		taskDb.GetDoneTasks()        // filter
		taskDb.UpdateTasksAs("done") // map multiple items

	}

	// for serializing into json, convert []byte to string

}
