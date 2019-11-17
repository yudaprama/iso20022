package model

// Information about the refusal of a demand.
type DemandRefusal1 struct {

	// Details related to the identification of the undertaking.
	UndertakingIdentification *Undertaking9 `xml:"UdrtkgId"`

	// Unique and unambiguous identifier assigned by the advising party to the undertaking.
	AdvisingPartyReferenceNumber *Max35Text `xml:"AdvsgPtyRefNb,omitempty"`

	// Unique and unambiguous identifier assigned by the second advising party to the undertaking.
	SecondAdvisingPartyReferenceNumber *Max35Text `xml:"ScndAdvsgPtyRefNb,omitempty"`

	// Unique and unambiguous identifier assigned by the confirmer to the undertaking.
	ConfirmerReferenceNumber *Max35Text `xml:"CnfrmrRefNb,omitempty"`

	// Details related to the demand.
	DemandDetails *Demand2 `xml:"DmndDtls"`

	// Expicit indication of 'REFUSED' as the processing status reported by the issuer.
	Status *Refused7Text `xml:"Sts"`

	// Details related to the discrepancies.
	Discrepancy []*Discrepancy1 `xml:"Dscrpncy,omitempty"`

	// Indication of how the demand presentation documents will be handled as a consequence of the demand refusal.
	DispositionOfDocuments []*Max2000Text `xml:"DspstnOfDocs,omitempty"`

	// Additional information related to the notification.
	AdditionalInformation []*Max2000Text `xml:"AddtlInf,omitempty"`
}

func (d *DemandRefusal1) AddUndertakingIdentification() *Undertaking9 {
	d.UndertakingIdentification = new(Undertaking9)
	return d.UndertakingIdentification
}

func (d *DemandRefusal1) SetAdvisingPartyReferenceNumber(value string) {
	d.AdvisingPartyReferenceNumber = (*Max35Text)(&value)
}

func (d *DemandRefusal1) SetSecondAdvisingPartyReferenceNumber(value string) {
	d.SecondAdvisingPartyReferenceNumber = (*Max35Text)(&value)
}

func (d *DemandRefusal1) SetConfirmerReferenceNumber(value string) {
	d.ConfirmerReferenceNumber = (*Max35Text)(&value)
}

func (d *DemandRefusal1) AddDemandDetails() *Demand2 {
	d.DemandDetails = new(Demand2)
	return d.DemandDetails
}

func (d *DemandRefusal1) SetStatus(value string) {
	d.Status = (*Refused7Text)(&value)
}

func (d *DemandRefusal1) AddDiscrepancy() *Discrepancy1 {
	newValue := new(Discrepancy1)
	d.Discrepancy = append(d.Discrepancy, newValue)
	return newValue
}

func (d *DemandRefusal1) AddDispositionOfDocuments(value string) {
	d.DispositionOfDocuments = append(d.DispositionOfDocuments, (*Max2000Text)(&value))
}

func (d *DemandRefusal1) AddAdditionalInformation(value string) {
	d.AdditionalInformation = append(d.AdditionalInformation, (*Max2000Text)(&value))
}
