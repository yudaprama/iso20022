package model

// Parties used for acting parties that apply either to the whole message or to individual sides.
type ConfirmationParties2 struct {

	// Party that buys goods or services, or a financial instrument.
	Buyer *ConfirmationPartyDetails2 `xml:"Buyr,omitempty"`

	// Party that has applied, met specific requirements, and received a monetary or securities loan from a lender. The party initiating the request signs a promissory note agreeing to pay the lien holder back during a specified timeframe for the entire loan amount plus any additional fees. The borrower is legally responsible for repayment of the loan and is subject to any penalties for not repaying the loan back based on the lending terms agreed upon.
	Borrower *ConfirmationPartyDetails2 `xml:"Brrwr,omitempty"`

	// Party that sells goods or services, or a financial instrument.
	Seller *ConfirmationPartyDetails2 `xml:"Sellr,omitempty"`

	// A private, public or institutional entity which makes funds available to others to borrow.
	Lender *ConfirmationPartyDetails2 `xml:"Lndr,omitempty"`

	// Brokerage firm which is the commissioned broker in a multi-broker trade.
	BrokerOfCredit *ConfirmationPartyDetails1 `xml:"BrkrOfCdt,omitempty"`

	// Broker or other intermediary with the closest association with the investor.
	IntroducingFirm *ConfirmationPartyDetails1 `xml:"IntrdcgFirm,omitempty"`

	// Brokerage firm assigned to take credit on the trade from the step-out brokerage firm.
	StepInFirm *ConfirmationPartyDetails1 `xml:"StepInFirm,omitempty"`

	// Brokerage firm that executes an order, but gives other firms credit and some of the commission for the trade.
	StepOutFirm *ConfirmationPartyDetails1 `xml:"StepOutFirm,omitempty"`

	// Party, also know as take up broker, that settles security transactions from another broker for a fee.
	ClearingFirm *ConfirmationPartyDetails5 `xml:"ClrFirm,omitempty"`

	// Party responsible for executing an order (for example, an executing or give-up broker).  Usually a commission is charged to the client for executing an order.
	ExecutingBroker *ConfirmationPartyDetails5 `xml:"ExctgBrkr,omitempty"`

	// Party sending the message to the CMU (Central Matching Utility) to identify the actual business unit as known to the CMU (Central Matching Utility).
	CMUParty *ConfirmationPartyDetails1 `xml:"CMUPty,omitempty"`

	// Actual business unit of the counterparty to the sender of the  message to the CMU (Central Matching Utility) as known to the CMU (Central Matching Utility).
	CMUCounterparty *ConfirmationPartyDetails1 `xml:"CMUCtrPty,omitempty"`

	// Party (buyer or seller) that positively affirms the details of a previously agreed security trade confirmation.
	AffirmingParty *ConfirmationPartyDetails1 `xml:"AffrmgPty,omitempty"`

	// Party involved  in a legal proceeding, agreement or other transaction.
	TradeBeneficiaryParty *ConfirmationPartyDetails3 `xml:"TradBnfcryPty,omitempty"`
}

func (c *ConfirmationParties2) AddBuyer() *ConfirmationPartyDetails2 {
	c.Buyer = new(ConfirmationPartyDetails2)
	return c.Buyer
}

func (c *ConfirmationParties2) AddBorrower() *ConfirmationPartyDetails2 {
	c.Borrower = new(ConfirmationPartyDetails2)
	return c.Borrower
}

func (c *ConfirmationParties2) AddSeller() *ConfirmationPartyDetails2 {
	c.Seller = new(ConfirmationPartyDetails2)
	return c.Seller
}

func (c *ConfirmationParties2) AddLender() *ConfirmationPartyDetails2 {
	c.Lender = new(ConfirmationPartyDetails2)
	return c.Lender
}

func (c *ConfirmationParties2) AddBrokerOfCredit() *ConfirmationPartyDetails1 {
	c.BrokerOfCredit = new(ConfirmationPartyDetails1)
	return c.BrokerOfCredit
}

func (c *ConfirmationParties2) AddIntroducingFirm() *ConfirmationPartyDetails1 {
	c.IntroducingFirm = new(ConfirmationPartyDetails1)
	return c.IntroducingFirm
}

func (c *ConfirmationParties2) AddStepInFirm() *ConfirmationPartyDetails1 {
	c.StepInFirm = new(ConfirmationPartyDetails1)
	return c.StepInFirm
}

func (c *ConfirmationParties2) AddStepOutFirm() *ConfirmationPartyDetails1 {
	c.StepOutFirm = new(ConfirmationPartyDetails1)
	return c.StepOutFirm
}

func (c *ConfirmationParties2) AddClearingFirm() *ConfirmationPartyDetails5 {
	c.ClearingFirm = new(ConfirmationPartyDetails5)
	return c.ClearingFirm
}

func (c *ConfirmationParties2) AddExecutingBroker() *ConfirmationPartyDetails5 {
	c.ExecutingBroker = new(ConfirmationPartyDetails5)
	return c.ExecutingBroker
}

func (c *ConfirmationParties2) AddCMUParty() *ConfirmationPartyDetails1 {
	c.CMUParty = new(ConfirmationPartyDetails1)
	return c.CMUParty
}

func (c *ConfirmationParties2) AddCMUCounterparty() *ConfirmationPartyDetails1 {
	c.CMUCounterparty = new(ConfirmationPartyDetails1)
	return c.CMUCounterparty
}

func (c *ConfirmationParties2) AddAffirmingParty() *ConfirmationPartyDetails1 {
	c.AffirmingParty = new(ConfirmationPartyDetails1)
	return c.AffirmingParty
}

func (c *ConfirmationParties2) AddTradeBeneficiaryParty() *ConfirmationPartyDetails3 {
	c.TradeBeneficiaryParty = new(ConfirmationPartyDetails3)
	return c.TradeBeneficiaryParty
}
