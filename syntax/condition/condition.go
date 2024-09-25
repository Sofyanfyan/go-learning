package condition

func Condition() {
	// If statements are used to perform different actions based on different conditions.
	// The if statement executes a block of code if a specified condition is true

	// Example:
	var x int = 10
	var y int = 5
	if x > y {
		println("n is greater than y")
	} else {
		println("n is less than y")
	}

	// Switch statements are used to select one of many code blocks to be executed.

	// Example:

	switch x {

	case 1, 2, 3, 4, 5:
		println("weekdays")
	case 6, 7:
		println("weekends")

	default:
		println("invalid day")
	}
}