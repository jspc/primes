default: kernel.elf

kernel.elf: export GOROOT = $(shell go1.16.13 env GOROOT)
kernel.elf: *.go go.*
	egg build -o $@

primes.iso: export GOROOT = $(shell go1.16.13 env GOROOT)
primes.iso: kernel.elf
	egg pack -o $@ -k $<

.PHONY: run
run: primes.iso
	egg  run -p 11111:11111 $<
