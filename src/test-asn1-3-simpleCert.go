package main

import (
	"encoding/asn1"
	"io/ioutil"
	"fmt"
	"encoding/base64"
	"encoding/hex"
	"time"
)




type SM2pkSeq struct {
	EncOid SM2pkOid
	asn1.BitString
}

type SM2pkOid struct{
	Enc asn1.ObjectIdentifier
	EncEnt asn1.ObjectIdentifier
}



type IssNv struct {
	Type  asn1.ObjectIdentifier
	Value interface{}
}

type IssSETs struct {
	CN IssNv `asn1:"set"`
	S IssNv `asn1:"set"`
	O IssNv `asn1:"set"`
	C IssNv `asn1:"set"`
}

type SubjectSETs struct {
	CN IssNv `asn1:"set"`
	O IssNv `asn1:"set"`
	C IssNv `asn1:"set"`

}

type Valid struct {
	S time.Time
	E time.Time
}

type Tbs struct {
	Iss IssSETs
	Vt Valid
	Own SubjectSETs
	Pkseq SM2pkSeq

}



func main(){


	issuer := IssSETs{
		IssNv{
			asn1.ObjectIdentifier{2,5,4,6},
			"CN",
		},
		IssNv{
			asn1.ObjectIdentifier{2,5,4,7},
			"beijing",
		},
		IssNv{
			asn1.ObjectIdentifier{2,5,4,10},
			"godoit.l.t",
		},
		IssNv{
			asn1.ObjectIdentifier{2,5,4,3},
			"名称",
		},
	}

	owner := SubjectSETs{
		IssNv{
			asn1.ObjectIdentifier{2,5,4,6},
			"CN",
		},
		IssNv{
			asn1.ObjectIdentifier{2,5,4,10},
			"test",
		},
		IssNv{
			asn1.ObjectIdentifier{2,5,4,3},
			"名称1",
		},
	}

	valid := Valid{
		time.Now().UTC(),
		//time.Date(2020, 4, 16, 12, 30, 00, 0, time.UTC),
		time.Now().AddDate(2, 0,0).UTC(),
	}

	hpk := "044D4F588B76888D9D74785DB87A18FD346743602070DD21D824B8E4027452D68D90B6339CF86E48278ABD7B2FC249094FF31CD45C59EDCE0B21169AF505FA0ED8"
	pk, err := hex.DecodeString(hpk)
	if err !=nil{
		panic(err)
		return
	}
	encOid := asn1.ObjectIdentifier{ 1,2,840,10045,2,1}
	encEntOid := asn1.ObjectIdentifier{1,2,156,10197,1,301}
	pkSeq := SM2pkSeq{
		SM2pkOid{
			encOid,
			encEntOid,
		},
		asn1.BitString{
			pk,
			len(pk),
		},
	}

	name := Tbs{
		issuer,
		valid,
		owner,
		pkSeq,
	}



	data, err := asn1.Marshal(name)
	if err != nil {
		panic(err)
		return
	}



	fmt.Println(hex.EncodeToString(data))


	data = [] byte(base64.StdEncoding.EncodeToString(data))
	if ioutil.WriteFile("/home/gotmp/asn1-3.cer.bin", data, 0644) == nil{
		fmt.Println("ioutil.WriteFile write asn1.bin success")
	}


	t1 := time.Now()
	fmt.Println(t1)
	t1 = t1.AddDate(10, 0, 0)
	fmt.Println(t1)

}
