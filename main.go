package main

import "fmt"


type Student struct {
	Name   string
	Score  []float64
}

func calculateaverage(scores []float64) float64 {
	total := 0.0
	for _, score := range scores {
		total += score
	}
	return total / float64(len(scores))
}

func main() {
     students := []Student{
          {
			Name: "kamsy", Score: []float64{80, 90, 75},
		  },

	      {
			Name: "Andre", Score: []float64{60, 70, 85},
		  },

		  {
			Name: "Nora", Score: []float64{95, 88, 52},
		  },
	 }

		  for _, s := range students {
			avg
			fmt.Println(s.Name, s.Score)
		  }

}