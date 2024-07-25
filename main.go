package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// Define structs to match the JSON structure
type Language struct {
	LanguageCode string `json:"language"`
}

type Data struct {
	Languages []Language `json:"languages"`
}

type Response struct {
	Data Data `json:"data"`
}

const spoofApiCalls = true

func getLanguages() ([]Language, error) {
	// 'data' will contain the JSON response from the API
	var data string

	if spoofApiCalls {
		data = `{"data":{"languages":[{"language":"af"},{"language":"ak"},{"language":"am"},{"language":"ar"},{"language":"as"},{"language":"ay"},{"language":"az"},{"language":"be"},{"language":"bg"},{"language":"bho"},{"language":"bm"},{"language":"bn"},{"language":"bs"},{"language":"ca"},{"language":"ceb"},{"language":"ckb"},{"language":"co"},{"language":"cs"},{"language":"cy"},{"language":"da"},{"language":"de"},{"language":"doi"},{"language":"dv"},{"language":"ee"},{"language":"el"},{"language":"en"},{"language":"eo"},{"language":"es"},{"language":"et"},{"language":"eu"},{"language":"fa"},{"language":"fi"},{"language":"fr"},{"language":"fy"},{"language":"ga"},{"language":"gd"},{"language":"gl"},{"language":"gn"},{"language":"gom"},{"language":"gu"},{"language":"ha"},{"language":"haw"},{"language":"he"},{"language":"hi"},{"language":"hmn"},{"language":"hr"},{"language":"ht"},{"language":"hu"},{"language":"hy"},{"language":"id"},{"language":"ig"},{"language":"ilo"},{"language":"is"},{"language":"it"},{"language":"iw"},{"language":"ja"},{"language":"jv"},{"language":"jw"},{"language":"ka"},{"language":"kk"},{"language":"km"},{"language":"kn"},{"language":"ko"},{"language":"kri"},{"language":"ku"},{"language":"ky"},{"language":"la"},{"language":"lb"},{"language":"lg"},{"language":"ln"},{"language":"lo"},{"language":"lt"},{"language":"lus"},{"language":"lv"},{"language":"mai"},{"language":"mg"},{"language":"mi"},{"language":"mk"},{"language":"ml"},{"language":"mn"},{"language":"mni-Mtei"},{"language":"mr"},{"language":"ms"},{"language":"mt"},{"language":"my"},{"language":"ne"},{"language":"nl"},{"language":"no"},{"language":"nso"},{"language":"ny"},{"language":"om"},{"language":"or"},{"language":"pa"},{"language":"pl"},{"language":"ps"},{"language":"pt"},{"language":"qu"},{"language":"ro"},{"language":"ru"},{"language":"rw"},{"language":"sa"},{"language":"sd"},{"language":"si"},{"language":"sk"},{"language":"sl"},{"language":"sm"},{"language":"sn"},{"language":"so"},{"language":"sq"},{"language":"sr"},{"language":"st"},{"language":"su"},{"language":"sv"},{"language":"sw"},{"language":"ta"},{"language":"te"},{"language":"tg"},{"language":"th"},{"language":"ti"},{"language":"tk"},{"language":"tl"},{"language":"tr"},{"language":"ts"},{"language":"tt"},{"language":"ug"},{"language":"uk"},{"language":"ur"},{"language":"uz"},{"language":"vi"},{"language":"xh"},{"language":"yi"},{"language":"yo"},{"language":"zh"},{"language":"zh-CN"},{"language":"zh-TW"},{"language":"zu"}]}}`
	} else {
		key, err := os.ReadFile("api.key")
		if err != nil {
			return nil, err
		}

		url := "https://google-translate1.p.rapidapi.com/language/translate/v2/languages"

		req, _ := http.NewRequest("GET", url, nil)

		req.Header.Add("x-rapidapi-key", string(key))
		req.Header.Add("x-rapidapi-host", "google-translate1.p.rapidapi.com")
		req.Header.Add("Accept-Encoding", "application/gzip")

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}

		defer res.Body.Close()
		body, _ := io.ReadAll(res.Body)
		data = string(body)
	}

	// Parse the JSON object
	var response Response
	err := json.Unmarshal([]byte(data), &response)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	// Return the array of language codes
	return response.Data.Languages, nil
}

func translate(text string, from string, to string) (string, error) {
	translated := from + "--" + text + "--" + to
	return translated, nil
}

func main() {
	/* If a list of languages is requested
	languages, _ := getLanguages()
	for _, language := range languages {
		fmt.Println(language.LanguageCode)
	}
	*/

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <text>")
		os.Exit(1)
	}

	// Build a sample request
	result, err := translate(os.Args[1], "en", "es")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	println(result)
}

/*
	url := "https://google-translate1.p.rapidapi.com/language/translate/v2/languages"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-key", key)
	req.Header.Add("x-rapidapi-host", "google-translate1.p.rapidapi.com")
	req.Header.Add("Accept-Encoding", "application/gzip")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
	var data := string(body)
*/
/*
package main

import (
	"fmt"
	"strings"
	"net/http"
	"io"
)

func main() {

	url := "https://google-translate1.p.rapidapi.com/language/translate/v2/detect"

	payload := strings.NewReader("-----011000010111000001101001\r\nContent-Disposition: form-data; name=\"q\"\r\n\r\nEnglish is hard, but detectably so\r\n-----011000010111000001101001--\r\n\r\n")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("x-rapidapi-key", key)
	req.Header.Add("x-rapidapi-host", "google-translate1.p.rapidapi.com")
	req.Header.Add("Content-Type", "multipart/form-data; boundary=---011000010111000001101001")
	req.Header.Add("Accept-Encoding", "application/gzip")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
*/
