package main

import "fmt"

/*
一个医院应用的责任链模式例子。 医院中会有多个部门， 如：
	前台
	医生
	药房
	收银
病人来访时， 首先都会去前台， 然后是看医生、 取药， 最后结账。
*/

type DepartMent interface {
	Execute(patinet *Patient)
	Next(department DepartMent)
}

type Patient struct {
	name       string
	registered bool
	doctor     bool
	medicine   bool
	payment    bool
}

//前台接待，病人挂号
type Reception struct {
	next DepartMent
}

func (r *Reception) Execute(patinet *Patient) {
	if patinet.registered {
		fmt.Println("Patient registration already done")
		r.next.Execute(patinet)
		return
	}
	fmt.Println("Reception registering patient")
	patinet.registered = true

	r.next.Execute(patinet)
}

func (r *Reception) Next(next DepartMent) {
	r.next = next
}

// 医生看诊
type Doctor struct {
	next DepartMent
}

func (d *Doctor) Execute(patinet *Patient) {
	if patinet.doctor {
		fmt.Println("Patient docker already done")
		d.next.Execute(patinet)
		return
	}
	fmt.Println("Docker checking patient ")
	d.next.Execute(patinet)
	patinet.doctor = true
}

func (d *Doctor) Next(next DepartMent) {
	d.next = next
}

//机器治疗病人
type Medical struct {
	next DepartMent
}

func (m *Medical) Execute(p *Patient) {
	if p.medicine {
		fmt.Println("Medicine already given to patient")
		m.next.Execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.medicine = true
	m.next.Execute(p)
}

func (m *Medical) Next(next DepartMent) {
	m.next = next
}

//缴费处收银
type Cashier struct {
	next DepartMent
}

func (c *Cashier) Execute(p *Patient) {
	if p.payment {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient patient")
}

func (c *Cashier) Next(next DepartMent) {
	c.next = next
}

func main() {

	cashier := &Cashier{}

	//Set next for medical department
	medical := &Medical{}
	medical.Next(cashier)

	//Set next for doctor department
	doctor := &Doctor{}
	doctor.Next(medical)

	//Set next for reception department
	reception := &Reception{}
	reception.Next(doctor)

	p := Patient{name: "tom"}
	reception.Execute(&p)

	//output:
	/*
		Reception registering patient
		Docker checking patient
		Medical giving medicine to patient
		Cashier getting money from patient patient
	*/
}
