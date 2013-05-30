package main

import "log"

type HWORD uint16
type WORD uint32
type DWORD uint64

const GP = 28
const SP = 29

type CPU struct {
	GPR, FP_GPR, CP0   []DWORD
	Mem                *Memory
	PC, MultHI, MultLO DWORD
	nextInstr          Instruction
}

func NewCPU() *CPU {
	cpu := new(CPU)
	cpu.GPR = make([]DWORD, 32)
	cpu.FP_GPR = make([]DWORD, 32)
	cpu.CP0 = make([]DWORD, 33)
	cpu.Mem = new(Memory)

	cpu.ResetCPU()
	return cpu
}

func (cpu *CPU) Tick() {
	// reset gpr[0] because it should be zero.
	cpu.GPR[0] = 0
	cpu.Execute(cpu.nextInstr)
	cpu.PC += 4
	cpu.Fetch()
}

func (cpu *CPU) Execute(i Instruction) {
	op := i.OP()
	log.Printf("Executing: %v", i)

	if OpTable[op] == nil {
		log.Panicf("Unrecongised opcode! (code: %v)", op)
	} else {
		OpTable[op](i, cpu)
	}
}

func (cpu *CPU) Fetch() {
	cpu.nextInstr = Instruction(cpu.Mem.LoadWord(cpu.PC))
}

func (cpu *CPU) ResetCPU() {
	cpu.PC = 0x0

	for i, _ := range cpu.GPR {
		cpu.GPR[i] = 0
	}

	cpu.GPR[GP] = 0x10008000
	cpu.GPR[SP] = 0x7ffffffc
	cpu.Mem.Init()
	cpu.Fetch()
}

var GPR_NAMES = []string{"r0", "at", "v0", "v1", "a0", "a1", "a2", "a3", "t0", "t1", "t2", "t3", "t4", "t5", "t6", "t7", "s0", "s1", "s2", "s3", "s4", "s5", "s6", "s7", "t8", "t9", "k0", "k1", "gp", "sp", "s8", "ra"}
var FPR_Name = []string{"f0", "f1", "f2", "f3", "f4", "f5", "f6", "f7", "f8", "f9", "f10", "f11", "f12", "f13", "f14", "f15", "f16", "f17", "f18", "f19", "f20", "f21", "f22", "f23", "f24", "f25", "f26", "f27", "f28", "f29", "f30", "f31"}
var Cop0_Name = []string{"Index", "Random", "EntryLo0", "EntryLo1", "Context", "PageMask", "Wired", "", "BadVAddr", "Count", "EntryHi", "Compare", "Status", "Cause", "EPC", "PRId", "Config", "LLAddr", "WatchLo", "WatchHi", "XContext", "", "", "", "", "", "ECC", "CacheErr", "TagLo", "TagHi", "ErrEPC", ""}
