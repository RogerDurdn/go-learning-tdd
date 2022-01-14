package main

import "fmt"

type Employee struct {
 Name string
 ID   int
}

func (e Employee) Describe() string {
 return fmt.Sprintf("%s (%d)", e.Name, e.ID)
}

type Manager struct {
 Employee
 Collaborators []Employee
}

func (m *Manager) AddEmployee(e Employee) {
 m.Collaborators = append(m.Collaborators, e)

}

func (m Manager) IsManagerOf(employee Employee) bool {
 for _, e := range m.Collaborators {
  if e == employee {
   return true
  }
 }
 return false
}

func main() {
 roger := Employee{
  "roger", 123,
 }

 man := Manager{
  Employee: Employee{
   "The manager", 22,
  },
 }

 fmt.Println(man.IsManagerOf(roger))
 man.AddEmployee(roger)
 fmt.Println(man.IsManagerOf(roger))

}
