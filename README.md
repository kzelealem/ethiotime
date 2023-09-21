# ethiotime: Ethiopian Date and Time Converter for Go

ethiotime is a specialized GoLang package tailored for effortless computation and presentation of the Ethiopian date and time system, offering seamless conversion of the standard time.Time values and customizable formatting.

# 🚀 Installation
To install the package, use the following:
```
go get github.com/Kzelealem/ethiotime
```

# 📖 Usage
## 1. Importing the package
```
import (
  ethiotime "github.com/Kzelealem/ethiotime"
)
```

## 2. Getting Current Ethiopian Time
``` go
// Retrieve current Ethiopian time
t := ethiotime.Now()

```

## 3. Converting from GoLang's time.Time
If you have an existing GoLang time.Time instance and you'd like to convert it:
```
tt := time.Now()
et := ethiotime.Date(tt)
```

## Formatting to Ethiopian Time String
Customize your date format using standard layout strings:
```
formattedTime := t.Format("January 01, 2006 03:04 PM")
fmt.Println(formattedTime)
```

