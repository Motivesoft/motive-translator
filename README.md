# motive-translator
Simple command line translation service

Built, for now, using [Deep Translate](https://rapidapi.com/gatzuma/api/deep-translate1)

## Build instructions
```shell
go build
```

A `bool` flag is available that uses hard-coded test strings during development to reduce API calls. Set `spoofApiCalls` as required.

## Runtime requirements
A file containing the API key for the translation calls is required. This must be placed in the directory containing the executable, with the same filename but with the file extension `.key`. 

Example:
```
C:\Utilities\translate.exe
C:\Utilities\translate.key

/usr/local/bin/translate
/use/local/bin/translate.key
```

The key file requires an entry for a [RapidAPI](https://rapidapi.com/hub) key:
```properties
RAPIDAPI_KEY=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
```

This file is not committed into the repository and needs to be created manually and stored locally by whoever uses this.

## Usage
Run the executable with a combination of options and text to translate. Quotes around the text to translate are optional.

```shell
translate.exe [options] "text to translate"
./translate   [options] "text to translate"
```

Where `[options]` can be:
```shell
-h          show help
-v          show version
-l          list supported languages
-s          source language (default: es)
-t          target language (default: en)
```
