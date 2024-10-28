package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Token struct {
	Token string
}

func jsonTokenDecoder() string {
	file, err := os.Open("token.json")

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	var token Token

	err = json.Unmarshal(bytes, &token)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	return token.Token
}
