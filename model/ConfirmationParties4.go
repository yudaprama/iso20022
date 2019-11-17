package model

// Parties used for acting parties that apply either to the whole message or to individual sides.
type ConfirmationParties4 struct {

	// Party that identifies the underlying investor.
	Investor []*PartyIdentificationAndAccount79 `xml:"Invstr,omitempty"`

	// Party that buys goods or services, or a financial instrument.
	Buyer *ConfirmationPartyDetails2 `xml:"Buyr,omitempty"`

	// Party that has applied, met specific requirements, and received a monetary or securities loan from a lender. The party initiating the request signs a promissory note agreeing to pay the lien holder back during a specified timeframe for the entire loan amount plus any additional fees. The borrower is legally responsible for repayment of the loan and is subject to any penalties for not repaying the loan back based on the lending terms agreed upon.
	Borrower *ConfirmationPartyDetails2 `xml:"Brrwr,omitempty"`

	// Party that sells goods or services, or a financial instrument.
	Seller *ConfirmationPartyDetails2 `xml:"Sellr,omitempty"`

	// A private, public or institutional entity which makes funds available to others to borrow.
	Lender *ConfirmationPartyDetails2 `xml:"Lndr,omitempty"`

	// Brokerage firm which is the commissioned broker in a multi-broker trade.
	BrokerOfCredit *ConfirmationPartyDetails3 `xml:"BrkrOfCdt,omitempty"`

	// Broker or other intermediary with the closest association with the investor.
	IntroducingFirm *ConfirmationPartyDetails3 `xml:"IntrdcgFirm,omitempty"`

	// Brokerage firm assigned to take credit on the trade from the step-out brokerage firm.
	StepInFirm *ConfirmationPartyDetails1 `xml:"StepInFirm,omitempty"`

	// Brokerage firm that executes an order, but gives other firms credit and some of the commission for the trade.
	StepOutFirm *ConfirmationPartyDetails1 `xml:"StepOutFirm,omitempty"`

	// Party, also know as take up broker, that settles security transactions from another broker for a fee.
	ClearingFirm *ConfirmationPartyDetails6 `xml:"ClrFirm,omitempty"`

	// Party responsible for executing an order (for example, an executing or give-up broker).  Usually a commission is charged to the client for executing an order.
	ExecutingBroker *ConfirmationPartyDetails6 `xml:"ExctgBrkr,omitempty"`

	// Party (buyer or seller) that positively affirms the details of a previously agreed security trade confirmation.
	AffirmingParty *ConfirmationPartyDetails3 `xml:"AffrmgPty,omitempty"`

	// Party involved  in a legal proceeding, agreement, or other transaction.
	TradeBeneficiaryParty *ConfirmationPartyDetails3 `xml:"TradBnfcryPty,omitempty"`
}

func (c *ConfirmationParties4) AddInvestor() *PartyIdentificationAndAccount79 {
	newValue := new(PartyIdentificationAndAccount79)
	c.Investor = append(c.Investor, newValue)
	return newValue
}

func (c *ConfirmationParties4) AddBuyer() *ConfirmationPartyDetails2 {
	c.Buyer = new(ConfirmationPartyDetails2)
	return c.Buyer
}

func (c *ConfirmationParties4) AddBorrower() *ConfirmationPartyDetails2 {
	c.Borrower = new(ConfirmationPartyDetails2)
	return c.Borrower
}

func (c *ConfirmationParties4) AddSeller() *ConfirmationPartyDetails2 {
	c.Seller = new(ConfirmationPartyDetails2)
	return c.Seller
}

func (c *ConfirmationParties4) AddLender() *ConfirmationPartyDetails2 {
	c.Lender = new(ConfirmationPartyDetails2)
	return c.Lender
}

func (c *ConfirmationParties4) AddBrokerOfCredit() *ConfirmationPartyDetails3 {
	c.BrokerOfCredit = new(ConfirmationPartyDetails3)
	return c.BrokerOfCredit
}

func (c *ConfirmationParties4) AddIntroducingFirm() *ConfirmationPartyDetails3 {
	c.IntroducingFirm = new(ConfirmationPartyDetails3)
	return c.IntroducingFirm
}

func (c *ConfirmationParties4) AddStepInFirm() *ConfirmationPartyDetails1 {
	c.StepInFirm = new(ConfirmationPartyDetails1)
	return c.StepInFirm
}

func (c *ConfirmationParties4) AddStepOutFirm() *ConfirmationPartyDetails1 {
	c.StepOutFirm = new(ConfirmationPartyDetails1)
	return c.StepOutFirm
}

func (c *ConfirmationParties4) AddClearingFirm() *ConfirmationPartyDetails6 {
	c.ClearingFirm = new(ConfirmationPartyDetails6)
	return c.ClearingFirm
}

func (c *ConfirmationParties4) AddExecutingBroker() *ConfirmationPartyDetails6 {
	c.ExecutingBroker = new(ConfirmationPartyDetails6)
	return c.ExecutingBroker
}

func (c *ConfirmationParties4) AddAffirmingParty() *ConfirmationPartyDetails3 {
	c.AffirmingParty = new(ConfirmationPartyDetails3)
	return c.AffirmingParty
}

func (c *ConfirmationParties4) AddTradeBeneficiaryParty() *ConfirmationPartyDetails3 {
	c.TradeBeneficiaryParty = new(ConfirmationPartyDetails3)
	return c.TradeBeneficiaryParty
}
