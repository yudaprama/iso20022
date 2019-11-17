package model

// Choice between default processing or standing instruction.
type DefaultProcessingOrStandingInstruction1Choice struct {

	// Indicates whether the option, for example, currency option, will be selected by default if no instruction is provided by the account owner.
	DefaultOptionIndicator *YesNoIndicator `xml:"DfltOptnInd"`

	// Indicates whether an account owner has placed a standing order to select this corporate action option.
	StandingInstructionIndicator *YesNoIndicator `xml:"StgInstrInd"`
}

func (d *DefaultProcessingOrStandingInstruction1Choice) SetDefaultOptionIndicator(value string) {
	d.DefaultOptionIndicator = (*YesNoIndicator)(&value)
}

func (d *DefaultProcessingOrStandingInstruction1Choice) SetStandingInstructionIndicator(value string) {
	d.StandingInstructionIndicator = (*YesNoIndicator)(&value)
}
