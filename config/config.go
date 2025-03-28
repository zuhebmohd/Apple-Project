package config

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var Properties map[string]string

func LoadProperties(filename string) {
	Properties = make(map[string]string)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening %s: %v", filename, err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			Properties[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading %s: %v", filename, err)
	}
}
