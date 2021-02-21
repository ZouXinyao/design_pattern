package chain

import "fmt"

type department interface {
	Execute(*Patient)
	AddNext(department)
}

type Reception struct {
	next department
}

func NewReception() *Reception {
	return &Reception{}
}

func (r *Reception) Execute(p *Patient) {
	if p.registrationDone {
		fmt.Println("Patient registration already done")
		r.next.Execute(p)
		return
	}
	fmt.Println("Reception registering patient")
	p.registrationDone = true
	r.next.Execute(p)
}

func (r *Reception) AddNext(next department) {
	r.next = next
}

type Doctor struct {
	next department
}

func NewDoctor() *Doctor {
	return &Doctor{}
}

func (d *Doctor) NeedMedicine(p *Patient) {
	p.medicine.needOrNot = true
}

func (d *Doctor) NotNeedMedicine(p *Patient) {
	p.medicine.needOrNot = false
}


func (d *Doctor) Execute(p *Patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.next.Execute(p)
		return
	}
	fmt.Println("Doctor checking patient")
	p.doctorCheckUpDone = true
	d.next.Execute(p)
}

func (d *Doctor) AddNext(next department) {
	d.next = next
}

type Pay struct {
	next department
}

func NewPay() *Pay {
	return &Pay{}
}

func (y *Pay) Execute(p *Patient) {
	if p.medicine.needOrNot {
		if p.medicine.paymentDone {
			fmt.Println("Payment Done")
		} else {
			p.medicine.paymentDone = true
			fmt.Println("Paying Now")
		}
	} else {
		fmt.Println("This patient doesn't need medicine")
	}
	y.next.Execute(p)
}

func (y *Pay) AddNext(next department) {
	y.next = next
}

type Medical struct {
	next department
}

func NewMedical() *Medical {
	return &Medical{}
}

func (m *Medical) Execute(p *Patient) {
	if p.medicine.needOrNot {
		if p.medicine.medicineDone {
			fmt.Println("Medicine already given to patient")
		} else {
			p.medicine.medicineDone = true
			fmt.Println("Medical giving medicine to patient")
		}
	} else {
		fmt.Println("This patient doesn't need medicine")
	}
}

func (m *Medical) AddNext(next department) {
	m.next = next
}