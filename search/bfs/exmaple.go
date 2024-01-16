/*
input:
{
  "performers":[{"score":99.9},{"score":3.1},{"score":1.0}],
  "venues":[{"score":88.1},{"score":12.43}],
  "events":[{"score":100.0},{"score":90.9},{"score":11.32}]
}

output:
{
  "performers":[{"score":99.9, "sort_order": 1},{"score":3.1},{"score":1.0}],
  "venues":[{"score":88.1, "sort_order": 3},{"score":12.43}],
  "events":[{"score":100.0, "sort_order": 0},{"score":90.9, "sort_order": 2},{"score":11.32}]
}


three array of inputs [performers, venues, events]
new_array = base_on_score_rate descendent
performers =[{"score": 88.32}, {"score": 99.9}]

1. arrays is already sorted
2. size is not the same
3. add sort_order based on the global ranking


ScoreStc {
  Score
  SortOrder
}

IntermediateStruct struct {
  ScoreStc
  Datatype string
}

func assignSortOrder(performers []ScoreStc, venues []ScoreStc, events []ScoreStc) {
  bigArray = []IntermediateStruct -> size = len(performers) + len(venues) + len(events)
  // loop through performers and add to bigArray
  // loop through venus and add to bigArray
  // loop through events and add to bigArray

  // sort bigArray

  // loop through bigArray and create output object

}

var memRanking []Score

func saveInMem(input []Score, origin string) {
  if len(memRanking)
    saveAllItemsInMemRanking

  for score in scores

}

performers = [0, 1, 2, 3, 4, 5, ..., n] => mem_ranking
venues = [perf_0, ven_0, ...] => mem_ranking
event = []

[0, 1, 2]


*/

package bfs

import (
	"sort"
)

type Score struct {
	Score     float64
	SortOrder int
	DataType  string
}

func assignCombinedSortOrder(performers, venues, events []Score) {
	allScores := append(append([]Score{}, performers...), append(venues[:], events[:]...)...) // Combine all scores

	// Sort the combined scores based on the Score field
	sort.Slice(allScores, func(i, j int) bool {
		return allScores[i].Score > allScores[j].Score // Sorting in descending order
	})

	// Assign SortOrder based on the combined sorted order
	for i, score := range allScores {
		if scoreInSlice(score, performers) {
			for j, performer := range performers {
				if score == performer {
					performers[j].SortOrder = i
					break
				}
			}
		} else if scoreInSlice(score, venues) {
			for j, venue := range venues {
				if score == venue {
					venues[j].SortOrder = i
					break
				}
			}
		} else {
			for j, event := range events {
				if score == event {
					events[j].SortOrder = i
					break
				}
			}
		}
	}
}

func scoreInSlice(score Score, scores []Score) bool {
	for _, s := range scores {
		if s == score {
			return true
		}
	}
	return false
}
