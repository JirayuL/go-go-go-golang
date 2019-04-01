package main

import (
	"fmt"
	"time"
)

const N = 1000

func mainHH() {
	var a [N][N]int
	var b [N][N]int
	var c [N][N]int

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			a[i][j] = 1
			b[i][j] = 1
		}
	}

	startTime := time.Now()

	// ver 1
	// for i := 0; i < N; i++ {
	// 	for j := 0; j < N; j++ {
	// 		for k := 0; k < N; k++ {
	// 			c[i][j] += a[i][k] * b[k][j]
	// 		}
	// 	}
	// }

	// ver 2
	// for i := 0; i < N; i++ {
	// 	for j := 0; j < N; j++ {
	// 		go func(i, j int) {
	// 			for k := 0; k < N; k++ {
	// 				c[i][j] += a[i][k] * b[k][j]
	// 			}
	// 		}(i, j)
	// 	}
	// }

	// ver 3
	// for i := 0; i < N; i++ {
	// 	go func(i int) {
	// 		for j := 0; j < N; j++ {
	// 			for k := 0; k < N; k++ {
	// 				c[i][j] += a[i][k] * b[k][j]
	// 			}
	// 		}
	// 	}(i)
	// }

	// ver 4
	// for i := 0; i < N; i++ {
	// 	go func(i int) {
	// 		for j := 0; j < N; j++ {
	// 			num := 0
	// 			for k := 0; k < N; k++ {
	// 				num += a[i][k] * b[k][j]
	// 			}
	// 			c[i][j] = num
	// 		}
	// 	}(i)
	// }

	// ver 5
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			b[i][j] = b[j][i]
		}
	}
	for i := 0; i < N; i++ {
		go func(i int) {
			for j := 0; j < N; j++ {
				num := 0
				for k := 0; k < N; k++ {
					num += a[i][k] * b[j][k]
				}
				c[i][j] = num
			}
		}(i)
	}

	endTime := time.Since(startTime)

	// for i := 0; i < N; i++ {
	// 	fmt.Println(c[i])
	// }

	fmt.Println(endTime)
}
