package main

import (
	"fmt"
	"github.com/ZouXinyao/design_pattern/code_go/15_chain/chain"
)

func main() {
	reception := chain.NewReception()
	doctor := chain.NewDoctor()
	pay := chain.NewPay()
	medical := chain.NewMedical()

	reception.AddNext(doctor)
	doctor.AddNext(pay)
	pay.AddNext(medical)

	patient01 := chain.NewPatient("num01")
	process01 := reception
	doctor.NotNeedMedicine(patient01)
	process01.Execute(patient01)
	fmt.Println("Have Seen a doctor: ", patient01.HavenSeenADoctor())

	patient02 := chain.NewPatient("num02")
	process02 := reception
	doctor.NeedMedicine(patient02)
	process02.Execute(patient02)
	fmt.Println("Have Seen a doctor: ", patient02.HavenSeenADoctor())



}
