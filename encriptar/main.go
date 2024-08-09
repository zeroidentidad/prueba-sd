package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(encryptMsg("dcj", "I love prOgrAmming!"))
}

// encryptMsg retorna un texto encriptado
func encryptMsg(key, msg string) string {
	if msg == "" {
		return ""
	}

	if key == "" {
		key = "DCJ" // default uppercase key
	}

	var (
		charset = "aeiouAEIOU"
		result  string
	)

	for _, char := range msg {
		if strings.Contains(charset, string(char)) {
			result += key + string(char) // cast rune as string
		} else {
			result += string(char)
		}
	}

	return result
}
