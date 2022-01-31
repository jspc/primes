default: kernel.elf

kernel.elf: *.go go.*
	egg build -o kernel.elf

primes.iso: kernel.elf
	egg pack -o eggos.iso -k kernel.elf
