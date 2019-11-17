package model

// Automatic extension information.
type AutoExtension1 struct {

	// Indicates that the undertaking is automatically extendable and the period of extension.
	Period *AutoExtend1Choice `xml:"Prd,omitempty"`

	// Final expiry date after which the undertaking will no longer be subject to automatic extension.
	FinalExpiryDate *ISODate `xml:"FnlXpryDt,omitempty"`

	// Details related to the notification of the end of the period for notification of non-extension of the expiry date.
	NonExtensionNotification []*NonExtension1 `xml:"NonXtnsnNtfctn,omitempty"`
}

func (a *AutoExtension1) AddPeriod() *AutoExtend1Choice {
	a.Period = new(AutoExtend1Choice)
	return a.Period
}

func (a *AutoExtension1) SetFinalExpiryDate(value string) {
	a.FinalExpiryDate = (*ISODate)(&value)
}

func (a *AutoExtension1) AddNonExtensionNotification() *NonExtension1 {
	newValue := new(NonExtension1)
	a.NonExtensionNotification = append(a.NonExtensionNotification, newValue)
	return newValue
}
