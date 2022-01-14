package main

import (
 "errors"
 "fmt"
 "net/http"
)

func LogOutput(message string) {
 fmt.Println(message)
}


func (l LoggerAdapter) Log(message string) {
 l(message)
}

type LoggerAdapter func(message string)

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
 return SimpleDataStorage{
  userData: map[string]string{
   "1": "juan",
   "2": "ramona",
   "3": "alfred",
  },
 }
}

func (s SimpleDataStorage) UserNameForID(id string) (string, bool) {
 name, ok := s.userData[id]
 return name, ok
}

type SimpleLogic struct {
 l  Logger
 ds DataStorage
}

func NewSimpleLogic(l Logger, ds DataStorage) SimpleLogic {
 return SimpleLogic{
  l:  l,
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
type Controller struct {
 l     Logger
 logic Logic
}

func NewController(l Logger, logic Logic) Controller {
 return Controller{
  l:     l,
  logic: logic,
 }
}

func (c Controller) SayHello(w http.ResponseWriter, r *http.Request) {
 c.l.Log("In SayHello")
 userID := r.URL.Query().Get("user_id")
 message, err := c.logic.SayHello(userID)
 if err != nil {
  w.WriteHeader(http.StatusBadRequest)
  w.Write([]byte(err.Error()))
  return
 }
 w.Write([]byte(message))
}

func main() {
  l := LoggerAdapter(LogOutput)
    ds := NewSimpleDataStorage()
    logic := NewSimpleLogic(l, ds)
    c := NewController(l, logic)
    http.HandleFunc("/hello", c.SayHello)
    http.ListenAndServe(":8080", nil)
}
