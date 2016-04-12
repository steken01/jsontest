package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func getReddit(subreddit string) []string {

	var urlArray = []string{}
	c := &http.Client{}
	resp, err := c.Get(subreddit)
	if err != nil {
		fmt.Println("there was an error")
	} else {
		defer resp.Body.Close()
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("there was an error in reading the body")
		}
		s, err := getImages([]byte(contents))
		if err != nil {
			fmt.Println("error getting images", err)
		}
		fmt.Println("entering forloop now")
		for _, v := range s.Data.Children {

			urlArray = append(urlArray, v.Data.URL)
			fmt.Printf("Appended %s to array\n", v.Data.URL)
		}
	}
	return urlArray
}

func getImages(body []byte) (*RedditJson, error) {
	s := new(RedditJson)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("Error unmarshalling", err)
	}
	return s, err
}

func downloadImages(s []string) {
	// loop over the slice of strings
	for _, v := range s {
		//chop the slice to create the filename
		filename := "/downloads/" + GetMD5Hash(v) + string(v[len(v)-4])
		//create a file for writing returnin a pointer to a writable file
		out, err := os.Create(filename)
		if err != nil {
			fmt.Printf("Problem creating file", err)
		}
		defer out.Close()
		//Get the file using the url
		fmt.Printf("retreiving %s", v)
		resp, err := http.Get(v)
		if err != nil {
			fmt.Println("there was an error retreiving files", err)
		}
		defer resp.Body.Close()
		//Write the body to file
		_, writeerror := io.Copy(out, resp.Body)
		if err != nil {
			fmt.Printf("there was an error writing to file", writeerror)
		}
	}
}
func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
func main() {
	res := getReddit("http://www.reddit.com/r/funny.json")
	downloadImages(res)
}
