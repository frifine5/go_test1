package main

import (
	"encoding/asn1"
	"fmt"
	"io/ioutil"
	"encoding/base64"
	"encoding/hex"
)




type Seq002 struct {
	asn1.RawContent
	EncOid Seq0020
	asn1.BitString
}

type Seq0020 struct{
	H asn1.RawContent
	Enc asn1.ObjectIdentifier
	EncEnt asn1.ObjectIdentifier
}




func main(){

	toid := asn1.ObjectIdentifier{1,2,840,10045,2,1}
	did, err := asn1.Marshal(toid)
	if err !=nil{
		fmt.Println("1")
		panic(err)
		return
	}
	fmt.Println(hex.EncodeToString(did))

	hpk := "044D4F588B76888D9D74785DB87A18FD346743602070DD21D824B8E4027452D68D90B6339CF86E48278ABD7B2FC249094FF31CD45C59EDCE0B21169AF505FA0ED8"
	pk, err := hex.DecodeString(hpk)
	if err !=nil{
		fmt.Println("2")
		panic(err)
		return
	}
	encOid := asn1.ObjectIdentifier{ 1,2,840,10045,2,1}
	encEntOid := asn1.ObjectIdentifier{1,2,156,10197,1,301}
	seqPkOid := Seq0020{
		asn1.RawContent{},
		encOid,
		encEntOid,
	}
	//seqPkOid.enc = encOid
	//seqPkOid.enc = encEntOid

	outData, err := asn1.Marshal(seqPkOid)
	if err !=nil{
		fmt.Println("3")
		panic(err)
		return
	}

	bitSpk := asn1.BitString{
		pk,
		len(pk),
	}
	seqPk := Seq002{
		asn1.RawContent{},
		seqPkOid,
		bitSpk,
	}
	outData, err = asn1.Marshal(seqPk)
	if err !=nil{
		panic(err)
		return
	}
	fmt.Println(seqPk)

	outData = [] byte(base64.StdEncoding.EncodeToString(outData))
	if ioutil.WriteFile("/home/gotmp/asn2.bin", outData, 0644) == nil{
		fmt.Println("ioutil.WriteFile write asn1.bin success")
	}

}
