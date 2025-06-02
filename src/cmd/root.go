package cmd

import (
	"flag"
	"os"
)

var (
	Help         bool
	Output       string
	Bulk         string
	Generate     bool
	GeneratePath string
	TokenLength  uint
	Count        int
)

func FlagArgs() []string {
	return flag.Args()
}

func ParseFlags() {
	flag.BoolVar(&Help, "help", false, "Show help message")
	flag.StringVar(&Output, "output", "", "Save results to JSON file")
	flag.StringVar(&Bulk, "bulk", "", "Path to tokens file")
	flag.BoolVar(&Generate, "generate", false, "Generate tokens")
	flag.StringVar(&GeneratePath, "generatepath", "", "Output file path for generated tokens (optional)")
	flag.UintVar(&TokenLength, "tokenlength", 70, "Length of generated tokens (59 or 70)")
	flag.IntVar(&Count, "count", 10, "Number of tokens to generate")
	flag.Parse()

	if Help {
		ShowHelp()
		os.Exit(0)
	}
}
