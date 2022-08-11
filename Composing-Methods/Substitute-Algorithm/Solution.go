package Substitute_Algorithm

import "gomodules.xyz/x/arrays"

// Solution: Replace the body of the method that implements the algorithm with a new algorithm.

func foundPersonSol(people []string) string {
	candidates := []string{"Don", "John", "Kent"}

	for i := range people {
		found, _ := arrays.Contains(candidates, people[i])
		if found {
			return people[i]
		}
	}

	return ""
}
