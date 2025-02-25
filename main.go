package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Channel Channel  `xml:"channel"`
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

type Config struct {
	Feeds []Feed `json:"feeds"`
}

type Feed struct {
	Name     string `json:"name"`
	URL      string `json:"url"`
	Category string `json:"category,omitempty"`
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

func readConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %v", err)
	}
	defer file.Close()
	// Create a new struct to hold the decoded config
	var config Config
	// Create a new JSON decoder and unmarshal the config into the struct
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, fmt.Errorf("failed to decode config: %v", err)
	}

	return &config, nil
}

func main() {
	// Read config
	fmt.Println("Reading config...")
	config, err := readConfig("config.json")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("... 	 success.")

	// TODO: For each feed in config:
	// - Show header for current feed
	// - Fetch and display feed content
	// - Add visual separation between feeds
	for _, feed := range config.Feeds {

		// TODO: Improve error handling:
		// - Track successful vs failed fetches
		// - Continue processing if one feed fails
		feedURL := feed.URL
		fmt.Printf("Fetching RSS feed from: %s\n\n", feedURL)

		rss, err := fetchRSS(feedURL)
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
}
