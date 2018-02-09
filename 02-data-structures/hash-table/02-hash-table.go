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

	//Create slice to hold counts
	buckets := make([]int, 200)

	//Loop over the words
	for scanner.Scan() {
		n := HashBucket(scanner.Text())
		/*
			That will count how many words start with a specific letter
			buckets[65] = buckets[65] + 1
			buckets[65] += buckets[65]
		*/
		buckets[n]++
	}

	//That will print: letter 65 have 203 words starting with this letter
	fmt.Println(buckets[65:123])
}

/*
	Convert word in number
	Look https://pt.wikipedia.org/wiki/ASCII
	Look the example 07-map.go to see the letter version
*/
func HashBucket(word string) int {
	return int(word[0])
}
