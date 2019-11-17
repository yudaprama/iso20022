package model

// Document that a user must file with an authorized servicer for each contract that involves foreign currency transactions with non residents.
type RegisteredContract3 struct {

	// Unique and unambiguous identification of the registered contract amendment request.
	RegisteredContractAmendmentIdentification *Max35Text `xml:"RegdCtrctAmdmntId"`

	// Identification of the original contract registration, as registered with the registration agent.
	OriginalRegisteredContractIdentification *Max35Text `xml:"OrgnlRegdCtrctId"`

	// Priority requested for the amendment of the registered contract.
	Priority *Priority2Code `xml:"Prty"`

	// Amendment details of the registered contract for the registered contract.
	Contract *UnderlyingContract1Choice `xml:"Ctrct"`

	// Contract balance on date of contract registration.
	ContractBalance []*ContractBalance1 `xml:"CtrctBal,omitempty"`

	// Type of the payment schedule provided in the contract.
	PaymentScheduleType *PaymentScheduleType1Choice `xml:"PmtSchdlTp,omitempty"`

	// Further additional information on the amendment.
	AdditionalInformation *Max1025Text `xml:"AddtlInf,omitempty"`

	// Documents provided as attachments to the contract registration amendment request.
	Attachment []*DocumentGeneralInformation3 `xml:"Attchmnt,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (r *RegisteredContract3) SetRegisteredContractAmendmentIdentification(value string) {
	r.RegisteredContractAmendmentIdentification = (*Max35Text)(&value)
}

func (r *RegisteredContract3) SetOriginalRegisteredContractIdentification(value string) {
	r.OriginalRegisteredContractIdentification = (*Max35Text)(&value)
}

func (r *RegisteredContract3) SetPriority(value string) {
	r.Priority = (*Priority2Code)(&value)
}

func (r *RegisteredContract3) AddContract() *UnderlyingContract1Choice {
	r.Contract = new(UnderlyingContract1Choice)
	return r.Contract
}

func (r *RegisteredContract3) AddContractBalance() *ContractBalance1 {
	newValue := new(ContractBalance1)
	r.ContractBalance = append(r.ContractBalance, newValue)
	return newValue
}

func (r *RegisteredContract3) AddPaymentScheduleType() *PaymentScheduleType1Choice {
	r.PaymentScheduleType = new(PaymentScheduleType1Choice)
	return r.PaymentScheduleType
}

func (r *RegisteredContract3) SetAdditionalInformation(value string) {
	r.AdditionalInformation = (*Max1025Text)(&value)
}

func (r *RegisteredContract3) AddAttachment() *DocumentGeneralInformation3 {
	newValue := new(DocumentGeneralInformation3)
	r.Attachment = append(r.Attachment, newValue)
	return newValue
}

func (r *RegisteredContract3) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	r.SupplementaryData = append(r.SupplementaryData, newValue)
	return newValue
}
