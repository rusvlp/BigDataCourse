package main

import (
	"encoding/json"
	"fmt"
	"json2file/internal/entity"
	"os"
	"strconv"
	"time"
)

func main() {
	generateCsvArray()
}

func generateCsvArray() {
	_, gen := entity.NewSiteGen()

	maxIts := 10

	if len(os.Args) >= 2 {
		maxIts, _ = strconv.Atoi(os.Args[1])
	}

	fmt.Println(entity.GetSiteHeaders())

	for i := 0; i < maxIts; i++ {
		_, site := gen.GenerateSite()
		fmt.Println(site.GetValues())
	}

	_, site := gen.GenerateSite()

	fmt.Println(site.GetValues())
}

func generateSiteJsonArray() {
	_, gen := entity.NewSiteGen()

	maxIts := 0

	if len(os.Args) >= 2 {
		maxIts, _ = strconv.Atoi(os.Args[1])
	}

	fmt.Print("[")
	for i := 0; i < maxIts; i++ {

		_, site := gen.GenerateSite()

		siteJson, _ := json.Marshal(site)

		if i == maxIts-1 {
			fmt.Print(string(siteJson))
		} else {
			fmt.Print(string(siteJson) + ",")
		}

	}
	fmt.Print("]")

}

func generateJsonToFile() {
	args := os.Args

	err, skinGen := entity.NewGenerator()

	if err != nil {
		fmt.Println(err)
	}

	var path string
	if len(args) <= 2 {
		path = "./"
	} else {
		path = args[1]
	}

	var timeout int
	if len(args) <= 3 {
		timeout = 5
	} else {
		timeout, err = strconv.Atoi(args[2])
	}

	fmt.Println(path + " " + strconv.Itoa(timeout))

	for i := 0; ; i++ {

		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			errDir := os.MkdirAll(path, 0755)
			if errDir != nil {
				fmt.Println(errDir)
			}
		}

		fileName := "SKIN_" + strconv.Itoa(i) + ".json"

		filePath := path + fileName

		fmt.Println(filePath)

		file, err := os.Create(filePath)
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()

		err, skin := skinGen.GenerateSkinPrice()

		if err != nil {
			fmt.Println(err)
		}

		jsonSkin, err := json.Marshal(skin)

		_, err = file.Write(jsonSkin)

		if err != nil {
			fmt.Println(err)
		}

		time.Sleep(time.Duration(int64(timeout)) * time.Second)
	}
}
