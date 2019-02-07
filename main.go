package main

import(
	"fmt"
	"syscall/js"
)

func print(i []js.Value){
	fmt.Println(i)
}

func registerCallbacks() {
	js.Global().Set("print", js.NewCallback(print))
}

func main(){
	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c
}