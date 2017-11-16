package main

type myError struct {
	Msg string
}

func (m *myError) Error() string {
	return m.Msg
}

func thereIsNoError() error {
	var me *myError = nil
	return me
}

func main() {
	e := thereIsNoError()
	if e != nil {
		panic("something really wrong")
	}
	println("now we are good")
}
