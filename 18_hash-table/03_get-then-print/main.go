package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//Go grab the Moby Dick txt file
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	//Error reporting
	if err != nil {
		log.Fatal(err)
	}
	//Set the body from the response of the Get Request of Moby Dick and set it to var page
	page, err := ioutil.ReadAll(res.Body)
	//Close the body request
	res.Body.Close()
	//Error reporting
	if err != nil {
		log.Fatal(err)
	}
	//Print out the text of mody dick
	fmt.Printf("%s", page)
}
