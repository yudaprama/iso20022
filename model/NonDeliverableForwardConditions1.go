package model

// Specifies the opening and valuation conditions for the non deliverable forward
type NonDeliverableForwardConditions1 struct {

	// Specifies whether the instruction is an NDF opening or fixing.
	OpeningIndicator *YesNoIndicator `xml:"OpngInd"`

	// Specifies either the conditions for an NDF oepning or an NDF fixing confirmation.
	OpeningFixingConditions *NDFOpeningFixing1Choice `xml:"OpngFxgConds"`
}

func (n *NonDeliverableForwardConditions1) SetOpeningIndicator(value string) {
	n.OpeningIndicator = (*YesNoIndicator)(&value)
}

func (n *NonDeliverableForwardConditions1) AddOpeningFixingConditions() *NDFOpeningFixing1Choice {
	n.OpeningFixingConditions = new(NDFOpeningFixing1Choice)
	return n.OpeningFixingConditions
}
