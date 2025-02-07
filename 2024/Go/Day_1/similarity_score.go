package main

func calculateSimilarity(column1 []int, column2 []int) int {
	similarityScore := 0
	frequencyMap := make(map[int]int)

	// iterate through the first list
	for i := 0; i < len(column1); i++ {
		value := column1[i]
		// check how many times the element is in column2
		// if it exists calculate the score based on the map
		_, exists := frequencyMap[value]
		if exists {
			similarityScore += (value * frequencyMap[value])
		} else {
			freq := containsElement(value, column2)
			frequencyMap[value] = freq
			similarityScore += (value * freq)
		}
	}

	return similarityScore
}

func containsElement(value int, column2 []int) int {
	freq := 0
	for i := 0; i < len(column2); i++ {
		if column2[i] == value {
			freq += 1
		}
	}

	return freq
}
