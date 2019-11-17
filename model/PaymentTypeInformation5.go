package model

// Set of elements that further details the information related to the type of payment.
type PaymentTypeInformation5 struct {

	// Indicator of the urgency or order of importance that the instructing party would like the instructed party to apply to the processing of the instruction.
	InstructionPriority *Priority2Code `xml:"InstrPrty,omitempty"`

	// Agreement under which or rules under which the transaction should be processed.
	ServiceLevel *RestrictedProprietaryChoice `xml:"SvcLvl,omitempty"`

	// Specifies the clearing channel to be used for the instruction.
	ClearingChannel *ClearingChannel2Code `xml:"ClrChanl,omitempty"`

	// User community specific instrument required for use within that user community.
	//
	// Usage : When available, codes provided by local authorities should be used.
	LocalInstrument *RestrictedProprietaryChoice `xml:"LclInstrm,omitempty"`
}

func (p *PaymentTypeInformation5) SetInstructionPriority(value string) {
	p.InstructionPriority = (*Priority2Code)(&value)
}

func (p *PaymentTypeInformation5) AddServiceLevel() *RestrictedProprietaryChoice {
	p.ServiceLevel = new(RestrictedProprietaryChoice)
	return p.ServiceLevel
}

func (p *PaymentTypeInformation5) SetClearingChannel(value string) {
	p.ClearingChannel = (*ClearingChannel2Code)(&value)
}

func (p *PaymentTypeInformation5) AddLocalInstrument() *RestrictedProprietaryChoice {
	p.LocalInstrument = new(RestrictedProprietaryChoice)
	return p.LocalInstrument
}
