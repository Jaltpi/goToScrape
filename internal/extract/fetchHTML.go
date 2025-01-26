package extract

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func FetchHTML(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	fmt.Printf("Response returned with status code: %d", resp.StatusCode)
	fmt.Println("Returning Response Body:")
	return string(body)
}
