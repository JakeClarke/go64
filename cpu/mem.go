package cpu

type Memory [0x8000]byte

func (m *Memory) StoreWord(addr DWORD, value WORD) {
	m.StoreHalfWord(addr+2, HWORD(value>>16))
	m.StoreHalfWord(addr, HWORD(value))
}

func (m *Memory) StoreHalfWord(addr DWORD, value HWORD) {
	m[addr+1] = byte(value >> 8)
	m[addr] = byte(value)
}

func (m *Memory) LoadWord(addr DWORD) WORD {
	var res WORD

	res = WORD(m.LoadHalfWord(addr+2)) << 16
	res += WORD(m.LoadHalfWord(addr))

	return res
}

func (m *Memory) LoadHalfWord(addr DWORD) HWORD {
	var res HWORD

	res = HWORD(m[addr+1]) << 8
	res += HWORD(m[addr])

	return res
}

func (m *Memory) LoadByteU(addr DWORD) uint8 {
	return uint8(m[addr])
}

func (m *Memory) Init() {
	for index, _ := range m {
		m[index] = 0x00
	}
}

func (m *Memory) SetBytes(addr DWORD, bytes []uint8) {
	for index, v := range bytes {
		m[addr+DWORD(index)] = v
	}
}
