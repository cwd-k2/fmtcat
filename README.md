# fmtcat

## Usage

Currently, available format specifier is `"%s"` alone.

```
$ fmtcat <format> [filenames...]
```

## What?

Almost same as:

```
# in zsh
$ printf <format> "$(< filename1)" "$(< filename2)" ...
```
