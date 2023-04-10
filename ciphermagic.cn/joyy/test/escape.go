package test

var z *int

func escape() {
	a := 1
	z = &a
}
