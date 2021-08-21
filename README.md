# ethiotime

ethiotime is a golang package for computing and displaying ethiopian date and time.

# Install
```
go get github.com/Kzelealem/ethiotime
```

# Usage
## Import
```
import (
  ethiotime "github.com/Kzelealem/ethiotime"
)
```

## get current time
```
// where you want to use
// for currernt time
t := ethiotime.Now()
// if you have the time already
```

## get ethiotime from golang time.Time
```
tt := time.Now()
et := ethiotime.Date(tt)
```

## format to ethiotime string
```
// for formatting
fmt.Println(t.Format("January 01, 2006 03:04 PM"))
```

