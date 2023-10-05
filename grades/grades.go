package grades

import (
	"fmt"
	"sync"
)

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

var (
	students Students

	// 保证并发访问安全
	studentsMutex sync.Mutex
)

func (ss Students) GetByID(id int) (*Student, error) {
	for i := range ss {
		if ss[i].ID == id {
			return &ss[i], nil
		}
	}
	return nil, fmt.Errorf("student with ID %d not fount", id)
}

type Grade struct {
	Title string
	Score float32
}
