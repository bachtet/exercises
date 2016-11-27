// Towers of Hanoi, print out length.
package main

import "fmt"

var towers [3][]int

func main() {
	disks := 5
	setUp(disks)
	move(disks, 0, 1, 2)
}

func setUp(n int) {
	towers[0] = make([]int, n)
	for i := range towers[0] {
		towers[0][i] = n - i
	}
}

func move(n, a, b, c int) {
	if n <= 0 {
		return
	}
	move(n-1, a, c, b)
	disk1 := towers[a][len(towers[a])-1]
	towers[a] = towers[a][:len(towers[a])-1]
	towers[c] = append(towers[c], disk1)
	fmt.Printf("Tower a %d, Tower b %d, Tower c %d \n", len(towers[0]), len(towers[1]), len(towers[2]))
	move(n-1, b, a, c)
}
