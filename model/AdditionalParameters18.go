package model

// Specifies additional parameters to the message or transaction.
type AdditionalParameters18 struct {

	// Specifies whether there exists a pre-confirmation.
	PreConfirmation *PreConfirmation1Code `xml:"PreConf,omitempty"`

	// Specifies partial settlement information.
	PartialSettlement *PartialSettlement2Code `xml:"PrtlSttlm,omitempty"`

	// Identification of the confirmation previously sent to confirm the partial settlement of a transaction.
	PreviousPartialConfirmationIdentification *Max35Text `xml:"PrvsPrtlConfId,omitempty"`
}

func (a *AdditionalParameters18) SetPreConfirmation(value string) {
	a.PreConfirmation = (*PreConfirmation1Code)(&value)
}

func (a *AdditionalParameters18) SetPartialSettlement(value string) {
	a.PartialSettlement = (*PartialSettlement2Code)(&value)
}

func (a *AdditionalParameters18) SetPreviousPartialConfirmationIdentification(value string) {
	a.PreviousPartialConfirmationIdentification = (*Max35Text)(&value)
}
