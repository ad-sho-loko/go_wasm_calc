package main

import(
	"syscall/js"
)

func manipulateDom(i []js.Value){
	js.Global().Get("document").Call("getElementById", "result").Set("innerHTML", "WebAssembly!")
}

func registerCallbacks() {
	js.Global().Set("manipulateDom", js.NewCallback(manipulateDom))
}

func main(){
	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c
}