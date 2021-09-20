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
