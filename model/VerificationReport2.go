package model

// Provides the details of the verification of identification data for which verification was requested.
type VerificationReport2 struct {

	// Unique identification, as assigned by a sending party, to unambiguously identify the party and account identification information group within the original message.
	OriginalIdentification *Max35Text `xml:"OrgnlId"`

	// Identifies whether the party and/or account information received is correct.
	Verification *IdentificationVerificationIndicator `xml:"Vrfctn"`

	// Specifies the reason why the verified identification information is incorrect.
	Reason *VerificationReason1Choice `xml:"Rsn,omitempty"`

	// Provides party and/or account identification information as given in the original message.
	OriginalPartyAndAccountIdentification *IdentificationInformation2 `xml:"OrgnlPtyAndAcctId,omitempty"`

	// Provides party and/or account identification information.
	UpdatedPartyAndAccountIdentification *IdentificationInformation2 `xml:"UpdtdPtyAndAcctId,omitempty"`
}

func (v *VerificationReport2) SetOriginalIdentification(value string) {
	v.OriginalIdentification = (*Max35Text)(&value)
}

func (v *VerificationReport2) SetVerification(value string) {
	v.Verification = (*IdentificationVerificationIndicator)(&value)
}

func (v *VerificationReport2) AddReason() *VerificationReason1Choice {
	v.Reason = new(VerificationReason1Choice)
	return v.Reason
}

func (v *VerificationReport2) AddOriginalPartyAndAccountIdentification() *IdentificationInformation2 {
	v.OriginalPartyAndAccountIdentification = new(IdentificationInformation2)
	return v.OriginalPartyAndAccountIdentification
}

func (v *VerificationReport2) AddUpdatedPartyAndAccountIdentification() *IdentificationInformation2 {
	v.UpdatedPartyAndAccountIdentification = new(IdentificationInformation2)
	return v.UpdatedPartyAndAccountIdentification
}
