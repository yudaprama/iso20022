package model

// Set of elements used  to provide information concerning the verification of identification data for which verification was requested.
type VerificationReport1 struct {

	// Unique identification, as assigned by a sending party, to unambiguously identify the party and account identification information group within the original message.
	OriginalIdentification *Max35Text `xml:"OrgnlId"`

	// Identifies whether the party and/or account information received is correct.
	Verification *IdentificationVerificationIndicator `xml:"Vrfctn"`

	// Specifies the reason why the verified identification information is incorrect.
	Reason *VerificationReason1Choice `xml:"Rsn,omitempty"`

	// Provides party and/or account identification information as given in the original message.
	OriginalPartyAndAccountIdentification *IdentificationInformation1 `xml:"OrgnlPtyAndAcctId,omitempty"`

	// Provides party and/or account identification information.
	UpdatedPartyAndAccountIdentification *IdentificationInformation1 `xml:"UpdtdPtyAndAcctId,omitempty"`
}

func (v *VerificationReport1) SetOriginalIdentification(value string) {
	v.OriginalIdentification = (*Max35Text)(&value)
}

func (v *VerificationReport1) SetVerification(value string) {
	v.Verification = (*IdentificationVerificationIndicator)(&value)
}

func (v *VerificationReport1) AddReason() *VerificationReason1Choice {
	v.Reason = new(VerificationReason1Choice)
	return v.Reason
}

func (v *VerificationReport1) AddOriginalPartyAndAccountIdentification() *IdentificationInformation1 {
	v.OriginalPartyAndAccountIdentification = new(IdentificationInformation1)
	return v.OriginalPartyAndAccountIdentification
}

func (v *VerificationReport1) AddUpdatedPartyAndAccountIdentification() *IdentificationInformation1 {
	v.UpdatedPartyAndAccountIdentification = new(IdentificationInformation1)
	return v.UpdatedPartyAndAccountIdentification
}
