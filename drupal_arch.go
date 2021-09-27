package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	help := flag.String("help", "", "This listing.")
	read_dir := flag.String("read_dir", ".", "Directory containing config files. Do not include trailing slash. Defaults to current directory.")
	write_dir := flag.String("write_dir", ".", "Directory to write CSVs to. Do not include trailing slash. Defaults to current directory")
	flag.Parse()

	if *help == "help" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	var content_types []string
	var someStruct map[string]interface{}

	// Walk the read_dir and check each file to see if it belongs in one of our slices.
	filepath.Walk(*read_dir, func(path string, file os.FileInfo, _ error) error {
		if !file.IsDir() {
			// Any file that begins with "node.type" is a content type definition.
			r, err := regexp.MatchString("node.type*", file.Name())
			if err == nil && r {
				content_types = append(content_types, file.Name())
			} else {
				checkError(err)
			}
		}
		return nil
	})

	file, err := os.Create(*write_dir + "/content_types.csv")
	checkError(err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Parse content type files and create array of records
	for _, file := range content_types {
		data, err := ioutil.ReadFile(*read_dir + "/" + file)
		checkError(err)

		jsonDoc, err := yaml.YAMLToJSON(data)
		checkError(err)

		err = json.Unmarshal(jsonDoc, &someStruct)
		checkError(err)
		record := []string{someStruct["type"].(string), someStruct["name"].(string), someStruct["description"].(string)}

		err = writer.Write(record)
		checkError(err)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
