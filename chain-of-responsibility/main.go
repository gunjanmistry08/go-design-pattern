package main

import "fmt"

type Candidate struct {
	Name  string
	Marks float64
}

type Shift struct {
	ShiftName  string
	Candidates []Candidate
	Mean       float64
}

//Calculate Mean Function
func (s *Shift) CalculateMean() {
	var sum float64
	for _, candidate := range s.Candidates {
		sum += candidate.Marks
	}
	s.Mean = sum / float64(len(s.Candidates))
}

func (s *Shift) GetMaxMarks() float64 {
	var maxMarks float64
	for _, candidate := range s.Candidates {
		if candidate.Marks > maxMarks {
			maxMarks = candidate.Marks
		}
	}
	return maxMarks
}

// Stringer function for shift
func (s Shift) String() string {
	return fmt.Sprintf("Shift Name: %s, Mean: %f, Candidates: %v\n", s.ShiftName, s.Mean, s.Candidates)
}

type ResponsibilityFunc func([]Shift) []Shift

// Respomsibility Structure
type Responsibility interface {
	SetNext(ResponsibilityFunc)
	ResponsibilityWork([]Shift) []Shift
}

type BaseResponsibility struct {
	nextResponsibility ResponsibilityFunc
}

func (r *BaseResponsibility) SetNext(responsibility ResponsibilityFunc) {
	r.nextResponsibility = responsibility
}

// Highest Present Count Function
type HighestPresentCount struct {
	BaseResponsibility
}

func (r *HighestPresentCount) ResponsibilityWork(shifts []Shift) []Shift {
	var maxCount int
	var maxShift []Shift
	for _, shift := range shifts {
		if len(shift.Candidates) > maxCount {
			maxCount = len(shift.Candidates)
			maxShift = []Shift{shift}
		} else if len(shift.Candidates) == maxCount {
			maxShift = append(maxShift, shift)
		}
	}
	if len(maxShift) == 1 {
		return maxShift

	} else {
		return r.nextResponsibility(maxShift)
	}
}

// Highest Individual Marks Function
type HighestIndividualMarks struct {
	BaseResponsibility
}

func (r *HighestIndividualMarks) ResponsibilityWork(shifts []Shift) []Shift {
	var maxMarks float64
	var maxShift []Shift
	for _, shift := range shifts {
		if shift.GetMaxMarks() > maxMarks {
			maxMarks = shift.GetMaxMarks()
			maxShift = []Shift{shift}
		} else if shift.GetMaxMarks() == maxMarks {
			maxShift = append(maxShift, shift)
		}
	}
	if len(maxShift) == 1 {
		return maxShift

	} else {
		return r.nextResponsibility(maxShift)
	}
}

// Highest Mean Function
type HighestMeanCount struct {
	BaseResponsibility
}

func (r *HighestMeanCount) ResponsibilityWork(shifts []Shift) []Shift {
	var maxMean float64
	var maxShift []Shift
	for _, shift := range shifts {
		if shift.Mean > maxMean {
			maxMean = shift.Mean
			maxShift = []Shift{shift}
		} else if shift.Mean == maxMean {
			maxShift = append(maxShift, shift)
		}
	}
	if len(maxShift) == 1 {
		return maxShift

	} else {
		fmt.Println("Mean is same for multiple shifts")
		return r.nextResponsibility(maxShift)
	}
}

func main() {

	fmt.Println("Hello, World!")

	candidates1 := []Candidate{
		{Name: "Alice", Marks: 85.5},
		{Name: "Bob", Marks: 78.0},
		{Name: "Charlie", Marks: 92.3},
		{Name: "Diana", Marks: 88.7},
		{Name: "Luffy", Marks: 88.7},
	}

	// Assign candidates to a shift
	shift1 := Shift{
		ShiftName:  "Morning Shift",
		Candidates: candidates1,
	}

	shift1.CalculateMean()

	candidates2 := []Candidate{
		{Name: "Eve", Marks: 90.1},
		{Name: "Frank", Marks: 82.4},
		{Name: "Grace", Marks: 87.6},
		{Name: "Hank", Marks: 79.5},
	}

	shift2 := Shift{
		ShiftName:  "Evening Shift",
		Candidates: candidates2,
	}

	shift2.CalculateMean()

	candidates3 := []Candidate{
		{Name: "Ivy", Marks: 86.5},
		{Name: "Jack", Marks: 77.0},
		{Name: "Ken", Marks: 93.3},
		{Name: "Liam", Marks: 89.76},
	}

	shift3 := Shift{
		ShiftName:  "Night Shift",
		Candidates: candidates3,
	}

	shift3.CalculateMean()

	highestMean := HighestMeanCount{}
	highestIndividualMarks := HighestIndividualMarks{}
	highestPresentCount := HighestPresentCount{}

	highestMean.SetNext(highestIndividualMarks.ResponsibilityWork)
	highestIndividualMarks.SetNext(highestPresentCount.ResponsibilityWork)

	shifts := []Shift{shift1, shift2, shift3}

	fmt.Println(shifts)
	fmt.Println(highestMean.ResponsibilityWork(shifts))
}
