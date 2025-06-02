package utils

import (
	"encoding/json"
	"fmt"
	"github.com/dev-nuII/tempus/src/helper"
	"os"
)

func SaveJSON(results []helper.TokenResult, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(results)
}

func PrintResults(results []helper.TokenResult) {
	for _, r := range results {
		fmt.Printf("%-70s : %s\n", r.Token, r.Status)
	}
}
