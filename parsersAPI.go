package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Person struct {
	Name     string   `json:"name"`
	Status   string   `json:"status"`
	Species  string   `json:"species"`
	Gender   string   `json:"gender"`
	Origin   Ori      `json:"origin"`
	Location Loc      `json:"location"`
	Image    string   `json:"image"`
	Episode  []string `json:"episode"`
}

type PersonJ struct {
	Name     string `json:"name"`
	Status   string `json:"status"`
	Species  string `json:"species"`
	Gender   string `json:"gender"`
	Origin   string `json:"origin"`
	Location string `json:"location"`
	Image    string `json:"image"`
	Episode  string `json:"episode"`
}
type FatJson struct {
	Result []Person `json:"results"`
}
type OutputFatJson struct {
	Result []PersonJ `json:"results"`
}

type Ori struct {
	Name string `json:"name"`
}
type Loc struct {
	Name string `json:"name"`
}

func parseHTTP(requestURL string) []byte {
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		log.Fatalf("client: could not create request: %s\n", err)
		os.Exit(1)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	return resBody
}

func linkCollector(name, status string) string {
	baseURL := "https://rickandmortyapi.com/api/character/"
	params := url.Values{}
	params.Add("name", name)
	params.Add("status", status)
	fullUrl := baseURL + "?" + params.Encode()
	return fullUrl
}

func parsFatJson(i []byte) OutputFatJson {
	var s FatJson
	ans := make([]PersonJ, 0)
	if err := json.Unmarshal(i, &s); err != nil {
		log.Fatalf("Eror func parsJSON: %s", err)
	}
	for _, z := range s.Result {
		ans = append(ans, PersonJ{
			Name:     z.Name,
			Status:   z.Status,
			Species:  z.Species,
			Gender:   z.Gender,
			Origin:   z.Origin.Name,
			Location: z.Location.Name,
			Image:    z.Image,
			Episode:  parseEp(z)})

	}
	return OutputFatJson{Result: ans}
}

func parseEp(a Person) string {
	i := a.Episode
	ans := ""
	k := 0
	for x, n := range i {
		k += 1
		if x == (len(i) - 1) {
			ans += (strings.Trim(n, "https://rickandmortyapi.com/api/episode/"))
		} else if k <= 4 {
			ans = ((strings.Trim(n, "https://rickandmortyapi.com/api/episode/")) + ",")
		} else {
			ans = ((strings.Trim(n, "https://rickandmortyapi.com/api/episode/")) + "," + "\n")
			k = 0
		}
	}
	return ans
}
