package model

// Identifies the security instrument by its name and typical characteristics.
type SecurityInstrumentDescription2 struct {

	// Description of the security.
	Description *Max350Text `xml:"Desc,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI) codification, eg, common share with voting rights, fully paid, or registered.
	ClassificationType *SecurityClassificationType1Choice `xml:"ClssfctnTp,omitempty"`

	// Provides the place of listing using a market identifier code (MIC).
	PlaceOfListing *MICIdentifier `xml:"PlcOfListg,omitempty"`

	// Exercise date/time of a derivative contract.
	ExerciseDate *ISODate `xml:"ExrcDt,omitempty"`

	// Maturity date/time at which an interest bearing security becomes due.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Specifies whether it is a call option (right to purchase a specific underlying asset) or a put option (right to sell a specific underlying asset).
	OptionType *OptionTypeCode `xml:"OptnTp,omitempty"`

	// Predetermined price at which the holder will have to buy or sell the underlying instrument.
	StrikePrice *PriceRateOrAmountChoice `xml:"StrkPric,omitempty"`

	// Indicates the ratio or multiplying factor used to convert one contract into a quantity. In the case of an equity or a bond, the price multiplier is 1.
	Multiplier *BaseOneRate `xml:"Mltplr,omitempty"`
}

func (s *SecurityInstrumentDescription2) SetDescription(value string) {
	s.Description = (*Max350Text)(&value)
}

func (s *SecurityInstrumentDescription2) AddClassificationType() *SecurityClassificationType1Choice {
	s.ClassificationType = new(SecurityClassificationType1Choice)
	return s.ClassificationType
}

func (s *SecurityInstrumentDescription2) SetPlaceOfListing(value string) {
	s.PlaceOfListing = (*MICIdentifier)(&value)
}

func (s *SecurityInstrumentDescription2) SetExerciseDate(value string) {
	s.ExerciseDate = (*ISODate)(&value)
}

func (s *SecurityInstrumentDescription2) SetMaturityDate(value string) {
	s.MaturityDate = (*ISODate)(&value)
}

func (s *SecurityInstrumentDescription2) SetOptionType(value string) {
	s.OptionType = (*OptionTypeCode)(&value)
}

func (s *SecurityInstrumentDescription2) AddStrikePrice() *PriceRateOrAmountChoice {
	s.StrikePrice = new(PriceRateOrAmountChoice)
	return s.StrikePrice
}

func (s *SecurityInstrumentDescription2) SetMultiplier(value string) {
	s.Multiplier = (*BaseOneRate)(&value)
}
