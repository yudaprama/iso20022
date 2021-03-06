package model

// Provides corporate action option details about total instructed balance.
type InstructedCorporateActionOption4 struct {

	// Number identifying the available corporate action options.
	OptionNumber *Exact3NumericText `xml:"OptnNb,omitempty"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption10Choice `xml:"OptnTp"`

	// Balance of instructed position.
	InstructedBalance *BalanceFormat1Choice `xml:"InstdBal"`

	// Indicates the default action related to a corporate action event.
	DefaultAction *DefaultProcessingOrStandingInstruction1Choice `xml:"DfltActn,omitempty"`

	// Any deadline chosen by the account servicer (service level agreement).
	DeadlineDateTime *ISODateTime `xml:"DdlnDtTm"`

	// Specifies the type of deadline for instructing.
	DeadlineType *DeadlineCode1Choice `xml:"DdlnTp"`
}

func (i *InstructedCorporateActionOption4) SetOptionNumber(value string) {
	i.OptionNumber = (*Exact3NumericText)(&value)
}

func (i *InstructedCorporateActionOption4) AddOptionType() *CorporateActionOption10Choice {
	i.OptionType = new(CorporateActionOption10Choice)
	return i.OptionType
}

func (i *InstructedCorporateActionOption4) AddInstructedBalance() *BalanceFormat1Choice {
	i.InstructedBalance = new(BalanceFormat1Choice)
	return i.InstructedBalance
}

func (i *InstructedCorporateActionOption4) AddDefaultAction() *DefaultProcessingOrStandingInstruction1Choice {
	i.DefaultAction = new(DefaultProcessingOrStandingInstruction1Choice)
	return i.DefaultAction
}

func (i *InstructedCorporateActionOption4) SetDeadlineDateTime(value string) {
	i.DeadlineDateTime = (*ISODateTime)(&value)
}

func (i *InstructedCorporateActionOption4) AddDeadlineType() *DeadlineCode1Choice {
	i.DeadlineType = new(DeadlineCode1Choice)
	return i.DeadlineType
}
