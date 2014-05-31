package main

import (
	"fmt"

	"./fuga"
)

func main() {
	fmt.Println("Hello, 世界")

	// read local file
	/*
		file, err := os.Open("file")
		if err != nil {
			log.Fatal(err)
		}
	*/

	// do my package method
	a := fuga.Fuga("name")
	fmt.Println(a)

	// cannot do private method
	//a = fuga.fuga()

	// get my package instance
	f := fuga.NewFuga("struct name")
	fmt.Println(f.A)

	// below a's scope is if else scope only
	if a := f.A; a == "struct name" {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	fmt.Println("for")
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	fmt.Println("like while")
	i := 0
	for true {
		i += 1
		fmt.Println(i)
		if i > 10 {
			break
		}
	}

	fmt.Println("for _ v")
	// 使わない変数は_で潰せる
	for _, v := range []int{10, 20, 30} {
		// i = 0,1,2
		// v = 10,20,30
		fmt.Println(v)
	}

	fmt.Println("slice map")
	var s []int
	// make(type, initial size, capacity)
	s = make([]int, 1, 100)
	size := len(s) // 1
	fmt.Println(size)
	s = []int{1, 2, 3}
	fmt.Println(s)

	// get sub range
	s2 := s[0:2]
	fmt.Println(s2)

	// append element
	s3 := append(s, 4)
	fmt.Println(s3)

	// append few elements
	ss := []int{4, 5, 6}
	s = append(s, ss...)
	fmt.Println(s)

	// map
	// map[key]value
	m := make(map[string]int)
	m["hoge"] = 100
	fmt.Println(m)
	fmt.Println(m["hoge"])

	// check if key include of m
	v, ok := m["hoge"]
	fmt.Println(v, ok)

	// if hoge exist, print ok
	if _, ok := m["hoge"]; ok {
		fmt.Println("ok")
	}

	// init map
	m = map[string]int{
		"hoge": 100,
		"fuga": 200,
	}
	fmt.Println(m)
}
