default: kernel.elf

kernel.elf: *.go go.*
	egg build -o $@

primes.iso: kernel.elf
	egg pack -o $@ -k $<
