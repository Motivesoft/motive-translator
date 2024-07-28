package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	flag "github.com/spf13/pflag"
)

// Define structs to match the JSON structure
type Language struct {
	LanguageCode string `json:"language"`
}

type Translation struct {
	TranslatedText string `json:"translatedText"`
}

type Data struct {
	Languages    []Language  `json:"languages"`
	Translations Translation `json:"translations"`
}

type Response struct {
	Data Data `json:"data"`
}

// Define the structure for the JSON data
type TranslationRequest struct {
	Q      string `json:"q"`
	Source string `json:"source"`
	Target string `json:"target"`
}

const spoofApiCalls = false

const versionString = "0.0.4"

func getLanguages(apiKey string) ([]Language, error) {
	// 'data' will contain the JSON response from the API
	var data string

	if spoofApiCalls {
		data = `{"data":{"languages":[{"language":"af"},{"language":"ak"},{"language":"am"},{"language":"ar"},{"language":"as"},{"language":"ay"},{"language":"az"},{"language":"be"},{"language":"bg"},{"language":"bho"},{"language":"bm"},{"language":"bn"},{"language":"bs"},{"language":"ca"},{"language":"ceb"},{"language":"ckb"},{"language":"co"},{"language":"cs"},{"language":"cy"},{"language":"da"},{"language":"de"},{"language":"doi"},{"language":"dv"},{"language":"ee"},{"language":"el"},{"language":"en"},{"language":"eo"},{"language":"es"},{"language":"et"},{"language":"eu"},{"language":"fa"},{"language":"fi"},{"language":"fr"},{"language":"fy"},{"language":"ga"},{"language":"gd"},{"language":"gl"},{"language":"gn"},{"language":"gom"},{"language":"gu"},{"language":"ha"},{"language":"haw"},{"language":"he"},{"language":"hi"},{"language":"hmn"},{"language":"hr"},{"language":"ht"},{"language":"hu"},{"language":"hy"},{"language":"id"},{"language":"ig"},{"language":"ilo"},{"language":"is"},{"language":"it"},{"language":"iw"},{"language":"ja"},{"language":"jv"},{"language":"jw"},{"language":"ka"},{"language":"kk"},{"language":"km"},{"language":"kn"},{"language":"ko"},{"language":"kri"},{"language":"ku"},{"language":"ky"},{"language":"la"},{"language":"lb"},{"language":"lg"},{"language":"ln"},{"language":"lo"},{"language":"lt"},{"language":"lus"},{"language":"lv"},{"language":"mai"},{"language":"mg"},{"language":"mi"},{"language":"mk"},{"language":"ml"},{"language":"mn"},{"language":"mni-Mtei"},{"language":"mr"},{"language":"ms"},{"language":"mt"},{"language":"my"},{"language":"ne"},{"language":"nl"},{"language":"no"},{"language":"nso"},{"language":"ny"},{"language":"om"},{"language":"or"},{"language":"pa"},{"language":"pl"},{"language":"ps"},{"language":"pt"},{"language":"qu"},{"language":"ro"},{"language":"ru"},{"language":"rw"},{"language":"sa"},{"language":"sd"},{"language":"si"},{"language":"sk"},{"language":"sl"},{"language":"sm"},{"language":"sn"},{"language":"so"},{"language":"sq"},{"language":"sr"},{"language":"st"},{"language":"su"},{"language":"sv"},{"language":"sw"},{"language":"ta"},{"language":"te"},{"language":"tg"},{"language":"th"},{"language":"ti"},{"language":"tk"},{"language":"tl"},{"language":"tr"},{"language":"ts"},{"language":"tt"},{"language":"ug"},{"language":"uk"},{"language":"ur"},{"language":"uz"},{"language":"vi"},{"language":"xh"},{"language":"yi"},{"language":"yo"},{"language":"zh"},{"language":"zh-CN"},{"language":"zh-TW"},{"language":"zu"}]}}`
	} else {
		url := "https://deep-translate1.p.rapidapi.com/language/translate/v2/languages"

		req, _ := http.NewRequest("GET", url, nil)

		req.Header.Add("x-rapidapi-key", apiKey)
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
		return nil, err
	}

	// Return the array of language codes
	return response.Data.Languages, nil
}

