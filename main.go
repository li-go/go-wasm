package main

import (
	"fmt"
	"runtime"
	"syscall/js"
)

var (
	count int
	wait  = make(chan interface{})
)

func main() {
	// start
	fmt.Println("Hello, wasm!")

	// add event listener to dom
	onClickBodyCb := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		count += 1
		if count == 1 {
			this.Set("innerHTML", fmt.Sprintf("<h1>Clicked %d time</h1>", count))
		} else {
			this.Set("innerHTML", fmt.Sprintf("<h1>Clicked %d times</h1>", count))
		}
		return nil
	})
	defer onClickBodyCb.Release()

	js.Global().Get("document").
		Call("getElementsByTagName", "body").Index(0).
		Call("addEventListener", "click", onClickBodyCb)

	// expose function to JavaScript
	getCpuNumCb := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		fmt.Println("getCpuNum called!")
		return runtime.NumCPU()
	})
	defer getCpuNumCb.Release()

	js.Global().Set("getCpuNum", getCpuNumCb)

	// wait for callback's lifetime
	<-wait
}
