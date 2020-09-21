package main

/*
	This program takes in a text file and takes a count of each word as it takes in each word from the file.

	Written by: Huy Le
*/

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {

	// ask user for an input
	fmt.Println("Enter in file location: ")
	var location string
	fmt.Scanln(&location) // Scanln takes in the input, the & symbol is a pointer to where variable file is stored
	fmt.Println()

	storyFile, err := ioutil.ReadFile(location) // read in a text file
	if err != nil {
		log.Fatal(err)
	}
	storyText := string(storyFile) // turns content in text file into a string

	// string processing
	reg, err := regexp.Compile("[^a-zA-Z0-9]+") // create a regexp variable to remove symbols, spaces are included
	if err != nil {
		log.Fatal(err)
	}
	processedText := reg.ReplaceAllString(strings.ToLower(storyText), " ") // replace each symbol with a space and lower case all letters

	// split the string into a slice and then remove its duplicates
	splitText := strings.Split(processedText, " ") // split string by a space and cast into an array
	uniqueText := removeDuplicate(splitText)

	// copy the elements from the uniqueText into a map
	textMap := make(map[string]int)        // make an empty map of string keys and int values
	for i := 0; i < len(uniqueText); i++ { // copy uniqueText into the map and assign its values to 0
		textMap[uniqueText[i]] = 0
	}

	// compare the keys of the map to the slice of texts from the file
	for key := range textMap {
		for i := 0; i < len(splitText); i++ {
			if key == splitText[i] { // increment its value everytime there is a duplicate string in the text
				textMap[key]++
			}
		}
	}

	// create a new file and print to it
	var outputLocation string = "textcount.txt"
	newFile, err := os.Create(outputLocation)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer newFile.Close()

	writer := bufio.NewWriter(newFile) // creates writer object to wrote to a specific location
	// print out keys and its values into console and output location
	for key, values := range textMap {
		fmt.Printf("%s %d\n", key, values)
		fmt.Fprintf(writer, "%s %d\n", key, values)
	}

	fmt.Fprintf(writer, "\nProgram is complete. End of file")
	writer.Flush() // ensures that data is flushed in case code stops writing
	newFile.Close()

	fmt.Println("\nProgram is complete and written to: ", outputLocation)
}

// function remove duplicate strings within a slice
func removeDuplicate(mainSlice []string) []string {
	duplicate := make(map[string]bool) // create an empty map of keys string and values of bool
	uniqueSlice := []string{}          // create an empty slice of strings
	for i := range mainSlice {         // take slice and copy into map; note: maps will replace any duplicate keys
		duplicate[mainSlice[i]] = true // assign the keys to true; the values are not important
	}
	for key, _ := range duplicate { // note: for key, _ 	==		for string, bool
		uniqueSlice = append(uniqueSlice, key) // append to a slice with the element key
	}
	return uniqueSlice // this slice contains a copy of the string elements key from map
}