func translate(text string, from string, to string, apiKey string) (string, error) {
	var data string
	if spoofApiCalls {
		data = `{"data":{"translations":{"translatedText":"\u00a1Hola Mundo!"}}}`
	} else {
		url := "https://deep-translate1.p.rapidapi.com/language/translate/v2"

		// Create an instance of the request structure so we can encode the translation query
		request := TranslationRequest{
			Q:      text,
			Source: from,
			Target: to,
		}

		// Encode the structure to JSON
		jsonData, err := json.Marshal(request)
		if err != nil {
			return "", fmt.Errorf("error encoding request: %v", err)
		}

		// Make the JSON string available as a reader
		payload := strings.NewReader(string(jsonData))

		req, err := http.NewRequest("POST", url, payload)
		if err != nil {
			return "", err
		}

		req.Header.Add("x-rapidapi-key", apiKey)
		req.Header.Add("x-rapidapi-host", "deep-translate1.p.rapidapi.com")
		req.Header.Add("Content-Type", "application/json")

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return "", err
		}

		defer res.Body.Close()
		body, _ := io.ReadAll(res.Body)
		data = string(body)
	}

	var response Response
	err := json.Unmarshal([]byte(data), &response)
	if err != nil {
		return "", err
	}

	// Return the array of language codes
	return response.Data.Translations.TranslatedText, nil
}

func main() {
	// Command line:
	// -listLanguages

	var listLanguages bool
	var showVersion bool
	var showHelp bool
	var sourceLanguage string
	var targetLanguage string

	//flag.String("", "", "Input text to translate")

	flag.BoolVarP(&listLanguages, "listLanguages", "l", false, "List all available languages.")
	flag.BoolVarP(&showVersion, "version", "v", false, "Show version information.")
	flag.BoolVarP(&showHelp, "help", "h", false, "Show this help.")
	flag.StringVarP(&sourceLanguage, "sourceLanguage", "s", "es", "Source language.")
	flag.StringVarP(&targetLanguage, "targetLanguage", "t", "en", "Target language.")

	flag.Parse()

	if spoofApiCalls {
		println("** Developer mode. Results are spoofed **")
	}

	if showHelp {
		flag.Usage()
		os.Exit(0)
	} else if showVersion {
		fmt.Println("motive-translator v" + versionString)
		os.Exit(0)
	}

	// Everything below this point requires an api key
	executable := os.Args[0]
	apiKeyFile := executable[:len(executable)-len(filepath.Ext(executable))] + ".key"
	apiKey, err := getApiKey(apiKeyFile)
	if err != nil {
		fmt.Println("API key unavailable:", err)
		os.Exit(2)
	}

	if listLanguages {
		languages, _ := getLanguages(apiKey)
		for _, language := range languages {
			fmt.Println(language.LanguageCode)
		}
	} else {
		if len(flag.Args()) < 2 {
			fmt.Println("Usage: go run main.go <text>")
			flag.Usage()
			os.Exit(1)
		}

		// Build the translation request from each word on the command line
		result, err := translate(strings.Join(flag.Args(), " "), sourceLanguage, targetLanguage, apiKey)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		println(result)
	}
}

func getApiKey(keyFile string) (string, error) {
	file, err := os.Open(keyFile)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var apiKey string
	scanner := bufio.NewScanner(file)

	// Search for "RAPIDAPI_KEY=xxx"
	for scanner.Scan() {
		line := scanner.Text()
		line = fmt.Sprint(line)

		if strings.HasPrefix(line, "RAPIDAPI_KEY=") {
			apiKey = strings.Split(line, "=")[1]
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	if apiKey == "" {
		return "", fmt.Errorf("keyfile does not contain an API key")
	}

	return apiKey, nil
}
