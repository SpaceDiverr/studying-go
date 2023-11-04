package intmath

type Sample struct {
	Box Item
}

type Item struct {
	Contents []string
}
func Abs(n int) (int) {
	if n < 0 {return -n}
	return n
}

func Pow(n, pow int) (res int ) {
	if n & pow == 1 || pow == 0 { return 1 }
	if n == 0 { return 0 }

	res = 1
	for i := 0; i < pow; i++ {
		res *= n
	}
	return
}




