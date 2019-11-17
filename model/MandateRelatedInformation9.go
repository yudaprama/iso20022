package model

// Provides further details related to a direct debit mandate signed between the creditor and the debtor.
type MandateRelatedInformation9 struct {

	// Unique identification, as assigned by the creditor, to unambiguously identify the mandate.
	MandateIdentification *Max35Text `xml:"MndtId,omitempty"`

	// Date on which the direct debit mandate has been signed by the debtor.
	DateOfSignature *ISODate `xml:"DtOfSgntr,omitempty"`

	// Indicator notifying whether the underlying mandate is amended or not.
	AmendmentIndicator *TrueFalseIndicator `xml:"AmdmntInd,omitempty"`

	// List of mandate elements that have been modified.
	AmendmentInformationDetails *AmendmentInformationDetails9 `xml:"AmdmntInfDtls,omitempty"`

	// Additional security provisions, such as a digital signature, as provided by the debtor.
	ElectronicSignature *Max1025Text `xml:"ElctrncSgntr,omitempty"`

	// Date of the first collection of a direct debit as per the mandate.
	FirstCollectionDate *ISODate `xml:"FrstColltnDt,omitempty"`

	// Date of the final collection of a direct debit as per the mandate.
	FinalCollectionDate *ISODate `xml:"FnlColltnDt,omitempty"`

	// Regularity with which direct debit instructions are to be created and processed.
	Frequency *Frequency6Code `xml:"Frqcy,omitempty"`
}

func (m *MandateRelatedInformation9) SetMandateIdentification(value string) {
	m.MandateIdentification = (*Max35Text)(&value)
}

func (m *MandateRelatedInformation9) SetDateOfSignature(value string) {
	m.DateOfSignature = (*ISODate)(&value)
}

func (m *MandateRelatedInformation9) SetAmendmentIndicator(value string) {
	m.AmendmentIndicator = (*TrueFalseIndicator)(&value)
}

func (m *MandateRelatedInformation9) AddAmendmentInformationDetails() *AmendmentInformationDetails9 {
	m.AmendmentInformationDetails = new(AmendmentInformationDetails9)
	return m.AmendmentInformationDetails
}

func (m *MandateRelatedInformation9) SetElectronicSignature(value string) {
	m.ElectronicSignature = (*Max1025Text)(&value)
}

func (m *MandateRelatedInformation9) SetFirstCollectionDate(value string) {
	m.FirstCollectionDate = (*ISODate)(&value)
}

func (m *MandateRelatedInformation9) SetFinalCollectionDate(value string) {
	m.FinalCollectionDate = (*ISODate)(&value)
}

func (m *MandateRelatedInformation9) SetFrequency(value string) {
	m.Frequency = (*Frequency6Code)(&value)
}
