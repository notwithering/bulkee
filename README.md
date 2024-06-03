> [!WARNING]
> Renaming is not yet implemented, attempting to do so will result in loss of data.

# Bulkee

[![MIT License](https://img.shields.io/badge/License-MIT-a10b31)](LICENSE)

**Bulkee** is a program that lets you quickly remove and create files easily similar to the "manually select packages" option for pacman.

```bash
$ ls
go.mod  go.sum  LICENSE  main.go  README.md
$ bulkee
(*open text editor*)
  1 LICENSE
  2
  3 go.mod
  4 go.sum
  5 main.go
  6 yep.txt
(*:wq*)
:: Entries will be removed:
 - README.md

:: Entries will be created:
 + yep.txt

:: Proceed with execution? [Y/n]
>>
$ ls
go.mod  go.sum  LICENSE  main.go  yep.txt
```

## Installation

### Installing

```bash
go install github.com/notwithering/bulkee@latest
```

### Testing

```bash
go run github.com/notwithering/bulkee@latest
```
