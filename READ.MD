# Base Converter (Go)

Simple interactive CLI tool written in Go that converts numbers from
hexadecimal, binary, or octal notation into decimal representation.

Supports arbitrarily large numbers using `math/big`.

---

## Features

- Convert numbers with prefixes:
    - `0x` — hexadecimal
    - `0b` — binary
    - `0o` — octal
- Default base (without prefix): **hexadecimal**
- Unlimited integer size (`math/big.Int`)
- Interactive input loop
- Clean exit with `stop`

---

## Usage

### Build and run

```bash
go run main.go
