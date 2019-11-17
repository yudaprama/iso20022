package model

// Loan offered by a group of lenders (called a syndicate) who work together to lend an amount of money to a single borrower.
type SyndicatedLoan1 struct {

	// Party which obtains the loan.
	Borrower *TradeParty2 `xml:"Brrwr"`

	// Party which provides an amount of money available to others to borrow.
	Lender *TradeParty2 `xml:"Lndr,omitempty"`

	// Amount of the part in the syndicated loan.
	Amount *ActiveCurrencyAndAmount `xml:"Amt,omitempty"`

	// Share of the part in the syndicated loan.
	Share *Percentage `xml:"Shr,omitempty"`

	// Provides details on the currency exchange rate and contract.
	ExchangeRateInformation *ExchangeRate1 `xml:"XchgRateInf,omitempty"`
}

func (s *SyndicatedLoan1) AddBorrower() *TradeParty2 {
	s.Borrower = new(TradeParty2)
	return s.Borrower
}

func (s *SyndicatedLoan1) AddLender() *TradeParty2 {
	s.Lender = new(TradeParty2)
	return s.Lender
}

func (s *SyndicatedLoan1) SetAmount(value, currency string) {
	s.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SyndicatedLoan1) SetShare(value string) {
	s.Share = (*Percentage)(&value)
}

func (s *SyndicatedLoan1) AddExchangeRateInformation() *ExchangeRate1 {
	s.ExchangeRateInformation = new(ExchangeRate1)
	return s.ExchangeRateInformation
}
