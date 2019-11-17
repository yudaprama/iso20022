package model

// Provides the history of status and reasons for a pending, posted or cancelled transaction.
type StatusTrail2 struct {

	// Date and time at which the status was assigned.
	StatusDate *ISODateTime `xml:"StsDt"`

	// Unique and unambiguous way to identify the organisation that sent the message instance.
	SendingOrganisationIdentification *OrganisationIdentification7 `xml:"SndgOrgId,omitempty"`

	// Unique and unambiguous way to identify the user that created the message instance.
	UserIdentification *Max35Text `xml:"UsrId,omitempty"`

	// Provides details on the processing status of the transaction.
	ProcessingStatus *ProcessingStatus19Choice `xml:"PrcgSts,omitempty"`

	// Provides the matching status of an instruction as per the account servicer or the Market Infrastructure based on an allegement. At this time no matching took place on the market (at the CSD/ICSD/MI).
	InferredMatchingStatus *MatchingStatus7Choice `xml:"IfrrdMtchgSts,omitempty"`

	// Provides the matching status of the instruction.
	MatchingStatus *MatchingStatus7Choice `xml:"MtchgSts,omitempty"`

	// Provides the status of settlement of a transaction.
	SettlementStatus *SettlementStatus7Choice `xml:"SttlmSts,omitempty"`

	// Provides details on the modification processing status of the transaction.
	ModificationProcessingStatus *ModificationProcessingStatus2Choice `xml:"ModPrcgSts,omitempty"`

	// Provides details on the processing status of the cancellation request.
	CancellationStatus *ProcessingStatus20Choice `xml:"CxlSts,omitempty"`

	// Status is settled.
	Settled *ProprietaryReason1 `xml:"Sttld,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *StatusTrail2) SetStatusDate(value string) {
	s.StatusDate = (*ISODateTime)(&value)
}

func (s *StatusTrail2) AddSendingOrganisationIdentification() *OrganisationIdentification7 {
	s.SendingOrganisationIdentification = new(OrganisationIdentification7)
	return s.SendingOrganisationIdentification
}

func (s *StatusTrail2) SetUserIdentification(value string) {
	s.UserIdentification = (*Max35Text)(&value)
}

func (s *StatusTrail2) AddProcessingStatus() *ProcessingStatus19Choice {
	s.ProcessingStatus = new(ProcessingStatus19Choice)
	return s.ProcessingStatus
}

func (s *StatusTrail2) AddInferredMatchingStatus() *MatchingStatus7Choice {
	s.InferredMatchingStatus = new(MatchingStatus7Choice)
	return s.InferredMatchingStatus
}

func (s *StatusTrail2) AddMatchingStatus() *MatchingStatus7Choice {
	s.MatchingStatus = new(MatchingStatus7Choice)
	return s.MatchingStatus
}

func (s *StatusTrail2) AddSettlementStatus() *SettlementStatus7Choice {
	s.SettlementStatus = new(SettlementStatus7Choice)
	return s.SettlementStatus
}

func (s *StatusTrail2) AddModificationProcessingStatus() *ModificationProcessingStatus2Choice {
	s.ModificationProcessingStatus = new(ModificationProcessingStatus2Choice)
	return s.ModificationProcessingStatus
}

func (s *StatusTrail2) AddCancellationStatus() *ProcessingStatus20Choice {
	s.CancellationStatus = new(ProcessingStatus20Choice)
	return s.CancellationStatus
}

func (s *StatusTrail2) AddSettled() *ProprietaryReason1 {
	s.Settled = new(ProprietaryReason1)
	return s.Settled
}

func (s *StatusTrail2) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
