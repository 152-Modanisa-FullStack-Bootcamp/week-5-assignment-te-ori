package assignment

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddUint32(t *testing.T) {
	/*
		Sum uint32 numbers, return uint32 sum value and boolean overflow flag
		cases need to pass:
			math.MaxUint32, 1 => 0, true
			1, 1 => 2, false
			42, 2701 => 2743, false
			42, math.MaxUint32 => 41, true
			4294967290, 5 => 4294967295, false
			4294967290, 6 => 0, true
			4294967290, 10 => 4, true
	*/

	cases := []struct {
		op1, op2, sum uint32
		overflow      bool
	}{
		{math.MaxUint32, 1, 0, true},
		{1, 1, 2, false},
		{42, 2701, 2743, false},
		{42, math.MaxUint32, 41, true},
		{4294967290, 5, 4294967295, false},
		{4294967290, 6, 0, true},
		{4294967290, 10, 4, true},
		{4294967290, 4294967290, 4294967284, true},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c), func(t *testing.T) {
			sum, overflow := AddUint32(c.op1, c.op2)

			assert.Equal(t, c.sum, sum)
			assert.Equal(t, c.overflow, overflow)
		})

	}
}

func TestCeilNumber(t *testing.T) {
	/*
		Ceil the number within 0.25
		cases need to pass:	*/
	cases := []struct {
		input    float64
		expected float64
	}{
		{42.42, 42.50},
		{42, 42},
		{42.01, 42.25},
		{42.24, 42.25},
		{42.25, 42.25},
		{42.26, 42.50},
		{42.55, 42.75},
		{42.75, 42.75},
		{42.76, 43},
		{42.99, 43},
		{43.13, 43.25},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c), func(t *testing.T) {
			fmt.Printf("%v", c)
			point := CeilNumber(float64(c.input))

			assert.Equal(t, c.expected, point)
		})

	}
}

func TestAlphabetSoup(t *testing.T) {
	/*
		String with the letters in alphabetical order.
		cases need to pass:
		 	"hello" => "ehllo"
			"" => ""
			"h" => "h"
			"ab" => "ab"
			"ba" => "ab"
			"bac" => "abc"
			"cba" => "abc"
	*/

	cases := []struct {
		input, expected string
	}{
		{"hello", "ehllo"},
		{"", ""},
		{"h", "h"},
		{"ab", "ab"},
		{"ba", "ab"},
		{"bac", "abc"},
		{"cba", "abc"},
	}

	for _, c := range cases {
		result := AlphabetSoup(c.input)

		assert.Equal(t, c.expected, result)
	}
}

func TestStringMask(t *testing.T) {

	cases := []struct {
		S        string
		N        uint
		Expected string
	}{
		{"!mysecret*", 2, "!m********"},
		{"", 1, "*"},
		{"a", 1, "*"},
		{"string", 0, "******"},
		{"string", 3, "str***"},
		{"string", 5, "strin*"},
		{"string", 6, "******"},
		{"string", 7, "******"},
		{"s*r*n*", 3, "s*r***"},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c), func(t *testing.T) {
			result := StringMask(c.S, c.N)
			assert.Equal(t, c.Expected, result)
		})
	}
}

func TestWordSplit(t *testing.T) {
	words := "apple,bat,cat,goodbye,hello,yellow,why,helloo,at"
	/*
		Your goal is to determine if the first element in the array can be split into two words,
		where both words exist in the dictionary(words variable) that is provided in the second element of array.

		cases need to pass:
			[2]string{"hellocat",words} => hello,cat
			[2]string{"catbat",words} => cat,bat
			[2]string{"yellowapple",words} => yellow,apple
			[2]string{"",words} => not possible
			[2]string{"notcat",words} => not possible
			[2]string{"bootcamprocks!",words} => not possible
	*/
	cases := []struct {
		input  [2]string
		expect string
	}{
		{[2]string{"hellooat", words}, "helloo,at"},
		{[2]string{"hellocat", words}, "hello,cat"},
		{[2]string{"catbat", words}, "cat,bat"},
		{[2]string{"yellowapple", words}, "yellow,apple"},
		{[2]string{"", words}, "not possible"},
		{[2]string{"notcat", words}, "not possible"},
		{[2]string{"bootcamprocks!", words}, "not possible"},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c), func(t *testing.T) {
			result := WordSplit(c.input)

			assert.Equal(t, c.expect, result)

		})

	}

}

func TestVariadicSet(t *testing.T) {
	/*
		FINAL BOSS ALERT :)
		Tip: Learn and apply golang variadic functions(search engine -> "golang variadic function" -> WOW You can really dance! )

		Convert inputs to set(no duplicate element)
		cases need to pass:
			4,2,5,4,2,4 => []interface{4,2,5}
			"bootcamp","rocks!","really","rocks! => []interface{"bootcamp","rocks!","really"}
			1,uint32(1),"first",2,uint32(2),"second",1,uint32(2),"first" => []interface{1,uint32(1),"first",2,uint32(2),"second"}
	*/

	cases := []struct {
		input, expected []interface{}
	}{
		{[]interface{}{"bootcamp", "rocks!", "really", "rocks!"}, []interface{}{"bootcamp", "rocks!", "really"}},
		{[]interface{}{1, uint32(1), "first", 2, uint32(2), "second", 1, uint32(2), "first"}, []interface{}{1, uint32(1), "first", 2, uint32(2), "second"}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c), func(t *testing.T) {
			set := VariadicSet(c.input...)
			m := make(map[interface{}]struct{})
			for _, element := range c.expected {
				m[element] = struct{}{}
			}

			for _, element := range set {
				_, isExist := m[element]
				assert.Equal(t, isExist, true)
			}
		})
	}
}
