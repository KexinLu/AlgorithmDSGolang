package strings

import (
	"bytes"
	"fmt"
)

// golang string is an array-like linear data structure

// most common way to deal with strings includes looping through string
// Printing a string
// using buffer to build a string

// loop through a string

//
/**
fmt
int
%b base 2
%c  the character represented by the corresponding Unicode code point
%d  base 10
%o  base 8
%O  base 8 with 0o prefix
%q  a single-quoted character literal safely escaped with Go syntax.
%x  base 16, with lower-case letters for a-f
%X  base 16, with upper-case letters for A-F
%U  Unicode format: U+1234; same as "U+%04X"

float
%b  decimalless scientific notation with exponent a power of two,
in the manner of strconv.FormatFloat with the 'b' format,
e.g. -123456p-78
%e  scientific notation, e.g. -1.234456e+78
%E  scientific notation, e.g. -1.234456E+78
%f  decimal point but no exponent, e.g. 123.456
%F  synonym for %f
%g  %e for large exponents, %f otherwise. Precision is discussed below.
%G  %E for large exponents, %F otherwise
%x  hexadecimal notation (with decimal power of two exponent), e.g. -0x1.23abcp+20
%X  upper-case hexadecimal notation, e.g. -0X1.23ABCP+20

string
%s  the uninterpreted bytes of the string or slice
%q  a double-quoted string safely escaped with Go syntax
%x  base 16, lower-case, two characters per byte
%X  base 16, upper-case, two characters per byte

slice
%p  address of 0th element in base 16 notation, with leading 0x

pointer
%p  base 16 notation, with leading 0x
The %b, %d, %o, %x and %X verbs also work with pointers,
formatting the value exactly as if it were an integer.

**/
// fmt is a set of helpers to help you manipulate strings
func Fmt(input string) string {
	return fmt.Sprintf("the value we have is %v", input)
}

// strings are cut into "rune"s if you loop through it
func LoopThrough(input string) {
	for idx, char := range input {
		fmt.Printf("character at index %d is %c, with Unicode %U \n", idx, char, char)
	}
}

// only print odd position chars
// byte(unit8) -> ASCII
// rune(int32) -> Unicode with UTF-8
func StringBuilder(input string) {
	var buffer bytes.Buffer
	// byteVal here is of type Rune
	for idx, byteVal := range input {
		if idx&1 == 0 {
			buffer.WriteRune(byteVal)
		}
	}
	for i := 0; i < len(input); i++ {
		if i&1 == 0 {
			buffer.WriteByte(input[i])
		}
	}

	fmt.Println(buffer.String())
}
