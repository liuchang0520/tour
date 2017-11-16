package main

type thing struct {
	value string
}

func (t thing) SetValue(val string) {
	t.value = val
}

func main() {
	th := thing{value: "testing"}
	th.SetValue("haha")
	if th.value != "haha" {
		panic("something not right!")
	}
}
