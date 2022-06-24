package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const imageCID = "bafybeibtbsrwnifl5drdzmu2v5fgcyq4othrlx3h3shcvua2om25gnfmvy"

func main() {
	fmt.Println("hello github actions")

	//TODO create histogram metric

	client := &http.Client{}

	url := fmt.Sprintf("https://api.nft.storage/%s", imageCID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("err: %w\n", err)
		return
	}

	tok := os.Getenv("NFT_STORAGE_TOKEN")
	req.Header.Add("Authorization", tok)

	//TODO start timer

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("err: %w\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("err: %w\n", err)
		return
	}

	fmt.Println(body)

	//TODO finish timer

	//TODO push metric

}
