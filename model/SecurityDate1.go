package model

// Specifies security date details.
type SecurityDate1 struct {

	// Date of the posting (credit or debit) to the account.
	PostingDate *DateAndDateTimeChoice `xml:"PstngDt"`

	// Date/time at which securities become available for trading, for example first dealing date.
	AvailableDate *DateFormat6Choice `xml:"AvlblDt,omitempty"`

	// Date/time at which security will assimilate, become fungible, or have the same rights to dividends as the parent issue.
	PariPassuDate *DateFormat6Choice `xml:"PrpssDt,omitempty"`

	// Date/time at which a security will be entitled to a dividend.
	DividendRankingDate *DateFormat6Choice `xml:"DvddRnkgDt,omitempty"`

	// Date/time at which a payment can be made, for example, if payment date is a non-business day or to indicate the first payment date of an offer.
	EarliestPaymentDate *DateFormat6Choice `xml:"EarlstPmtDt,omitempty"`

	// Date/time at which the movement is due to take place (cash and/or securities).
	PaymentDate *DateFormat6Choice `xml:"PmtDt,omitempty"`
}

func (s *SecurityDate1) AddPostingDate() *DateAndDateTimeChoice {
	s.PostingDate = new(DateAndDateTimeChoice)
	return s.PostingDate
}

func (s *SecurityDate1) AddAvailableDate() *DateFormat6Choice {
	s.AvailableDate = new(DateFormat6Choice)
	return s.AvailableDate
}

func (s *SecurityDate1) AddPariPassuDate() *DateFormat6Choice {
	s.PariPassuDate = new(DateFormat6Choice)
	return s.PariPassuDate
}

func (s *SecurityDate1) AddDividendRankingDate() *DateFormat6Choice {
	s.DividendRankingDate = new(DateFormat6Choice)
	return s.DividendRankingDate
}

func (s *SecurityDate1) AddEarliestPaymentDate() *DateFormat6Choice {
	s.EarliestPaymentDate = new(DateFormat6Choice)
	return s.EarliestPaymentDate
}

func (s *SecurityDate1) AddPaymentDate() *DateFormat6Choice {
	s.PaymentDate = new(DateFormat6Choice)
	return s.PaymentDate
}
