package model

// Choice of leg.
type LegIdentification1Choice struct {

	// Unique technical identifier for the instance of the leg within a switch.
	RedemptionLegIdentification *Max35Text `xml:"RedLegId"`

	// Unique technical identifier for the instance of the leg within a switch.
	SubscriptionLegIdentification *Max35Text `xml:"SbcptLegId"`
}

func (l *LegIdentification1Choice) SetRedemptionLegIdentification(value string) {
	l.RedemptionLegIdentification = (*Max35Text)(&value)
}

func (l *LegIdentification1Choice) SetSubscriptionLegIdentification(value string) {
	l.SubscriptionLegIdentification = (*Max35Text)(&value)
}
