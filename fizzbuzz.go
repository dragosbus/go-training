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
