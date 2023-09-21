# Colorise

Color patterns in your command output.

Uses escape sequence colors. So the color `32` is used as `\e[32m`, the color
`38;5;219` is used as `\e[38;5;219m`, etc...

## Installation

```
make
```

Then place the program `colorise` in your $PATH.

## Usage

```
echo "Hello my World" | colorise Hello 33 World 31
```
