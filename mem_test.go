package main

import (
	"testing"
)

func TestStoreAndLoad(t *testing.T) {
	const val = WORD(0xFFFF)
	const storeAddr = DWORD(0xFF)
	mem := new(Memory)
	mem.StoreWord(storeAddr, val)
	if x := mem.LoadWord(storeAddr); x != val {
		t.Errorf("Store val(%v, addr %v), loaded wrong (%v)", val, storeAddr, x)
	}
}
