package cpu

type Memory [0x8000]byte

// Store word stores a word to the memory.
func (m *Memory) StoreWord(addr DWORD, value WORD) {
	m.StoreHalfWord(addr+2, HWORD(value>>16))
	m.StoreHalfWord(addr, HWORD(value))
}

// Store halfword stores a halfword to the memory.
func (m *Memory) StoreHalfWord(addr DWORD, value HWORD) {
	m[addr+1] = byte(value >> 8)
	m[addr] = byte(value)
}

// Load word loads a word from the memory.
func (m *Memory) LoadWord(addr DWORD) WORD {
	var res WORD

	res = WORD(m.LoadHalfWord(addr+2)) << 16
	res += WORD(m.LoadHalfWord(addr))

	return res
}

// Load halfword loads a halfword from the memory.
func (m *Memory) LoadHalfWord(addr DWORD) HWORD {
	var res HWORD

	res = HWORD(m[addr+1]) << 8
	res += HWORD(m[addr])

	return res
}

// Load bytes unsigned loads an unsigned byte.
func (m *Memory) LoadByteU(addr DWORD) uint8 {
	return uint8(m[addr])
}

// Init resets all memory locations to 0.
func (m *Memory) Init() {
	for index, _ := range m {
		m[index] = 0x00
	}
}

// SetBytes sets large numbers of bytes.
func (m *Memory) SetBytes(addr DWORD, bytes []uint8) {
	for index, v := range bytes {
		m[addr+DWORD(index)] = v
	}
}
