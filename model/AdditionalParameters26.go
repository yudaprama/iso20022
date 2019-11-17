package model

// Specifies additional parameters to the message or transaction.
type AdditionalParameters26 struct {

	// Specifies whether there exists a pre-confirmation.
	PreConfirmation *PreConfirmation1Code `xml:"PreConf,omitempty"`

	// Specifies partial settlement information.
	PartialSettlement *PartialSettlement2Code `xml:"PrtlSttlm,omitempty"`

	// Identification of the confirmation previously sent to confirm the partial settlement of a transaction.
	PreviousPartialConfirmationIdentification *RestrictedFINXMax16Text `xml:"PrvsPrtlConfId,omitempty"`
}

func (a *AdditionalParameters26) SetPreConfirmation(value string) {
	a.PreConfirmation = (*PreConfirmation1Code)(&value)
}

func (a *AdditionalParameters26) SetPartialSettlement(value string) {
	a.PartialSettlement = (*PartialSettlement2Code)(&value)
}

func (a *AdditionalParameters26) SetPreviousPartialConfirmationIdentification(value string) {
	a.PreviousPartialConfirmationIdentification = (*RestrictedFINXMax16Text)(&value)
}
