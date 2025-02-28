package main

import (
	"Encoder/assembler"
	"Encoder/encoder"
	"Encoder/lexer"
	// "fmt"
)

func main() {
	tokens := lexer.GetTokens("../assembly_file.txt")

	// for _, token := range tokens {
	// 	fmt.Printf("Token: %+v\n", token)
	// }

	assembler.GetBinary(tokens)

	encoder.RunBinary("output.mem")
}
