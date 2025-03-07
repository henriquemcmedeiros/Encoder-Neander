package main

import (
	"fmt"
	"Encoder/lexer"
	"Encoder/parser"
)

func main() {
	// Lexer
	tokens := lexer.GetTokens("assembly_file.txt")

	// Parser
	parser := parser.NovoParser(tokens)
	instrucoes, variaveis, err := parser.Parse()
	if err != nil {
		fmt.Println("Erro ao fazer parsing:", err)
		return
	}

	err = parser.ResolverSimbolos()
	if err != nil {
		fmt.Println("Erro ao resolver símbolos:", err)
		return
	}

	fmt.Println("Instruções após resolução de símbolos:")
	for _, inst := range instrucoes {
		fmt.Printf("%s %s\n", inst.Opcode, inst.Operand)
	}

	fmt.Println("Variáveis:")
	for nome, variavel := range variaveis {
		if variavel.Tipo == "DB" {
			fmt.Printf("%s: DB %02X (endereço: %02X)\n", nome, *variavel.Valor, *variavel.Valor)
		} else if variavel.Tipo == "DS" {
			fmt.Printf("%s: DS (endereço: %02X)\n", nome, *variavel.Valor)
		}
	}
}