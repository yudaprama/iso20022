package model

// Specifies the opening and valuation conditions for the non deliverable forward
type NonDeliverableForwardConditions2 struct {

	// Provides the opening information associated with an NDF trade.
	OpeningConditions *OpeningConditions1 `xml:"OpngConds"`

	// Provides the additional information for an NDF as supplied on a fixing instruction.
	FixingConditions *FixingConditions1 `xml:"FxgConds,omitempty"`
}

func (n *NonDeliverableForwardConditions2) AddOpeningConditions() *OpeningConditions1 {
	n.OpeningConditions = new(OpeningConditions1)
	return n.OpeningConditions
}

func (n *NonDeliverableForwardConditions2) AddFixingConditions() *FixingConditions1 {
	n.FixingConditions = new(FixingConditions1)
	return n.FixingConditions
}
