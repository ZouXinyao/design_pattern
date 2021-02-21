package chain

type Patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicine          needMedicine
}

type needMedicine struct {
	needOrNot         bool
	paymentDone		  bool
	medicineDone      bool
}

func NewPatient(name string) *Patient {
	return &Patient{
		name: name,
	}
}

func (p *Patient) HavenSeenADoctor() bool {
	if p.doctorCheckUpDone {
		if p.medicine.needOrNot {
			return p.medicine.medicineDone
		}
		return true
	}
	return false
}
