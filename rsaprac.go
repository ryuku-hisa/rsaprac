package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	var mode int
	fmt.Println("Enter 0 or 1")
	fmt.Println("Encyption: 0, Decryption: 1")
	fmt.Print("mode: ")
	fmt.Scan(&mode)

	var data, p, q, E int

	fmt.Print("p = ")
	fmt.Scan(&p)
	fmt.Print("q = ")
	fmt.Scan(&q)
	fmt.Print("E = ")
	fmt.Scan(&E)

	N := p * q
	println("n =", N)
	L := LCM(p-1, q-1)
	println("L =", L)

	n := 1
	for {
		if (L*n+1)%E != 0 {
			n++
		} else {
			println("n =", n)
			break
		}
	}
	D := (L*n + 1) / E
	println("D =", D)

	fmt.Println("* press ctrl+c if you want to quit *")
	var crypto, plain int
	if mode == 0 {
		for {
			fmt.Print("Data = ")
			fmt.Scan(&data)
			crypto = int(math.Pow(float64(data), float64(E))) % N
			fmt.Println("Cryptogram is", crypto)
		}
	} else if mode == 1 {
		for {
			fmt.Print("Data = ")
			fmt.Scan(&data)
			plain = int(math.Pow(float64(data), float64(D))) % N
			fmt.Println("PlainText is", plain)
		}
	} else {
		log.Panic("non proper code")
	}
}

// following code is from : https://play.golang.org/p/SmzvkDjYlb

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
