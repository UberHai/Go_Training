package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	//Go grab the Sherlock Homes txt file - page 1661
	//res, err := http.Get("http://www.gutenberg.org/cache/epub/1661/pg1661.txt")

	//Grab the Moby Dick txt file
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")

	//Error reporting
	if err != nil {
		log.Fatal(err)
	}
	//Scan the page
	scanner := bufio.NewScanner(res.Body)
	//Close the scanner at end of function
	defer res.Body.Close()
	//Set the split function to words for the scanning operation
	scanner.Split(bufio.ScanWords)
	//Create slice of slice of string to hold slices of words
	//12 Buckets
	buckets := make([][]string, 12)
	//Create slices to hold words
	for i := 0; i < 12; i++ {
		buckets = append(buckets, []string{})
	}
	//Loop through the words
	for scanner.Scan() {
		word := scanner.Text()
		n := HashBucket(word, 12)
		buckets[n] = append(buckets[n], word)
	}
	//Print the length of each bucket
	for i := 0; i < 12; i++ {
		fmt.Println(i, " - ", len(buckets[i]))
	}
	//Print the words in one bucket
	fmt.Println(buckets[11]) //7th Bucket
}

func HashBucket(word string, buckets int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % buckets
}
