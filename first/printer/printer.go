package printer

import (
	"fmt"
	"runtime"
)

// Hello is an exported function
func Hello() {
	fmt.Println("Exported Hello")
}

// Version shows go runtime version
func Version() string {
	return runtime.Version()
}
