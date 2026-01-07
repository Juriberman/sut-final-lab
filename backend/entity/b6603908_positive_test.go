package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestValid(t *testing.T) {
	g := NewGomegaWithT(t)

	employee := Employees{
		Name:         "Hello",
		Salary:       50000,
		EmployeeCode: "GG-1234",
	}

	ok, err := govalidator.ValidateStruct(employee)

	g.Expect(ok).To(BeTrue())
	g.Expect(err).To(BeNil())
}
