package model

// Set of elements used to provide further details related to a direct debit mandate signed between the creditor and the debtor.
type MandateRelatedInformation1 struct {

	// Reference of the direct debit mandate that has been signed between by the debtor and the creditor.
	MandateIdentification *Max35Text `xml:"MndtId,omitempty"`

	// Date on which the direct debit mandate has been signed by the debtor.
	DateOfSignature *ISODate `xml:"DtOfSgntr,omitempty"`

	// Indicator notifying whether the underlying mandate is amended or not.
	AmendmentIndicator *TrueFalseIndicator `xml:"AmdmntInd,omitempty"`

	// List of direct debit mandate elements that have been modified.
	AmendmentInformationDetails *AmendmentInformationDetails1 `xml:"AmdmntInfDtls,omitempty"`

	// Additional security provisions, such as a digital signature, as provided by the debtor.
	ElectronicSignature *Max1025Text `xml:"ElctrncSgntr,omitempty"`

	// Date of the first collection of a direct debit as per the mandate.
	FirstCollectionDate *ISODate `xml:"FrstColltnDt,omitempty"`

	// Date of the final collection of a direct debit as per the mandate.
	FinalCollectionDate *ISODate `xml:"FnlColltnDt,omitempty"`

	// Regularity with which direct debit instructions are to be created and processed.
	Frequency *Frequency1Code `xml:"Frqcy,omitempty"`
}

func (m *MandateRelatedInformation1) SetMandateIdentification(value string) {
	m.MandateIdentification = (*Max35Text)(&value)
}

func (m *MandateRelatedInformation1) SetDateOfSignature(value string) {
	m.DateOfSignature = (*ISODate)(&value)
}

func (m *MandateRelatedInformation1) SetAmendmentIndicator(value string) {
	m.AmendmentIndicator = (*TrueFalseIndicator)(&value)
}

func (m *MandateRelatedInformation1) AddAmendmentInformationDetails() *AmendmentInformationDetails1 {
	m.AmendmentInformationDetails = new(AmendmentInformationDetails1)
	return m.AmendmentInformationDetails
}

func (m *MandateRelatedInformation1) SetElectronicSignature(value string) {
	m.ElectronicSignature = (*Max1025Text)(&value)
}

func (m *MandateRelatedInformation1) SetFirstCollectionDate(value string) {
	m.FirstCollectionDate = (*ISODate)(&value)
}

func (m *MandateRelatedInformation1) SetFinalCollectionDate(value string) {
	m.FinalCollectionDate = (*ISODate)(&value)
}

func (m *MandateRelatedInformation1) SetFrequency(value string) {
	m.Frequency = (*Frequency1Code)(&value)
}
