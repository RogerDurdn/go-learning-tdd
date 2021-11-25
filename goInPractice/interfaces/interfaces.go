package main

import "log"

type LogicProviderA struct {
}


type LogicProviderB struct {
}

func (l LogicProviderA) Process(data string) string  {
 return data + " from providerA"
}

func (l LogicProviderB) Process(data string) string  {
 return data + " from providerB"
}

type Logic interface {
 Process(data string) string 
}

type Client struct {
 L Logic
}

func (c *Client) SetLogic(l Logic)  {
 c.L = l
}

func (c Client) Program()  {
 log.Println(c.L.Process("client says:"))
}


func main() {

 c := Client{}
 c.SetLogic(LogicProviderA{})
 c.Program()
 c.SetLogic(LogicProviderB{})
 c.Program()
}
