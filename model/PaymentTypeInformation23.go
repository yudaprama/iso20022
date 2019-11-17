package model

// Set of elements used to provide further details of the type of payment.
type PaymentTypeInformation23 struct {

	// Indicator of the urgency or order of importance that the instructing party would like the instructed party to apply to the processing of the instruction.
	InstructionPriority *Priority2Code `xml:"InstrPrty,omitempty"`

	// Specifies the clearing channel to be used to process the payment instruction.
	ClearingChannel *ClearingChannel2Code `xml:"ClrChanl,omitempty"`

	// Agreement under which or rules under which the transaction should be processed.
	ServiceLevel *ServiceLevel8Choice `xml:"SvcLvl,omitempty"`

	// User community specific instrument.
	// Usage: This element is used to specify a local instrument, local clearing option and/or further qualify the service or service level.
	LocalInstrument *LocalInstrument2Choice `xml:"LclInstrm,omitempty"`
}

func (p *PaymentTypeInformation23) SetInstructionPriority(value string) {
	p.InstructionPriority = (*Priority2Code)(&value)
}

func (p *PaymentTypeInformation23) SetClearingChannel(value string) {
	p.ClearingChannel = (*ClearingChannel2Code)(&value)
}

func (p *PaymentTypeInformation23) AddServiceLevel() *ServiceLevel8Choice {
	p.ServiceLevel = new(ServiceLevel8Choice)
	return p.ServiceLevel
}

func (p *PaymentTypeInformation23) AddLocalInstrument() *LocalInstrument2Choice {
	p.LocalInstrument = new(LocalInstrument2Choice)
	return p.LocalInstrument
}
