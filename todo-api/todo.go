package main

import (
    "fmt"
    "net/http"
    "strconv"
    "encoding/json"

    "github.com/gorilla/mux"
)

type Task struct {
    Id int
    Todo string
    Done bool
}

var tasks = [10]Task {
    {0 , "zero",  true},
    {1 , "one",   false},
    {2 , "two",   false},
    {3 , "three", true},
    {4 , "four",  false},
    {5 , "five",  false},
    {6 , "six",   true},
    {7 , "seven", false},
    {8 , "eight", false},
    {9 , "nine",  false},
}



func getTasks(w http.ResponseWriter, r *http.Request) { 
    for i := 0; i < 10; i++ {
        fmt.Fprintf(w, "%+v\n", tasks[i])
    }
}

func getTask(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    u64, err := strconv.ParseUint(vars["id"], 10, 32)

    if err != nil {
        fmt.Fprint(w, "Invalid task id\n")
        return 
    }

    id := uint(u64)

    if id < 0 || id > 9 {
        fmt.Fprint(w, "Id must be from 0 to 9\n")
        return 
    }
    
    fmt.Fprintf(w, "Task %ui: %+v\n", id, tasks[id])
}

func createTask(w http.ResponseWriter, r *http.Request) {
    // get message body
    // find empty slot
    // add 

}

// edits todo message
func editTask(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)  
    u64, err := strconv.ParseUint(vars["id"], 10, 32)

    if err != nil {
        fmt.Fprint(w, "Invalid id\n")
        return
    }

    id := uint(u64)

    if id < 0 || id > 9 {
        fmt.Fprint(w, "Id must be from 0 to 9\n")
        return 
    }

    type S struct {
        Todo string
    }

    var s S 

    error := json.NewDecoder(r.Body).Decode(&s)

    if error != nil {
        fmt.Fprint(w, "Incorrect JSON body\n")
        return
    }

    tasks[id].Todo = s.Todo

    fmt.Fprintf(w, "Task %ui updated: %+v\n", id, tasks[id])
}

func markTaskAsDone(w http.ResponseWriter, r *http.Request) {
}

func deleteTask(w http.ResponseWriter, r *http.Request) {

}

func main() {
    router := mux.NewRouter()

    router.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "You've requested %s\n", r.URL.Path)
    })

    router.HandleFunc("/tasks/all",         getTasks).Methods("GET")
    router.HandleFunc("/tasks/{id}",        getTask).Methods("GET")
    router.HandleFunc("/tasks/{id}",        createTask).Methods("POST")
    router.HandleFunc("/tasks/{id}",        editTask).Methods("PUT")
    router.HandleFunc("/tasks/{id}/done",   markTaskAsDone).Methods("PUT")
    router.HandleFunc("/tasks/{id}/undone", markTaskAsUndone).Methods("PUT")
    router.HandleFunc("/tasks/{id}",        deleteTask).Methods("DELETE")

    http.ListenAndServe(":8080", router)
}
