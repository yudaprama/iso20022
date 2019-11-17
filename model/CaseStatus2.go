package model

// Defines the status of an investigation case.
type CaseStatus2 struct {

	// Date and time of the status.
	DateTime *ISODateTime `xml:"DtTm"`

	// Status of the case.
	CaseStatus *CaseStatus2Code `xml:"CaseSts"`

	// Free text justification of the status.
	Reason *Max140Text `xml:"Rsn,omitempty"`
}

func (c *CaseStatus2) SetDateTime(value string) {
	c.DateTime = (*ISODateTime)(&value)
}

func (c *CaseStatus2) SetCaseStatus(value string) {
	c.CaseStatus = (*CaseStatus2Code)(&value)
}

func (c *CaseStatus2) SetReason(value string) {
	c.Reason = (*Max140Text)(&value)
}
