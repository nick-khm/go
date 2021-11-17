package cert

import (
	"fmt"
	"strings"
	"time"
)

var MaxLenCourse = 20
var MaxLenName = 10

type Cert struct {
	Course string
	Name   string
	Date   time.Time

	LabelTitle         string
	LabelCompletion    string
	LabelPressented    string
	LabelParticipation string
	LabelDate          string
}

type Saver interface {
	Save(c Cert) error
}

func New(course, name, date string) (*Cert, error) {
	c, err := ValidateCourse(course)
	n, err := ValidateName(name)
	d, err := parseDate(date)
	if err != nil {
		return nil, err
	}

	cert := &Cert{
		Course:             c,
		Name:               n,
		Date:               d,
		LabelTitle:         fmt.Sprintf("%v Certificate - %v", c, n),
		LabelCompletion:    "Certificate of Completion",
		LabelPressented:    "This Certificate is Presented To",
		LabelParticipation: fmt.Sprintf("For participation in the %v", c),
		LabelDate:          fmt.Sprintf("Date: %v", d.Format("02/01/2021")),
	}
	return cert, nil
}

func ValidateCourse(course string) (string, error) {
	c, err := validateStr(course, MaxLenCourse)
	if err != nil {
		return "", err
	}
	if !strings.HasSuffix(c, " course") {
		c = c + " course"
	}
	return strings.ToTitle(c), nil
}

func validateStr(str string, maxLen int) (string, error) {
	str = strings.TrimSpace(str)
	if len(str) <= 0 {
		return str, fmt.Errorf("Course can't be empty")
	} else if len(str) >= maxLen {
		return str, fmt.Errorf("Invalid string, got len=%d", len(str))
	}
	return str, nil
}

func ValidateName(name string) (string, error) {
	n, err := validateStr(name, MaxLenName)
	if err != nil {
		return "", err
	}
	return strings.ToTitle(n), nil
}

func parseDate(date string) (time.Time, error) {
	t, err := time.Parse("2021-01-02", date)
	if err != nil {
		return t, err
	}
	return t, nil
}
