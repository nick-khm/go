package cert

import "testing"

func TestValidCertData(t *testing.T) {
	c, err := New("Golang", "Bob", "2018-05-31")
	if err != nil {
		t.Errorf("Cert data should be valid. err=%v", err)
	}
	if c == nil {
		t.Errorf("Cert should be a valid reference, got=nil")
	}

	if c.Course != "GOLANG COURSE" {
		t.Errorf("Course name is not valid; expected='GOLANG COURSE', got=%v", c.Course)
	}
}

func TestCourseEmptyValue(t *testing.T) {
	_, err := New("", "Bob", "2018-05-31")
	if err == nil {
		t.Errorf("Error should be returned on an empty course")
	}
}

func TestCourseTooLong(t *testing.T) {
	course := "azesjfozijfozidfijzdjfzjfizjfizjfizjdfizjfiglgggloremazesjfozijfozidfijzdjfzjfizjfizjfizjdfizjfiglgggloremazesjfozijfozidfijzdjfzjfizjfizjfizjdfizjfiglgggloremazesjfozijfozidfijzdjfzjfizjfizjfizjdfizjfiglgggloremazesjfozijfozidfijzdjfzjfizjfizjfizjdfizjfiglgggloremazesjfozijfozidfijzdjfzjfizjfizjfizjdfizjfiglggglorem"
	_, err := New(course, "Bob", "2021-11-14")
	if err == nil {
		t.Errorf("Error should be returned on a too long course (course=%s)", course)
	}
}

func TestNameEmptyValue(t *testing.T) {
	_, err := New("course", "", "2021-11-15")
	if err == nil {
		t.Errorf("Error should be raised when the name of a student is empty")
	}
}

func TestNameTooLong(t *testing.T) {
	name := "verylongnameofstudentverylongnameofstudentverylongnameofstudent"
	_, err := New("course", name, "2021-11-15")
	if err == nil {
		t.Errorf("Error should be raised when the name of student is too long. Got=(%s)", name)
	}
}
