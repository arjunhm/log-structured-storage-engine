package main

import (
	"fmt"
	"lsm"
)

func main() {
	kv := lsm.NewKV("name", "jun", false)
	fmt.Println(kv.Key, kv.Value);
	kv.Display();
}
