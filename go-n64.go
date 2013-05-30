package main

import "fmt"

func main() {
	fmt.Println("Test test")
	cpu := NewCPU()
	fmt.Printf("CP: %v\n", cpu.PC)
	cpu.Tick()
	fmt.Printf("CP: %v\n", cpu.PC)
	fmt.Println("Registers:")

	for i, _ := range cpu.GPR {
		fmt.Printf("%s: %v\n", GPR_NAMES[i], cpu.GPR[i])
	}
}
