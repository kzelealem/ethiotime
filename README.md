# ethiotime

[![Go Reference](https://pkg.go.dev/badge/github.com/kzelealem/ethiotime.svg)](https://pkg.go.dev/github.com/kzelealem/ethiotime)
[![Go Report Card](https://goreportcard.com/badge/github.com/kzelealem/ethiotime)](https://goreportcard.com/report/github.com/kzelealem/ethiotime)
[![Go Version](https://img.shields.io/github/go-mod/go-version/kzelealem/ethiotime)](https://golang.org/)
[![License](https://img.shields.io/github/license/kzelealem/ethiotime)](https://github.com/kzelealem/ethiotime/blob/main/LICENSE)

A robust Go package for Ethiopian calendar date and time conversion, offering seamless integration with Go's standard `time.Time` package.

## Features

- 🔄 Convert between Gregorian and Ethiopian calendar dates
- ⏰ Ethiopian time representation (6 hours behind UTC)
- 📅 Support for Ethiopian months and weekdays in Amharic
- 🎨 Flexible date and time formatting
- 🌐 Full compatibility with Go's time.Time

## Installation

```bash
go get github.com/kzelealem/ethiotime
```

## Quick Start

```go
package main

import (
    "fmt"
    "time"
    ethiotime "github.com/kzelealem/ethiotime"
)

func main() {
    // Get current Ethiopian date and time
    now := ethiotime.Now()
    fmt.Println("Current Ethiopian date:", now.String())

    // Convert from Gregorian date
    gregorianDate := time.Date(2024, 1, 1, 14, 30, 0, 0, time.UTC)
    ethiopianDate := ethiotime.Date(gregorianDate)
    
    // Format the date
    formatted := ethiopianDate.Format("Monday, January 02, 2006 03:04 PM")
    fmt.Println("Formatted date:", formatted)
}
```

## Usage Examples

### 1. Getting Ethiopian Date Components

```go
year, month, day := ethiopianDate.Date()
fmt.Printf("Year: %d, Month: %s, Day: %d\n", year, month, day)
```

### 2. Working with Time

```go
// Get hour, minute, second in Ethiopian time
hour, min, sec := ethiopianDate.Clock()
fmt.Printf("Time: %02d:%02d:%02d\n", hour, min, sec)
```

### 3. Formatting Options

```go
// Different formatting examples
fmt.Println(ethiopianDate.Format("January 02, 2006"))     // Date only
fmt.Println(ethiopianDate.Format("03:04 PM"))           // Time only
fmt.Println(ethiopianDate.Format("Monday, January 02"))  // Weekday and date
fmt.Println(ethiopianDate.Format("01/02/2006"))         // Numeric date
```

## API Reference

### Types

- `Time`: Represents an Ethiopian calendar date and time
- `Month`: Specifies a month of the Ethiopian calendar (Meskerem = 1, ...)
- `Weekday`: Specifies a day of the week (Sunday = 0, ...)

### Main Functions

- `Now() Time`: Returns the current time in Ethiopian calendar
- `Date(time.Time) Time`: Converts Gregorian time to Ethiopian time
- `(Time) Format(layout string) string`: Formats the time according to the layout
- `(Time) Clock() (hour, min, sec int)`: Returns Ethiopian time components
- `(Time) Weekday() Weekday`: Returns the day of the week
- `(Time) GetGregorian() time.Time`: Returns the Gregorian calendar equivalent

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Ethiopian Calendar conversion algorithms
- Ethiopian time system (6 hours behind UTC)
- Amharic month and weekday names

