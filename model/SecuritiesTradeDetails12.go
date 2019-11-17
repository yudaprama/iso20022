package model

// Details of the securities trade.
type SecuritiesTradeDetails12 struct {

	// Specifies the date/time on which the trade was executed.
	TradeDate *DateAndDateTimeChoice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	OpeningSettlementDate *DateAndDateTimeChoice `xml:"OpngSttlmDt"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	InstructionProcessingAdditionalDetails *RestrictedFINXMax350Text `xml:"InstrPrcgAddtlDtls,omitempty"`
}

func (s *SecuritiesTradeDetails12) AddTradeDate() *DateAndDateTimeChoice {
	s.TradeDate = new(DateAndDateTimeChoice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails12) AddOpeningSettlementDate() *DateAndDateTimeChoice {
	s.OpeningSettlementDate = new(DateAndDateTimeChoice)
	return s.OpeningSettlementDate
}

func (s *SecuritiesTradeDetails12) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails12) SetInstructionProcessingAdditionalDetails(value string) {
	s.InstructionProcessingAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}
