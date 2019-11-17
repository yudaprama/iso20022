package model

// Provides corporate action option details about total instructed balance.
type InstructedCorporateActionOption6 struct {

	// Number identifying the available corporate action options.
	OptionNumber *Exact3NumericText `xml:"OptnNb,omitempty"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption18Choice `xml:"OptnTp"`

	// Balance of instructed position.
	InstructedBalance *BalanceFormat5Choice `xml:"InstdBal"`

	// Indicates the default action related to a corporate action event.
	DefaultAction *DefaultProcessingOrStandingInstruction1Choice `xml:"DfltActn,omitempty"`

	// Any deadline chosen by the account servicer (service level agreement).
	DeadlineDateTime *ISODateTime `xml:"DdlnDtTm"`

	// Specifies the type of deadline for instructing.
	DeadlineType *DeadlineCode3Choice `xml:"DdlnTp"`
}

func (i *InstructedCorporateActionOption6) SetOptionNumber(value string) {
	i.OptionNumber = (*Exact3NumericText)(&value)
}

func (i *InstructedCorporateActionOption6) AddOptionType() *CorporateActionOption18Choice {
	i.OptionType = new(CorporateActionOption18Choice)
	return i.OptionType
}

func (i *InstructedCorporateActionOption6) AddInstructedBalance() *BalanceFormat5Choice {
	i.InstructedBalance = new(BalanceFormat5Choice)
	return i.InstructedBalance
}

func (i *InstructedCorporateActionOption6) AddDefaultAction() *DefaultProcessingOrStandingInstruction1Choice {
	i.DefaultAction = new(DefaultProcessingOrStandingInstruction1Choice)
	return i.DefaultAction
}

func (i *InstructedCorporateActionOption6) SetDeadlineDateTime(value string) {
	i.DeadlineDateTime = (*ISODateTime)(&value)
}

func (i *InstructedCorporateActionOption6) AddDeadlineType() *DeadlineCode3Choice {
	i.DeadlineType = new(DeadlineCode3Choice)
	return i.DeadlineType
}
