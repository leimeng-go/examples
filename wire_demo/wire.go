//+build wireinject

package main

import(
	"fmt"
	"github.com/google/wire"
)

//Message A message for a greeter
type Message string

//Greeter A greeter who converys that message
type Greeter struct{
   Message Message
}

//Greet a greet method
func (g Greeter)Greet()Message{
	return g.Message
}

//Event an event that starts with the greeter greeting guests.
type Event struct{
   Greeter Greeter
}

//NewMessage a simple initializer that always returns a hard-codeed message
func NewMessage() Message{
	return Message("Hi there!")
}

//NewGreeter crate a g
func NewGreeter(m Message)Greeter{
      return Greeter{Message: m}
}

//NewEvent crate event instance
func NewEvent(g Greeter)Event{
	return Event{Greeter: g}
}

//Start a
func (e Event)Start(){
   msg:=e.Greeter.Greet()
   fmt.Println(msg)
}

//go:generate echo GoGoGo!
//go:generate go run main.go
//go:generate echo $GOARCH $GOOS $GOFILE $GOLINE $GOPACKAGE

//InitializeEvent a
func InitializeEvent()Event{
   wire.Build(NewEvent,NewGreeter,NewMessage)
   return Event{}
}