# Color Utility

## Goal

Escape codes have always baffled me. I wanted to write a small tool to add colored output to messages or the contents of a file so they would be a bit clearer.

## Description

This utility formats text output with color using ANSI escape codes. It supports reading text from standard input, a file, or a command-line argument. 

## Usage

```bash
color -color=<color> [-verbose=true] ( -file="<filepath>" | "<message>" )
```

### Parameters

| **Flag**         | **Description**                                             | **Default**     | **Required** |
|-------------------|-------------------------------------------------------------|-----------------|--------------|
| `-color=<color>` | Specify the color of the text output (see below for options) | `white`         | Yes          |
| `-file=<path>`   | Path to a file containing the message to colorize            | Not specified   | Either `-file` or `<message>` is required |
| `<message>`      | Message to colorize (if not using `-file`)                   | Not specified   | Either `<message>` or `-file` is required |
| `-verbose=true`  | Enable verbose logging to show processing details            | `false`         | No           |
| `-help`          | Display help information                                     | `false`         | No           |


### Valid Colors

- `black`
- `red`
- `green`
- `yellow`
- `blue`
- `purple`
- `cyan`
- `white`


### Examples

#### Colorize a Message
```bash
color -color=blue "Hello, World!"
```

#### Colorize a File's Contents
```bash
color -color=red -file="example.txt"
```

#### Display Verbose Output
```bash
color -color=green -verbose=true "This is a verbose message"
```

#### Display Help
```bash
color -help
```


## Error Handling

- If no message or file is provided, the program will exit and display usage instructions.
- If both a message and a file are provided, the program will display a usage error.
- If an invalid color is specified, the program defaults to `white`.

