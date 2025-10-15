package migratory_birds

func migratoryBirds(arr []int32) int32 {
	birdIds := make(map[int32]int)
	id, count := 0, 0

	for _, birdId := range arr {
		birdIds[birdId]++
		occurrences := birdIds[birdId]

		if id == 0 || (occurrences > count && int(birdId) != id) || (occurrences == count && int(birdId) < id) {
			id = int(birdId)
		}

		if count == 0 || occurrences > count {
			count = occurrences
		}
	}

	return int32(id)
}
