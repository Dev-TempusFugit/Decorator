package main

import "github.com/Dev-TempusFugit/Decorator/functions"

func execute(name string, f func(string)){
	f(name)
}


func main(){
	name := "Maxim"

	execute(name, functions.Saludar)

}