package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime/pprof"
)

func insert(path string) {
	fmt.Println(path)
	auth := "admin:Complexpass#123"

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	reader = bufio.NewReaderSize(reader, 2048*1024)

	req, err := http.NewRequest("POST", "http://localhost:4080/api/_bulk", reader)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(auth)))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}

func main() {
	// Create a file for CPU profiling.
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal(err)
	}

	// Start CPU profiling.
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	// file-reading folder needs to be created manually
	root := "../file-reading/parsed_files"

	// walk through the folder and insert each file
	err = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		// skip directories
		if d.IsDir() {
			return nil
		}
		insert(path)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("filepath.WalkDir() returned %v\n", err)

}
