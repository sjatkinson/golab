// RView - remote log viewer. Serves up log files to remote clients via a http interface.

// TODO: use a optional config file to allow multiple sandboxes,
//       set the log file matchers, maybe allow a optional filter
//       when displaying the log, add a flag to allow setting the port

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// TODO: make this a command line param
var sandbox string = os.Getenv("DEVROOT")

func expand(base, rel string) string {
	return filepath.Join(base, rel)
}

func allowedDir(sandbox, path string) bool {
	path, err := filepath.Abs(path)
        if err != nil {
		return false
	}
	if sandbox, err = filepath.Abs(sandbox); err != nil {
		return false
	}
	return strings.HasPrefix(path, sandbox)
}

func allowedFile(sandbox, path string) bool {
	return allowedDir(sandbox, filepath.Dir(path)) &&
		filepath.Ext(path) == ".log"
}

func inDir(path, dir string) bool {
	return strings.HasPrefix(path, dir)
}

func handleDir(sandbox, path string, w http.ResponseWriter) {
	if allowedDir(sandbox, path) {
		fmt.Fprintf(w, "<H1>%s</H1>", path)
		displayLogFiles(path, w)
	} else {
		http.Error(w, "Access denied", http.StatusUnauthorized)
	}
}

func handleFile(sandbox, path string, w http.ResponseWriter) {
	if !allowedFile(sandbox, path) {
		http.Error(w, "Not allowed", http.StatusForbidden)
	} else {
		displayFile(path, w)
	}

}

func displayLogFiles(path string, w io.Writer) {
	filepath.Walk(path, func(file string, info os.FileInfo, err error) error {
		if filepath.Ext(file) == ".log" {
			relative, _ := filepath.Rel(path, file)
			fmt.Fprintf(w, "<a href=\"%s\">%s</a><br>", file, relative)
		}
		return nil
	})
}

func displayFile(f string, w io.Writer) {
	body, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Fprintf(w, "Error reading file %v", err)
	} else {
		fmt.Fprintf(w, "%s", body)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == "/" || !(inDir(path, sandbox) || inDir(sandbox, path)) {
		path = expand(sandbox, path)
	}
	info, err := os.Stat(path)
	if err != nil {
		fmt.Println("Not found!! ", path)
		http.NotFound(w, r)
	} else {
		if info.IsDir() {
			handleDir(sandbox, path, w)
		} else {
			handleFile(sandbox, path, w)
		}
	}
}

func main() {
	fmt.Println("Starting...")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
