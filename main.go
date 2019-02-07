package main

import(
	"strconv"
	"syscall/js"
)

// 現在入力中の数字
var inputtingNum = ""

// 保存している値
var accumulator = ""

// 選択中の演算子
var operator = None

type Operator int

const (
	Plus Operator = iota
	Sub
	Mul
	Div
	None
)

func doOperate(i []js.Value){
	accumulator += inputtingNum
	inputtingNum = ""
	js.Global().Get("document").Call("getElementById", "result").Set("innerHTML", "0")

	// 演算子ごとに処理する
	switch i[0].String() {
	case "+":
		operator = Plus
	case "-":
		operator = Sub
	case "*":
		operator = Mul
	case "/":
		operator = Div
	default:
		operator = None
	}
}

func doEqual(i []js.Value){
	// String -> int
	int1, _ := strconv.Atoi(inputtingNum)
	int2, _ := strconv.Atoi(accumulator)

	accumulator = ""

	// 演算子ごとに計算する
	var result int
	switch operator{
	case Plus:
		result = int1 + int2
	case Sub:
		result = int2 - int1
	case Mul:
		result = int1 * int2
	case Div:
		result = int2 / int1
	}

	js.Global().Get("document").Call("getElementById", "result").Set("innerHTML", result)
}

func inputNum(i []js.Value){
	inputtingNum += i[0].String()
	js.Global().Get("document").Call("getElementById", "result").Set("innerHTML", inputtingNum)
}

func clearNum(i []js.Value){
	inputtingNum = ""
	accumulator = ""
	js.Global().Get("document").Call("getElementById", "result").Set("innerHTML", "0")
}

func registerCallbacks() {
	js.Global().Set("inputNum", js.NewCallback(inputNum))
	js.Global().Set("doOperate", js.NewCallback(doOperate))
	js.Global().Set("doEqual", js.NewCallback(doEqual))
	js.Global().Set("clearNum", js.NewCallback(clearNum))
}

func main(){
	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c
}