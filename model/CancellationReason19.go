package model

// Specifies the reason why the instruction or request is cancelled.
type CancellationReason19 struct {

	// Specifies the reason why the instruction is cancelled.
	Code *CancellationReason24Choice `xml:"Cd"`

	// Provides the corporate action event identification of the event that triggered the cancellation.
	CorporateActionEventIdentification *RestrictedFINMax16Text `xml:"CorpActnEvtId,omitempty"`
}

func (c *CancellationReason19) AddCode() *CancellationReason24Choice {
	c.Code = new(CancellationReason24Choice)
	return c.Code
}

func (c *CancellationReason19) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*RestrictedFINMax16Text)(&value)
}
