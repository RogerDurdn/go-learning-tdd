package main

import "fmt"

const englishHelloWord = "hello"

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
