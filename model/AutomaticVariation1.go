package model

// Variation parameters and triggers.
type AutomaticVariation1 struct {

	// Identification of the automatic variation.
	Identification *Max35Text `xml:"Id"`

	// Type of variation.
	Type *VariationType1Code `xml:"Tp"`

	// Details related to the variation amount and its trigger.
	AmountAndTrigger []*AmountAndTrigger1 `xml:"AmtAndTrggr"`

	// Additional information related to the variation.
	AdditionalInformation []*Max2000Text `xml:"AddtlInf,omitempty"`
}

func (a *AutomaticVariation1) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AutomaticVariation1) SetType(value string) {
	a.Type = (*VariationType1Code)(&value)
}

func (a *AutomaticVariation1) AddAmountAndTrigger() *AmountAndTrigger1 {
	newValue := new(AmountAndTrigger1)
	a.AmountAndTrigger = append(a.AmountAndTrigger, newValue)
	return newValue
}

func (a *AutomaticVariation1) AddAdditionalInformation(value string) {
	a.AdditionalInformation = append(a.AdditionalInformation, (*Max2000Text)(&value))
}
