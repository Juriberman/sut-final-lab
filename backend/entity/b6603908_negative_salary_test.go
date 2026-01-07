package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestNegativeSalary(t *testing.T) {
	g := NewGomegaWithT(t)

	employee := Employees{
		Name:         "Hello",
		Salary:       100,
		EmployeeCode: "GG-1234",
	}

	ok, err := govalidator.ValidateStruct(employee)

	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("Salary must be between 15000 and 200000"))
}
