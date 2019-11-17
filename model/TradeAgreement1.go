package model

// Date and identification of a trade.
type TradeAgreement1 struct {

	// Date at which the trading parties agree on a treasury trade.
	TradeDate *ISODate `xml:"TradDt"`

	// Identification of a notification.This identification must be unique amongst all notifications of same type confirmed by the same party.
	NotificationIdentification *Max35Text `xml:"NtfctnId"`

	// Reference common to the parties of a trade.
	CommonReference *Max35Text `xml:"CmonRef,omitempty"`
}

func (t *TradeAgreement1) SetTradeDate(value string) {
	t.TradeDate = (*ISODate)(&value)
}

func (t *TradeAgreement1) SetNotificationIdentification(value string) {
	t.NotificationIdentification = (*Max35Text)(&value)
}

func (t *TradeAgreement1) SetCommonReference(value string) {
	t.CommonReference = (*Max35Text)(&value)
}
