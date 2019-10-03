package main

import (
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

func doOperate(this js.Value, i []js.Value) interface{} {
	accumulator += inputtingNum
	inputtingNum = ""
	js.Global().Get("document").Call("getElementById", "result").Set("textContent", "0")

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
	return nil
}

func doEqual(this js.Value, i []js.Value) interface{} {
	// String -> int
	int1, _ := strconv.Atoi(inputtingNum)
	int2, _ := strconv.Atoi(accumulator)

	accumulator = ""

	// 演算子ごとに計算する
	var result int
	switch operator {
	case Plus:
		result = int1 + int2
	case Sub:
		result = int2 - int1
	case Mul:
		result = int1 * int2
	case Div:
		result = int2 / int1
	}

	js.Global().Get("document").Call("getElementById", "result").Set("textContent", result)
	return nil
}

func inputNum(this js.Value, i []js.Value) interface{} {
	inputtingNum += i[0].String()
	js.Global().Get("document").Call("getElementById", "result").Set("textContent", inputtingNum)
	return nil
}

func clearNum(this js.Value, i []js.Value) interface{} {
	inputtingNum = ""
	accumulator = ""
	js.Global().Get("document").Call("getElementById", "result").Set("textContent", "0")
	return nil
}

func registerCallbacks() {
	js.Global().Set("inputNum", js.FuncOf(inputNum))
	js.Global().Set("doOperate", js.FuncOf(doOperate))
	js.Global().Set("doEqual", js.FuncOf(doEqual))
	js.Global().Set("clearNum", js.FuncOf(clearNum))
}

func main() {
	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c
}
