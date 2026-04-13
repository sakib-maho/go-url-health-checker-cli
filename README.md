# URL Health Checker (Go)

<!-- BrandCloud:readme-standard -->
[![Maintained](https://img.shields.io/badge/Maintained-yes-brightgreen.svg)](#)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Showcase](https://img.shields.io/badge/Portfolio-Showcase-blue.svg)](#)

_Part of the `sakib-maho` project showcase series with consistent documentation and quality standards._

Modern Go CLI project for checking the health of multiple URLs concurrently and exporting results in JSON or CSV.

## Features

- Concurrent HTTP checks with worker-pool style goroutines
- Configurable timeout and retry count
- Input from CLI args or text file
- Output formats:
  - console table
  - JSON file
  - CSV file
- Unit tests for parsing and aggregation logic

## Quick Start

```bash
git clone https://github.com/sakib-maho/go-url-health-checker-cli.git
cd go-url-health-checker-cli
go run ./cmd/healthcheck --urls https://github.com,https://example.com
```

## Tests

```bash
go test ./...
```

## License

MIT - see [LICENSE](LICENSE).

## Project Structure

```text
.
├── cmd/healthcheck/main.go
├── internal/
│   ├── checker/checker.go
│   ├── model/model.go
│   ├── output/output.go
│   └── parse/parse.go
├── tests/sample_urls.txt
└── go.mod
```

## Usage

### Check URLs from comma-separated list

```bash
go run ./cmd/healthcheck \
  --urls https://github.com,https://example.com \
  --timeout 5 \
  --retries 2
```

### Check URLs from file

```bash
go run ./cmd/healthcheck \
  --file tests/sample_urls.txt \
  --json-out output/results.json \
  --csv-out output/results.csv
```

## Flags

- `--urls`: comma-separated URL list
- `--file`: path to newline-separated URL file
- `--timeout`: request timeout in seconds (default `5`)
- `--retries`: retry count for failed requests (default `1`)
- `--json-out`: optional JSON output path
- `--csv-out`: optional CSV output path

## Example Output

```text
URL                          STATUS   LATENCY_MS   ERROR
https://github.com           200      142          -
https://example.com          200      84           -
https://invalid.local        0        0            dial tcp: lookup failed
```
