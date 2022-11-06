package main

import (
	"fmt"
	"os"
	"strconv"
)

var aArr []float64
var bArr []float64
var cArr []float64
var fArr []float64
var alpha []float64
var beta []float64
var matrix [][]float64
var n, m int
var eps, gamma float64
var s []string
var xArr []float64

const (
	test0 = iota
	test1
	test2
	test3
)

func main() {

	if len(os.Args) != 5 {
		fmt.Println("Input args: n eps gamma testNumber")
		return
		//bytes, err := os.ReadFile("data")
		//if err != nil {
		//	log.Fatal(err)
		//}
		//fileText := string(bytes[:])
		//s = strings.Split(fileText, ",")
		//filter()
		//n, m = getMatrixSizeFromFile(s)
		//if n != m {
		//	fmt.Println("n!=m")
		//	return
		//}
		//matrix = getMatrixFromFile(n)
		//getFFromFile(n)
		//getABCFromMatrix(n)
	} else {
		n, _ = strconv.Atoi(os.Args[1])
		eps, _ = strconv.ParseFloat(os.Args[2], 8)
		gamma, _ = strconv.ParseFloat(os.Args[3], 8)
		test, _ := strconv.Atoi(os.Args[4])
		aArr = make([]float64, 0)
		bArr = make([]float64, 0)
		cArr = make([]float64, 0)
		fArr = make([]float64, 0)
		switch test {
		case test1:
			for i := 0; i < n; i++ {
				if i != n-1 {
					cArr = append(cArr, -1)
				}
				if i != 0 {
					aArr = append(aArr, -1)
				}
				bArr = append(bArr, 2)
			}
			for i := 0; i < n; i++ {
				fArr = append(fArr, 2)
			}
		case test2:
			for i := 0; i < n; i++ {
				if i != n-1 {
					cArr = append(cArr, -1)
				}
				if i != 0 {
					aArr = append(aArr, -1)
				}
				bArr = append(bArr, 2)
			}
			for i := 0; i < n; i++ {
				fArr = append(fArr, 2+eps)
			}
		case test3:
			for i := 1; i <= n; i++ {
				if i != n {
					cArr = append(cArr, -1)
				}
				if i != 1 {
					aArr = append(aArr, -1)
				}
				bArr = append(bArr, float64(2*i)+gamma)
			}
			for i := 1; i <= n; i++ {
				fArr = append(fArr, float64(2*(i+1))+gamma)
			}
		}
	}

	fmt.Println(aArr, bArr, cArr, fArr)
	fmt.Println()
	getAlphaBeta(n)
	fmt.Println(alpha, beta)
	fmt.Println()
	runX(n)
	fmt.Println("x:")
	fmt.Println(xArr)

}

func runX(size int) {
	xArr = make([]float64, 0)
	xArr = append(xArr, beta[len(beta)-1])
	for i := size - 2; i >= 0; i-- {
		x := alpha[i]*xArr[size-1-i-1] + beta[i]
		xArr = append(xArr, x)
	}
}

func getAlphaBeta(size int) {
	alpha = make([]float64, 0)
	beta = make([]float64, 0)
	y := bArr[0]
	alpha = append(alpha, -cArr[0]/y)
	beta = append(beta, fArr[0]/y)

	for i := 1; i < size; i++ {
		y = bArr[i] + aArr[i-1]*alpha[i-1]
		if i != size-1 {
			alpha = append(alpha, -cArr[i]/y)
		}
		beta = append(beta, (fArr[i]/y)-(aArr[i-1]*beta[i-1]/y))
	}
}

func getABCFromMatrix(size int) {
	aArr = make([]float64, 0)
	bArr = make([]float64, 0)
	cArr = make([]float64, 0)

	for i := 0; i < size; i++ {
		if i != size-1 {
			cArr = append(cArr, matrix[i][i+1])
		}
		if i != 0 {
			aArr = append(aArr, matrix[i][i-1])
		}
		bArr = append(bArr, matrix[i][i])
	}
}

func filter() {
	for i, substr := range s {
		for _, e := range substr {
			if e == '\n' {
				remove(s, i)
				break
			}
		}
	}
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func getMatrixSizeFromFile(s []string) (int, int) {
	n, _ = strconv.Atoi(s[0])
	m, _ = strconv.Atoi(s[1])
	return n, m
}

func getFFromFile(size int) {
	fArr = make([]float64, 0)
	for j := 0; j < size; j++ {
		el, _ := strconv.Atoi(s[2+j])
		fArr = append(fArr, float64(el))
	}
}

func getMatrixFromFile(size int) [][]float64 {
	m := make([][]float64, size)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			el, _ := strconv.Atoi(s[2+n+(i*n+j)])
			m[i] = append(m[i], float64(el))
		}
	}
	return m
}
