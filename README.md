# motive-translator
Simple command line translation service

Built, for now, using [Deep Translate](https://rapidapi.com/gatzuma/api/deep-translate1)

## Build instructions
```shell
go build
```

A `bool` flag is available that uses hard-coded test strings during development to reduce API calls. Set `spoofApiCalls` as required.

## Runtime requirements
A file called `api.key` in the current directory containing the API key for the translation calls. This file is not in the repository and needs to be created manually by whoever uses this.

## Usage
Run the executable with a combination of options and text to translate.
```shell
translate.exe [options] text to translate
./translate   [options] text to translate
```

Where `[options]` can be:
```shell
-h          show help
-v          show version
-l          list supported languages
-s          source language (default: es)
-t          target language (default: en)
```
