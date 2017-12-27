package main

import (
	"bytes"
	"fmt"
)

const UINT_SIZE = 32 << (^uint(0) >> 63)

// IntSetは負はない小さな整数のセットです。
// そのゼロ値は空セットを表しています。
type IntSet struct {
	words []uint
}

// Hasは筆はない値xをセットが含んでいるか否かを報告します。
func (s *IntSet) Has(x int) bool {
	word, bit := x/UINT_SIZE, uint(x%UINT_SIZE)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Addはセットに筆はない値xを追加します。
func (s *IntSet) Add(x int) {
	word, bit := x/UINT_SIZE, uint(x%UINT_SIZE)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWithは、sとtの和集合をsに設定します。
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// Stringは"{1 2 3}"の形式の文字列としてセットを返します。
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < UINT_SIZE; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", UINT_SIZE*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// 要素数を返す
func (s *IntSet) Len() int {
	var length int
	for _, word := range s.words {
		for j := 0; j < UINT_SIZE; j++ {
			if word&(1<<uint(j)) != 0 {
				length += 1
			}
		}
	}
	return length
}

// セットからxを取り除きます
func (s *IntSet) Remove(x int) {
	word, bit := x/UINT_SIZE, uint(x%UINT_SIZE)
	if word >len(s.words) {
		return
	}
	s.words[word] &^= 1 << bit
}

// セットから全ての要素を取り除きます
func (s *IntSet) Clear() {
	s.words = []uint{}
}

// セットのコピーを返す
func (s *IntSet) Copy() *IntSet {
	var c IntSet
	c.words = make([]uint, len(s.words))
	copy(c.words, s.words)
	return &c
}

func (s *IntSet) AddAll(numbers ...int) {
	for _, number := range numbers {
		s.Add(number)
	}
}

func (s *IntSet) IntersectWith(t *IntSet) {
	for i, word := range t.words {
		if i < len(s.words) {
			s.words[i] &= word
		}
	}

	for i := len(t.words); i < len(s.words); i++ {
		s.words[i] = 0
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, word := range t.words {
		if i < len(s.words) {
			s.words[i] &^= word
		}
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, word := range t.words {
		if i < len(s.words) {
			s.words[i] ^= word
		} else {
			s.words = append(s.words, word)
		}
	}
}

func (s *IntSet) Elems() []int {
	length := s.Len()
	if length == 0 {
		return []int{}
	}

	elems := make([]int, 0, length)
	for i, word := range s.words {
		for bit := uint(0); bit < UINT_SIZE; bit++ {
			if word&(1<<bit) != 0 {
				elems = append(elems, i*UINT_SIZE+int(bit))
			}
		}
	}
	return elems
}

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String())

	y.Add(2)
	y.Add(9)
	y.Add(42)
	fmt.Println(y.String())

	x.UnionWith(&y)
	fmt.Println(x.String())

	fmt.Println(x.Has(9), x.Has(123))

	fmt.Println(x.Len())

	c := x.Copy()

	x.Remove(1)
	fmt.Println(x.String())

	x.Clear()
	fmt.Println(c.String())

	x.AddAll(1, 2, 3)
	fmt.Println(x.String())

	x.IntersectWith(&y)
	fmt.Println(x.String())

	y.DifferenceWith(&x)
	fmt.Println(y.String())

	x.SymmetricDifference(&y)
	fmt.Println(x.String())

	fmt.Println(x.Elems())
}
