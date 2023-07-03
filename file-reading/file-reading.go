package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"
)

var (
	counter = 0
)

type Mail struct {
	MessageID               string `json:"Message-ID"`
	Date                    string `json:"Date"`
	From                    string `json:"From"`
	To                      string `json:"To"`
	Subject                 string `json:"Subject,omitempty"`
	Cc                      string `json:"Cc,omitempty"`
	MimeVersion             string `json:"Mime-Version,omitempty"`
	ContentType             string `json:"Content-Type,omitempty"`
	ContentTransferEncoding string `json:"Content-Transfer-Encoding,omitempty"`
	Bcc                     string `json:"Bcc,omitempty"`
	XFrom                   string `json:"X-From,omitempty"`
	XTo                     string `json:"X-To,omitempty"`
	XCc                     string `json:"X-Cc,omitempty"`
	XBcc                    string `json:"X-Bcc,omitempty"`
	XFolder                 string `json:"X-Folder,omitempty"`
	XOrigin                 string `json:"X-Origin,omitempty"`
	XFileName               string `json:"X-FileName,omitempty"`
	Message                 string `json:"Message,omitempty"`
}

func visit(path string, di fs.DirEntry, err error) (error, Mail) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	mail := Mail{}
	currentKey := "Message-ID"

	scanner := bufio.NewScanner(file)
	buf := []byte{}
	scanner.Buffer(buf, 1024*1024)
	for scanner.Scan() {
		currentLine := scanner.Text()
		currentKey = parseLine(currentLine, &mail, currentKey)
	}
	if err := scanner.Err(); err != nil {
		return err, mail
	}
	counter++
	return nil, mail
}

func parseLine(line string, mail *Mail, currentKey string) string {
	line = strings.TrimSpace(line)
	if strings.Contains(line, ":") && currentKey != "Message" {
		keyValues := strings.Split(line, ":")
		if len(keyValues) > 2 && keyValues[0] != "Date" {
			keyValues[1] = strings.Join(keyValues[1:], ": ")
		}
		if len(keyValues) == 1 {
			keyValues = append(keyValues, "")
		}
		keyValues[1] = strings.TrimSpace(keyValues[1])
		switch keyValues[0] {
		case "Message-ID":
			mail.MessageID = keyValues[1]
			currentKey = "Message-ID"
		case "Date":
			keyValues[1] = strings.Join(keyValues[1:], ":")
			//remove (PDT) from date
			keyValues[1] = strings.Split(keyValues[1], " (")[0]
			//convert string to time
			t, err := time.Parse("Mon, 2 Jan 2006 15:04:05 -0700", keyValues[1])
			if err != nil {
				fmt.Println(err)
			}
			mail.Date = t.Format("Mon, 02 Jan 2006 15:04:05 -0700")
			currentKey = "Date"
		case "From":
			mail.From = keyValues[1]
			currentKey = "From"
		case "To":
			mail.To = keyValues[1]
			currentKey = "To"
		case "Subject":
			mail.Subject = keyValues[1]
			currentKey = "Subject"
		case "Cc":
			mail.Cc = keyValues[1]
			currentKey = "Cc"
		case "Mime-Version":
			mail.MimeVersion = keyValues[1]
			currentKey = "Mime-Version"
		case "Content-Type":
			mail.ContentType = keyValues[1]
			currentKey = "Content-Type"
		case "Content-Transfer-Encoding":
			mail.ContentTransferEncoding = keyValues[1]
			currentKey = "Content-Transfer-Encoding"
		case "Bcc":
			mail.Bcc = keyValues[1]
			currentKey = "Bcc"
		case "X-From":
			mail.XFrom = keyValues[1]
			currentKey = "X-From"
		case "X-To":
			mail.XTo = keyValues[1]
			currentKey = "X-To"
		case "X-Cc":
			mail.XCc = keyValues[1]
			currentKey = "X-Cc"
		case "X-Bcc":
			mail.XBcc = keyValues[1]
			currentKey = "X-Bcc"
		case "X-Folder":
			mail.XFolder = keyValues[1]
			currentKey = "X-Folder"
		case "X-Origin":
			mail.XOrigin = keyValues[1]
			currentKey = "X-Origin"
		case "X-FileName":
			mail.XFileName = keyValues[1]
			currentKey = "X-FileName"
		default:
			value := reflect.ValueOf(mail).Elem().FieldByName(currentKey)
			if value.IsValid() {
				value.SetString(value.String() + " " + line)
			}
		}
	} else {
		if line == "" {
			currentKey = "Message"
		}
		value := reflect.ValueOf(mail).Elem().FieldByName(currentKey)
		if value.IsValid() {
			value.SetString(value.String() + " " + line)
		}
	}

	return currentKey
}

func getKeyValues(path string, di fs.DirEntry, err error, keyValues map[string]int, fileNewValues map[string]string) error {
	if di.IsDir() {
		return nil
	}
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	buf := []byte{}
	scanner.Buffer(buf, 1024*1024)
	for scanner.Scan() {
		cl := scanner.Text()
		if cl == "" || cl[0] == ' ' || cl[0] == '\t' || cl[0] == '\n' {
			return nil
		}
		if strings.Contains(cl, ":") {
			values := strings.Split(cl, ": ")
			if len(values) > 2 {
				values[1] = strings.Join(values[1:], ": ")
			}
			if value, ok := keyValues[values[0]]; ok {
				keyValues[values[0]] = value + 1
			} else {
				keyValues[values[0]] = 1
				fileNewValues[path] = values[0]
			}
		}
	}
	return nil
}

func writeToFile(root string, current string) {
	newFilePath := "parsed_files" + "/" + current + ".ndjson"
	// if err := os.Mkdir(newFilePath, os.ModePerm); err != nil {
	//     log.Fatal(err)
	// }
	// newFilePath = newFilePath
	newFile, err := os.Create(newFilePath)

	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	fmt.Println(root)

	err = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if d.IsDir() {
			return nil
		}

		info, _ := d.Info()
		if info.Size() > 1000000 {
			fmt.Printf("Skipping %s\n file bigger than 1 mb", path)
			return nil
		}

		err, mail := visit(path, d, err)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		// turn struct into json
		jsonMail, err := json.Marshal(mail)
		if err != nil {
			log.Fatal(err)
		}

		newFile.WriteString(`{ "index" : { "_index" : "enron_mail" } }` + "\n")
		newFile.Write(jsonMail)
		newFile.WriteString("\n")
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	root := "../enron_mail_20110402"
	// keyValues := make(map[string]int)
	// fileNewValues := make(map[string]string)

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if d.IsDir() {
			// get directory parent
			dir, current := filepath.Split(path)
			_, father := filepath.Split(dir[0 : len(dir)-1])
			if father == "maildir" {
				writeToFile(path, current)
			}

			return nil
		}

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(keyValues, len(keyValues))
	fmt.Println(counter)
	fmt.Printf("filepath.WalkDir() returned %v\n", err)
}
