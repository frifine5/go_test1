package main

import (
	"encoding/asn1"
	"fmt"
	"io/ioutil"
	"encoding/base64"
)


type Seq struct{
	asn1.RawContent		// if sequence use it
	asn1.BitString
}




func main(){

	data := []byte("test bit string")

	b1 := asn1.BitString{
		data,
		len(data),
	}
	fmt.Println(b1)
	fmt.Println(string(b1.Bytes))


	seqB1 := Seq{
		asn1.RawContent{},
		b1,
	}

	outData, err := asn1.Marshal(seqB1)
	if err !=nil{
		panic(err)
	}

	outData = [] byte(base64.StdEncoding.EncodeToString(outData))
	if ioutil.WriteFile("/home/gotmp/asn1.bin", outData, 0644) == nil{
		fmt.Println("ioutil.WriteFile write asn1.bin success")
	}

}
