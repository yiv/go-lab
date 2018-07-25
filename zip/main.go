package main

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	r, err := zip.OpenReader("./app.apk")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	// Iterate through the files in the archive,
	// printing some of their contents.
	for _, f := range r.File {
		//fmt.Println(f.Name)
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		c, err := ioutil.ReadAll(rc)
		if err != nil {
			log.Fatal(err)
		}
		dir := "main/"

		if filepath.Dir(f.Name) != "." {
			dir += filepath.Dir(f.Name)
		}
		if _, err = os.Stat(dir); os.IsNotExist(err) {
			if err := os.MkdirAll(dir, 0755); err != nil {
				log.Fatal("1" + err.Error())
			}
		}
		fmt.Println(dir + f.Name)
		abs, _ := filepath.Abs(dir + f.Name)
		fmt.Println(abs)
		if err = ioutil.WriteFile(abs, c, 0755); err != nil {
			log.Fatal("2" + err.Error())
		}
	}
}
