package Substitute_Algorithm

// So you want to replace an existing algorithm with a new one?

// Why Refactor:
// Sometimes a method is so cluttered with issues that it’s easier to tear down the method and start fresh.
//And perhaps you have found an algorithm that’s much simpler and more efficient.
//If this is the case, you should simply replace the old algorithm with the new one.

func foundPerson(people []string) string {
	for i := range people {
		if people[i] == "Don" {
			return "Don"
		}
		if people[i] == "John" {
			return "John"
		}
		if people[i] == "Kent" {
			return "Kent"
		}
	}

	return "";
}