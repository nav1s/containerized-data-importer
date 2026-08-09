package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
  envFile := flag.String("env-file", "/shared/env", "file to create for saving the image environment variables")
	flag.Parse()
  log.Printf("Extracting environment variables to %s", *envFile)
	var envString strings.Builder
	for _, env := range os.Environ() {
		key, value, _ := strings.Cut(env, "=")
		fmt.Fprintf(&envString, "%s=%s\n", key, value)
	}
	if err := os.WriteFile(*envFile, []byte(envString.String()), 0600); err != nil {
		log.Fatalf("Failed creating \"env\" file: %v", err)
	}

	log.Printf("Finished extracting environment variables to %s", *envFile)
}
