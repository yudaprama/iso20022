package model

// Provides information about the corporate action security option.
type SecuritiesOption53 struct {

	// Identification of the financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`

	// Specifies whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Quantity of securities that have been posted (credit or debit) to the safekeeping account.
	PostingQuantity *Quantity6Choice `xml:"PstngQty"`

	// Date of the posting (credit or debit) to the account.
	PostingDate *DateAndDateTimeChoice `xml:"PstngDt"`

	// Date/Time of the posting (credit or debit) to the account that was initially communicated in the confirmation.
	OriginalPostingDate *DateAndDateTimeChoice `xml:"OrgnlPstngDt,omitempty"`
}

func (s *SecuritiesOption53) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesOption53) SetCreditDebitIndicator(value string) {
	s.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (s *SecuritiesOption53) AddPostingQuantity() *Quantity6Choice {
	s.PostingQuantity = new(Quantity6Choice)
	return s.PostingQuantity
}

func (s *SecuritiesOption53) AddPostingDate() *DateAndDateTimeChoice {
	s.PostingDate = new(DateAndDateTimeChoice)
	return s.PostingDate
}

func (s *SecuritiesOption53) AddOriginalPostingDate() *DateAndDateTimeChoice {
	s.OriginalPostingDate = new(DateAndDateTimeChoice)
	return s.OriginalPostingDate
}
