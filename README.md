# SiteChecker

A simple Go program to monitor changes on a specified web page and notify you in the console when changes occur.

## Features

* Continuously monitors a specified URL.
* Alerts the user when the HTML content of the page changes.
* Allows customizing the checking interval in seconds.

## Requirements

* Go 1.20+ (or latest version)

## Installation / Download

### Build from Source

1. Clone the repository:

```bash
git clone https://github.com/Alperosci/SiteChecker.git
cd SiteChecker
```

2. Run the program:

```bash
go run main.go
```

3. Or build it:

```bash
go build main.go
```

### Pre-built Release

You can also download pre-built binaries from the [Releases](https://github.com/username/web-change-detector/releases) section and run them directly without building from source.

## Usage

When you start the program, you will be prompted for two inputs:

1. **URL:** The web page you want to monitor.
2. **Wait time:** The checking interval in seconds.

Example:

```text
URL : https://example.com
Wait time (second) : 10
Waiting started
Waiting 10 seconds ...
Waiting 10 seconds ...
A change detected!
```

## How It Works

* The program first stores the initial HTML content of the page.
* It then repeatedly checks the page at the specified interval.
* If the content changes, it prints `"A change detected!"` to the console.

## Contributing

* Fork the repository.
* Commit your changes.
* Submit a pull request.
