# ethiotime

```
import (
  ethiotime "github.com/Kzelealem/ethiotime"
)
...
// where you want to use
// for currernt time
t := ethiotime.Now()
// if you have the time already
tt := time.Now()
et := ethiotime.Date(tt)

// for formatting
fmt.Println(t.Format("January 01, 2006 03:04 PM"))
```
