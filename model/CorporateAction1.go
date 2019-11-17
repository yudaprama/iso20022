package model

// An event determined by a corporation's board of directors, that changes the existing corporate capital structure or financial condition.
type CorporateAction1 struct {

	// Specifies the code of corporate action event, in free-text format.
	Code *Max35Text `xml:"Cd,omitempty"`

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	Number *Max35Text `xml:"Nb,omitempty"`

	// Proprietary corporate action event information.
	Proprietary *Max35Text `xml:"Prtry,omitempty"`
}

func (c *CorporateAction1) SetCode(value string) {
	c.Code = (*Max35Text)(&value)
}

func (c *CorporateAction1) SetNumber(value string) {
	c.Number = (*Max35Text)(&value)
}

func (c *CorporateAction1) SetProprietary(value string) {
	c.Proprietary = (*Max35Text)(&value)
}
