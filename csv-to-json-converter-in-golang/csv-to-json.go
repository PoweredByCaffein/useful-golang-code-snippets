package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Path to your sample file
	path := "sample.csv"

	// Open the file
	f, err := os.Open(path)
	if err != nil {
		log.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	// Create a new scanner
	scanner := bufio.NewScanner(f)

	getHeaders := true // flag to read headers
	var csvHeaders []string
	var jsonData []interface{}
	var linesRead int
	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) == "" {
			continue
		}

		if getHeaders {
			csvHeaders = fetchHeaders(line)
			getHeaders = false
			continue
		}

		// It is not the first line as we don't want to count the empty lines and the header
		linesRead++

		// This returns the JSON for the given line as a map[string]string
		jsonLine := convertLineToJSON(csvHeaders, line)
		jsonData = append(jsonData, jsonLine)
	}

	output, err := json.Marshal(jsonData)
	if err != nil {
		log.Println("Failed to Marshal JSON:", err)
	}

	fmt.Println(strings.Repeat("=", 100))
	fmt.Println("JSON String")
	fmt.Println(strings.Repeat("=", 100))
	fmt.Println(string(output))
	fmt.Println(strings.Repeat("=", 100))
	fmt.Println("Total Lines Read (excluding header):\t", linesRead)
	fmt.Println("Total JSON values created:\t\t", len(jsonData))
}

func fetchHeaders(line string) []string {
	headers := strings.Split(line, ",")
	for k, header := range headers {
		headers[k] = strings.TrimSpace(header)
	}
	return headers
}

func convertLineToJSON(headers []string, line string) map[string]string {
	jsonInterface := make(map[string]string)

	values := strings.Split(line, ",")
	if len(values) != len(headers) {
		log.Println("Invalid line: the number of headers don't match number of values in the line.")
		return nil
	}

	for k, value := range values {
		values[k] = strings.TrimSpace(value)
	}

	for k, header := range headers {
		jsonInterface[header] = values[k]
	}

	return jsonInterface

	// Use this to print the JSON data
	// jsonData, err := json.Marshal(jsonInterface)
	// if err != nil {
	// 	log.Println("Failed to Marshal JSON:", err)
	// }

	// return string(jsonData)

}
