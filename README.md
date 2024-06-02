# Bulkee

[![MIT License](https://img.shields.io/badge/License-MIT-a10b31)](LICENSE)

**Bulkee** is a completely horrid program that I mocked up in 1 hour to let me quickly bulk remove and create files similar to the "manually select packages" option for pacman.

Please for the love of god do not actually use unless you are forced to, one day I will fix it to have less crap error messages and make it look pretty but for now, it is absolutely rancid to use.

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
Remove:
README.md
Create:
yep.txt
Executing in 3 seconds...
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
