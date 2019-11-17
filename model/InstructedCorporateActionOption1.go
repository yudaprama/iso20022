package model

// Provides corporate action option details about total instructed balance.
type InstructedCorporateActionOption1 struct {

	// Number identifying the available corporate action options.
	OptionNumber *Exact3NumericText `xml:"OptnNb,omitempty"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption2Choice `xml:"OptnTp"`

	// Balance of instructed position.
	InstructedBalance *BalanceFormat1Choice `xml:"InstdBal"`

	// Indicates the default action related to a corporate action event.
	DefaultAction *DefaultProcessingOrStandingInstruction1Choice `xml:"DfltActn,omitempty"`

	// Any deadline chosen by the account servicer (service level agreement).
	DeadlineDateTime *ISODateTime `xml:"DdlnDtTm"`

	// Specifies the type of deadline for instructing.
	DeadlineType *DeadlineCode1Choice `xml:"DdlnTp"`
}

func (i *InstructedCorporateActionOption1) SetOptionNumber(value string) {
	i.OptionNumber = (*Exact3NumericText)(&value)
}

func (i *InstructedCorporateActionOption1) AddOptionType() *CorporateActionOption2Choice {
	i.OptionType = new(CorporateActionOption2Choice)
	return i.OptionType
}

func (i *InstructedCorporateActionOption1) AddInstructedBalance() *BalanceFormat1Choice {
	i.InstructedBalance = new(BalanceFormat1Choice)
	return i.InstructedBalance
}

func (i *InstructedCorporateActionOption1) AddDefaultAction() *DefaultProcessingOrStandingInstruction1Choice {
	i.DefaultAction = new(DefaultProcessingOrStandingInstruction1Choice)
	return i.DefaultAction
}

func (i *InstructedCorporateActionOption1) SetDeadlineDateTime(value string) {
	i.DeadlineDateTime = (*ISODateTime)(&value)
}

func (i *InstructedCorporateActionOption1) AddDeadlineType() *DeadlineCode1Choice {
	i.DeadlineType = new(DeadlineCode1Choice)
	return i.DeadlineType
}
