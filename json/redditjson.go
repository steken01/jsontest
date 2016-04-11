package main

type RedditJson struct {
	Data struct {
		After    string      `json:"after"`
		Before   interface{} `json:"before"`
		Children []struct {
			Data struct {
				//Title string `json:"title"`
				URL string `json:"url"`
			} `json:"data"`
		} `json:"children"`
	} `json:"data"`
}
