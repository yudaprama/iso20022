package model

// Specifies elements related to the notification (or warn) sent by the central counterparty to the clearing member in the context of the buy in process.
type BuyIn4 struct {

	// Indicates whether the message is a warning only or a notification.
	WarningIndicator *YesNoIndicator `xml:"WrngInd,omitempty"`

	// Provides the date at which the buy-in will occur.
	ExpectedBuyInDate *DateFormat15Choice `xml:"XpctdBuyInDt"`

	// Identifies the latest date by which the buy-in operation can be cancelled.
	CancellationLimitDate *ISODate `xml:"CxlLmtDt,omitempty"`

	// Identifies the date by which the buy-in operation is reversed by the CCP.
	BuyInReversionDate *ISODate `xml:"BuyInRvrsnDt,omitempty"`
}

func (b *BuyIn4) SetWarningIndicator(value string) {
	b.WarningIndicator = (*YesNoIndicator)(&value)
}

func (b *BuyIn4) AddExpectedBuyInDate() *DateFormat15Choice {
	b.ExpectedBuyInDate = new(DateFormat15Choice)
	return b.ExpectedBuyInDate
}

func (b *BuyIn4) SetCancellationLimitDate(value string) {
	b.CancellationLimitDate = (*ISODate)(&value)
}

func (b *BuyIn4) SetBuyInReversionDate(value string) {
	b.BuyInReversionDate = (*ISODate)(&value)
}
