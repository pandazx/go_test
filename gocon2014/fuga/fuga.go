package fuga

//package main
// main()はpackage mainでのみ実行できる
// これがないと、go run: cannot run non-main package と言われる

import "fmt"

// 構造体の定義
type FugaStruct struct {
	// field	type
	A string // public
	b int    // private
}

// public method
func Fuga(name string) string {
	return "public fuga: " + name + "," + fuga()
}

// private method
func fuga() string {
	return "private fuga"
}

// return pointer
// goのdefaultは値渡しになる
func NewFuga(a string) *FugaStruct {
	//return &Fuga{a, 100}
	return &FugaStruct{
		b: 100,
		A: a,
	}
}

type FugaKai struct {
	*FugaStruct
	C string
}

// hogeからimportされると最初に実行される
func init() {
	fmt.Println("fuga init")
	fuga := &FugaStruct{"hoge", 100}
	fugakai := &FugaKai{fuga, "Cstr"}
	fugakai.A = "hogekai"
	fmt.Println(fugakai.A, fugakai.C)
}

type Hex int

// String()はtoString()のこと
func (h Hex) String() string {
	return fmt.Sprintf("0x%x", int(h))
}

func main() {
	fmt.Println("fuga main")
	fmt.Println(Hex(15))
}
