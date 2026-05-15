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

func letterGrade (avg float64) string {
	if avg >= 70 {
		return "A"
	} else if avg >= 60 {
		return "B"
	} else if avg >= 50 {
		return "C"
	} else if avg >= 40 {
		return  "D"
	} else {
		return "F"
	}
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
			Name: "Norah", Score: []float64{95, 88, 52},
		  },
	 }

		  for _, s := range students {
			avg := calculateaverage(s.Score)
			grade := letterGrade(avg)
			fmt.Printf("%s - Average:%.2f - Grade: %s\n", s.Name, avg, grade)
		  }

}