package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"log"
	"os"
)

func DecodeJson(b *bytes.Buffer, i interface{}) {
	decoder := json.NewDecoder(b)

	if err := decoder.Decode(i); err != nil {
		log.Println("Error decoding buffer: ", err)
	}
}

func LoadJson(s string, i interface{}) {
	file, err := os.Open(s)

	if err != nil {
		log.Println("Error opening file: ", err)
		return
	}

	defer file.Close()
	buffer := new(bytes.Buffer)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		buffer.Write(scanner.Bytes())
	}

	DecodeJson(buffer, i)
}
