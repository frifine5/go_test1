package main


import (
	"fmt"
	"os"
	"path/filepath"
	"io/ioutil"
	"time"
)

func main() {
	dir := "C:\\Users\\49762\\Desktop\\itext"
/*	list, err := getDirList(dir)
	if err != nil {
		fmt.Println(err)
		return
	}*/

	list2, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return
	}

	flist := []string{}
	for _, v := range list2 {
		if v.IsDir(){
			fmt.Println(v)
		}else{
			fmt.Printf(v.Name() + "  ")
			flist = append(flist, v.Name())
		}
	}
	fmt.Println("---------------------")
	fmt.Println(len(flist))
	for _, v:= range flist{
		fmt.Printf(v + "  ")
	}

	if ioutil.WriteFile( "C:\\Users\\49762\\Desktop\\finish-record",
		[]byte(  "finish time: "+time.Now().Format("2006-01-02 15:04:05")), 0644) == nil {
		fmt.Println("write err")
	}else{
		fmt.Println("not nil")
	}

}

func getDirList(dirpath string) ([]string, error) {
	var dir_list []string
	dir_err := filepath.Walk(dirpath,
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if f.IsDir() {
				dir_list = append(dir_list, path)
				return nil
			}

			return nil
		})
	return dir_list, dir_err
}