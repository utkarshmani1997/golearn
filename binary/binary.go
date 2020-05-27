package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

//this type represnts a record with three fields
type payload struct {
	One   float32
	Two   float64
	Three uint32
}

func main() {
	file, err := os.Create("test.bin")
	defer file.Close()
	if err != nil {
		log.Fatal("main:", err)
	}

	readFile()
	writeFile(file)
}

func readFile() {
	file, err := os.Open("test.bin")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	m := payload{}
	for i := 0; i < 10; i++ {
		//data := readNextBytes(file, 16)
		buffer := bufio.NewReaderSize(file, 8096)
		err = binary.Read(buffer, binary.BigEndian, &m)
		if err != nil {
			log.Fatal("binary.Read failed", err)
		}
		fmt.Println(m)
	}
}

func readNextBytes(file *os.File, number int) []byte {
	bytes := make([]byte, number)

	_, err := file.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}

	return bytes
}

func writeFile(file *os.File) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 10; i++ {

		s := &payload{
			r.Float32(),
			r.Float64(),
			r.Uint32(),
		}
		var bin_buf bytes.Buffer
		binary.Write(&bin_buf, binary.BigEndian, s)
		//b :=bin\_buf.Bytes()
		//l := len(b)
		//fmt.Println(l)
		writeNextBytes(file, bin_buf.Bytes())

	}
}
func writeNextBytes(file *os.File, bytes []byte) {

	_, err := file.Write(bytes)

	if err != nil {
		log.Fatal(err)
	}

}
