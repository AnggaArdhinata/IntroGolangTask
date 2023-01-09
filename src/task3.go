package src

import "strconv"

var data = []int{1, 7, 3, 4, 8, 9}

func FlightDuration(duration int) string {

	result := ""

	for i , v := range data {
		for j := 0; j < i; j++ {
			if sum := v + data[j]; sum == duration {
				data1 := strconv.Itoa(v)
				data2 := strconv.Itoa(data[j])

				result += data1 + " dan " + data2 + "\n"
			}

		}
			
	}
	return result
}
	
	

