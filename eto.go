package main

var eto [12]string

func init() {
	eto = [12]string{
		"申年",
		"酉年",
		"戌年",
		"亥年",
		"子年",
		"丑年",
		"寅年",
		"卯年",
		"辰年",
		"巳年",
		"午年",
		"未年",
	}

}

func Calc(y uint64) string {
	return eto[y%12]
}
