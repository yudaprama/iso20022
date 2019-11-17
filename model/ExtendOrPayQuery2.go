package model

// Undertaking extend or pay query details.
type ExtendOrPayQuery2 struct {

	// Details related to the identification of the undertaking.
	UndertakingIdentification *Undertaking9 `xml:"UdrtkgId"`

	// Details related to the demand.
	DemandDetails *Demand4 `xml:"DmndDtls"`

	// Processing status reported by the applicant.
	Status *DemandStatus1Code `xml:"Sts"`
}

func (e *ExtendOrPayQuery2) AddUndertakingIdentification() *Undertaking9 {
	e.UndertakingIdentification = new(Undertaking9)
	return e.UndertakingIdentification
}

func (e *ExtendOrPayQuery2) AddDemandDetails() *Demand4 {
	e.DemandDetails = new(Demand4)
	return e.DemandDetails
}

func (e *ExtendOrPayQuery2) SetStatus(value string) {
	e.Status = (*DemandStatus1Code)(&value)
}
