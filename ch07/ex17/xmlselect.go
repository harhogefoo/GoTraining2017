// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 214.
//!+

// Xmlselect prints the text of selected elements of an XML document.
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
	"bytes"
)

func main() {
	e := parseArgs()

	dec := xml.NewDecoder(os.Stdin)
	var stack []xml.StartElement
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if containsAll(stack, e) {
				fmt.Printf("%s: %s\n", joinStack(stack), tok)
			}
		}
	}
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(stack []xml.StartElement, arg []*element) bool {
	for len(arg) <= len(stack) {
		if len(arg) == 0 {
			return true
		}
		if stack[0].Name.Local == arg[0].name {
			if containsAllAttributes(stack[0].Attr, arg[0].attributes) {
				arg = arg[1:]
			}
		}
		stack = stack[1:]
	}
	return false
}

func joinStack(stack []xml.StartElement) string {
	var result []string

	for _, elem := range stack {
		result = append(result, elem.Name.Local)
	}
	return strings.Join(result, " ")
}

func containsAllAttributes(stack []xml.Attr, arg []attribute) bool {
	for _, argAttr := range arg {
		matched := false
		for _, stackAttr := range stack {
			if stackAttr.Name.Local == argAttr.name {
				if stackAttr.Value != argAttr.value {
					return false
				} else {
					matched = true
				}
			}
		}
		if !matched {
			return false
		}
	}
	return true
}

//!-

type attribute struct {
	name string
	value string
}

func (a *attribute) String() string {
	return fmt.Sprintf("%s=\"%s\"", a.name, a.value)
}

type element struct {
	name string
	attributes []attribute
}

func  (e *element) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("<")
	buffer.WriteString(e.name)
	for _, attr := range e.attributes {
		buffer.WriteByte(' ')
		buffer.WriteString(attr.String())
	}
	buffer.WriteString(">")
	return buffer.String()
}

func parseArgs() []*element {
	var result []*element
	var e *element
	for _, arg := range os.Args[1:] {
		if strings.Contains(arg, "=") {
			if e == nil {
				fmt.Printf("element name is not specified: [%s]  ignoreed\n", arg)
				continue
			}
			name := strings.Split(arg, "=")
			if len(name) != 2 {
				fmt.Printf("Illegal format: [%s] ignored\n", arg)
			}
			attr := attribute{name[0], name[1]}
			e.attributes = append(e.attributes, attr)
			continue
		}
		if e != nil {
			result = append(result, e)
		}
		e = &element{name: arg}
	}
	result = append(result, e)
	return result
}