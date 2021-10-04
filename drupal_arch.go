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
	"strings"
)

var write_dir *string
var read_dir *string

func main() {
	help := flag.String("help", "", "This listing.")
	read_dir = flag.String("read_dir", ".", "Directory containing config files. Do not include trailing slash. Defaults to current directory.")
	write_dir = flag.String("write_dir", ".", "Directory to write CSVs to. Do not include trailing slash. Defaults to current directory")
	flag.Parse()

	if *help == "help" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	var contentTypes []string
	var taxonomies []string
	var views []string

	contentTypes = filterDirectoryList("node\\.type*")
	taxonomies = filterDirectoryList("taxonomy\\.vocabulary*")
	views = filterDirectoryList("views\\.view*")

	handleContentTypes(contentTypes)
	handleTaxonomies(taxonomies)
	handleViews(views)
}

// Do all the work to write out the content_types csv
func handleContentTypes(contentTypes []string) {
	// Header row for the CSV
	var header = []string{"Type", "Name", "Description"}
	var typeHeader = []string{"Label", "Machine Name", "Type", "Description", "Required", "Cardinality", "Translatable"}
	var fields []string
	var fieldName string

	// Create the file
	file, err := os.Create(*write_dir + "/content_types.csv")
	checkError(err)
	defer file.Close()

	// Open it for writing
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the header
	err = writer.Write(header)
	checkError(err)

	// Iterate content types array
	for _, contentTypeFile := range contentTypes {
		// Write a record to the main content types file for this content type
		configData := getConfigData(*read_dir + "/" + contentTypeFile)
		record := []string{configData["type"].(string), configData["name"].(string), configData["description"].(string)}
		err = writer.Write(record)
		checkError(err)

		// Extract the actual content type machine name from the file name
		parts := strings.Split(contentTypeFile, ".")
		contentTypeName := parts[2]

		// Create the type-specific file
		typeFile, err := os.Create(*write_dir + "/content_type_" + contentTypeName + ".csv")
		checkError(err)
		defer typeFile.Close()

		// Open it for writing
		typeWriter := csv.NewWriter(typeFile)
		defer typeWriter.Flush()

		// Write the header
		err = typeWriter.Write(typeHeader)
		checkError(err)

		// Find the fields for this content type
		fields = filterDirectoryList("field.field.node." + contentTypeName + "*")
		for _, field := range fields {
			// Extract the actual field machine name from the file name
			parts = strings.Split(field, ".")
			fieldName = parts[4]

			// Get the content-type-specific config data from this field
			typeData := getConfigData(*read_dir + "/" + field)

			// Get the storage-specific config data for this field
			storageData := getConfigData(*read_dir + "/" + "field.storage.node." + fieldName + ".yml")

			// Cardinality -1 = Unlimited
			cardinality := fmt.Sprintf("%v", storageData["cardinality"])
			if cardinality == "-1" {
				cardinality = "Unlimited"
			}

			// Write the row
			record := []string{typeData["label"].(string), fieldName, storageData["type"].(string), typeData["description"].(string), fmt.Sprintf("%v", typeData["required"]), cardinality, fmt.Sprintf("%v", typeData["translatable"])}
			err = typeWriter.Write(record)
			checkError(err)
		}
	}
}

// Do all the work to write out the taxonomies csv
func handleTaxonomies(taxonomies []string) {
	var header = []string{"Type", "Name", "Description"}

	file, err := os.Create(*write_dir + "/taxonomies.csv")
	checkError(err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(header)
	checkError(err)

	// Parse content type files and create array of records
	for _, file := range taxonomies {
		configData := getConfigData(*read_dir + "/" + file)

		record := []string{configData["vid"].(string), configData["name"].(string), configData["description"].(string)}

		err = writer.Write(record)
		checkError(err)
	}
}

// Do all the work to write out the taxonomies csv
func handleViews(views []string) {
	var header = []string{"Label", "Description"}

	file, err := os.Create(*write_dir + "/views.csv")
	checkError(err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(header)
	checkError(err)

	// Parse content type files and create array of records
	for _, file := range views {
		configData := getConfigData(*read_dir + "/" + file)

		record := []string{configData["label"].(string), configData["description"].(string)}

		err = writer.Write(record)
		checkError(err)
	}
}

// Given a filename (with full path), open the file, parse the yaml
// and return a map.
func getConfigData(file string) map[string]interface{} {
	var configData map[string]interface{}

	data, err := ioutil.ReadFile(file)
	checkError(err)

	jsonDoc, err := yaml.YAMLToJSON(data)
	checkError(err)

	err = json.Unmarshal(jsonDoc, &configData)
	checkError(err)

	return configData
}

// Simple and dumb error handler
func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Given a regular expression, return a slice containing the files that match in readDir.
func filterDirectoryList(regex string) []string {
	var results []string

	// Walk the read_dir and check each file to see if it belongs in one of our slices.
	filepath.Walk(*read_dir, func(path string, file os.FileInfo, _ error) error {
		if !file.IsDir() {
			r, err := regexp.MatchString(regex, file.Name())
			if err == nil && r {
				results = append(results, file.Name())
			} else {
				checkError(err)
			}
		}
		return nil
	})

	return results
}
