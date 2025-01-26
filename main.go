package main

import (
	"fmt"
	"os"

	"github.com/Jaltpi/goToScrape/internal/extract"
)

func main() {

	for _, url := range os.Args[1:] {
		fmt.Println("Scraping: ", url)
		fmt.Println(extract.FetchHTML(url))
	}

	fmt.Println("End of Scraping", os.Args[1:])

}
