package main

import (
	"fmt"
	"miiky976/Godis/kv"
)

func main() {
	kv.Test()
	fmt.Println(kv.KV)
	fmt.Println(kv.GET("name"))
	kv.SET("name", "Jose")
	kv.SET("age", 12)
	fmt.Println(kv.KV)
	kv.DEL("name")
	kv.INCR("age", 5)
	fmt.Println(kv.KV)
	kv.DEL("age")
	fmt.Println(kv.KV)
}
