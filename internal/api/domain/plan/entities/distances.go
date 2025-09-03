package entities

var (
	HalfMarathonLongRuns12Week = []float32{6.00, 7.00, 8.00, 9.00, 10.00, 11.50, 13.00, 14.50, 16.00, 17.00, 13.00}
	HalfMarathonEasyRuns12Week = []float32{4.00, 4.00, 4.00, 5.00, 5.00, 6.00, 6.00, 7.00, 5.00, 5.00, 3.00}
	HalfMarathonLongRuns8Week  = []float32{7.00, 8.50, 10.00, 12.00, 15.00, 17.00, 13.00}
	HalfMarathonEasyRuns8Week  = []float32{4.00, 5.00, 5.00, 5.00, 6.00, 4.00, 4.00}
)

var HalfMarathonLongDistancesMap = map[int][]float32{
	8:  HalfMarathonLongRuns8Week,
	12: HalfMarathonLongRuns12Week,
}

var HalfMarathonEasyDistancesMap = map[int][]float32{
	8:  HalfMarathonEasyRuns8Week,
	12: HalfMarathonEasyRuns12Week,
}
