# whatos

Package whatos provides the splendid function `TellMe()`, which returns the string value: "macOS", "linux", or "other" depending on which operating system it is run on.  
  
## usage

```go
package main

import "github.com/ogryzek/whatos"

func main() {
	myOs := whatos.TellMe()
	fmt.Println(myOs)
}

```
  
[Go Doc](https://godoc.org/github.com/ogryzek/whatos)  

