> [!WARNING]
> Renaming is not yet implemented, attempting to do so will result in loss of data.

# Bulkee

[![MIT License](https://img.shields.io/badge/License-MIT-a10b31)](LICENSE)

**Bulkee** is a program that lets you quickly remove and create files easily similar to the "manually select packages" option for pacman.

```ruby
$ bulkee
```

_(bulkee opens default text editor)_

1. _go.mod_
2. _go.sum_
3. ~~junkdir/~~
4. ~~junkfile~~
5. _locale/_
6. _main.go_
7. **README.md**
8. **LICENSE**

```ruby
:: Entries will be removed:
 - junkdir/
 - junkfile

:: Entries will be created:
 + README.md
 + LICENSE

:: Proceed with execution? [Y/n]
>> y
$ ls
go.mod  go.sum  LICENSE  locale  main.go  README.md
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
