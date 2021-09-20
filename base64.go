// https://gobyexample.com/base64-encoding
// https://gobyexample.com/command-line-arguments
// https://gobyexample.com/if-else
// https://www.base64decode.org/

// https://blog.filippo.io/shrink-your-go-binaries-with-this-one-weird-trick/
// go build -ldflags="-s -w" base64.go
// https://github.com/xaionaro/documentation/blob/master/golang/reduce-binary-size.md
// go build -a -gcflags=all="-l -B -wb=false" -ldflags="-w -s"; stat -c %s base64.go
// UPX https://softfamous.com/upx/
// https://www.cockroachlabs.com/blog/go-file-size/

package main

import (
    b64 "encoding/base64"
    "fmt"
	"os"
)

func main() {
	
	num_args := len(os.Args)
	
	//fmt.Println(num_args)
	
	if num_args==3 {
	
		enc_dec := os.Args[1]
		//fmt.Println(enc_dec)
		data := os.Args[2]

		if (enc_dec=="-e"){
			sEnc := b64.StdEncoding.EncodeToString([]byte(data))
			fmt.Println(sEnc)
		}

		if (enc_dec=="-d"){
			uDec, _ := b64.StdEncoding.DecodeString(data)
			fmt.Println(string(uDec))
			fmt.Println()
		}

	}
}