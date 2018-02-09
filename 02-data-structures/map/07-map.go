package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	//get the book
	res, err := http.Get("https://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatalln(err)
	}

	//scan the page
	scanner := bufio.NewScanner(res.Body)

	//defer always run at the end. Close the scanner
	defer res.Body.Close()

	//set the split function for the scanning operation
	scanner.Split(bufio.ScanWords)

	//Create map to hold counts
	buckets := make(map[string]int)

	//Loop over the words
	for scanner.Scan() {
		n := HashBucket(scanner.Text())
		//That will count how many words start with a specific letter
		buckets[n] += 1
	}

	fmt.Println(buckets)
}

//Return the first letter in the word
func HashBucket(word string) string {
	return string(word[0])
}
