package parser

import (
	"fmt"
	"Encoder/lexer"
	"strconv"
)

type Instrucao struct {
	Opcode  string
	Operand string
}

type Variavel struct {
	Nome  string
	Valor *int
	Tipo  string
}

type Parser struct {
	tokens     []lexer.Token
	pos        int
	instrucoes []Instrucao
	variaveis  map[string]Variavel
}

// NovoParser cria um novo parser
func NovoParser(tokens []lexer.Token) *Parser {
	return &Parser{
		tokens:     tokens,
		pos:        0,
		instrucoes: make([]Instrucao, 0),
		variaveis:  make(map[string]Variavel),
	}
}

func (p *Parser) ResolverSimbolos() error {
	tamanhoPrograma := len(p.instrucoes)

	enderecoVar := tamanhoPrograma
	for nome, variavel := range p.variaveis {
		if enderecoVar > 0xFF {
			return fmt.Errorf("memória insuficiente: endereço %d ultrapassa 0xFF", enderecoVar)
		}

		novaVariavel := Variavel{
			Nome:  nome,
			Valor: new(int),
			Tipo:  variavel.Tipo,
		}
		*novaVariavel.Valor = enderecoVar

		p.variaveis[nome] = novaVariavel

		enderecoVar++
	}

	for i, inst := range p.instrucoes {
		if variavel, ok := p.variaveis[inst.Operand]; ok {
			p.instrucoes[i].Operand = fmt.Sprintf("%02X", *variavel.Valor)
		}
	}

	return nil
}

func (p *Parser) Parse() ([]Instrucao, map[string]Variavel, error) {
	for p.pos < len(p.tokens) {
		token := p.tokens[p.pos]

		switch token.Tipo {
		case lexer.TOKEN_INSTR:
			err := p.parseInstrucao()
			if err != nil {
				return nil, nil, err
			}
		case lexer.TOKEN_VAR:
			err := p.parseVariavel()
			if err != nil {
				return nil, nil, err
			}
		case lexer.TOKEN_SECTION:
			err := p.parseSecao()
			if err != nil {
				return nil, nil, err
			}
		case lexer.TOKEN_EOF:
			return p.instrucoes, p.variaveis, nil
		default:
			return nil, nil, fmt.Errorf("token inesperado: %s", token.Valor)
		}

		p.pos++
	}

	return p.instrucoes, p.variaveis, nil
}

func (p *Parser) parseInstrucao() error {
	instrucao := Instrucao{Opcode: p.tokens[p.pos].Valor}

	// Avança para o próximo token (operando)
	p.pos++
	if p.pos < len(p.tokens) && (p.tokens[p.pos].Tipo == lexer.TOKEN_NUMBER || p.tokens[p.pos].Tipo == lexer.TOKEN_VAR) {
		instrucao.Operand = p.tokens[p.pos].Valor
	}

	p.instrucoes = append(p.instrucoes, instrucao)
	return nil
}

func (p *Parser) parseVariavel() error {
	// Nome da variável
	nomeVar := p.tokens[p.pos].Valor

	// Avança para o próximo token (DB ou DS)
	p.pos++
	if p.pos >= len(p.tokens) || p.tokens[p.pos].Tipo != lexer.TOKEN_DEFINE {
		return fmt.Errorf("esperado DB ou DS após nome da variável")
	}
	diretiva := p.tokens[p.pos].Valor

	// Processa DB ou DS
	if diretiva == "DB" {
		// Avança para o próximo token (valor)
		p.pos++
		if p.pos >= len(p.tokens) || p.tokens[p.pos].Tipo != lexer.TOKEN_NUMBER {
			return fmt.Errorf("esperado valor após DB")
		}

		// Converte o valor para inteiro (hexadecimal)
		valor, err := strconv.ParseInt(p.tokens[p.pos].Valor, 16, 64)
		if err != nil {
			return fmt.Errorf("valor inválido após DB: %s", p.tokens[p.pos].Valor)
		}

		// Armazena a variável com o valor
		p.variaveis[nomeVar] = Variavel{Nome: nomeVar, Valor: new(int), Tipo: "DB"}
		*p.variaveis[nomeVar].Valor = int(valor)
	} else if diretiva == "DS" {
		// DS não tem valor, apenas reserva espaço
		p.variaveis[nomeVar] = Variavel{Nome: nomeVar, Valor: nil, Tipo: "DS"}
	} else {
		return fmt.Errorf("diretiva desconhecida: %s", diretiva)
	}

	return nil
}

func (p *Parser) parseSecao() error {
	// Implemente a lógica para processar seções, se necessário
	return nil
}