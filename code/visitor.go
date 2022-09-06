package main

import "fmt"

type Customer interface {
	Accept(Visitor)
}

type Visitor interface {
	Visit(Customer)
}

type EnterpriseCustomer struct {
	name string
}

type CustomerCommon struct {
	customers []Customer
}

func (c *CustomerCommon) Add(customer Customer) {
	c.customers = append(c.customers, customer)
}

func (c *CustomerCommon) Accept(visitor Visitor) {
	for _, customer := range c.customers {
		visitor.Visit(customer)
	}
}

func NewEnterpriseCustomer(name string) *EnterpriseCustomer {
	return &EnterpriseCustomer{
		name: name,
	}
}

func (e *EnterpriseCustomer) Accept(visitor Visitor) {
	visitor.Visit(e)
}

type IndividualCustomer struct {
	name string
}

func NewIndividualCustomer(name string) *IndividualCustomer {
	return &IndividualCustomer{
		name: name,
	}
}

func (i *IndividualCustomer) Accept(visitor Visitor) {
	visitor.Visit(i)
}

type ServiceRequestVisitor struct {
}

func (*ServiceRequestVisitor) Visit(customer Customer) {
	switch c := customer.(type) {
	case *IndividualCustomer:
		fmt.Printf("serving Individual customer %s\n", c.name)
	case *EnterpriseCustomer:
		fmt.Printf("serving enterprise customer %s\n", c.name)
	}
}

// only for enterprise
type AnalysisVisitor struct{}

func (*AnalysisVisitor) Visit(customer Customer) {
	switch c := customer.(type) {
	case *EnterpriseCustomer:
		fmt.Printf("analysis enterprise customer %s\n", c.name)
	}
}

func TestRequestVisitor() {
	c := &CustomerCommon{}
	c.Add(NewIndividualCustomer("bob"))
	c.Add(NewEnterpriseCustomer("A company"))
	c.Add(NewEnterpriseCustomer("B company"))
	c.Accept(&ServiceRequestVisitor{})
}

func TestAnalysisVisitor() {
	c := &CustomerCommon{}
	c.Add(NewIndividualCustomer("bob"))
	c.Add(NewEnterpriseCustomer("A company"))
	c.Add(NewEnterpriseCustomer("B company"))

	c.Accept(&AnalysisVisitor{})
}

func main() {
	TestRequestVisitor()
	TestAnalysisVisitor()
}
