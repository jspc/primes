package main

import (
	"context"
	"fmt"
	"math/big"
	"net"
	"sync"

	"github.com/icexin/eggos/inet"
	"golang.org/x/sync/semaphore"
)

var (
	sem      = semaphore.NewWeighted(int64(6))
	locker   = new(sync.Mutex)
	largestI = int64(0)

	zero = big.NewInt(0)
	one  = big.NewInt(1)
	two  = big.NewInt(2)

	start = two
	step  = one
)

func main() {
	inet.Init()

	fmt.Println("hello world <3")
	defer func() {
		fmt.Println("bye bye")
	}()

	go tcpServer()

	ctx := context.Background()

	for i := int64(0); ; i++ {
		sem.Acquire(ctx, 1)
		go func(i int64) {
			defer sem.Release(1)

			if primeish(uint(i)) {
				locker.Lock()
				defer locker.Unlock()

				largestI = i
			}
		}(i)
	}
}

func tcpServer() {
	l, err := net.Listen("tcp", "0.0.0.0:11111")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}

		conn.Write(formatLargest())
		conn.Close()
	}
}

// this is hand crafted json-alike
func formatLargest() []byte {
	iI := big.NewInt(largestI)

	v := iI.Exp(two, iI, nil)
	v.Sub(v, one)

	return []byte(fmt.Sprintf(`{"i":%d,"v":%s,"len(v)":%v}`, largestI, v, len(v.String())))
}

// primeish is largely taken from *somewhere* but for the life of me
// I can't remember where- I suspect if I remembered what this algorithm
// is called (two names, first designed some/most of the algo in the.... 1700s?
// and the second name improved it in the 1980s?) I'd be able to follow the same
// google strategy I took to get here.
//
// as soon as I figure it out, I'll update this atrribution/ remove the code,
// depending on how it's licenced
func primeish(p uint) bool {
	var (
		dummy1, dummy2 big.Int
	)

	s := big.NewInt(4)
	m := big.NewInt(0)
	m = m.Sub(m.Lsh(one, p), one) // = (1 << p) - 1

	for i := 0; i < int(p)-2; i++ {
		s = s.Sub(s.Mul(s, s), two)

		for s.Cmp(m) == 1 {
			s.Add(dummy1.And(s, m), dummy2.Rsh(s, p))
		}

		if s.Cmp(m) == 0 {
			s = zero
		}
	}

	return s.Cmp(zero) == 0
}
