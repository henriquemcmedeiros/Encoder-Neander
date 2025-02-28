package assembler

import (
	"Encoder/lexer"
	"fmt"
	"log"
	"os"
	"strconv"
)

const TOTAL_SIZE = 516

// var Instructions = map[string]uint8{
// 	"NOP": 0x00, "STA": 0x10, "LDA": 0x20, "ADD": 0x30,
// 	"OR": 0x40, "AND": 0x50, "NOT": 0x60, "JMP": 0x80,
// 	"JN": 0x90, "JZ": 0xA0, "HLT": 0xF0,
// }

var InstructionsRequireNumber = map[string]bool{
	"STA": true, "LDA": true, "ADD": true, "OR": true, "AND": true, "JMP": true, "JN": true, "JZ": true,
}

func assemble(tokens []lexer.Token) []byte {
	memory := make([]byte, TOTAL_SIZE)

	// Bytes Iniciais Obrigatórios
	memory[0] = 0x03
	memory[1] = 0x4E
	memory[2] = 0x44
	memory[3] = 0x52

	pc := 4

	for i := 0; i < len(tokens); i++ {
		// Fazer validação do espaços .CODE/.DATA/.ORG
		// Fazer validação pra quando não puder fazer parse

		switch tokens[i].Tipo {
		case "INSTRUCTION":
			if opcode, exists := lexer.Instructions[tokens[i].Valor]; exists {
				memory[pc] = opcode
				pc += 2
			} else {
				log.Printf("Instrução desconhecida: %s", tokens[i].Valor)
			}
		case "NUMBER":
			val, err := strconv.ParseUint(tokens[i].Valor, 16, 8)
			if err == nil {
				memory[pc] = byte(val)
				pc += 2
			} else {
				log.Printf("Número inválido: %s", tokens[i].Valor)
			}

		case "VARIABLE":
			// Fazer a lógica
		}
	}

	return memory
}

func GetBinary(tokens []lexer.Token) {
	memory := assemble(tokens)

	err := os.WriteFile("output.mem", memory, 0644)
	if err != nil {
		log.Fatalf("Erro ao salvar arquivo .mem: %v", err)
	}

	fmt.Println("Arquivo output.mem gerado com sucesso!")
}
