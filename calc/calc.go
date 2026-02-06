package calc

func Add(x, y int) int { //首字母大写代表公有方法，可以在其他包进行访问
	return x + y
}

func Sub(x, y int) int {
	return x - y
}
