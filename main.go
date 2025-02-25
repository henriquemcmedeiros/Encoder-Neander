package main

import (
	"fmt"
	"os"
)

const (
	TOTAL_SIZE = 516

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

func flagZero(AC int) bool {
	return AC == 0x00
}

func flagNeg(AC int) bool {
	return AC&0x80 != 0
}

func main() {
	AC := 0
	PC := 0x04

	file, err := os.Open("/Users/hmmedeiros/CLionProjects/untitled/TESTEGERAL.mem")
	if err != nil {
		fmt.Println("Não foi possível ler o arquivo!")
		return
	}
	defer file.Close()

	memory := make([]byte, TOTAL_SIZE)
	_, err = file.Read(memory)
	if err != nil {
		fmt.Println("Erro ao ler o arquivo!")
		return
	}

	memory[0x80] = 0x05
	memory[0x81] = 0x03
	memory[0x83] = 0x02
	memory[0x84] = 0x01
	memory[0x86] = 0x06
	memory[0x87] = 0x03
	memory[0x89] = 0xFC
	memory[0x8A] = 0x03

	posicao := 0

	for memory[PC] != HLT && PC <= 0xFF {
		fmt.Printf("AC: %2x PC: %2x FZ: %5t FN: %5t INSTRUCAO: %2x CONTEUDO: %2x\n", AC & 0xFF, PC, flagZero(AC), flagNeg(AC), memory[PC], memory[PC+2])

		switch memory[PC] {
			case STA:
				PC += 2
				posicao = int(memory[PC])
				memory[posicao] = byte(AC)
				PC += 2
			case LDA:
				PC += 2
				posicao = int(memory[PC])
				AC = int(memory[posicao])
				PC += 2
			case ADD:
				PC += 2
				posicao = int(memory[PC])
				AC += int(memory[posicao])
				PC += 2
			case OR:
				PC += 2
				posicao = int(memory[PC])
				AC |= int(memory[posicao])
				PC += 2
			case AND:
				PC += 2
				posicao = int(memory[PC])
				AC &= int(memory[posicao])
				PC += 2
			case NOT:
				AC = ^AC
				PC += 2
			case JMP:
				PC += 2
				PC = int(memory[PC]) * 2 + 4
			case JN:
				PC += 2
				if flagNeg(AC) {
					PC = int(memory[PC]) * 2 + 4
				} else {
					PC += 2
				}
			case JZ:
				PC += 2
				if flagZero(AC) {
					PC = int(memory[PC]) * 2 + 4
				} else {
					PC += 2
				}
			default:
				PC += 2
		}
	}

	fmt.Println("========== Retorno de Memória ===========")
	for i := 0; i < TOTAL_SIZE; i++ {
		fmt.Printf("%3x:%3x ", i, memory[i])
		if i%16 == 15 {
			fmt.Println()
		}
	}
}
