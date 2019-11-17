package model

// Specifies elements related to the confirmation sent by the central counterparty to the clearing member in the context of the buy in process.
type BuyIn2 struct {

	// Indicates the reference of the BuyInNotification message.
	BuyInNotificationIdentification *Max35Text `xml:"BuyInNtfctnId,omitempty"`

	// Indicates the reference id of the buy in.
	BuyInIdentification *Max35Text `xml:"BuyInId"`

	// Provides the date at which the buy occured.
	Date *ISODate `xml:"Dt"`

	// Provides the price of the buy-in.
	Price *Price4 `xml:"Pric,omitempty"`

	// Specifies the elements related to the securities that the central counterparty had to buy in the context of the buy-in process.
	SecuritiesBuyIn *SecuritiesCompensation1 `xml:"SctiesBuyIn,omitempty"`

	// Provides details about the cash compensation required, in case the central counterparty could not buy the securities to cover the trade(s) that failed.
	RequiredCashCompensation *CashCompensation1 `xml:"ReqrdCshCompstn,omitempty"`
}

func (b *BuyIn2) SetBuyInNotificationIdentification(value string) {
	b.BuyInNotificationIdentification = (*Max35Text)(&value)
}

func (b *BuyIn2) SetBuyInIdentification(value string) {
	b.BuyInIdentification = (*Max35Text)(&value)
}

func (b *BuyIn2) SetDate(value string) {
	b.Date = (*ISODate)(&value)
}

func (b *BuyIn2) AddPrice() *Price4 {
	b.Price = new(Price4)
	return b.Price
}

func (b *BuyIn2) AddSecuritiesBuyIn() *SecuritiesCompensation1 {
	b.SecuritiesBuyIn = new(SecuritiesCompensation1)
	return b.SecuritiesBuyIn
}

func (b *BuyIn2) AddRequiredCashCompensation() *CashCompensation1 {
	b.RequiredCashCompensation = new(CashCompensation1)
	return b.RequiredCashCompensation
}
