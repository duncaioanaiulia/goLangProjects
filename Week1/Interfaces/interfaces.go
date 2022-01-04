package interfaces

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	fmt.Println("TRAINING INTERFACES")
	BasicSyntax()
	InterfaceIncrement()
	ComposingInterfaces()
	EmptyInterface()
	TypeSwitches()
	ValuesVSPointers()
	BestPractices()
}

type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

type WriteCloser interface {
	Writer
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

type ConsoleWriter struct{}

func BasicSyntax() {
	fmt.Println("Interfaces describe behaviors not data, contains data definitions")

	fmt.Println("Inplicet implementations:")
	var w Writer = ConsoleWriter{}
	w.Write([]byte("EPAM"))
}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

func InterfaceIncrement() {
	intCounter := IntCounter(0)
	var inc Incrementer = &intCounter
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}
}

func ComposingInterfaces() {
	var wc WriteCloser = NewBufferedWriterClosed()
	wc.Write([]byte("Hello EPAM, is my first week"))
	wc.Close()

	bwc := wc.(*BufferedWriterCloser)
	fmt.Println(bwc)
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}

	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if (err) != nil {
			return err
		}
	}
	return nil
}

func NewBufferedWriterClosed() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

func EmptyInterface() {
	fmt.Println("--------Empty Interface:")
	var obj interface{} = NewBufferedWriterClosed()
	if wc, ok := obj.(WriteCloser); ok {
		wc.Write([]byte("Exemplu Interfaces"))
		wc.Close()
	}
	r, ok := obj.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Conversion failed")
	}

	fmt.Println("Type:")
	var i interface{} = 0
	switch i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("??")
	}
}

func TypeSwitches() {

}

func ValuesVSPointers() {

}

func BestPractices() {
	fmt.Println("create small interfaces")
	fmt.Println("sigle method interfaces are some of the most powerful and flexible ")
	fmt.Println("io.writer, io.reader, interface{}")

}
