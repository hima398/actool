package main

import (
	"bufio"
	"flag"
	"math/rand"
	"os"
	"strconv"
)

const name = "ac"
const version = "0.0.1"

// TODO: randを引数で渡す
func CreateTestData(n, max int, filename string) {
	const interval = 1000
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return
	}
	defer file.Close()
	w := bufio.NewWriter(file)
	w.WriteString(strconv.Itoa(n) + "\n")

	modulus := 2*max + 1
	s := strconv.Itoa(rand.Int()%modulus - max)

	for i := 1; i < n; i++ {
		v := rand.Int()%modulus - max
		s += " " + strconv.Itoa(v)
		if i%interval == 0 {
			w.WriteString(s)
			w.Flush()
		}
	}
	w.WriteString(s)
	w.Flush()
}

func main() {

	// argments
	var ndata int
	var max int
	var filename string

	flag.IntVar(&ndata, "n", 10, "")
	flag.IntVar(&max, "x", 100, "")
	flag.StringVar(&filename, "f", "in.txt", "")

	flag.Parse()

	CreateTestData(ndata, max, filename)
}
