package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"bytes"
)

func extractVars(){

  var envString strings.Builder
	envFile := "env-file"
	for _, env := range os.Environ() {
		key, value, _ := strings.Cut(env, "=")
		fmt.Fprintf(&envString,"%s = %s\n", key, value)
	}
	if err := os.WriteFile(envFile, []byte(envString.String()), 0644); err != nil {
		log.Fatalf("Failed creating \"env\" file: %v", err)
	}

}

func main() {
	  labels := map[string]string{"instancetype.kubevirt.io/default-instancetype":"u1.big"}
    data, err := os.ReadFile("env-file")
    if err != nil {
			return 
    }

      scanner := bufio.NewScanner(bytes.NewReader(data))
      for scanner.Scan() {
          line := scanner.Text()
          key, value, ok := strings.Cut(line, " = ")
          if ok {
              labels[strings.TrimSpace(key)] = strings.TrimSpace(value)
          }
	  }
	fmt.Println(labels)

}
