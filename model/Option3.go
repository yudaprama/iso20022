package model

// List of elements which provide the parameters of an option trade.
type Option3 struct {

	// Specifies the call and the put amount of the underlying foreign exchange trade.
	OptionAmounts *AmountsAndValueDate2 `xml:"OptnAmts"`

	// Specifies the rate of exchange at which the foreign exchange option has been struck.
	StrikePrice *AgreedRate1 `xml:"StrkPric"`

	// Defines how an option can be exercised.
	ExerciseStyle *OptionStyle2Code `xml:"ExrcStyle"`

	// First date on which an american option can be exercised.
	EarliestExerciseDate *ISODate `xml:"EarlstExrcDt,omitempty"`

	// Date on which a privilege (eg, option, right, warrant,...) expires. If it is an European option, the option holder can only exercise the right or let it lapse on expiry date. If it is an American option, the option holder can exercise the right up to the expiry date.
	ExpiryDateAndTime *ISODateTime `xml:"XpryDtAndTm"`

	// Financial center where option expires.
	ExpiryLocation *Max4AlphaNumericText `xml:"XpryLctn"`

	// Indicates whether the trade is to be settled as principal or netted off against another trade.
	SettlementType *SettlementType1Code `xml:"SttlmTp"`

	// Free format text that may contain information on the option.
	AdditionalOptionInformation *Max140Text `xml:"AddtlOptnInf,omitempty"`

	// Specifies the amount of the premium of a foreign exchange option trade and its settlement place.
	Premium *PremiumAmount2 `xml:"Prm"`
}

func (o *Option3) AddOptionAmounts() *AmountsAndValueDate2 {
	o.OptionAmounts = new(AmountsAndValueDate2)
	return o.OptionAmounts
}

func (o *Option3) AddStrikePrice() *AgreedRate1 {
	o.StrikePrice = new(AgreedRate1)
	return o.StrikePrice
}

func (o *Option3) SetExerciseStyle(value string) {
	o.ExerciseStyle = (*OptionStyle2Code)(&value)
}

func (o *Option3) SetEarliestExerciseDate(value string) {
	o.EarliestExerciseDate = (*ISODate)(&value)
}

func (o *Option3) SetExpiryDateAndTime(value string) {
	o.ExpiryDateAndTime = (*ISODateTime)(&value)
}

func (o *Option3) SetExpiryLocation(value string) {
	o.ExpiryLocation = (*Max4AlphaNumericText)(&value)
}

func (o *Option3) SetSettlementType(value string) {
	o.SettlementType = (*SettlementType1Code)(&value)
}

func (o *Option3) SetAdditionalOptionInformation(value string) {
	o.AdditionalOptionInformation = (*Max140Text)(&value)
}

func (o *Option3) AddPremium() *PremiumAmount2 {
	o.Premium = new(PremiumAmount2)
	return o.Premium
}
