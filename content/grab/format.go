package main

import "time"

// format methods takes 2 arguments and output a formatted string
// sample: format1(1,2.0) => "formated-int(1)-float(2.00)" the float number shall always have 2 digits after decimal point.

func format1(intValue int64, floatValue float64) string {
	panic("implement me")
	return ""
}

func format2(intValue int64, floatValue float64) string {
	panic("implement me")
	return ""
}

func format3(intValue int64, floatValue float64) string {
	panic("implement me")
	return ""
}

func main() {
	intValue := int64(100)
	floatValue := 3.1415926535897932384626
	strValue := "formated-int(100)-float(3.14)"
	var s1, s2, s3 string
	var t1, t2, t3 int64
	{
		ts := time.Now().UnixNano()
		s1 = format1(intValue, floatValue)
		t1 = time.Now().UnixNano() - ts
	}
	{
		ts := time.Now().UnixNano()
		s2 = format2(intValue, floatValue)
		t2 = time.Now().UnixNano() - ts
	}
	{
		ts := time.Now().UnixNano()
		s3 = format3(intValue, floatValue)
		t3 = time.Now().UnixNano() - ts
	}
	if s1 != strValue || s2 != strValue || s3 != strValue {
		panic("this is not right...")
	}
	println("t1:", t1)
	println("t2:", t2)
	println("t3:", t3)

}
