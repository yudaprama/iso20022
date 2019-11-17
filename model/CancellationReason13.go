package model

// Specifies the reason why the instruction or request is cancelled.
type CancellationReason13 struct {

	// Specifies the reason why the instruction is cancelled.
	Code *CancellationReason19Choice `xml:"Cd"`

	// Provides the corporate action event identification of the event that triggered the cancellation.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`
}

func (c *CancellationReason13) AddCode() *CancellationReason19Choice {
	c.Code = new(CancellationReason19Choice)
	return c.Code
}

func (c *CancellationReason13) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}
