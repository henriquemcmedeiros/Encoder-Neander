package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

const (
	TOKEN_SECTION string = "SECTION"
	TOKEN_EOF string = "EOF"
	TOKEN_INSTR string = "INSTRUCTION"
	TOKEN_NUMBER string = "NUMBER"
	TOKEN_VAR string = "VARIABLE"
	TOKEN_DEFINE string = "DEFINE"

	NOP = 0x00
	STA = 0x10
	LDA = 0x20
	ADD = 0x30
	OR  = 0x40
	AND = 0x50
	NOT = 0x60
	JMP = 0x80
	JN  = 0x90
	JZ  = 0xA0
	HLT = 0xF0
)

var (
	Instructions = map[string]int{
		"NOP": NOP, 
		"STA": STA, 
		"LDA": LDA, 
		"ADD": ADD, 
		"OR": OR, 
		"AND": AND,
		"NOT": NOT, 
		"JMP": JMP,
		"JN": JN, 
		"JZ": JZ, 
		"HLT": HLT,
	}

	Define = map[string]bool{
		"DB": true,
		"DS": false,
	}
)

type Token struct {
	tipo string
	valor string
}

		
func lexer(lexema string, posicao int, tamanho int) Token{
			
	if posicao >= tamanho {
		return Token{tipo: TOKEN_EOF, valor: ""}
	}
	
	if strings.HasPrefix(lexema, ".") {
		return Token{tipo: TOKEN_SECTION, valor: strings.Replace(lexema, ".", "", 1)}
	}
	
	if _, existe := Instructions[lexema]; existe{
		return Token{tipo: TOKEN_INSTR, valor: lexema}
	}

	if _, existe := Define[lexema]; existe{
		return Token{tipo: TOKEN_DEFINE, valor: lexema}
	}

	if _, err := fmt.Sscanf(lexema, "%d", new(int)); err == nil {
		return Token{tipo: TOKEN_NUMBER, valor: lexema}
	}
		
	return Token{tipo: TOKEN_VAR, valor: lexema}
}
		
func main() {
			
	arquivo, err := os.ReadFile("assembly_file.txt")
			
	if err != nil {
		log.Fatalf("Não foi possível ler o arquivo!")
		return
	}
			
	re := regexp.MustCompile(`\s+`)
	lexemas := re.Split(string(arquivo), - 1)
			
	var tokens []Token
			
	for index, lexema := range lexemas {
		tokens = append(tokens, lexer(lexema, index + 1, len(lexemas)))
	}
			
	for _, token := range tokens {
		fmt.Printf("Token: %+v\n", token)
	}
		
}
		