# primes

A proof-of-concept app for [eggos](https://github.com/icexin/eggos) which tries to compute mersenne primes on bare metal, exposing a tcp server to return something vaguely jsony.

## Why?

It's a cool, easily verifiable, resource intensive task which would be great as a unikernel. Eggos is cool, and does some nice stuff, and writing a non-trivial complex application is a cool way to try it out.

## Building

This project provides a `Makefile` which wraps the `egg` command from [eggos](https://github.com/icexin/eggos)

```bash
$ make               # analogous to make kernel.elf
$ make kernel.elf    # compiles a valid multiboot kernel
$ make primes.iso    # compiles kernel.elf, and packs into an iso which can be booted from a usb
```

## Running on Qemu

The project `Makefile` can be used to pack an iso file, which it will then run via qemu:

```bash
$ make run
```

## Accessing primes

The primes unikernel exposes a tcp socket on port 11111. When the app boots, it solicits a DHCP lease which it then prints to screen.

Accessing this IP (whatever it may be) on port 11111 will return a json string which looks like:

```bash
$ nc localhost 11111 | jq '.'
{
  "i": 2281,
  "v": 1.7976931348623157e+308,
  "len(v)": 687
}
```

The values have the following meanings:

  * `i` is the value used in `2^i - 1` to calculate a prime
  * `v` is the actual computed prime number
  * `len(v)` is the number of digits in this prime number

## Licence

MIT License

Copyright (c) 2022 James Condron ('jspc')

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
