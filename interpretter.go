package main

import "log"

type Operation func(Instruction, *CPU)

var OpTable [0x3F]Operation

const SPEC_ADD = 0x20
const SPEC_ADDU = 0x20

func init() {
	OpTable[0] = opSpecial
	OpTable[0x2] = opJump
	OpTable[0x3] = opJal
	OpTable[0x4] = opBeq
	OpTable[0x5] = opBne
	OpTable[0x8] = opAddi
	OpTable[0x9] = opAddiU
	OpTable[0xF] = opLoadUpperI

	OpTable[0xc] = opAndi

	OpTable[0x23] = opLoadWord
	OpTable[0x24] = opLoadByteU
	OpTable[0x25] = opLoadHalfWord

	OpTable[0x2B] = opStoreWord
	OpTable[0x29] = opStoreHalfWord
	OpTable[0x30] = opLoadLinked
}

func opSpecial(i Instruction, cpu *CPU) {
	funct := i.Funct()
	if SpecOpTable[funct] == nil {
		log.Panicf("Unrecongised special opcode! (code: %v)", funct)
	} else {
		SpecOpTable[funct](i, cpu)
	}
}

func opAddi(i Instruction, cpu *CPU) {
	cpu.GPR[i.RD()] = DWORD(int64(cpu.GPR[i.RS()]) + int64(i.Immediate()))
}

func opAddiU(i Instruction, cpu *CPU) {
	cpu.GPR[i.RD()] = cpu.GPR[i.RS()] + DWORD(i.Immediate())
}

func opAndi(i Instruction, cpu *CPU) {
	cpu.GPR[i.RD()] = cpu.GPR[i.RS()] & DWORD(i.Immediate())
}

func opJump(i Instruction, cpu *CPU) {
	cpu.PC = jumpAddr(i, cpu)
}

func opJal(i Instruction, cpu *CPU) {
	cpu.GPR[31] = cpu.PC + 8
	cpu.PC = jumpAddr(i, cpu)
}

func opBeq(i Instruction, cpu *CPU) {
	if cpu.GPR[i.RS()] == cpu.GPR[i.RT()] {
		cpu.PC = branchAddr(i, cpu)
	}
}

func opBne(i Instruction, cpu *CPU) {
	if cpu.GPR[i.RS()] != cpu.GPR[i.RT()] {
		cpu.PC = branchAddr(i, cpu)
	}
}

func opLoadWord(i Instruction, cpu *CPU) {
	cpu.GPR[i.RT()] = DWORD(cpu.Mem.LoadWord(cpu.GPR[i.RT()] + DWORD(i.Immediate())))
}

func opLoadHalfWord(i Instruction, cpu *CPU) {
	cpu.GPR[i.RT()] = DWORD(cpu.Mem.LoadHalfWord(cpu.GPR[i.RS()] + DWORD(i.Immediate())))
}

func opStoreWord(i Instruction, cpu *CPU) {
	cpu.Mem.StoreWord(cpu.GPR[i.RS()]+DWORD(i.Immediate()), WORD(cpu.GPR[i.RT()]))
}

func opStoreHalfWord(i Instruction, cpu *CPU) {
	cpu.Mem.StoreHalfWord(cpu.GPR[i.RS()]+DWORD(i.Immediate()), HWORD(cpu.GPR[i.RT()]))
}

func opLoadByteU(i Instruction, cpu *CPU) {
	cpu.GPR[i.RT()] = DWORD(cpu.Mem.LoadByteU(cpu.GPR[i.RS()] + DWORD(i.Immediate()&0xFF)))
}

func opLoadLinked(i Instruction, cpu *CPU) {
	log.Panic("Load linked not implemented")
}

func opLoadUpperI(i Instruction, cpu *CPU) {
	cpu.GPR[i.RT()] = DWORD(i.Immediate()) << 16
}

func opOrI(i Instruction, cpu *CPU) {
	cpu.GPR[i.RT()] = cpu.GPR[i.RS()] | DWORD(zeroExtImm(i))
}

// util ops

func jumpAddr(i Instruction, cpu *CPU) DWORD {
	return (cpu.PC & 0xF0000000) + (DWORD(i.Addr()) << 2)
}

func branchAddr(i Instruction, cpu *CPU) DWORD {
	return cpu.PC + (DWORD(i.Addr()) << 2) + 4
}

func zeroExtImm(i Instruction) HWORD {
	return i.Immediate() & 0x7FFF
}
