package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	TOKEN_SECTION  = "SECTION"
	TOKEN_EOF      = "EOF"
	TOKEN_INSTR    = "INSTRUCTION"
	TOKEN_NUMBER   = "NUMBER"
	TOKEN_VAR      = "VARIABLE"
	TOKEN_DEFINE   = "DEFINE"
	TOKEN_UNKNOWN  = "UNKNOWN"
)

var (
	Instructions = map[string]uint8{
		"NOP": 0x00, "STA": 0x10, "LDA": 0x20, "ADD": 0x30,
		"OR": 0x40, "AND": 0x50, "NOT": 0x60, "JMP": 0x80,
		"JN": 0x90, "JZ": 0xA0, "HLT": 0xF0,
	}

	Define = map[string]bool{
		"DB": true, "DS": false,
	}
)

type Token struct {
	tipo  string
	valor string
}

func isInstruction(lexema string) bool {
	_, existe := Instructions[lexema]
	return existe
}

func isDefine(lexema string) bool {
	_, existe := Define[lexema]
	return existe
}

func isNumber(lexema string) bool {
	if _, err := strconv.ParseInt(lexema, 16, 64); err == nil {
		return true
	}
	return false
}

func isVariable(lexema string) bool {
	return regexp.MustCompile(`^[A-Za-z_][A-Za-z0-9_]*$`).MatchString(lexema)
}

func lexer(lexema string) Token {
	switch {
	case strings.HasPrefix(lexema, "."):
		return Token{tipo: TOKEN_SECTION, valor: strings.TrimPrefix(lexema, ".")}
	case isInstruction(lexema):
		return Token{tipo: TOKEN_INSTR, valor: lexema}
	case isDefine(lexema):
		return Token{tipo: TOKEN_DEFINE, valor: lexema}
	case isNumber(lexema):
		return Token{tipo: TOKEN_NUMBER, valor: lexema}
	case isVariable(lexema):
		return Token{tipo: TOKEN_VAR, valor: lexema}
	default:
		log.Printf("Token desconhecido: %s", lexema)
		return Token{tipo: TOKEN_UNKNOWN, valor: lexema}
	}
}

func main() {
	arquivo, err := os.ReadFile("assembly_file.txt")
	if err != nil {
		log.Fatalf("Não foi possível ler o arquivo: %v", err)
	}

	re := regexp.MustCompile(`\s+`)
	lexemas := re.Split(string(arquivo), -1)

	var tokens []Token
	for _, lexema := range lexemas {
		lexema = strings.TrimSpace(lexema)
		if lexema != "" {
			tokens = append(tokens, lexer(lexema))
		}
	}

	tokens = append(tokens, Token{tipo: TOKEN_EOF, valor: ""})
	for _, token := range tokens {
		fmt.Printf("Token: %+v\n", token)
	}
}
