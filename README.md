# Tempus - Fast As Fuck Discord Token Checker & Generator

## Description

Tempus is a powerful and efficient Discord token tool designed for both token validation and mass token generation.

It features a high-performance token generator with a dual algorithm:

- For small batch generations (<250k tokens), it uses a fast, single-threaded method optimized for minimal memory overhead.
- For large batches (≥250k tokens), Tempus automatically switches to a highly concurrent, parallelized generation process utilizing all CPU cores for maximum throughput.

The token checker performs concurrent HTTP requests with controlled concurrency to quickly verify token validity against Discord's API while handling rate limits gracefully.

---

## Build & Installation

### Prerequisites

- [Go](https://golang.org/dl/) 1.18 or higher installed and configured.
- Git installed (optional, for cloning repo).

### Clone & Build from Source

```bash
git clone [https://github.com/dev-nuII/tempus.git](https://github.com/dev-nuII/tempus.git)
cd tempus
go build -o tempus main.go
```

This creates an executable called `tempus` (or `tempus.exe` on Windows) in the current directory.

### Add to PATH / Environment Variables

To use `tempus` globally from any terminal session, add its folder to your system PATH:

#### Windows

1. Move `tempus.exe` to a folder, e.g., `C:\Tools\tempus\`.
2. Open **System Properties** > **Advanced** > **Environment Variables**.
3. Under **System variables**, select **Path** > **Edit**.
4. Click **New** and add `C:\Tools\tempus\`.
5. Click OK and restart your terminal or system to apply.

#### Linux / macOS

1. Move the `tempus` executable to `/usr/local/bin/` or any directory in your PATH:
2. Verify by running:

If you prefer not to move the binary, you can temporarily add its folder to your PATH for the current session:

---

## Usage

```bash
tempus [--bulk tokens.txt] [--output results.json] [TOKEN]
tempus --generate [FILE] [--count NUM] [--tokenlength 59|70]
```
---

## Options

| Option | Description |
| --- | --- |
| `--help` | Show a help message |
| `--bulk FILE` | Check tokens from a file (one token per line) |
| `--output FILE` | Save checking results to JSON file. If no file specified, defaults to `output/YYYY-MM-DD_HH-MM-SS.json` |
| `--generate [FILE]` | Generate tokens and save to FILE. If no file given, saves to `tokens/YYYY-MM-DD_HH-MM-SS.txt` |
| `--count NUM` | Number of tokens to generate (default 10) |
| `--tokenlength NUM` | Length of generated tokens: 59 or 70 (default 70) |

---

## Algorithm & Performance Highlights

### Token Generation

- Uses a high-performance random byte filling algorithm optimized with batch writes.
- Automatically switches between a simple buffered write and a parallelized worker pool for very large token counts.
- Employs concurrency proportional to CPU cores to maximize throughput on multi-core systems.

### Token Checking

- Concurrent HTTP requests with adjustable concurrency to balance speed and rate limiting.
- Gracefully handles HTTP errors, including rate limiting (HTTP 429), invalid tokens, and network issues.
- Results include detailed statuses for easy filtering and post-processing.

---

## Remember

Use Tempus responsibly. Excessive token generation or validation requests may be flagged or blocked by Discord. Always respect their Terms of Service.
