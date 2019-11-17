package model

// Choice between formats for the identification of a financial instrument.
type SecurityIdentification3Choice struct {

	// International Securities Identification Number (ISIN).  A numbering system designed by the United Nation's International Organisation for Standardisation (ISO). The ISIN is composed of a 2-character prefix representing the country of issue, followed by the national security number (if one exists), and a check digit. Each country has a national numbering agency that assigns ISIN numbers for securities in that country.
	ISIN *ISINIdentifier `xml:"ISIN"`

	// Stock Exchange Daily Official List (SEDOL) number.  A code used by the London Stock Exchange to identify foreign stocks, especially those that aren't actively traded in the US and don't have a CUSIP number.
	SEDOL *SEDOLIdentifier `xml:"SEDOL"`

	// Committee on Uniform Securities and Identification Procedures (CUSIP). The standards body that created and maintains the securities classification system in the US. The CUSIP is composed of a 9-character number that uniquely identifies a particular security.  Non-US securities have a similar number called the CINS number.
	CUSIP *CUSIPIdentifier `xml:"CUSIP"`

	// Reuters Identification Code (RIC). A numbering system used within the Reuters system to identify instruments worldwide. The RIC contains an X-character market specific code (can be the CUSIP or EPIC codes) followed by a full stop, then the two-digit ISO country code, eg, IBM in UK is IBM.UK.
	RIC *RICIdentifier `xml:"RIC"`

	// Letters that identify a stock traded on a stock exchange. The Ticker Symbol is a short and convenient way of identifying a stock, eg, RTR.L for Reuters quoted in London.
	TickerSymbol *TickerIdentifier `xml:"TckrSymb"`

	// Identifier of a security assigned by the Bloomberg organisation.
	Bloomberg *BloombergIdentifier `xml:"Blmbrg"`

	// Identifier of a security assigned by the Consolidated Tape Association.
	CTA *ConsolidatedTapeAssociationIdentifier `xml:"CTA"`

	// Identifier of a security assigned by the Japanese QUICK identification scheme for financial instruments.
	QUICK *QUICKIdentifier `xml:"QUICK"`

	// Wertpapier Kenn-nummer.  A number issued in Germany by the Wertpapier Mitteilungen. The Wertpapier Kenn-nummer, sometimes called WPK, contains 6-digits, but no check digit. There are different ranges of numbers representing different classes of securities.
	Wertpapier *WertpapierIdentifier `xml:"Wrtppr"`

	// Identifier for Dutch securities.
	Dutch *DutchIdentifier `xml:"Dtch"`

	// Identifier for Swiss securities assigned by Telekurs Financial, the Swiss numbering agency.
	Valoren *ValorenIdentifier `xml:"Vlrn"`

	// Identifier for French securities assigned by the Societe Interprofessionnelle Pour La Compensation des Valeurs Mobilieres in France.  The Sicovam is composed of 5-digits.
	Sicovam *SicovamIdentifier `xml:"SCVM"`

	// Identifier for Belgian securities.
	Belgian *BelgianIdentifier `xml:"Belgn"`

	// Identifier of securities issued in Luxembourg.  The common code is a 9-digit code that replaces the CEDEL (Clearstream) and Euroclear codes.
	Common *EuroclearClearstreamIdentifier `xml:"Cmon"`

	// Proprietary identification of a security assigned by an institution or organisation.
	OtherProprietaryIdentification *AlternateSecurityIdentification1 `xml:"OthrPrtryId"`
}

func (s *SecurityIdentification3Choice) SetISIN(value string) {
	s.ISIN = (*ISINIdentifier)(&value)
}

func (s *SecurityIdentification3Choice) SetSEDOL(value string) {
	s.SEDOL = (*SEDOLIdentifier)(&value)
}

func (s *SecurityIdentification3Choice) SetCUSIP(value string) {
	s.CUSIP = (*CUSIPIdentifier)(&value)
}

func (s *SecurityIdentification3Choice) SetRIC(value string) {
	s.RIC = (*RICIdentifier)(&value)
}

func (s *SecurityIdentification3Choice) SetTickerSymbol(value string) {
	s.TickerSymbol = (*TickerIdentifier)(&value)
}

func (s *SecurityIdentification3Choice) SetBloomberg(value string) {
	s.Bloomberg = (*BloombergIdentifier)(&value)
}

func (s *SecurityIdentification3Choice) SetCTA(value string) {
	s.CTA = (*ConsolidatedTapeAssociationIdentifier)(&value)
}

func (s *SecurityIdentification3Choice) SetQUICK(value string) {
	s.QUICK = (*QUICKIdentifier)(&value)
}

func (s *SecurityIdentification3Choice) SetWertpapier(value string) {
	s.Wertpapier = (*WertpapierIdentifier)(&value)
}

func (s *SecurityIdentification3Choice) SetDutch(value string) {
	s.Dutch = (*DutchIdentifier)(&value)
}

func (s *SecurityIdentification3Choice) SetValoren(value string) {
	s.Valoren = (*ValorenIdentifier)(&value)
}

func (s *SecurityIdentification3Choice) SetSicovam(value string) {
	s.Sicovam = (*SicovamIdentifier)(&value)
}

func (s *SecurityIdentification3Choice) SetBelgian(value string) {
	s.Belgian = (*BelgianIdentifier)(&value)
}

func (s *SecurityIdentification3Choice) SetCommon(value string) {
	s.Common = (*EuroclearClearstreamIdentifier)(&value)
}

func (s *SecurityIdentification3Choice) AddOtherProprietaryIdentification() *AlternateSecurityIdentification1 {
	s.OtherProprietaryIdentification = new(AlternateSecurityIdentification1)
	return s.OtherProprietaryIdentification
}
