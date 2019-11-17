package model

// Specifies security date details.
type SecurityDate12 struct {

	// Date/time at which the movement is due to take place (cash and/or securities).
	PaymentDate *DateFormat31Choice `xml:"PmtDt"`

	// Date/time at which securities become available for trading, for example first dealing date.
	AvailableDate *DateFormat31Choice `xml:"AvlblDt,omitempty"`

	// Date/time at which a security will be entitled to a dividend.
	DividendRankingDate *DateFormat31Choice `xml:"DvddRnkgDt,omitempty"`

	// Date/time at which a payment can be made, for example, if payment date is a non-business day or to indicate the first payment date of an offer.
	EarliestPaymentDate *DateFormat31Choice `xml:"EarlstPmtDt,omitempty"`

	// Date/time at which security will assimilate, become fungible, or have the same rights to dividends as the parent issue.
	PariPassuDate *DateFormat31Choice `xml:"PrpssDt,omitempty"`

	// Date/time at which the securities to be reorganised will cease to be tradeable.
	LastTradingDate *DateFormat31Choice `xml:"LastTradgDt,omitempty"`
}

func (s *SecurityDate12) AddPaymentDate() *DateFormat31Choice {
	s.PaymentDate = new(DateFormat31Choice)
	return s.PaymentDate
}

func (s *SecurityDate12) AddAvailableDate() *DateFormat31Choice {
	s.AvailableDate = new(DateFormat31Choice)
	return s.AvailableDate
}

func (s *SecurityDate12) AddDividendRankingDate() *DateFormat31Choice {
	s.DividendRankingDate = new(DateFormat31Choice)
	return s.DividendRankingDate
}

func (s *SecurityDate12) AddEarliestPaymentDate() *DateFormat31Choice {
	s.EarliestPaymentDate = new(DateFormat31Choice)
	return s.EarliestPaymentDate
}

func (s *SecurityDate12) AddPariPassuDate() *DateFormat31Choice {
	s.PariPassuDate = new(DateFormat31Choice)
	return s.PariPassuDate
}

func (s *SecurityDate12) AddLastTradingDate() *DateFormat31Choice {
	s.LastTradingDate = new(DateFormat31Choice)
	return s.LastTradingDate
}
