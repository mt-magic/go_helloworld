package main

// import "fmt"
import (
	"fmt"

	calculator "github.com/mt-magic/go_calculator"
	"rsc.io/quote"
)

func main() {
    // fmt.Println("Hello World!")

    // update
    total := calculator.Sum(30, 55)
    fmt.Println(total)
    fmt.Println("Version: ", calculator.Version)

    // update add
    fmt.Println(quote.Hello())

    // Error
    // total := go_calculator.internalSum(5)
    // fmt.Println(total)
    // fmt.Println("Version: ", go_calculator.logMessage)
}
