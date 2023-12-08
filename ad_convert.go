type Data struct {
	a, b, c, d bool
}

func (v Data) Get() {
	fmt.Println(v)
}

func (v *Data) Set(a, b, c, d bool) {
	/*
		input will by float number
		I function will test if it is lowwer than 7,75
		Than convert it to 5 digit binary value, where:
		- the first three digits contains real number
		- the last two digits express decimalpart

		Detailed conversion table is bellow

		0		000			00   ... 0,0
		1		001         01   ... 0,25
		2		010         10   ... 0,5
		3		011         11   ... 0,75
		4		100
		5		101
		6		110
		7		111

		Example:
		 3,7  = 011 11
		 4.3  = 100 01
		 ...
	*/
	v.a = a
	v.b = b
	v.c = c
	v.d = d
}


func main() {
	x := Data{true, false, false, true}
	x.Get()
	x.Set(false, false, true, false)
	x.Get()
	var y Data
	y.Set(true, false, true, false)
	y.Get()
}
