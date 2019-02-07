package main

import(
	"strconv"
	"syscall/js"
)

// 現在入力中の数字
var inputtingNum = ""

// 保存している値
var accumulator = ""

func doPlus(i []js.Value){
	accumulator += inputtingNum
	inputtingNum = ""
	js.Global().Get("document").Call("getElementById", "result").Set("innerHTML", "0")
}

func doEqual(i []js.Value){
	// String -> int
	int1, _ := strconv.Atoi(inputtingNum)
	int2, _ := strconv.Atoi(accumulator)

	accumulator = ""

	js.Global().Get("document").Call("getElementById", "result").Set("innerHTML", int1 + int2)
}

func inputNum(i []js.Value){
	inputtingNum += i[0].String()
	js.Global().Get("document").Call("getElementById", "result").Set("innerHTML", inputtingNum)
}

func registerCallbacks() {
	js.Global().Set("inputNum", js.NewCallback(inputNum))
	js.Global().Set("doPlus", js.NewCallback(doPlus))
	js.Global().Set("doEqual", js.NewCallback(doEqual))
}

func main(){
	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c
}