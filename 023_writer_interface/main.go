package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type taco struct {
	name string
}

func (t taco) writeOut(w io.Writer) error {
	_, err := w.Write([]byte(t.name))
	return err
}

func WriterInterfacePractice() {

	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer f.Close()

	s := []byte("Hello Gophers")

	_, err = f.Write(s)
	if err != nil {
		log.Fatalf("error %s", err)
	}

	var byteBuff bytes.Buffer

	t := taco{
		name: "gordita crunch",
	}

	//writing to the file
	t.writeOut(f)
	//writing to the bytes buffer
	t.writeOut(&byteBuff)

	fmt.Println(byteBuff.String())

	//bytes buffer practice
	b := bytes.NewBufferString("Hello ")
	fmt.Println(b.String())

	b.WriteString("gophers")
	fmt.Println(b)

	b.Reset()
	fmt.Println(b)

	b.WriteString("tacos")
	fmt.Println(b)

	b.Write([]byte(" are delicious"))
	fmt.Println(b)

	fmt.Printf("Type of b = %T and value = %v\n", b, b)
	fmt.Printf("Type of b.String() = %T and value = %v\n", b.String(), b.String())

}

func main() {
	WriterInterfacePractice()
}
