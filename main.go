package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Entry point
func main() {
	if os.Args[1] == "gen" {
		generateRelease()
	} else if os.Args[1] == "new" {
		createNewVersion()
	} else {
		fmt.Println("Invalid Command")
	}
}

// Creates a new version
func createNewVersion() {
	oldRelease := os.Args[2]
	newRelease := os.Args[3]
	buildOutputFolder(newRelease + "/")
	files, err := ioutil.ReadDir("./" + oldRelease + "/")

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fileContents := getFileContents(oldRelease + "/" + f.Name())
		createOutputFile(newRelease+"/"+f.Name(), fileContents)
	}
}

// Generates a release
func generateRelease() {
	release := os.Args[2]
	configFile := release + "/" + os.Args[3] + ".conf"
	config := getFileContents(configFile)
	outputPath := "target/" + os.Args[3] + "/" + release + "/"
	buildOutputFolder(outputPath)
	files, err := ioutil.ReadDir("./" + release + "/")

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if !strings.Contains(f.Name(), ".conf") {
			fmt.Println(f.Name())
			output := buildTemplateWithConfig(config, getFileContents(release+"/"+f.Name()))
			createOutputFile(outputPath+f.Name(), output)
		}
	}
}

// Returns a string will the templated variables in
// @param config is a multiline string which has keyvalue pairs seperated by an '=' sign (Templates key -> Value)
// @param template is a multiline string which has marked values which need to be replaced by config values
// @return a multiline string that is the template filled in
func buildTemplateWithConfig(config string, template string) string {
	output := template
	configScanner := bufio.NewScanner(strings.NewReader(config))
	for configScanner.Scan() {
		line := configScanner.Text()
		if line[0] != '#' {
			i := strings.Index(line, "=")
			key := fmt.Sprintf("{{%s}}", line[:i])
			value := line[i+1:]
			output = strings.Replace(output, key, value, -1)
		}
	}
	return output
}

// Creates a file with the given content
// @param file the file path
// @param the contents of the file
func createOutputFile(file string, contents string) {
	os.Remove(file)
	err := ioutil.WriteFile(file, []byte(contents), 0644)
	if err != nil {
		fmt.Println("Failed to save generated file: " + file)
		os.Exit(0)
	}
}

// Creates the output folder where all generated files will go
// @param outputFolder the path for the folder
func buildOutputFolder(outputFolder string) {
	_, err := os.Stat(outputFolder)
	if os.IsNotExist(err) {
		errMkdir := os.MkdirAll(outputFolder, 0755)
		if errMkdir != nil {
			panic("Failed to create the buildfolder. Check directory permissions.")
		}
	}
}

// Returns a string that contains everything in the specified file.
// @param file is the file path
// @return contents of the given file
func getFileContents(file string) string {
	templateBytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("File not found: " + file)
		os.Exit(0)
	}
	return string(templateBytes)
}
