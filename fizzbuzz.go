package main

func fizzbuzz(x int) interface{} {
	if x%3 == 0 && x%5 == 0 {
		return "fizzbuz"
	}
	if x%3 == 0 {
		return "fizz"
	}
	if x%5 == 0 {
		return "buzz"
	}

	return x
}

func fizzbuzSwitch(x int) interface{} {
	//if the cases are coerced to a bool, the expresion can be avoided
	switch {
	case x%3 == 0 && x%5 == 0:
		return "fizzbuz"
	case x%3 == 0:
		return "fizz"
	case x%5 == 0:
		return "buzz"
	default:
		return x
	}
}
