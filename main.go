package main

import (
	"log"
	"time"

	"github.com/bytecodealliance/wasmtime-go"
)

func main() {
	bs, err := wasmtime.Wat2Wasm(`
	  (func (export "loop")
	    (loop br 0))
	`)
	check(err)

	for i := 0; ; i++ {
		config := wasmtime.NewConfig()
		config.SetInterruptable(true)
		store := wasmtime.NewStore(wasmtime.NewEngineWithConfig(config))

		module, err := wasmtime.NewModule(store.Engine, bs)
		check(err)

		instance, err := wasmtime.NewInstance(store, module, nil)
		check(err)

		handle, err := store.InterruptHandle()
		check(err)

		go func() {
			time.Sleep(100 * time.Millisecond)
			handle.Interrupt()
		}()
		_, err = instance.GetFunc(store, "loop").Call(store)
		if err == nil {
			panic("expected error")
		}
		if t, ok := err.(*wasmtime.Trap); ok {
			log.Printf("%d: trap: %s", i, t.Message())
		}
	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
