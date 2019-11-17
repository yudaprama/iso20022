package model

// Set of elements that further details the information related to the type of payment.
type PaymentTypeInformation2 struct {

	// Indicator of the urgency or order of importance that the instructing party would like the instructed party to apply to the processing of the instruction.
	InstructionPriority *Priority2Code `xml:"InstrPrty,omitempty"`

	// Agreement under which or rules under which the transaction should be processed.
	ServiceLevel *ServiceLevel3Choice `xml:"SvcLvl,omitempty"`

	// Specifies the clearing channel to be used for the instruction.
	ClearingChannel *ClearingChannel2Code `xml:"ClrChanl,omitempty"`

	// User community specific instrument.
	//
	// Usage : When available, codes provided by local authorities should be used.
	LocalInstrument *LocalInstrument1Choice `xml:"LclInstrm,omitempty"`

	// Identifies the direct debit sequence, eg, first, recurrent, final or one-off.
	SequenceType *SequenceType1Code `xml:"SeqTp,omitempty"`

	// Specifies the high level purpose of the instruction based on a set of pre-defined categories.
	CategoryPurpose *PaymentCategoryPurpose1Code `xml:"CtgyPurp,omitempty"`
}

func (p *PaymentTypeInformation2) SetInstructionPriority(value string) {
	p.InstructionPriority = (*Priority2Code)(&value)
}

func (p *PaymentTypeInformation2) AddServiceLevel() *ServiceLevel3Choice {
	p.ServiceLevel = new(ServiceLevel3Choice)
	return p.ServiceLevel
}

func (p *PaymentTypeInformation2) SetClearingChannel(value string) {
	p.ClearingChannel = (*ClearingChannel2Code)(&value)
}

func (p *PaymentTypeInformation2) AddLocalInstrument() *LocalInstrument1Choice {
	p.LocalInstrument = new(LocalInstrument1Choice)
	return p.LocalInstrument
}

func (p *PaymentTypeInformation2) SetSequenceType(value string) {
	p.SequenceType = (*SequenceType1Code)(&value)
}

func (p *PaymentTypeInformation2) SetCategoryPurpose(value string) {
	p.CategoryPurpose = (*PaymentCategoryPurpose1Code)(&value)
}
