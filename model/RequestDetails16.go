package model

// Details of the settlement condition modification request
type RequestDetails16 struct {

	// References of the transaction for which the securities settlement condition modification is requested.
	Reference *References21 `xml:"Ref"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing11Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether the instruction due to expire is confirmed for settlement
	RetainIndicator *YesNoIndicator `xml:"RtnInd,omitempty"`

	// Specifies the type of linkage requested.
	Linkage *LinkageType4Choice `xml:"Lkg,omitempty"`

	// Specifies whether the transaction is to be executed with a high priority.
	Priority *PriorityNumeric5Choice `xml:"Prty,omitempty"`

	// Specifies another type of processing change request
	OtherProcessing []*GenericIdentification47 `xml:"OthrPrcg,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS5Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *HoldIndicator7 `xml:"HldInd,omitempty"`

	// Specifies the matching processing change requested.
	MatchingDenial *MatchingDenied4Choice `xml:"MtchgDnl,omitempty"`

	// Specifies that the transaction is requested to be unilaterally split.
	UnilateralSplit *UnilateralSplit4Choice `xml:"UnltrlSplt,omitempty"`

	// Information regarding the linkage requested.
	Linkages []*Linkages44 `xml:"Lnkgs,omitempty"`
}

func (r *RequestDetails16) AddReference() *References21 {
	r.Reference = new(References21)
	return r.Reference
}

func (r *RequestDetails16) AddAutomaticBorrowing() *AutomaticBorrowing11Choice {
	r.AutomaticBorrowing = new(AutomaticBorrowing11Choice)
	return r.AutomaticBorrowing
}

func (r *RequestDetails16) SetRetainIndicator(value string) {
	r.RetainIndicator = (*YesNoIndicator)(&value)
}

func (r *RequestDetails16) AddLinkage() *LinkageType4Choice {
	r.Linkage = new(LinkageType4Choice)
	return r.Linkage
}

func (r *RequestDetails16) AddPriority() *PriorityNumeric5Choice {
	r.Priority = new(PriorityNumeric5Choice)
	return r.Priority
}

func (r *RequestDetails16) AddOtherProcessing() *GenericIdentification47 {
	newValue := new(GenericIdentification47)
	r.OtherProcessing = append(r.OtherProcessing, newValue)
	return newValue
}

func (r *RequestDetails16) SetPartialSettlementIndicator(value string) {
	r.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (r *RequestDetails16) AddSecuritiesRTGS() *SecuritiesRTGS5Choice {
	r.SecuritiesRTGS = new(SecuritiesRTGS5Choice)
	return r.SecuritiesRTGS
}

func (r *RequestDetails16) AddHoldIndicator() *HoldIndicator7 {
	r.HoldIndicator = new(HoldIndicator7)
	return r.HoldIndicator
}

func (r *RequestDetails16) AddMatchingDenial() *MatchingDenied4Choice {
	r.MatchingDenial = new(MatchingDenied4Choice)
	return r.MatchingDenial
}

func (r *RequestDetails16) AddUnilateralSplit() *UnilateralSplit4Choice {
	r.UnilateralSplit = new(UnilateralSplit4Choice)
	return r.UnilateralSplit
}

func (r *RequestDetails16) AddLinkages() *Linkages44 {
	newValue := new(Linkages44)
	r.Linkages = append(r.Linkages, newValue)
	return newValue
}
