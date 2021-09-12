package staff

import "log"

var OverPaidLimit = 75000
var UnderPaidLimit = 20000

//upper case!!!!
type Employee struct {
	FirstName string
	LastName  string
	Salary    int
	FullTime  bool
}

type Office struct {
	AllStaff []Employee
}

func (e *Office) All() []Employee {
	return e.AllStaff
}

func (e *Office) Overpaid() []Employee {
	var overpaid []Employee

	for _, x := range e.AllStaff {
		if x.Salary > OverPaidLimit {
			overpaid = append(overpaid, x)
		}
	}
	return overpaid
}

func (e *Office) Underpaid() []Employee {
	var underpaid []Employee
	myFunction()

	for _, x := range e.AllStaff {
		if x.Salary < UnderPaidLimit {
			underpaid = append(underpaid, x)
		}
	}
	return underpaid
}

func (e *Office) notvisible() {
	log.Println("hello, World")
}

//package level function.. this has no access to reciever variable e
func myFunction() {
	log.Println("I am a function")
}
