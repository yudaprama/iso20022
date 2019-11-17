package model

// Amount and trigger information.
type AmountAndTrigger1 struct {

	// Identification of the amount and trigger details.
	Identification *Max35Text `xml:"Id"`

	// Choice between an amount and a percentage.
	AmountDetailsChoice *AmountOrPercentage1Choice `xml:"AmtDtlsChc"`

	// Trigger that causes the variation to come into effect.
	Trigger []*Trigger1 `xml:"Trggr"`
}

func (a *AmountAndTrigger1) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AmountAndTrigger1) AddAmountDetailsChoice() *AmountOrPercentage1Choice {
	a.AmountDetailsChoice = new(AmountOrPercentage1Choice)
	return a.AmountDetailsChoice
}

func (a *AmountAndTrigger1) AddTrigger() *Trigger1 {
	newValue := new(Trigger1)
	a.Trigger = append(a.Trigger, newValue)
	return newValue
}
