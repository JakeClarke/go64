package main

import "log"
import "time"
import pcpu "go64/cpu"

func main() {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds)
	cpu := pcpu.NewCPU()
	log.Printf("CP: %v\n", cpu.PC)
	cpu.Start()
	for cpu.State() != 0 {
		time.Sleep(100 * time.Millisecond)
		log.Printf("CPU state: %v\n", cpu.State())
	}

}
