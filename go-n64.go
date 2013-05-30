package main

import "fmt"
import "time"
import pcpu "go64/cpu"

func main() {
	fmt.Println("Test test")
	cpu := pcpu.NewCPU()
	fmt.Printf("CP: %v\n", cpu.PC)
	cpu.Start()
	for cpu.State() != 0 {
		time.Sleep(100 * time.Millisecond)
	}

}
