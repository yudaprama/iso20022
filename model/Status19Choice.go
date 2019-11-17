package model

// Choice of status.
type Status19Choice struct {

	// Status of affirmation of a trade.
	AffirmationStatus *AffirmationStatus8Choice `xml:"AffirmSts"`

	// Provides the status of allocation of collateral to cover the instruction.
	AllocationStatus *AllocationSatus3Choice `xml:"AllcnSts"`

	// Provides the status of the repurchase agreement call request.
	RepoCallRequestStatus *RepoCallRequestStatus8Choice `xml:"RepoCallReqSts"`

	// Provides the status of a corporate action or the status of a payment.
	CorporateActionEventProcessingStatus *CorporateActionEventProcessingStatus3Choice `xml:"CorpActnEvtPrcgSts"`

	// Stage in the corporate action event life cycle.
	CorporateActionEventStage *CorporateActionEventStage3Choice `xml:"CorpActnEvtStag"`

	// Provides the matching status of an instruction as known by the account servicer based on an allegement. At this time no matching took place on the market (at the CSD/ICSD).
	InferredMatchingStatus *MatchingStatus27Choice `xml:"IfrrdMtchgSts"`

	// Provides the status of an instruction.
	InstructionProcessingStatus *InstructionProcessingStatus23Choice `xml:"InstrPrcgSts"`

	// Provides the matching status of the instruction.
	MatchingStatus *MatchingStatus27Choice `xml:"MtchgSts"`

	// Provides the status of the registration processing.
	RegistrationProcessingStatus *RegistrationProcessingStatus3Choice `xml:"RegnPrcgSts"`

	// Provides the status of the received collateral message (collateral claim, a collateral proposal or a proposal/request for collateral substitution) from a collateral management perspective.
	ResponseStatus *ResponseStatus5Choice `xml:"RspnSts"`

	// Provides the processing status of the replacement request.
	ReplacementProcessingStatus *ReplacementProcessingStatus8Choice `xml:"RplcmntPrcgSts"`

	// Provides the status of a cancellation request.
	CancellationProcessingStatus *CancellationProcessingStatus7Choice `xml:"CxlPrcgSts"`

	// Provides the status of settlement of a transaction.
	SettlementStatus *SettlementStatus19Choice `xml:"SttlmSts"`

	// Provides the status of a securities settlement condition modification request.
	SettlementConditionModificationStatus *SettlementConditionModificationStatus3Choice `xml:"SttlmCondModSts"`
}

func (s *Status19Choice) AddAffirmationStatus() *AffirmationStatus8Choice {
	s.AffirmationStatus = new(AffirmationStatus8Choice)
	return s.AffirmationStatus
}

func (s *Status19Choice) AddAllocationStatus() *AllocationSatus3Choice {
	s.AllocationStatus = new(AllocationSatus3Choice)
	return s.AllocationStatus
}

func (s *Status19Choice) AddRepoCallRequestStatus() *RepoCallRequestStatus8Choice {
	s.RepoCallRequestStatus = new(RepoCallRequestStatus8Choice)
	return s.RepoCallRequestStatus
}

func (s *Status19Choice) AddCorporateActionEventProcessingStatus() *CorporateActionEventProcessingStatus3Choice {
	s.CorporateActionEventProcessingStatus = new(CorporateActionEventProcessingStatus3Choice)
	return s.CorporateActionEventProcessingStatus
}

func (s *Status19Choice) AddCorporateActionEventStage() *CorporateActionEventStage3Choice {
	s.CorporateActionEventStage = new(CorporateActionEventStage3Choice)
	return s.CorporateActionEventStage
}

func (s *Status19Choice) AddInferredMatchingStatus() *MatchingStatus27Choice {
	s.InferredMatchingStatus = new(MatchingStatus27Choice)
	return s.InferredMatchingStatus
}

func (s *Status19Choice) AddInstructionProcessingStatus() *InstructionProcessingStatus23Choice {
	s.InstructionProcessingStatus = new(InstructionProcessingStatus23Choice)
	return s.InstructionProcessingStatus
}

func (s *Status19Choice) AddMatchingStatus() *MatchingStatus27Choice {
	s.MatchingStatus = new(MatchingStatus27Choice)
	return s.MatchingStatus
}

func (s *Status19Choice) AddRegistrationProcessingStatus() *RegistrationProcessingStatus3Choice {
	s.RegistrationProcessingStatus = new(RegistrationProcessingStatus3Choice)
	return s.RegistrationProcessingStatus
}

func (s *Status19Choice) AddResponseStatus() *ResponseStatus5Choice {
	s.ResponseStatus = new(ResponseStatus5Choice)
	return s.ResponseStatus
}

func (s *Status19Choice) AddReplacementProcessingStatus() *ReplacementProcessingStatus8Choice {
	s.ReplacementProcessingStatus = new(ReplacementProcessingStatus8Choice)
	return s.ReplacementProcessingStatus
}

func (s *Status19Choice) AddCancellationProcessingStatus() *CancellationProcessingStatus7Choice {
	s.CancellationProcessingStatus = new(CancellationProcessingStatus7Choice)
	return s.CancellationProcessingStatus
}

func (s *Status19Choice) AddSettlementStatus() *SettlementStatus19Choice {
	s.SettlementStatus = new(SettlementStatus19Choice)
	return s.SettlementStatus
}

func (s *Status19Choice) AddSettlementConditionModificationStatus() *SettlementConditionModificationStatus3Choice {
	s.SettlementConditionModificationStatus = new(SettlementConditionModificationStatus3Choice)
	return s.SettlementConditionModificationStatus
}
