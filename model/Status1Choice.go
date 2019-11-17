package model

// Choice of status.
type Status1Choice struct {

	// Status of affirmation of a trade.
	AffirmationStatus *AffirmationStatus1Choice `xml:"AffirmSts"`

	// Provides the status of allocation of collateral to cover the instruction.
	AllocationStatus *AllocationSatus1Choice `xml:"AllcnSts"`

	// Provides the status of the repurchase agreement call request.
	RepoCallRequestStatus *RepoCallRequestStatus1Choice `xml:"RepoCallReqSts"`

	// Provides the status of a corporate action or the status of a payment.
	CorporateActionEventProcessingStatus *CorporateActionEventProcessingStatus1Choice `xml:"CorpActnEvtPrcgSts"`

	// Stage in the corporate action event life cycle.
	CorporateActionEventStage *CorporateActionEventStage1Choice `xml:"CorpActnEvtStag"`

	// Provides the matching status of an instruction as known by the account servicer based on an allegement. At this time no matching took place on the market (at the CSD/ICSD).
	InferredMatchingStatus *MatchingStatus1Choice `xml:"IfrrdMtchgSts"`

	// Provides the status of an instruction.
	InstructionProcessingStatus *InstructionProcessingStatus2Choice `xml:"InstrPrcgSts"`

	// Provides the matching status of the instruction.
	MatchingStatus *MatchingStatus1Choice `xml:"MtchgSts"`

	// Provides the status of the registration processing.
	RegistrationProcessingStatus *RegistrationProcessingStatus1Choice `xml:"RegnPrcgSts"`

	// Provides the status of the received collateral message (collateral claim, a collateral proposal or a proposal/request for collateral substitution) from a collateral management perspective.
	ResponseStatus *ResponseStatus1Choice `xml:"RspnSts"`

	// Provides the processing status of the replacement request.
	ReplacementProcessingStatus *ReplacementProcessingStatus1Choice `xml:"RplcmntPrcgSts"`

	// Provides the status of a cancellation request.
	CancellationProcessingStatus *CancellationProcessingStatus1Choice `xml:"CxlPrcgSts"`

	// Provides the status of settlement of a transaction.
	SettlementStatus *SettlementStatus1Choice `xml:"SttlmSts"`

	// Provides the status of a securities settlement condition modification request.
	SettlementConditionModificationStatus *SettlementConditionModificationStatus1Choice `xml:"SttlmCondModSts"`
}

func (s *Status1Choice) AddAffirmationStatus() *AffirmationStatus1Choice {
	s.AffirmationStatus = new(AffirmationStatus1Choice)
	return s.AffirmationStatus
}

func (s *Status1Choice) AddAllocationStatus() *AllocationSatus1Choice {
	s.AllocationStatus = new(AllocationSatus1Choice)
	return s.AllocationStatus
}

func (s *Status1Choice) AddRepoCallRequestStatus() *RepoCallRequestStatus1Choice {
	s.RepoCallRequestStatus = new(RepoCallRequestStatus1Choice)
	return s.RepoCallRequestStatus
}

func (s *Status1Choice) AddCorporateActionEventProcessingStatus() *CorporateActionEventProcessingStatus1Choice {
	s.CorporateActionEventProcessingStatus = new(CorporateActionEventProcessingStatus1Choice)
	return s.CorporateActionEventProcessingStatus
}

func (s *Status1Choice) AddCorporateActionEventStage() *CorporateActionEventStage1Choice {
	s.CorporateActionEventStage = new(CorporateActionEventStage1Choice)
	return s.CorporateActionEventStage
}

func (s *Status1Choice) AddInferredMatchingStatus() *MatchingStatus1Choice {
	s.InferredMatchingStatus = new(MatchingStatus1Choice)
	return s.InferredMatchingStatus
}

func (s *Status1Choice) AddInstructionProcessingStatus() *InstructionProcessingStatus2Choice {
	s.InstructionProcessingStatus = new(InstructionProcessingStatus2Choice)
	return s.InstructionProcessingStatus
}

func (s *Status1Choice) AddMatchingStatus() *MatchingStatus1Choice {
	s.MatchingStatus = new(MatchingStatus1Choice)
	return s.MatchingStatus
}

func (s *Status1Choice) AddRegistrationProcessingStatus() *RegistrationProcessingStatus1Choice {
	s.RegistrationProcessingStatus = new(RegistrationProcessingStatus1Choice)
	return s.RegistrationProcessingStatus
}

func (s *Status1Choice) AddResponseStatus() *ResponseStatus1Choice {
	s.ResponseStatus = new(ResponseStatus1Choice)
	return s.ResponseStatus
}

func (s *Status1Choice) AddReplacementProcessingStatus() *ReplacementProcessingStatus1Choice {
	s.ReplacementProcessingStatus = new(ReplacementProcessingStatus1Choice)
	return s.ReplacementProcessingStatus
}

func (s *Status1Choice) AddCancellationProcessingStatus() *CancellationProcessingStatus1Choice {
	s.CancellationProcessingStatus = new(CancellationProcessingStatus1Choice)
	return s.CancellationProcessingStatus
}

func (s *Status1Choice) AddSettlementStatus() *SettlementStatus1Choice {
	s.SettlementStatus = new(SettlementStatus1Choice)
	return s.SettlementStatus
}

func (s *Status1Choice) AddSettlementConditionModificationStatus() *SettlementConditionModificationStatus1Choice {
	s.SettlementConditionModificationStatus = new(SettlementConditionModificationStatus1Choice)
	return s.SettlementConditionModificationStatus
}
