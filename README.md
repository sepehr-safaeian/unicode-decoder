# Unicode Decoder for .properties Files

This Go project is designed to traverse a directory, identify `.properties` files, and decode any Unicode escape sequences (formatted as `\uXXXX`) into their corresponding characters. It is useful for internationalization scenarios where `.properties` files might contain Unicode representations that need to be converted back to human-readable format.

## Features

- **Directory Traversal**: The program walks through a given directory and processes all files with the `.properties` extension.
- **Unicode Decoding**: Converts Unicode escape sequences (e.g., `\u00E9`) within the files to the actual characters.
- **Efficient File Processing**: Reads each file line by line to avoid memory issues with large files.
- **Error Handling**: Robust error handling to ensure smooth processing, even when encountering problematic files.

## How to Use

1. Clone the repository:
   ```bash
   git clone https://github.com/sepehr-safaeian/unicode-decoder.git
   cd unicode-decoder
   go install
   go run .\main.go

## Example

1. Assume you have a .properties file with the following content:
   ```bash
    greeting=\u0048\u0065\u006C\u006C\u006F, \u0077\u006F\u0072\u006C\u0064\u0021
After running the program, this file will be updated to:
   ```bash
    greeting=Hello, world!

