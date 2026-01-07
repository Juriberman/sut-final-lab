package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestNegativeEmployeeCode(t *testing.T) {
	g := NewGomegaWithT(t)

	employee := Employees{
		Name:         "Hello",
		Salary:       50000,
		EmployeeCode: "gg1234",
	}

	ok, err := govalidator.ValidateStruct(employee)

	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("EmployeeCode must be 2 uppercase English letters (A-Z) followed by ‘-’ and 4 digits (0-9)"))
}
