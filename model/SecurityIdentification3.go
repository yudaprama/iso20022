package model

// Identification of a security by its symbol.
type SecurityIdentification3 struct {

	// International Securities Identification Number (ISIN).  A numbering system designed by the United Nation's International Organisation for Standardisation (ISO). The ISIN is composed of a 2-character prefix representing the country of issue, followed by the national security number (if one exists), and a check digit. Each country has a national numbering agency that assigns ISIN numbers for securities in that country.
	ISIN *ISINIdentifier `xml:"ISIN"`

	// Letters that identify a stock traded on a stock exchange. The Ticker Symbol is a short and convenient way of identifying a stock, eg, RTR.L for Reuters quoted in London.
	TickerSymbol *TickerIdentifier `xml:"TckrSymb,omitempty"`

	// Committee on Uniform Securities and Identification Procedures (CUSIP). The standards body that created and maintains the securities classification system in the US. The CUSIP is composed of a 9-character number that uniquely identifies a particular security.  Non-US securities have a similar number called the CINS number.
	CUSIP *CUSIPIdentifier `xml:"CUSIP,omitempty"`

	// Stock Exchange Daily Official List (SEDOL) number.  A code used by the London Stock Exchange to identify foreign stocks, especially those that aren't actively traded in the US and don't have a CUSIP number.
	SEDOL *SEDOLIdentifier `xml:"SEDOL,omitempty"`

	// Identifier of a security assigned by the Japanese QUICK identification scheme for financial instruments.
	QUICK *QUICKIdentifier `xml:"QUICK,omitempty"`

	// Proprietary identification of a security assigned by an institution or organisation.
	OtherIdentification *AlternateFinancialInstrumentIdentification1 `xml:"OthrId,omitempty"`
}

func (s *SecurityIdentification3) SetISIN(value string) {
	s.ISIN = (*ISINIdentifier)(&value)
}

func (s *SecurityIdentification3) SetTickerSymbol(value string) {
	s.TickerSymbol = (*TickerIdentifier)(&value)
}

func (s *SecurityIdentification3) SetCUSIP(value string) {
	s.CUSIP = (*CUSIPIdentifier)(&value)
}

func (s *SecurityIdentification3) SetSEDOL(value string) {
	s.SEDOL = (*SEDOLIdentifier)(&value)
}

func (s *SecurityIdentification3) SetQUICK(value string) {
	s.QUICK = (*QUICKIdentifier)(&value)
}

func (s *SecurityIdentification3) AddOtherIdentification() *AlternateFinancialInstrumentIdentification1 {
	s.OtherIdentification = new(AlternateFinancialInstrumentIdentification1)
	return s.OtherIdentification
}
