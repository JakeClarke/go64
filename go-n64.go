package main

import "fmt"
import pcpu "go64/cpu"

func main() {
	fmt.Println("Test test")
	cpu := pcpu.NewCPU()
	fmt.Printf("CP: %v\n", cpu.PC)
	cpu.Tick()
	fmt.Printf("CP: %v\n", cpu.PC)
	fmt.Println("Registers:")

	for i, _ := range cpu.GPR {
		fmt.Printf("%s: %v\n", pcpu.GPR_NAMES[i], cpu.GPR[i])
	}

	fmt.Printf("CP: %+v\n", cpu)
}
