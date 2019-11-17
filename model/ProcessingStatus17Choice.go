package model

// Choice of status for the processing.
type ProcessingStatus17Choice struct {

	// Trade is AcknowledgedAccepted.
	AcknowledgedAccepted *ProprietaryReason1 `xml:"AckdAccptd"`

	// Trade is AlreadyMatchedAndAffirmed.
	AlreadyMatchedAndAffirmed *ProprietaryReason1 `xml:"AlrdyMtchdAndAffrmd"`

	// Trade is DefaultAction.
	DefaultAction *ProprietaryReason1 `xml:"DfltActn"`

	// Trade is Done.
	Done *ProprietaryReason1 `xml:"Done"`

	// Trade is in forced rejection.
	ForcedRejection *ProprietaryReason1 `xml:"ForcdRjctn"`

	// The trade is fully executed and the confirmation is sent.
	FullyExecutedConfirmationSent *ProprietaryReason1 `xml:"FullyExctdConfSnt"`

	// Trade is future.
	Future *ProprietaryReason1 `xml:"Futr"`

	// Trade is generated.
	Generated *ProprietaryReason1 `xml:"Gnrtd"`

	// Trade is InRepair.
	InRepair *InstructionProcessingReason2Choice `xml:"InRpr"`

	// Trade is in no instruction.
	NoInstruction *ProprietaryReason1 `xml:"NoInstr"`

	// Trade is in OpenOrder
	OpenOrder *ProprietaryReason1 `xml:"OpnOrdr"`

	// Processing of the trade is pending .
	PendingProcessing *PendingProcessing1Choice `xml:"PdgPrcg"`

	// Trade is ReceivedAtIntermediary.
	ReceivedAtIntermediary *ProprietaryReason1 `xml:"RcvdAtIntrmy"`

	// Trade has been rejected.
	Rejected *InstructionProcessingReason1Choice `xml:"Rjctd"`

	// Settlement Instruction for the trade is sent.
	SettlementInstructionSent *ProprietaryReason1 `xml:"SttlmInstrSnt"`

	// Trade is standing instruction.
	StandingInstruction *ProprietaryReason1 `xml:"StgInstr"`

	// Trading is suspended by the stock exchange.
	TradingSuspendedByStockExchange *ProprietaryReason1 `xml:"TradgSspdByStockXchg"`

	// Trade is treated.
	Treated *ProprietaryReason1 `xml:"Trtd"`

	// Provides a proprietary status and a proprietary reason of the processing status of the trade.
	ProprietaryStatus *ProprietaryStatusAndReason1 `xml:"PrtrySts,omitempty"`
}

func (p *ProcessingStatus17Choice) AddAcknowledgedAccepted() *ProprietaryReason1 {
	p.AcknowledgedAccepted = new(ProprietaryReason1)
	return p.AcknowledgedAccepted
}

func (p *ProcessingStatus17Choice) AddAlreadyMatchedAndAffirmed() *ProprietaryReason1 {
	p.AlreadyMatchedAndAffirmed = new(ProprietaryReason1)
	return p.AlreadyMatchedAndAffirmed
}

func (p *ProcessingStatus17Choice) AddDefaultAction() *ProprietaryReason1 {
	p.DefaultAction = new(ProprietaryReason1)
	return p.DefaultAction
}

func (p *ProcessingStatus17Choice) AddDone() *ProprietaryReason1 {
	p.Done = new(ProprietaryReason1)
	return p.Done
}

func (p *ProcessingStatus17Choice) AddForcedRejection() *ProprietaryReason1 {
	p.ForcedRejection = new(ProprietaryReason1)
	return p.ForcedRejection
}

func (p *ProcessingStatus17Choice) AddFullyExecutedConfirmationSent() *ProprietaryReason1 {
	p.FullyExecutedConfirmationSent = new(ProprietaryReason1)
	return p.FullyExecutedConfirmationSent
}

func (p *ProcessingStatus17Choice) AddFuture() *ProprietaryReason1 {
	p.Future = new(ProprietaryReason1)
	return p.Future
}

func (p *ProcessingStatus17Choice) AddGenerated() *ProprietaryReason1 {
	p.Generated = new(ProprietaryReason1)
	return p.Generated
}

func (p *ProcessingStatus17Choice) AddInRepair() *InstructionProcessingReason2Choice {
	p.InRepair = new(InstructionProcessingReason2Choice)
	return p.InRepair
}

func (p *ProcessingStatus17Choice) AddNoInstruction() *ProprietaryReason1 {
	p.NoInstruction = new(ProprietaryReason1)
	return p.NoInstruction
}

func (p *ProcessingStatus17Choice) AddOpenOrder() *ProprietaryReason1 {
	p.OpenOrder = new(ProprietaryReason1)
	return p.OpenOrder
}

func (p *ProcessingStatus17Choice) AddPendingProcessing() *PendingProcessing1Choice {
	p.PendingProcessing = new(PendingProcessing1Choice)
	return p.PendingProcessing
}

func (p *ProcessingStatus17Choice) AddReceivedAtIntermediary() *ProprietaryReason1 {
	p.ReceivedAtIntermediary = new(ProprietaryReason1)
	return p.ReceivedAtIntermediary
}

func (p *ProcessingStatus17Choice) AddRejected() *InstructionProcessingReason1Choice {
	p.Rejected = new(InstructionProcessingReason1Choice)
	return p.Rejected
}

func (p *ProcessingStatus17Choice) AddSettlementInstructionSent() *ProprietaryReason1 {
	p.SettlementInstructionSent = new(ProprietaryReason1)
	return p.SettlementInstructionSent
}

func (p *ProcessingStatus17Choice) AddStandingInstruction() *ProprietaryReason1 {
	p.StandingInstruction = new(ProprietaryReason1)
	return p.StandingInstruction
}

func (p *ProcessingStatus17Choice) AddTradingSuspendedByStockExchange() *ProprietaryReason1 {
	p.TradingSuspendedByStockExchange = new(ProprietaryReason1)
	return p.TradingSuspendedByStockExchange
}

func (p *ProcessingStatus17Choice) AddTreated() *ProprietaryReason1 {
	p.Treated = new(ProprietaryReason1)
	return p.Treated
}

func (p *ProcessingStatus17Choice) AddProprietaryStatus() *ProprietaryStatusAndReason1 {
	p.ProprietaryStatus = new(ProprietaryStatusAndReason1)
	return p.ProprietaryStatus
}
