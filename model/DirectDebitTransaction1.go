package model

// Set of elements providing specific information on the direct debit transaction and the related mandate.
type DirectDebitTransaction1 struct {

	// Set of elements used to provide further details related to a direct debit mandate signed between the creditor and the debtor.
	//
	// Usage: Mandate related information is to be used only when the direct debit relates to a mandate signed between the debtor and the creditor.
	MandateRelatedInformation *MandateRelatedInformation1 `xml:"MndtRltdInf,omitempty"`

	// Credit party that signs the direct debit mandate.
	CreditorSchemeIdentification *PartyIdentification8 `xml:"CdtrSchmeId,omitempty"`

	// Unique and unambiguous identification of the pre-notification which is sent separately from the direct debit instruction.
	//
	// Usage: The direct debit pre-notification is used to reconcile separately sent collection information with the direct debit transaction information.
	PreNotificationIdentification *Max35Text `xml:"PreNtfctnId,omitempty"`

	// Date on which the creditor notifies the debtor about the amount and date on which the direct debit instruction will be presented to the debtor's agent.
	PreNotificationDate *ISODate `xml:"PreNtfctnDt,omitempty"`
}

func (d *DirectDebitTransaction1) AddMandateRelatedInformation() *MandateRelatedInformation1 {
	d.MandateRelatedInformation = new(MandateRelatedInformation1)
	return d.MandateRelatedInformation
}

func (d *DirectDebitTransaction1) AddCreditorSchemeIdentification() *PartyIdentification8 {
	d.CreditorSchemeIdentification = new(PartyIdentification8)
	return d.CreditorSchemeIdentification
}

func (d *DirectDebitTransaction1) SetPreNotificationIdentification(value string) {
	d.PreNotificationIdentification = (*Max35Text)(&value)
}

func (d *DirectDebitTransaction1) SetPreNotificationDate(value string) {
	d.PreNotificationDate = (*ISODate)(&value)
}
