package main

type Instruction WORD

const OP_MASK = Instruction(0xFC000000)
const R_SRC_MASK = Instruction(0x3E00000)
const R_TAR_MASK = Instruction(R_SRC_MASK >> 5)
const R_DST_MASK = Instruction(R_TAR_MASK >> 5)
const JUMP_MASK = Instruction(0x3FFFFFF)

const FUNCT_MASK = Instruction(0x3F)
const SHAMT_MASK = Instruction(0x3E0)

const IMMEDIATE_MASK = Instruction(0xFFFF)

func (i Instruction) OP() byte {
	return byte(i >> 26)
}

func (i Instruction) RS() byte {
	return byte((R_DST_MASK & i) >> 21)
}

func (i Instruction) RT() byte {
	return byte((R_DST_MASK & i) >> 16)
}

func (i Instruction) RD() byte {
	return byte((R_DST_MASK & i) >> 11)
}

func (i Instruction) Immediate() HWORD {
	return HWORD(i & Instruction(IMMEDIATE_MASK))
}

func (i Instruction) Addr() HWORD {
	return HWORD(i & JUMP_MASK)
}

func (i Instruction) Shamt() byte {
	return byte((i & SHAMT_MASK) >> 5)
}

func (i Instruction) Funct() byte {
	return byte((i & FUNCT_MASK))
}
