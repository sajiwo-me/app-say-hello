package main


import (
	"fmt"
	say "github.com/sajiwo-me/module-say-hello"
	gsreetings "github.com/sajiwo-me/module-greetings"
)


func main(){
	fmt.Println(say.SayHello("bagus"))
	fmt.Println(gsreetings.GoodMorning(" sajiwo"))
}