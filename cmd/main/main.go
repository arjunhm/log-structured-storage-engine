package main

import (
	"fmt"
	"lsm"
)

func main() {
	kv := lsm.NewKeyValue("name", "jun", false)
	fmt.Println(kv.Key, kv.Value);
	kv.Display();
}
