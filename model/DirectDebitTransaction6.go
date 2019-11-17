package model

// Set of elements used to provide specific information on the direct debit transaction and the related mandate.
type DirectDebitTransaction6 struct {

	// Set of elements used to provide further details of the direct debit mandate signed between the creditor and the debtor.
	MandateRelatedInformation *MandateRelatedInformation6 `xml:"MndtRltdInf,omitempty"`

	// Credit party that signs the mandate.
	CreditorSchemeIdentification *PartyIdentification32 `xml:"CdtrSchmeId,omitempty"`

	// Unique and unambiguous identification of the pre-notification which is sent separately from the direct debit instruction.
	//
	// Usage: The direct debit pre-notification is used to reconcile separately sent collection information with the direct debit transaction information.
	PreNotificationIdentification *Max35Text `xml:"PreNtfctnId,omitempty"`

	// Date on which the creditor notifies the debtor about the amount and date on which the direct debit instruction will be presented to the debtor's agent.
	PreNotificationDate *ISODate `xml:"PreNtfctnDt,omitempty"`
}

func (d *DirectDebitTransaction6) AddMandateRelatedInformation() *MandateRelatedInformation6 {
	d.MandateRelatedInformation = new(MandateRelatedInformation6)
	return d.MandateRelatedInformation
}

func (d *DirectDebitTransaction6) AddCreditorSchemeIdentification() *PartyIdentification32 {
	d.CreditorSchemeIdentification = new(PartyIdentification32)
	return d.CreditorSchemeIdentification
}

func (d *DirectDebitTransaction6) SetPreNotificationIdentification(value string) {
	d.PreNotificationIdentification = (*Max35Text)(&value)
}

func (d *DirectDebitTransaction6) SetPreNotificationDate(value string) {
	d.PreNotificationDate = (*ISODate)(&value)
}
