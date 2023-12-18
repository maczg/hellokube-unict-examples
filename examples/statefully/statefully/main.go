package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/user"
)

const filepath = "/tmp/benchy.log"

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		// user running the process
		usr, _ := user.Current()
		_, _ = fmt.Fprintf(w, fmt.Sprintf("hostname: %s , user: %s ", hostname, usr))

	})

	r.HandleFunc("/write", func(w http.ResponseWriter, r *http.Request) {
		n := r.URL.Query().Get("n")
		if n == "" {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = fmt.Fprintf(w, "missing parameter n\n")
			return
		}

		f, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = fmt.Fprintf(w, "error opening file: %s\n", err)
			return
		}
		defer f.Close()
		number, err := f.WriteString(n + "\n")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = fmt.Fprintf(w, "error writing to file: %s\n", err)
			return
		}
		_, _ = fmt.Fprintf(w, "wrote %d bytes\n", number)

	})

	r.HandleFunc("/read", func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open(filepath)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = fmt.Fprintf(w, "error opening file: %s\n", err)
			return
		}
		defer f.Close()

		fileScanner := bufio.NewScanner(f)
		fileScanner.Split(bufio.ScanLines)
		for fileScanner.Scan() {
			_, _ = fmt.Fprintf(w, "%s\n", fileScanner.Text())
		}

	})

	_ = http.ListenAndServe(":8080", r)

}
