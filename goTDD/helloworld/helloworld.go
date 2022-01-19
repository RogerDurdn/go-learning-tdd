package main

import "fmt"

const englishHelloWord = "hello"
const spanishHelloWord = "hola"
const frenchHelloWord = "bonjour"
var languages = map[string]string{
 "default": englishHelloWord,
 "spanish": spanishHelloWord,
 "french": frenchHelloWord,
}

func main() {
 fmt.Println(hello())
 fmt.Println(helloRecipient("world"))
 fmt.Println(helloDefault(""))
 fmt.Println(helloDefault("roger"))
}

func hello() string {
 return englishHelloWord + " world"
}

func helloRecipient(recipient string) string {
 return englishHelloWord +" "+ recipient
}

func helloDefault(recipient string) string {
 if recipient == "" {
  recipient = "world"
 }
 return englishHelloWord +" "+ recipient
}

func helloLanguage(recipient, language string) string {
 switch recipient {
 case "":
  recipient = "world"
 }
 greeting, ok := languages[language]
 if !ok {
  greeting = languages["default"]
 }
 return greeting + " " + recipient
}
