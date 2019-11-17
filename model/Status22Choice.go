package model

// Choice of status.
type Status22Choice struct {

	// Status of affirmation of a trade.
	AffirmationStatus *AffirmationStatus9Choice `xml:"AffirmSts"`

	// Provides the status of allocation of collateral to cover the instruction.
	AllocationStatus *AllocationSatus4Choice `xml:"AllcnSts"`

	// Provides the status of the repurchase agreement call request.
	RepoCallRequestStatus *RepoCallRequestStatus10Choice `xml:"RepoCallReqSts"`

	// Provides the status of a corporate action or the status of a payment.
	CorporateActionEventProcessingStatus *CorporateActionEventProcessingStatus4Choice `xml:"CorpActnEvtPrcgSts"`

	// Stage in the corporate action event life cycle.
	CorporateActionEventStage *CorporateActionEventStage4Choice `xml:"CorpActnEvtStag"`

	// Provides the matching status of an instruction as known by the account servicer based on an allegement. At this time no matching took place on the market (at the CSD/ICSD).
	InferredMatchingStatus *MatchingStatus28Choice `xml:"IfrrdMtchgSts"`

	// Provides the status of an instruction.
	InstructionProcessingStatus *InstructionProcessingStatus26Choice `xml:"InstrPrcgSts"`

	// Provides the matching status of the instruction.
	MatchingStatus *MatchingStatus28Choice `xml:"MtchgSts"`

	// Provides the status of the registration processing.
	RegistrationProcessingStatus *RegistrationProcessingStatus4Choice `xml:"RegnPrcgSts"`

	// Provides the status of the received collateral message (collateral claim, a collateral proposal or a proposal/request for collateral substitution) from a collateral management perspective.
	ResponseStatus *ResponseStatus7Choice `xml:"RspnSts"`

	// Provides the processing status of the replacement request.
	ReplacementProcessingStatus *ReplacementProcessingStatus9Choice `xml:"RplcmntPrcgSts"`

	// Provides the status of a cancellation request.
	CancellationProcessingStatus *CancellationProcessingStatus8Choice `xml:"CxlPrcgSts"`

	// Provides the status of settlement of a transaction.
	SettlementStatus *SettlementStatus25Choice `xml:"SttlmSts"`

	// Provides the status of a securities settlement condition modification request.
	SettlementConditionModificationStatus *SettlementConditionModificationStatus4Choice `xml:"SttlmCondModSts"`
}

func (s *Status22Choice) AddAffirmationStatus() *AffirmationStatus9Choice {
	s.AffirmationStatus = new(AffirmationStatus9Choice)
	return s.AffirmationStatus
}

func (s *Status22Choice) AddAllocationStatus() *AllocationSatus4Choice {
	s.AllocationStatus = new(AllocationSatus4Choice)
	return s.AllocationStatus
}

func (s *Status22Choice) AddRepoCallRequestStatus() *RepoCallRequestStatus10Choice {
	s.RepoCallRequestStatus = new(RepoCallRequestStatus10Choice)
	return s.RepoCallRequestStatus
}

func (s *Status22Choice) AddCorporateActionEventProcessingStatus() *CorporateActionEventProcessingStatus4Choice {
	s.CorporateActionEventProcessingStatus = new(CorporateActionEventProcessingStatus4Choice)
	return s.CorporateActionEventProcessingStatus
}

func (s *Status22Choice) AddCorporateActionEventStage() *CorporateActionEventStage4Choice {
	s.CorporateActionEventStage = new(CorporateActionEventStage4Choice)
	return s.CorporateActionEventStage
}

func (s *Status22Choice) AddInferredMatchingStatus() *MatchingStatus28Choice {
	s.InferredMatchingStatus = new(MatchingStatus28Choice)
	return s.InferredMatchingStatus
}

func (s *Status22Choice) AddInstructionProcessingStatus() *InstructionProcessingStatus26Choice {
	s.InstructionProcessingStatus = new(InstructionProcessingStatus26Choice)
	return s.InstructionProcessingStatus
}

func (s *Status22Choice) AddMatchingStatus() *MatchingStatus28Choice {
	s.MatchingStatus = new(MatchingStatus28Choice)
	return s.MatchingStatus
}

func (s *Status22Choice) AddRegistrationProcessingStatus() *RegistrationProcessingStatus4Choice {
	s.RegistrationProcessingStatus = new(RegistrationProcessingStatus4Choice)
	return s.RegistrationProcessingStatus
}

func (s *Status22Choice) AddResponseStatus() *ResponseStatus7Choice {
	s.ResponseStatus = new(ResponseStatus7Choice)
	return s.ResponseStatus
}

func (s *Status22Choice) AddReplacementProcessingStatus() *ReplacementProcessingStatus9Choice {
	s.ReplacementProcessingStatus = new(ReplacementProcessingStatus9Choice)
	return s.ReplacementProcessingStatus
}

func (s *Status22Choice) AddCancellationProcessingStatus() *CancellationProcessingStatus8Choice {
	s.CancellationProcessingStatus = new(CancellationProcessingStatus8Choice)
	return s.CancellationProcessingStatus
}

func (s *Status22Choice) AddSettlementStatus() *SettlementStatus25Choice {
	s.SettlementStatus = new(SettlementStatus25Choice)
	return s.SettlementStatus
}

func (s *Status22Choice) AddSettlementConditionModificationStatus() *SettlementConditionModificationStatus4Choice {
	s.SettlementConditionModificationStatus = new(SettlementConditionModificationStatus4Choice)
	return s.SettlementConditionModificationStatus
}
