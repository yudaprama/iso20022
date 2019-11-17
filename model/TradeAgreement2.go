package model

// Date and identification of a trade together with references to previous events in its life.
type TradeAgreement2 struct {

	// Date at which the trading parties have agreed to amend or cancel a treasury trade.
	TradeDate *ISODate `xml:"TradDt"`

	// Identification of a notification.This identification must be unique amongst all notifications of same type confirmed by the same party.
	NotificationIdentification *Max35Text `xml:"NtfctnId"`

	// Reference common to the parties of a trade.
	CommonReference *Max35Text `xml:"CmonRef,omitempty"`

	// Describes the reason for the cancellation or the amendment.
	AmendOrCancelReason *Max35Text `xml:"AmdOrCclRsn,omitempty"`

	// Refers to the identification of a previous event in the life of a trade which is amended or cancelled.
	RelatedReference *Max35Text `xml:"RltdRef"`
}

func (t *TradeAgreement2) SetTradeDate(value string) {
	t.TradeDate = (*ISODate)(&value)
}

func (t *TradeAgreement2) SetNotificationIdentification(value string) {
	t.NotificationIdentification = (*Max35Text)(&value)
}

func (t *TradeAgreement2) SetCommonReference(value string) {
	t.CommonReference = (*Max35Text)(&value)
}

func (t *TradeAgreement2) SetAmendOrCancelReason(value string) {
	t.AmendOrCancelReason = (*Max35Text)(&value)
}

func (t *TradeAgreement2) SetRelatedReference(value string) {
	t.RelatedReference = (*Max35Text)(&value)
}
