package model

// Provides further specific details on the direct debit transaction and the related mandate.
type DirectDebitTransaction9 struct {

	// Provides further details of the direct debit mandate signed between the creditor and the debtor.
	MandateRelatedInformation *MandateRelatedInformation11 `xml:"MndtRltdInf,omitempty"`

	// Credit party that signs the mandate.
	CreditorSchemeIdentification *PartyIdentification43 `xml:"CdtrSchmeId,omitempty"`

	// Unique and unambiguous identification of the pre-notification which is sent separately from the direct debit instruction.
	//
	// Usage: The direct debit pre-notification is used to reconcile separately sent collection information with the direct debit transaction information.
	PreNotificationIdentification *Max35Text `xml:"PreNtfctnId,omitempty"`

	// Date on which the creditor notifies the debtor about the amount and date on which the direct debit instruction will be presented to the debtor's agent.
	PreNotificationDate *ISODate `xml:"PreNtfctnDt,omitempty"`
}

func (d *DirectDebitTransaction9) AddMandateRelatedInformation() *MandateRelatedInformation11 {
	d.MandateRelatedInformation = new(MandateRelatedInformation11)
	return d.MandateRelatedInformation
}

func (d *DirectDebitTransaction9) AddCreditorSchemeIdentification() *PartyIdentification43 {
	d.CreditorSchemeIdentification = new(PartyIdentification43)
	return d.CreditorSchemeIdentification
}

func (d *DirectDebitTransaction9) SetPreNotificationIdentification(value string) {
	d.PreNotificationIdentification = (*Max35Text)(&value)
}

func (d *DirectDebitTransaction9) SetPreNotificationDate(value string) {
	d.PreNotificationDate = (*ISODate)(&value)
}
