package model

// Specifies elements related to the response sent by the clearing member to the central counterparty in the context of the buy in process.
type BuyIn3 struct {

	// Indicates the reference of the BuyInNotification message.
	BuyInNotificationIdentification *Max35Text `xml:"BuyInNtfctnId"`

	// Specific continuous net settlement case where the central counterparty can call for buy-in at a date anterior to "theoretical" buy-in date, the clearing member may request a delay.
	RequestForDelayIndicator *YesNoIndicator `xml:"ReqForDelyInd"`

	// Number of days associated to the request for delay.
	NumberOfDays *Number `xml:"NbOfDays"`

	// Buy in quantity called initially by the central counterparty.
	InitialQuantity *FinancialInstrumentQuantity1Choice `xml:"InitlQty"`

	// Quantity amount covered by the clearing member after notification.
	CoveredQuantity *FinancialInstrumentQuantity1Choice `xml:"CvrdQty"`

	// Quantity amount non covered by the clearing member after notification (this is, new buy in amount to be executed).
	UncoveredQuantity *FinancialInstrumentQuantity1Choice `xml:"UcvrdQty"`
}

func (b *BuyIn3) SetBuyInNotificationIdentification(value string) {
	b.BuyInNotificationIdentification = (*Max35Text)(&value)
}

func (b *BuyIn3) SetRequestForDelayIndicator(value string) {
	b.RequestForDelayIndicator = (*YesNoIndicator)(&value)
}

func (b *BuyIn3) SetNumberOfDays(value string) {
	b.NumberOfDays = (*Number)(&value)
}

func (b *BuyIn3) AddInitialQuantity() *FinancialInstrumentQuantity1Choice {
	b.InitialQuantity = new(FinancialInstrumentQuantity1Choice)
	return b.InitialQuantity
}

func (b *BuyIn3) AddCoveredQuantity() *FinancialInstrumentQuantity1Choice {
	b.CoveredQuantity = new(FinancialInstrumentQuantity1Choice)
	return b.CoveredQuantity
}

func (b *BuyIn3) AddUncoveredQuantity() *FinancialInstrumentQuantity1Choice {
	b.UncoveredQuantity = new(FinancialInstrumentQuantity1Choice)
	return b.UncoveredQuantity
}
