package model

// Parties used for acting parties that apply either to the whole message or to individual sides.
type ConfirmationParties3 struct {

	// Party (buyer or seller) that positively affirms the details of a previously agreed security trade confirmation.
	AffirmingParty *ConfirmationPartyDetails4 `xml:"AffrmgPty"`

	// Party that buys goods or services, or a financial instrument.
	Buyer *ConfirmationPartyDetails2 `xml:"Buyr,omitempty"`

	// Party that has applied, met specific requirements, and received a monetary or securities loan from a lender. The party initiating the request signs a promissory note agreeing to pay the lien holder back during a specified timeframe for the entire loan amount plus any additional fees. The borrower is legally responsible for repayment of the loan and is subject to any penalties for not repaying the loan back based on the lending terms agreed upon.
	Borrower *ConfirmationPartyDetails2 `xml:"Brrwr,omitempty"`

	// Party that sells goods or services, or a financial instrument.
	Seller *ConfirmationPartyDetails2 `xml:"Sellr,omitempty"`

	// A private, public or institutional entity which makes funds available to others to borrow.
	Lender *ConfirmationPartyDetails2 `xml:"Lndr,omitempty"`

	// Party involved  in a legal proceeding, agreement, or other transaction.
	TradeBeneficiaryParty *ConfirmationPartyDetails3 `xml:"TradBnfcryPty,omitempty"`
}

func (c *ConfirmationParties3) AddAffirmingParty() *ConfirmationPartyDetails4 {
	c.AffirmingParty = new(ConfirmationPartyDetails4)
	return c.AffirmingParty
}

func (c *ConfirmationParties3) AddBuyer() *ConfirmationPartyDetails2 {
	c.Buyer = new(ConfirmationPartyDetails2)
	return c.Buyer
}

func (c *ConfirmationParties3) AddBorrower() *ConfirmationPartyDetails2 {
	c.Borrower = new(ConfirmationPartyDetails2)
	return c.Borrower
}

func (c *ConfirmationParties3) AddSeller() *ConfirmationPartyDetails2 {
	c.Seller = new(ConfirmationPartyDetails2)
	return c.Seller
}

func (c *ConfirmationParties3) AddLender() *ConfirmationPartyDetails2 {
	c.Lender = new(ConfirmationPartyDetails2)
	return c.Lender
}

func (c *ConfirmationParties3) AddTradeBeneficiaryParty() *ConfirmationPartyDetails3 {
	c.TradeBeneficiaryParty = new(ConfirmationPartyDetails3)
	return c.TradeBeneficiaryParty
}
