# `bintxt` – A CLI for converting binary text

Convert utf8-encoded text to binary text, and vice versa.

⚠️ **Disclaimer:** I'm just tinkering. Use at your own risk!

## Features

- [x] Encode text to binary text
- [x] Decode binary text to human-readable text

## Usage

### Commands

- `encode`: Convert text to binary text
- `decode`: Convert binary text to readable text

### Examples

#### Encode text

```sh
bintxt encode hello
```

Output:

```text
01101000 01100101 01101100 01101100 01101111
```

#### Decode binary text

```sh
bintxt decode 01101000 01100101 01101100 01101100 01101111
```

Output:

```text
hello
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file
for details.
