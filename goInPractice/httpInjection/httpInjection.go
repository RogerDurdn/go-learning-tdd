package main

import (
 "errors"
 "fmt"
)

func LogOutput(message string) {
 fmt.Println(message)
}

type LoggerAdapter func(message string)

func (l LoggerAdapter) Log(message string)  {
 l.Log(message)
}

type Logger interface {
 Log(message string)
}

type SimpleDataStorage struct {
 userData map[string]string
}

type DataStorage interface {
 UserNameForID(id string) (string, bool)
}

func NewSimpleDataStorage() SimpleDataStorage {
 return SimpleDataStorage{userData: map[string]string{
  "1": "juan",
  "2": "ramona",
  "3": "alfred",
 }}
}

func (s SimpleDataStorage) UsernameForID(id string) (string, bool) {
 name, ok := s.userData[id]
 return name, ok
}

type SimpleLogic struct {
 l Logger
 ds DataStorage
}
func NewSimpleLogic(l Logger, ds DataStorage) SimpleLogic {
    return SimpleLogic{
        l:    l,
        ds: ds,
    }
}

func (sl SimpleLogic) SayHello(userID string) (string, error) {
    sl.l.Log("in SayHello for " + userID)
    name, ok := sl.ds.UserNameForID(userID)
    if !ok {
        return "", errors.New("unknown user")
    }
    return "Hello, " + name, nil
}

func (sl SimpleLogic) SayGoodbye(userID string) (string, error) {
    sl.l.Log("in SayGoodbye for " + userID)
    name, ok := sl.ds.UserNameForID(userID)
    if !ok {
        return "", errors.New("unknown user")
    }
    return "Goodbye, " + name, nil
}

type Logic interface {
    SayHello(userID string) (string, error)
}

func main() {

}
