package model

// Provides details on the settlement obligation report.
type Report5 struct {

	// Provides the identification for the non-clearing member. This is mandatory if the clearing member identification equals a general clearing member.
	NonClearingMember []*PartyIdentificationAndAccount31 `xml:"NonClrMmb,omitempty"`

	// Provides information about the settlement obligation details.
	SettlementObligationDetails []*SettlementObligation8 `xml:"SttlmOblgtnDtls"`
}

func (r *Report5) AddNonClearingMember() *PartyIdentificationAndAccount31 {
	newValue := new(PartyIdentificationAndAccount31)
	r.NonClearingMember = append(r.NonClearingMember, newValue)
	return newValue
}

func (r *Report5) AddSettlementObligationDetails() *SettlementObligation8 {
	newValue := new(SettlementObligation8)
	r.SettlementObligationDetails = append(r.SettlementObligationDetails, newValue)
	return newValue
}
