package fibonacci

func FibonacciIterative(n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	prev := 1
	cur := 1

	for i := 3; i <= n; i++ {
		next := prev + cur

		prev = cur
		cur = next
	}

	return cur
}

func FibonacciRecursive(n int) int {
	var recur func(prev, cur, counter int) int
	recur = func(prev, cur, counter int) int {
		if counter > n {
			return cur
		} else {
			next := prev + cur
			prev = cur
			cur = next

			counter++

			return recur(prev, cur, counter)
		}

	}

	if n == 1 || n == 2 {
		return 1
	} else {
		counter := 3
		return recur(1, 1, counter)

	}
}
