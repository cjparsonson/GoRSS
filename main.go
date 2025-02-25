package main

import (
	"fmt"
	"encoding/xml"
	"io"
	"net/http"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Version string  `xml:"version,attr"`
	Channel Channel `xml:"channel"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	PubDate     string `xml:"pubDate"`
	Description string `xml:"description"`
}

type Channel struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
}

func fetchRSS(url string) (*RSS, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch RSS: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	var rss RSS
	err = xml.Unmarshal(body, &rss)
	if err != nil {
		return nil, fmt.Errorf("failed to parse RSS: %v", err)
	}
	return &rss, nil
}

func main() {
	feedUrl := "https://news.ycombinator.com/rss"
	
	fmt.Printf("Fetching RSS feed from: %s\n\n", feedUrl)
	
	rss, err := fetchRSS(feedUrl)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	
	fmt.Printf("Channel: %s\n", rss.Channel.Title)
	fmt.Printf("Description: %s\n\n", rss.Channel.Description)
	
	fmt.Println("Latest Items:")
	fmt.Println("-------------")
	for i, item := range rss.Channel.Items {
		if i >= 3 {
			break
		}
		fmt.Printf("%d. %s\n", i+1, item.Title)
		fmt.Printf("   Published: %s\n", item.PubDate)
		fmt.Printf("   Link: %s\n\n", item.Link)
	}
}