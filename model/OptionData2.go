package model

// Set of data which contains the link to a previously notified option trade.
type OptionData2 struct {

	// Date at which the trading parties have agreed on an option trade.
	TradeDate *ISODate `xml:"TradDt"`

	// Refers to the identification of a trade assigned by the trading side of a foreign exchange option trade.
	NotificationIdentification *Max35Text `xml:"NtfctnId"`

	// Reference common to the parties of a trade.
	CommonReference *Max35Text `xml:"CmonRef,omitempty"`

	// Refers to the identification of a previous event in the life of a foreign exchange option trade.
	RelatedReference *Max35Text `xml:"RltdRef,omitempty"`

	// Describes the reason for the cancellation or the amendment.
	AmendOrCancelReason *Max35Text `xml:"AmdOrCclRsn,omitempty"`

	// Set of data defining a foreign exchange option sold.
	Option *Option3 `xml:"Optn"`
}

func (o *OptionData2) SetTradeDate(value string) {
	o.TradeDate = (*ISODate)(&value)
}

func (o *OptionData2) SetNotificationIdentification(value string) {
	o.NotificationIdentification = (*Max35Text)(&value)
}

func (o *OptionData2) SetCommonReference(value string) {
	o.CommonReference = (*Max35Text)(&value)
}

func (o *OptionData2) SetRelatedReference(value string) {
	o.RelatedReference = (*Max35Text)(&value)
}

func (o *OptionData2) SetAmendOrCancelReason(value string) {
	o.AmendOrCancelReason = (*Max35Text)(&value)
}

func (o *OptionData2) AddOption() *Option3 {
	o.Option = new(Option3)
	return o.Option
}
