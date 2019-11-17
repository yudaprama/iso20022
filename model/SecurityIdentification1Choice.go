package model

// Choice between formats for the identification of a financial instrument.
type SecurityIdentification1Choice struct {

	// International Securities Identification Number (ISIN).  A numbering system designed by the United Nation's International Organisation for Standardisation (ISO). The ISIN is composed of a 2-character prefix representing the country of issue, followed by the national security number (if one exists), and a check digit. Each country has a national numbering agency that assigns ISIN numbers for securities in that country.
	ISIN *ISINIdentifier `xml:"ISIN"`

	// Proprietary identification of a security assigned by an institution or organisation.
	AlternateIdentification *AlternateSecurityIdentification1 `xml:"AltrnId"`

	// Reuters Identification Code (RIC). A numbering system used within the Reuters system to identify instruments worldwide. The RIC contains an X-character market specific code (can be the CUSIP or EPIC codes) followed by a full stop, then the two-digit ISO country code, eg, IBM in UK is IBM.UK.
	RIC *RICIdentifier `xml:"RIC"`

	// Letters that identify a stock traded on a stock exchange. The Ticker Symbol is a short and convenient way of identifying a stock, eg, RTR.L for Reuters quoted in London.
	TickerSymbol *TickerIdentifier `xml:"TckrSymb"`

	// Identifier of a security assigned by the Bloomberg organisation.
	Bloomberg *BloombergIdentifier `xml:"Blmbrg"`

	// Identifier of a security assigned by the Consolidated Tape Association.
	CTA *ConsolidatedTapeAssociationIdentifier `xml:"CTA"`

	// Identifier of securities issued in Luxembourg.  The common code is a 9-digit code that replaces the CEDEL (Clearstream) and Euroclear codes.
	Common *EuroclearClearstreamIdentifier `xml:"Cmon"`
}

func (s *SecurityIdentification1Choice) SetISIN(value string) {
	s.ISIN = (*ISINIdentifier)(&value)
}

func (s *SecurityIdentification1Choice) AddAlternateIdentification() *AlternateSecurityIdentification1 {
	s.AlternateIdentification = new(AlternateSecurityIdentification1)
	return s.AlternateIdentification
}

func (s *SecurityIdentification1Choice) SetRIC(value string) {
	s.RIC = (*RICIdentifier)(&value)
}

func (s *SecurityIdentification1Choice) SetTickerSymbol(value string) {
	s.TickerSymbol = (*TickerIdentifier)(&value)
}

func (s *SecurityIdentification1Choice) SetBloomberg(value string) {
	s.Bloomberg = (*BloombergIdentifier)(&value)
}

func (s *SecurityIdentification1Choice) SetCTA(value string) {
	s.CTA = (*ConsolidatedTapeAssociationIdentifier)(&value)
}

func (s *SecurityIdentification1Choice) SetCommon(value string) {
	s.Common = (*EuroclearClearstreamIdentifier)(&value)
}
