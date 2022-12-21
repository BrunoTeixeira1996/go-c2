package webserver

import (
	"fmt"
    "net/http"
    "encoding/json"
    "go-c2/server/utils"
)

// Function that handles the server stuff
func server() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost  {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            return
        }

        data := &utils.Data{}
        err := json.NewDecoder(r.Body).Decode(data)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        w.WriteHeader(http.StatusCreated)
        fmt.Printf("%s\n", data.Result)
    })

    if err := http.ListenAndServe(":8080", nil); err != http.ErrServerClosed {
        panic(err)
    }
}


// Function to start the server as a go routine
func RunAPI() {
    server()
}
