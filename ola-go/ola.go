package main

import "fmt"

const prefixHey = "Hey "

func Ola(name string) string {
	if name == ""{
		name = "you"
	}
	
	return prefixHey + name
}
func main(){
	fmt.Println(Ola("world"))

}
