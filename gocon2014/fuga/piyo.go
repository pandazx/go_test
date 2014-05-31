package fuga

// interface
type Piyo interface {
	C() string
}

type PiyoImpl struct {
}

// 明示的にimplementsする必要はない
func (p *PiyoImpl) C() string {
	return "hoge"
}

type PiyoFunc func() string

func (f PiyoFunc) C() string {
	return f()
}

type OrenoReader interface {
	Read(p []byte) (n int, err error)
}

func init() {
	/*
		var piyo Piyo
		piyo = &PiyoImpl{}
	*/

	// 以下はエラーにならない
	/*
		var r io.Reader
		var or OrenoReader
		or = r
	*/

	// closureでPiyoFuncにmethod 追加
	/*
		f := func() string {
			return "closure"
		}
		piyo = PiyoFunc(f)
	*/

}
