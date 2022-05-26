
# golored

[![Release](https://img.shields.io/github/v/release/alexcoder04/golored)](https://github.com/alexcoder04/golored/releases/latest)
[![Top language](https://img.shields.io/github/languages/top/alexcoder04/golored)](https://github.com/alexcoder04/golored/search?l=go)
[![License](https://img.shields.io/github/license/alexcoder04/golored)](https://github.com/alexcoder04/golored/blob/main/LICENSE)
[![Issues](https://img.shields.io/github/issues/alexcoder04/golored)](https://github.com/alexcoder04/golored/issues)
[![Pull requests](https://img.shields.io/github/issues-pr/alexcoder04/golored)](https://github.com/alexcoder04/golored/pulls)

Simply color your shell scripts' output by piping it into this program.

## Installation

### Pre-built binaries

Simply download the binary from the [releases page](https://github.com/alexcoder04/golored/releases/latest).

### Build from source

```sh
git clone https://github.com/alexcoder04/golored.git
cd golored
go build .
go install # installs the binary to your $GOPATH
```

## Usage

| argument         | function                  |
|------------------|---------------------------|
| `-F FORMATTING`  | specify formatting option |
| `-b COLOR`       | specify background color  |
| `-f COLOR`       | specify foreground color  |
| `-h`             | print help                |
| `-i`             | print color codes         |
| `-s SOURCE_FILE` | read from file            |

### Examples:

```sh
golored -f red -F underline -s file.txt
```

Outputs the content of `file.txt` underlined and with red foreground.

```
ls | golored -b blue
```

Colors the output of `ls` with blue background.

### List of colors:

`black`, `blue`, `red`, `magenta`, `green`, `cyan`, `yellow`, `white`

### List of formatting options:

`bold`, `dim`, `italic`, `underline`

