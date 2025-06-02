package cmd

import "fmt"

const (
	Reset  = "\033[0m"
	Bold   = "\033[1m"
	Cyan   = "\033[36m"
	Yellow = "\033[33m"
	Gray   = "\033[90m"
)

func ShowHelp() {
	helpText := `
` + Bold + Cyan + `Tempus - Fast as Fuck Discord Token Checker & Generator` + Reset + `

` + Cyan + `Description:` + Reset + `
  Tempus is a powerful and efficient Discord token tool designed for both token validation and mass token generation.
  It features a high-performance token generator with a dual algorithm:
    - For small batch generations (<250k tokens), it uses a fast, single-threaded method optimized for minimal memory overhead.
    - For large batches (>=250k tokens), Tempus automatically switches to a highly concurrent, parallelized generation process utilizing all CPU cores for maximum throughput.
  The token checker performs concurrent HTTP requests with controlled concurrency to quickly verify token validity against Discord's API while handling rate limits gracefully.

` + Yellow + `Important Notes & Warnings:` + Reset + `
  - Valid token lengths are strictly 59 or 70 characters. The 70-length tokens start with the "mfa." prefix (Multi-Factor Authentication).
  - Generating very large amounts of tokens (especially >250k) may require significant system resources (CPU, RAM).
  - Excessive generation counts can lead to slowdowns or crashes if your system is underpowered.
  - Token checking respects concurrency limits but beware of Discord's rate limiting (HTTP 429). Tempus will report when tokens are rate limited.
  - Always use responsibly and within Discord's Terms of Service.

` + Cyan + `Usage:` + Reset + `
  tempus ` + Yellow + `[--bulk tokens.txt]` + Reset + ` ` + Yellow + `[--output results.json]` + Reset + ` ` + Yellow + `[TOKEN]` + Reset + `
  tempus ` + Yellow + `--generate` + Reset + ` [` + Yellow + `FILE` + Reset + `] [` + Yellow + `--count NUM` + Reset + `] [` + Yellow + `--tokenlength 59|70` + Reset + `]

` + Cyan + `Options:` + Reset + `
  ` + Yellow + `--help` + Reset + `               Show this help message
  ` + Yellow + `--bulk FILE` + Reset + `          Check tokens from a file (one token per line)
  ` + Yellow + `--output FILE` + Reset + `        Save checking results to JSON file. If no file specified, defaults to output/YYYY-MM-DD_HH-MM-SS.json
  ` + Yellow + `--generate [FILE]` + Reset + `    Generate tokens and save to FILE. If no file given, saves to tokens/YYYY-MM-DD_HH-MM-SS.txt
  ` + Yellow + `--count NUM` + Reset + `          Number of tokens to generate (default 10)
  ` + Yellow + `--tokenlength NUM` + Reset + `    Length of generated tokens: 59 or 70 (default 70)

` + Cyan + `Examples:` + Reset + `
  # Check tokens from a file and output results to JSON
  tempus --bulk tokens.txt --output results.json

  # Check a single token directly
  tempus your_token_here

  # Generate 100 tokens with default length (70) and save automatically
  tempus --generate --count 100

  # Generate 20 tokens of length 59 and save to a specific file
  tempus --generate tokens/generated.txt --count 20 --tokenlength 59

` + Cyan + `Algorithm & Performance Highlights:` + Reset + `
  - Token Generation:
      * Uses a high-performance random byte filling algorithm optimized with batch writes.
      * Automatically switches between a simple buffered write and a parallelized worker pool for very large token counts.
      * Employs concurrency proportional to CPU cores to maximize throughput on multi-core systems.
  - Token Checking:
      * Concurrent HTTP requests with adjustable concurrency to balance speed and rate limiting.
      * Gracefully handles HTTP errors, including rate limiting (HTTP 429), invalid tokens, and network issues.
      * Results include detailed statuses for easy filtering and post-processing.

` + Yellow + `Remember:` + Reset + `
  Use Tempus responsibly. Excessive token generation or validation requests may be flagged or blocked by Discord. Always respect their Terms of Service.

`
	fmt.Print(helpText)
}
