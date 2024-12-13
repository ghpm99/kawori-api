package runtime

import (
	"fmt"
	"runtime"
)

func ConfigRuntime() {
	numeroCpu := runtime.NumCPU()
	runtime.GOMAXPROCS(numeroCpu)
	fmt.Printf("Running with %d CPUs\n", numeroCpu)
}
