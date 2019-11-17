package model

// Provides the history of status and reasons for a pending, posted or cancelled transaction.
type StatusTrail6 struct {

	// Date and time at which the status was assigned.
	StatusDate *ISODateTime `xml:"StsDt"`

	// Unique and unambiguous way to identify the organisation that sent the message instance.
	SendingOrganisationIdentification *OrganisationIdentification7 `xml:"SndgOrgId,omitempty"`

	// Unique and unambiguous way to identify the user that created the message instance.
	UserIdentification *Max35Text `xml:"UsrId,omitempty"`

	// Provides details on the processing status of the transaction.
	ProcessingStatus *ProcessingStatus49Choice `xml:"PrcgSts,omitempty"`

	// Provides the matching status of an instruction as per the account servicer or the Market Infrastructure based on an allegement. At this time no matching took place on the market (at the CSD/ICSD/MI).
	InferredMatchingStatus *MatchingStatus25Choice `xml:"IfrrdMtchgSts,omitempty"`

	// Provides the matching status of the instruction.
	MatchingStatus *MatchingStatus25Choice `xml:"MtchgSts,omitempty"`

	// Provides the status of settlement of a transaction.
	SettlementStatus *SettlementStatus17Choice `xml:"SttlmSts,omitempty"`

	// Provides details on the modification processing status of the transaction.
	ModificationProcessingStatus *ModificationProcessingStatus7Choice `xml:"ModPrcgSts,omitempty"`

	// Provides details on the processing status of the cancellation request.
	CancellationStatus *ProcessingStatus53Choice `xml:"CxlSts,omitempty"`

	// Status is settled.
	Settled *ProprietaryReason4 `xml:"Sttld,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *StatusTrail6) SetStatusDate(value string) {
	s.StatusDate = (*ISODateTime)(&value)
}

func (s *StatusTrail6) AddSendingOrganisationIdentification() *OrganisationIdentification7 {
	s.SendingOrganisationIdentification = new(OrganisationIdentification7)
	return s.SendingOrganisationIdentification
}

func (s *StatusTrail6) SetUserIdentification(value string) {
	s.UserIdentification = (*Max35Text)(&value)
}

func (s *StatusTrail6) AddProcessingStatus() *ProcessingStatus49Choice {
	s.ProcessingStatus = new(ProcessingStatus49Choice)
	return s.ProcessingStatus
}

func (s *StatusTrail6) AddInferredMatchingStatus() *MatchingStatus25Choice {
	s.InferredMatchingStatus = new(MatchingStatus25Choice)
	return s.InferredMatchingStatus
}

func (s *StatusTrail6) AddMatchingStatus() *MatchingStatus25Choice {
	s.MatchingStatus = new(MatchingStatus25Choice)
	return s.MatchingStatus
}

func (s *StatusTrail6) AddSettlementStatus() *SettlementStatus17Choice {
	s.SettlementStatus = new(SettlementStatus17Choice)
	return s.SettlementStatus
}

func (s *StatusTrail6) AddModificationProcessingStatus() *ModificationProcessingStatus7Choice {
	s.ModificationProcessingStatus = new(ModificationProcessingStatus7Choice)
	return s.ModificationProcessingStatus
}

func (s *StatusTrail6) AddCancellationStatus() *ProcessingStatus53Choice {
	s.CancellationStatus = new(ProcessingStatus53Choice)
	return s.CancellationStatus
}

func (s *StatusTrail6) AddSettled() *ProprietaryReason4 {
	s.Settled = new(ProprietaryReason4)
	return s.Settled
}

func (s *StatusTrail6) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
