package model

// Provides corporate action option details about total instructed balance.
type InstructedCorporateActionOption7 struct {

	// Number identifying the available corporate action options.
	OptionNumber *Exact3NumericText `xml:"OptnNb,omitempty"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption23Choice `xml:"OptnTp"`

	// Balance of instructed position.
	InstructedBalance *BalanceFormat7Choice `xml:"InstdBal"`

	// Indicates the default action related to a corporate action event.
	DefaultAction *DefaultProcessingOrStandingInstruction1Choice `xml:"DfltActn,omitempty"`

	// Any deadline chosen by the account servicer (service level agreement).
	DeadlineDateTime *ISODateTime `xml:"DdlnDtTm"`

	// Specifies the type of deadline for instructing.
	DeadlineType *DeadlineCode4Choice `xml:"DdlnTp"`
}

func (i *InstructedCorporateActionOption7) SetOptionNumber(value string) {
	i.OptionNumber = (*Exact3NumericText)(&value)
}

func (i *InstructedCorporateActionOption7) AddOptionType() *CorporateActionOption23Choice {
	i.OptionType = new(CorporateActionOption23Choice)
	return i.OptionType
}

func (i *InstructedCorporateActionOption7) AddInstructedBalance() *BalanceFormat7Choice {
	i.InstructedBalance = new(BalanceFormat7Choice)
	return i.InstructedBalance
}

func (i *InstructedCorporateActionOption7) AddDefaultAction() *DefaultProcessingOrStandingInstruction1Choice {
	i.DefaultAction = new(DefaultProcessingOrStandingInstruction1Choice)
	return i.DefaultAction
}

func (i *InstructedCorporateActionOption7) SetDeadlineDateTime(value string) {
	i.DeadlineDateTime = (*ISODateTime)(&value)
}

func (i *InstructedCorporateActionOption7) AddDeadlineType() *DeadlineCode4Choice {
	i.DeadlineType = new(DeadlineCode4Choice)
	return i.DeadlineType
}
