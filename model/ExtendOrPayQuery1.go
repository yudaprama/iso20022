package model

// Undertaking extend or pay query details.
type ExtendOrPayQuery1 struct {

	// Details related to the identification of the undertaking.
	UndertakingIdentification *Undertaking9 `xml:"UdrtkgId"`

	// Details related to the demand.
	DemandDetails *Demand2 `xml:"DmndDtls"`

	// Requested new expiry date as an alternative to payment of the demand.
	RequestedExpiryDate *ISODate `xml:"ReqdXpryDt"`

	// Details of the instructions from the bank.
	BankInstructions *BankInstructions1 `xml:"BkInstrs,omitempty"`

	// Contact at the issuing bank.
	BankContact []*Contacts3 `xml:"BkCtct,omitempty"`

	// Document or template enclosed in the request.
	EnclosedFile []*Document9 `xml:"NclsdFile,omitempty"`

	// Additional information related to the request.
	AdditionalInformation []*Max2000Text `xml:"AddtlInf,omitempty"`
}

func (e *ExtendOrPayQuery1) AddUndertakingIdentification() *Undertaking9 {
	e.UndertakingIdentification = new(Undertaking9)
	return e.UndertakingIdentification
}

func (e *ExtendOrPayQuery1) AddDemandDetails() *Demand2 {
	e.DemandDetails = new(Demand2)
	return e.DemandDetails
}

func (e *ExtendOrPayQuery1) SetRequestedExpiryDate(value string) {
	e.RequestedExpiryDate = (*ISODate)(&value)
}

func (e *ExtendOrPayQuery1) AddBankInstructions() *BankInstructions1 {
	e.BankInstructions = new(BankInstructions1)
	return e.BankInstructions
}

func (e *ExtendOrPayQuery1) AddBankContact() *Contacts3 {
	newValue := new(Contacts3)
	e.BankContact = append(e.BankContact, newValue)
	return newValue
}

func (e *ExtendOrPayQuery1) AddEnclosedFile() *Document9 {
	newValue := new(Document9)
	e.EnclosedFile = append(e.EnclosedFile, newValue)
	return newValue
}

func (e *ExtendOrPayQuery1) AddAdditionalInformation(value string) {
	e.AdditionalInformation = append(e.AdditionalInformation, (*Max2000Text)(&value))
}
