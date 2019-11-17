package model

// Information about hold back and gating.
type HoldBackInformation2 struct {

	// Type of gating or a hold back.
	Type *GateHoldBack1Code `xml:"Tp"`

	// Value of the redemption amount subject to gating or a hold back.
	Amount *ActiveCurrencyAndAmount `xml:"Amt,omitempty"`

	// Date on which the gated amount or hold back amount is expected to be released.
	ExpectedReleaseDate *ISODate `xml:"XpctdRlsDt,omitempty"`

	// New identification of the security.
	FinancialInstrumentIdentification *SecurityIdentification25Choice `xml:"FinInstrmId,omitempty"`

	// New name of the security.
	FinancialInstrumentName *Max350Text `xml:"FinInstrmNm,omitempty"`

	// Specifies whether or not additional redemption order instructions are required in order for the redemption to be completed.
	RedemptionCompletion *RedemptionCompletion1Code `xml:"RedCmpltn,omitempty"`

	// Indicates whether or not this is the final redemption confirmation in the execution of a gated redemption.
	FinalConfirmation *YesNoIndicator `xml:"FnlConf,omitempty"`
}

func (h *HoldBackInformation2) SetType(value string) {
	h.Type = (*GateHoldBack1Code)(&value)
}

func (h *HoldBackInformation2) SetAmount(value, currency string) {
	h.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (h *HoldBackInformation2) SetExpectedReleaseDate(value string) {
	h.ExpectedReleaseDate = (*ISODate)(&value)
}

func (h *HoldBackInformation2) AddFinancialInstrumentIdentification() *SecurityIdentification25Choice {
	h.FinancialInstrumentIdentification = new(SecurityIdentification25Choice)
	return h.FinancialInstrumentIdentification
}

func (h *HoldBackInformation2) SetFinancialInstrumentName(value string) {
	h.FinancialInstrumentName = (*Max350Text)(&value)
}

func (h *HoldBackInformation2) SetRedemptionCompletion(value string) {
	h.RedemptionCompletion = (*RedemptionCompletion1Code)(&value)
}

func (h *HoldBackInformation2) SetFinalConfirmation(value string) {
	h.FinalConfirmation = (*YesNoIndicator)(&value)
}
