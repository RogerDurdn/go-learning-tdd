package main

import "fmt"

func main() {
 fmt.Println(hello())
}

func hello() string {
 return "hello world"
}

func helloRecipient(recipient string) string {
 return "hello " + recipient
}
