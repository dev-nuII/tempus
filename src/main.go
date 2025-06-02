package main

import (
	"bufio"
	"fmt"
	"github.com/dev-nuII/tempus/src/cmd"
	"github.com/dev-nuII/tempus/src/helper"
	"github.com/dev-nuII/tempus/src/utils"
	"os"
	"path/filepath"
	"time"
)

func main() {
	cmd.ParseFlags()

	if cmd.Generate {
		cmd.GenerateTokens()
		return
	}

	var tokens []string

	if cmd.Bulk != "" {
		file, err := os.Open(cmd.Bulk)
		if err != nil {
			fmt.Println("Error opening bulk file:", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			tokens = append(tokens, scanner.Text())
		}
	} else if args := cmd.FlagArgs(); len(args) > 0 {
		tokens = append(tokens, args[0])
	} else {
		fmt.Println("No token provided. Use --help for usage.")
		return
	}

	concurrency := 50

	results := helper.CheckToken(tokens, concurrency)

	if cmd.Output != "" {
		outputPath := cmd.Output
		dir := filepath.Dir(outputPath)

		if outputPath[len(outputPath)-1] == os.PathSeparator {
			dir = outputPath
			now := time.Now().Format("2006-01-02_15-04-05")
			outputPath = filepath.Join(dir, now+".json")
		}

		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			fmt.Println("Failed to create output directory:", err)
			return
		}

		fi, err := os.Stat(outputPath)
		if err == nil && fi.IsDir() {
			now := time.Now().Format("2006-01-02_15-04-05")
			outputPath = filepath.Join(outputPath, now+".json")
		}

		if err := utils.SaveJSON(results, outputPath); err != nil {
			fmt.Println("Error saving JSON:", err)
		} else {
			fmt.Println("Results saved to", outputPath)
		}
	} else {
		utils.PrintResults(results)
	}
}
