package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	c := &http.Client{}
	resp, err := c.Get("http://www.reddit.com/r/funny.json")
	if err != nil {
		fmt.Println("there was an error")
	} else {
		defer resp.Body.Close()
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("there was an error in reading the body")
		}
		//getimages should later be a interface method so that getimages can be used on any json implemetnted or something like that
		s, err := getImages([]byte(contents))
		fmt.Println(s.Data)
	}
}

func getImages(body []byte) (*RedditJson, error) {
	s := new(RedditJson)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("Error unmarshalling", err)
	}
	return s, err
}
