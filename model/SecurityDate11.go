package model

// Specifies security date details.
type SecurityDate11 struct {

	// Date of the posting (credit or debit) to the account.
	PostingDate *DateAndDateTimeChoice `xml:"PstngDt"`

	// Date/time at which securities become available for trading, for example first dealing date.
	AvailableDate *DateFormat31Choice `xml:"AvlblDt,omitempty"`

	// Date/time at which security will assimilate, become fungible, or have the same rights to dividends as the parent issue.
	PariPassuDate *DateFormat31Choice `xml:"PrpssDt,omitempty"`

	// Date/time at which a security will be entitled to a dividend.
	DividendRankingDate *DateFormat31Choice `xml:"DvddRnkgDt,omitempty"`

	// Date/time at which a payment can be made, for example, if payment date is a non-business day or to indicate the first payment date of an offer.
	EarliestPaymentDate *DateFormat31Choice `xml:"EarlstPmtDt,omitempty"`

	// Date/time at which the movement is due to take place (cash and/or securities).
	PaymentDate *DateFormat31Choice `xml:"PmtDt,omitempty"`
}

func (s *SecurityDate11) AddPostingDate() *DateAndDateTimeChoice {
	s.PostingDate = new(DateAndDateTimeChoice)
	return s.PostingDate
}

func (s *SecurityDate11) AddAvailableDate() *DateFormat31Choice {
	s.AvailableDate = new(DateFormat31Choice)
	return s.AvailableDate
}

func (s *SecurityDate11) AddPariPassuDate() *DateFormat31Choice {
	s.PariPassuDate = new(DateFormat31Choice)
	return s.PariPassuDate
}

func (s *SecurityDate11) AddDividendRankingDate() *DateFormat31Choice {
	s.DividendRankingDate = new(DateFormat31Choice)
	return s.DividendRankingDate
}

func (s *SecurityDate11) AddEarliestPaymentDate() *DateFormat31Choice {
	s.EarliestPaymentDate = new(DateFormat31Choice)
	return s.EarliestPaymentDate
}

func (s *SecurityDate11) AddPaymentDate() *DateFormat31Choice {
	s.PaymentDate = new(DateFormat31Choice)
	return s.PaymentDate
}
