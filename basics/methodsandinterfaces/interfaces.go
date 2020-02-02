package methodsandinterfaces

import (
	"fmt"
	"io"
	"strings"
)

type MyInterface interface {
	Foo()
}

//--------------------------

type MyStruct struct {
	A, B, C int
}

func (ms MyStruct) Foo() {
	fmt.Println("MyStruct (value reciever)", ms)
}

func (ms MyStruct) String() string {
	return fmt.Sprintf("\"MyStruct [%v %v %v]\"", ms.A, ms.B, ms.C)
}

func (ms MyStruct) WillReturnError() (int, error) {
	return 0, MyError{}
}

//--------------------------

type MyError struct {
}

func (e MyError) Error() string {
	return "my custom error"
}

//--------------------------

type MyInt int

func (mi MyInt) Foo() {
	fmt.Println("MyInt (value reciever)", mi)

}

//--------------------------

type AnotherStruct struct {
}

func (ms *AnotherStruct) Foo() {
	fmt.Println("AnotherStruct (pointer reciever)", ms)

}

//--------------------------

func Interfaces() {
	interfaces()
	typeAssertions()
	typeSwitch()
	stringers()
	errors()
	readers()
}

func interfaces() {
	var i MyInterface
	fmt.Println("nil interface", i)

	i = MyInt(3)
	i.Foo()

	i = MyStruct{}
	i.Foo()

	var nilMyStruct MyStruct
	i = nilMyStruct
	fmt.Printf("Interface with nil internal concrete value (%v, %T)\n", i, i)

	fmt.Println("A type only implements and interface if the method recievers (pointers or values) match the concrete item assigne to an interface variable.")
	fmt.Println("This is why it's recomended not to mix reciever types - as it will be impossible for types to implemnent interfaces if the type has mixed reciever types.")

	var j MyInterface
	j = &AnotherStruct{}
	j.Foo()

	fmt.Println("An empty interface may hold values of any type.")
	var k interface{}
	k = i
	fmt.Println("\tk", k)
	k = "foobar"
	fmt.Println("\tk", k)

}

func typeAssertions() {

	var i MyInterface
	i = MyStruct{}

	fmt.Println("A type assertion provides access to an interface value's underlying concrete value.")
	t := i.(MyStruct)
	fmt.Println("\tt is of the type asserted", t)

	t2, ok := i.(MyInt)
	fmt.Println("\tt2 is not of the type asserted", t2, ok)
}

func typeSwitch() {

	var i MyInterface
	i = MyStruct{}

	fmt.Println("A type switch is a construct that permits several type assertions in series.")
	switch v := i.(type) {
	case MyInt:
		fmt.Printf("\tMyInt type %T!\n", v)
	case MyStruct:
		fmt.Printf("\tMyStruct type %T!\n", v)
	default:
		fmt.Printf("\tI don't know about type %T!\n", v)
	}
}

func stringers() {

	fmt.Println("One of the most ubiquitous interfaces is Stringer defined by the fmt package.")
	fmt.Println("A Stringer is a type that can describe itself as a string.")

	var s = MyStruct{7, 8, 9}

	fmt.Println("\t see: func (ms MyStruct) String() string", s)

}

func errors() {

	fmt.Println("Go programs express error state with error values.")
	fmt.Println("The error type is a built-in interface.")

	var s = MyStruct{7, 8, 9}

	_, err := s.WillReturnError()
	if err != nil {
		fmt.Printf("\tError returned: %v\n", err)
		return
	}

}

func readers() {
	fmt.Println("The io package specifies the io.Reader interface, which represents the read end of a stream of data.")
	fmt.Println("func (T) Read(b []byte) (n int, err error)")

	b := make([]byte, 8)
	r := strings.NewReader("Hello, Reader!")
	for {
		n, err := r.Read(b)
		fmt.Printf("\tn = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("\tb[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

	r.Reset("Hello Runes")
	for {
		runeChar, size, err := r.ReadRune()
		fmt.Printf("\truneChar = %v\t%q\tsize = %v err = %v \n", runeChar, runeChar, size, err)
		if err == io.EOF {
			break
		}
	}

}
