package model

// Provides further details related to a direct debit mandate signed between the creditor and the debtor.
type MandateRelatedInformation11 struct {

	// Unique identification, as assigned by the creditor, to unambiguously identify the mandate.
	MandateIdentification *Max35Text `xml:"MndtId,omitempty"`

	// Date on which the direct debit mandate has been signed by the debtor.
	DateOfSignature *ISODate `xml:"DtOfSgntr,omitempty"`

	// Indicator notifying whether the underlying mandate is amended or not.
	AmendmentIndicator *TrueFalseIndicator `xml:"AmdmntInd,omitempty"`

	// List of mandate elements that have been modified.
	AmendmentInformationDetails *AmendmentInformationDetails11 `xml:"AmdmntInfDtls,omitempty"`

	// Additional security provisions, such as a digital signature, as provided by the debtor.
	ElectronicSignature *Max1025Text `xml:"ElctrncSgntr,omitempty"`

	// Date of the first collection of a direct debit as per the mandate.
	FirstCollectionDate *ISODate `xml:"FrstColltnDt,omitempty"`

	// Date of the final collection of a direct debit as per the mandate.
	FinalCollectionDate *ISODate `xml:"FnlColltnDt,omitempty"`

	// Regularity with which direct debit instructions are to be created and processed.
	Frequency *Frequency36Choice `xml:"Frqcy,omitempty"`

	// Reason for the direct debit mandate to allow the user to distinguish between different mandates for the same creditor.
	Reason *MandateSetupReason1Choice `xml:"Rsn,omitempty"`

	// Specifies the number of days the direct debit instruction must be tracked.
	TrackingDays *Exact2NumericText `xml:"TrckgDays,omitempty"`
}

func (m *MandateRelatedInformation11) SetMandateIdentification(value string) {
	m.MandateIdentification = (*Max35Text)(&value)
}

func (m *MandateRelatedInformation11) SetDateOfSignature(value string) {
	m.DateOfSignature = (*ISODate)(&value)
}

func (m *MandateRelatedInformation11) SetAmendmentIndicator(value string) {
	m.AmendmentIndicator = (*TrueFalseIndicator)(&value)
}

func (m *MandateRelatedInformation11) AddAmendmentInformationDetails() *AmendmentInformationDetails11 {
	m.AmendmentInformationDetails = new(AmendmentInformationDetails11)
	return m.AmendmentInformationDetails
}

func (m *MandateRelatedInformation11) SetElectronicSignature(value string) {
	m.ElectronicSignature = (*Max1025Text)(&value)
}

func (m *MandateRelatedInformation11) SetFirstCollectionDate(value string) {
	m.FirstCollectionDate = (*ISODate)(&value)
}

func (m *MandateRelatedInformation11) SetFinalCollectionDate(value string) {
	m.FinalCollectionDate = (*ISODate)(&value)
}

func (m *MandateRelatedInformation11) AddFrequency() *Frequency36Choice {
	m.Frequency = new(Frequency36Choice)
	return m.Frequency
}

func (m *MandateRelatedInformation11) AddReason() *MandateSetupReason1Choice {
	m.Reason = new(MandateSetupReason1Choice)
	return m.Reason
}

func (m *MandateRelatedInformation11) SetTrackingDays(value string) {
	m.TrackingDays = (*Exact2NumericText)(&value)
}
