package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"unicode"
)

type AsciiData struct {
	Text   string
	Banner string
	Output string
}

func main() {
	http.HandleFunc("/", indexHandler)

	fmt.Printf("Starting server at localhost:8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var tmpl *template.Template
	var err error
	if r.URL.Path != "/" {
		notFoundHandler(w, r)
		tmpl, err = template.ParseFiles("templates/error.html")
	} else {
		tmpl, err = template.ParseFiles("templates/index.html")
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost {
		data := AsciiData{
			Text:   r.FormValue("text"),
			Banner: r.FormValue("banner"),
		}
		if !isAscii(data.Text) || data.Text=="" {
			http.Error(w, "400 BAD REQUEST", http.StatusBadRequest)
			return
		} else {
			output, err := generateAsciiArt(data.Text, data.Banner)
			if err != nil {
				http.Error(w, "500 INTERNAL SERVER ERROR", http.StatusInternalServerError)
				return
			}

			data.Output = output
			tmpl.Execute(w, data)
		}
	} else {
		tmpl.Execute(w, nil)
	}
}

func generateAsciiArt(text, banner string) (string, error) {
	filename := "ArtStyles/" + banner + ".txt"
	var outputBuffer strings.Builder

	strArr := strings.Split(text, "\n")
	for i := 0; i <= len(strArr)-1; i++ {
		if i-1 >= 0 {
			if strArr[i] == "" && i == len(strArr)-1 {
				continue
			}
		}
		if strArr[i] == "" {
			outputBuffer.WriteString("\n")
			continue
		}
		runes := []rune(strArr[i])
		for j := 0; j <= 8; j++ {
			for k := 0; k <= len(runes)-1; k++ {
				line := 2 + 9*(int(runes[k])-32) + j
				err := printLine(filename, line, &outputBuffer)
				if err != nil {
					return "", err
				}
			}
			if j < 8 {
				outputBuffer.WriteString("\n")
			}
		}
		if i < len(strArr)-1 {
			outputBuffer.WriteString("\n")
		}
	}

	return outputBuffer.String(), nil
}

func printLine(filename string, line int, output *strings.Builder) error {
	styleFile, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer styleFile.Close()

	styleFile.Seek(0, 0)

	scanner := bufio.NewScanner(styleFile)
	counter := 0
	for scanner.Scan() {
		counter++
		if counter == line {
			output.WriteString(scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func isAscii(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}
