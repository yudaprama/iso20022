package model

// Payment instrument between a debtor and a creditor, which flows through one or more financial institutions or systems.
type CreditTransfer3 struct {

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts' receivable system.
	Reference *Max35Text `xml:"Ref,omitempty"`

	// Party that receives an amount of money from the debtor. The creditor is also the credit account owner.
	CreditorDetails *Creditor2 `xml:"CdtrDtls"`

	// Party that owes an amount of money to the creditor. The debtor is also the debit account owner.
	DebtorDetails *Debtor2 `xml:"DbtrDtls,omitempty"`
}

func (c *CreditTransfer3) SetReference(value string) {
	c.Reference = (*Max35Text)(&value)
}

func (c *CreditTransfer3) AddCreditorDetails() *Creditor2 {
	c.CreditorDetails = new(Creditor2)
	return c.CreditorDetails
}

func (c *CreditTransfer3) AddDebtorDetails() *Debtor2 {
	c.DebtorDetails = new(Debtor2)
	return c.DebtorDetails
}
