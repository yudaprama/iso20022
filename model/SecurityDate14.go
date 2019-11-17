package model

// Specifies security date details.
type SecurityDate14 struct {

	// Date/time at which the movement is due to take place (cash and/or securities).
	PaymentDate *DateFormat34Choice `xml:"PmtDt"`

	// Date/time at which securities become available for trading, for example first dealing date.
	AvailableDate *DateFormat34Choice `xml:"AvlblDt,omitempty"`

	// Date/time at which a security will be entitled to a dividend.
	DividendRankingDate *DateFormat34Choice `xml:"DvddRnkgDt,omitempty"`

	// Date/time at which a payment can be made, for example, if payment date is a non-business day or to indicate the first payment date of an offer.
	EarliestPaymentDate *DateFormat34Choice `xml:"EarlstPmtDt,omitempty"`

	// Date/time at which security will assimilate, become fungible, or have the same rights to dividends as the parent issue.
	PariPassuDate *DateFormat34Choice `xml:"PrpssDt,omitempty"`

	// Date/time at which the securities to be reorganised will cease to be tradeable.
	LastTradingDate *DateFormat34Choice `xml:"LastTradgDt,omitempty"`
}

func (s *SecurityDate14) AddPaymentDate() *DateFormat34Choice {
	s.PaymentDate = new(DateFormat34Choice)
	return s.PaymentDate
}

func (s *SecurityDate14) AddAvailableDate() *DateFormat34Choice {
	s.AvailableDate = new(DateFormat34Choice)
	return s.AvailableDate
}

func (s *SecurityDate14) AddDividendRankingDate() *DateFormat34Choice {
	s.DividendRankingDate = new(DateFormat34Choice)
	return s.DividendRankingDate
}

func (s *SecurityDate14) AddEarliestPaymentDate() *DateFormat34Choice {
	s.EarliestPaymentDate = new(DateFormat34Choice)
	return s.EarliestPaymentDate
}

func (s *SecurityDate14) AddPariPassuDate() *DateFormat34Choice {
	s.PariPassuDate = new(DateFormat34Choice)
	return s.PariPassuDate
}

func (s *SecurityDate14) AddLastTradingDate() *DateFormat34Choice {
	s.LastTradingDate = new(DateFormat34Choice)
	return s.LastTradingDate
}
