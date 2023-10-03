package business

import "fmt"

type Student struct {
	ID     int
	Name   string
	Grades []Grade
}

func (s Student) AverageScore() float32 {
	var r float32
	for _, grade := range s.Grades {
		r += grade.Score
	}
	return r / float32(len(s.Grades))
}

type Students []Student

var students Students

func (ss Students) GetById(id int) (*Student, error) {
	for i := range ss {
		if ss[i].ID == id {
			return &ss[i], nil
		}
	}
	return nil, fmt.Errorf("student with ID %d not fount.", id)
}

type Grade struct {
	Title string
	Score float32
}
