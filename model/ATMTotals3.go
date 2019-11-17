package model

// Totals of the ATM.
type ATMTotals3 struct {

	// Identification of the type of transaction. The following values are predefined: Withdrawal, QuickWithdrawal, WithdrawalForDisabledPerson, CashDeposit, Transfer, InternationalTransfer, PeriodicTransfer, CheckCommand, MobileTopUp, Moneo.
	Identification *Max70Text `xml:"Id"`

	// Additional identification of the type of transaction. The following values are predefined:  Vodaphone, TMobile, Verizon.
	AdditionalIdentification *Max70Text `xml:"AddtlId,omitempty"`

	// Period of computation for the counters.
	Period *ATMCounterType2Code `xml:"Prd"`

	// Currency of the totals.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Number of transaction with the defined currency.
	Count *Number `xml:"Cnt"`

	// Amount of transaction with the defined currency.
	Amount *ImpliedCurrencyAndAmount `xml:"Amt,omitempty"`
}

func (a *ATMTotals3) SetIdentification(value string) {
	a.Identification = (*Max70Text)(&value)
}

func (a *ATMTotals3) SetAdditionalIdentification(value string) {
	a.AdditionalIdentification = (*Max70Text)(&value)
}

func (a *ATMTotals3) SetPeriod(value string) {
	a.Period = (*ATMCounterType2Code)(&value)
}

func (a *ATMTotals3) SetCurrency(value string) {
	a.Currency = (*ActiveCurrencyCode)(&value)
}

func (a *ATMTotals3) SetCount(value string) {
	a.Count = (*Number)(&value)
}

func (a *ATMTotals3) SetAmount(value, currency string) {
	a.Amount = NewImpliedCurrencyAndAmount(value, currency)
}
