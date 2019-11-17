package model

// Specifies security date details.
type SecurityDate13 struct {

	// Date of the posting (credit or debit) to the account.
	PostingDate *DateAndDateTimeChoice `xml:"PstngDt"`

	// Date/time at which securities become available for trading, for example first dealing date.
	AvailableDate *DateFormat34Choice `xml:"AvlblDt,omitempty"`

	// Date/time at which security will assimilate, become fungible, or have the same rights to dividends as the parent issue.
	PariPassuDate *DateFormat34Choice `xml:"PrpssDt,omitempty"`

	// Date/time at which a security will be entitled to a dividend.
	DividendRankingDate *DateFormat34Choice `xml:"DvddRnkgDt,omitempty"`

	// Date/time at which a payment can be made, for example, if payment date is a non-business day or to indicate the first payment date of an offer.
	EarliestPaymentDate *DateFormat34Choice `xml:"EarlstPmtDt,omitempty"`

	// Date/time at which the movement is due to take place (cash and/or securities).
	PaymentDate *DateFormat34Choice `xml:"PmtDt,omitempty"`
}

func (s *SecurityDate13) AddPostingDate() *DateAndDateTimeChoice {
	s.PostingDate = new(DateAndDateTimeChoice)
	return s.PostingDate
}

func (s *SecurityDate13) AddAvailableDate() *DateFormat34Choice {
	s.AvailableDate = new(DateFormat34Choice)
	return s.AvailableDate
}

func (s *SecurityDate13) AddPariPassuDate() *DateFormat34Choice {
	s.PariPassuDate = new(DateFormat34Choice)
	return s.PariPassuDate
}

func (s *SecurityDate13) AddDividendRankingDate() *DateFormat34Choice {
	s.DividendRankingDate = new(DateFormat34Choice)
	return s.DividendRankingDate
}

func (s *SecurityDate13) AddEarliestPaymentDate() *DateFormat34Choice {
	s.EarliestPaymentDate = new(DateFormat34Choice)
	return s.EarliestPaymentDate
}

func (s *SecurityDate13) AddPaymentDate() *DateFormat34Choice {
	s.PaymentDate = new(DateFormat34Choice)
	return s.PaymentDate
}
