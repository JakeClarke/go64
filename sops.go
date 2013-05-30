package main

var SpecOpTable [0x3F]Operation

func init() {
	SpecOpTable[0x00] = specOpSll
	SpecOpTable[0x02] = specOpSrl
	SpecOpTable[0x08] = specOpJR

	SpecOpTable[0x20] = specOpAdd
	SpecOpTable[0x21] = specOpAddU

	SpecOpTable[0x25] = specOpOr

	SpecOpTable[0x27] = specOpNor

	SpecOpTable[0x2A] = specOpSlt
	SpecOpTable[0x2B] = specOpSltU
}

func specOpAdd(i Instruction, cpu *CPU) {
	cpu.GPR[i.RD()] = DWORD(int64(cpu.GPR[i.RS()]) + int64(cpu.GPR[i.RT()]))
}

func specOpAddU(i Instruction, cpu *CPU) {
	cpu.GPR[i.RD()] = cpu.GPR[i.RS()] + cpu.GPR[i.RT()]
}

func specOpJR(i Instruction, cpu *CPU) {
	cpu.PC = cpu.GPR[i.RS()]
}

func specOpNor(i Instruction, cpu *CPU) {
	cpu.GPR[i.RD()] = ^(cpu.GPR[i.RS()] | cpu.GPR[i.RT()])
}

func specOpOr(i Instruction, cpu *CPU) {
	cpu.GPR[i.RD()] = (cpu.GPR[i.RS()] | cpu.GPR[i.RT()])
}

func specOpSlt(i Instruction, cpu *CPU) {
	if int64(cpu.GPR[i.RS()]) < int64(cpu.GPR[i.RT()]) {
		cpu.GPR[i.RD()] = 0x1
	} else {
		cpu.GPR[i.RD()] = 0x0
	}
}

func specOpSltU(i Instruction, cpu *CPU) {
	if cpu.GPR[i.RS()] < cpu.GPR[i.RT()] {
		cpu.GPR[i.RD()] = 0x1
	} else {
		cpu.GPR[i.RD()] = 0x0
	}
}

func specOpSll(i Instruction, cpu *CPU) {
	cpu.GPR[i.RD()] = cpu.GPR[i.RT()] << i.Shamt()
}

func specOpSrl(i Instruction, cpu *CPU) {
	cpu.GPR[i.RD()] = cpu.GPR[i.RT()] >> i.Shamt()
}
