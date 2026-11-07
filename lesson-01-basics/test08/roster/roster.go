package roster

import "strings"
import "fmt"

type Student struct {
	ID    string
	Name  string
	Score int
}

type Roster struct {
	students    []Student
	comparisons int
}

func (r *Roster) Add(s Student) {
	r.students = append(r.students, s)
}

func (r *Roster) FindByID(id string) (Student, bool) {
	for _, s := range r.students {
		r.comparisons++
		if s.ID == id {
			return s, true
		}
	}
	return Student{}, false
}

func (r *Roster) TotalScoreByName(name string) int {
	sum := 0
	for _, s := range r.students {
		r.comparisons++
		if s.Name == name {
			sum += s.Score
		}
	}
	return sum
}

func FormatText(r *Roster) string {
	var builder strings.Builder
	for _, s := range r.students {
		line := fmt.Sprintf("%s %s %d\n", s.ID, s.Name, s.Score)
		builder.WriteString(line)
	}
	return builder.String()
}

func FormatCSV(r *Roster) string {
	var builder strings.Builder
	builder.WriteString("id,name,score\n")
	for _, s := range r.students {
		line := fmt.Sprintf("%s,%s,%d\n", s.ID, s.Name, s.Score)
		builder.WriteString(line)
	}
	return builder.String()
}


