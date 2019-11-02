
package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	"net/url"
)



func main(){


	url0 := "http://localhost:10090/makesealapi/api/gdata/pdf/oldprehash?"
	u, _ := url.Parse(url0)
	q := u.Query()
	q.Set("fileId", "e843453cbf8d45abae18a3f55fa52fb8.pdf")
	u.RawQuery = q.Encode()
	res, err := http.Get(u.String())
	if err != nil {
		log.Fatal(err)
		return
	}
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%s", result)

}