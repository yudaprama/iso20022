package model

// Provides the history of status and reasons for a pending, posted or cancelled transaction.
type StatusTrail7 struct {

	// Date and time at which the status was assigned.
	StatusDate *ISODateTime `xml:"StsDt"`

	// Unique and unambiguous way to identify the organisation that sent the message instance.
	SendingOrganisationIdentification *OrganisationIdentification9 `xml:"SndgOrgId,omitempty"`

	// Unique and unambiguous way to identify the user that created the message instance.
	UserIdentification *RestrictedFINXMax35Text `xml:"UsrId,omitempty"`

	// Provides details on the processing status of the transaction.
	ProcessingStatus *ProcessingStatus60Choice `xml:"PrcgSts,omitempty"`

	// Provides the matching status of an instruction as per the account servicer or the Market Infrastructure based on an allegement. At this time no matching took place on the market (at the CSD/ICSD/MI).
	InferredMatchingStatus *MatchingStatus30Choice `xml:"IfrrdMtchgSts,omitempty"`

	// Provides the matching status of the instruction.
	MatchingStatus *MatchingStatus30Choice `xml:"MtchgSts,omitempty"`

	// Provides the status of settlement of a transaction.
	SettlementStatus *SettlementStatus22Choice `xml:"SttlmSts,omitempty"`

	// Provides details on the modification processing status of the transaction.
	ModificationProcessingStatus *ModificationProcessingStatus8Choice `xml:"ModPrcgSts,omitempty"`

	// Provides details on the processing status of the cancellation request.
	CancellationStatus *ProcessingStatus61Choice `xml:"CxlSts,omitempty"`

	// Status is settled.
	Settled *ProprietaryReason5 `xml:"Sttld,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *StatusTrail7) SetStatusDate(value string) {
	s.StatusDate = (*ISODateTime)(&value)
}

func (s *StatusTrail7) AddSendingOrganisationIdentification() *OrganisationIdentification9 {
	s.SendingOrganisationIdentification = new(OrganisationIdentification9)
	return s.SendingOrganisationIdentification
}

func (s *StatusTrail7) SetUserIdentification(value string) {
	s.UserIdentification = (*RestrictedFINXMax35Text)(&value)
}

func (s *StatusTrail7) AddProcessingStatus() *ProcessingStatus60Choice {
	s.ProcessingStatus = new(ProcessingStatus60Choice)
	return s.ProcessingStatus
}

func (s *StatusTrail7) AddInferredMatchingStatus() *MatchingStatus30Choice {
	s.InferredMatchingStatus = new(MatchingStatus30Choice)
	return s.InferredMatchingStatus
}

func (s *StatusTrail7) AddMatchingStatus() *MatchingStatus30Choice {
	s.MatchingStatus = new(MatchingStatus30Choice)
	return s.MatchingStatus
}

func (s *StatusTrail7) AddSettlementStatus() *SettlementStatus22Choice {
	s.SettlementStatus = new(SettlementStatus22Choice)
	return s.SettlementStatus
}

func (s *StatusTrail7) AddModificationProcessingStatus() *ModificationProcessingStatus8Choice {
	s.ModificationProcessingStatus = new(ModificationProcessingStatus8Choice)
	return s.ModificationProcessingStatus
}

func (s *StatusTrail7) AddCancellationStatus() *ProcessingStatus61Choice {
	s.CancellationStatus = new(ProcessingStatus61Choice)
	return s.CancellationStatus
}

func (s *StatusTrail7) AddSettled() *ProprietaryReason5 {
	s.Settled = new(ProprietaryReason5)
	return s.Settled
}

func (s *StatusTrail7) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
