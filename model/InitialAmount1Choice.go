package model

// Choice of an initial amount or number of pre-paid instalments.
type InitialAmount1Choice struct {

	// Number of pre-paid instalment periods at the time the investment plan is created.
	InitialNumberOfInstalments *Number `xml:"InitlNbOfInstlmts"`

	// Amount to be paid or redeemed at the time the investment plan is created.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`
}

func (i *InitialAmount1Choice) SetInitialNumberOfInstalments(value string) {
	i.InitialNumberOfInstalments = (*Number)(&value)
}

func (i *InitialAmount1Choice) SetAmount(value, currency string) {
	i.Amount = NewActiveCurrencyAndAmount(value, currency)
}
