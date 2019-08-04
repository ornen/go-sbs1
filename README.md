# go-sbs1

This is a SBS1 parsing library for Go.

## Example
The following example connects to the SBS1 port of `dump1090` and prints the messages.
```go
package main

import (
	"github.com/ornen/go-sbs1"
	"io"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30003")

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	var reader = sbs1.NewReader(conn)

	for {
		var message, err = reader.Read()

		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Println(err)
				continue
			}
		}

		log.Println(message)
	}
}
```

## License

This code is licensed under the Apache License 2.0.
