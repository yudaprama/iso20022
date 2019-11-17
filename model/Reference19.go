package model

// Provides the references of the underlying trade leg(s) and/or the reference to the related NetPosition report message.
type Reference19 struct {

	// Reference allocated by the central counterparty - central counterpatry trade leg reference identification that uniquely identifies the trade.
	TradeLegNotificationIdentification []*Max35Text `xml:"TradLegNtfctnId,omitempty"`

	// After netting, reference that is common to a net transaction to settle and all its underlying trades
	NetPositionIdentification *Max35Text `xml:"NetPosId,omitempty"`
}

func (r *Reference19) AddTradeLegNotificationIdentification(value string) {
	r.TradeLegNotificationIdentification = append(r.TradeLegNotificationIdentification, (*Max35Text)(&value))
}

func (r *Reference19) SetNetPositionIdentification(value string) {
	r.NetPositionIdentification = (*Max35Text)(&value)
}
