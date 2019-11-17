package model

// Set of elements used to provide information on the underlying statement entry.
type UnderlyingStatementEntry1 struct {

	// Set of elements used to provide information on the original message.
	OriginalGroupInformation *OriginalGroupInformation3 `xml:"OrgnlGrpInf,omitempty"`

	// Unique identification, as assigned by the account servicer, to unambiguously identify the original statement.
	OriginalStatementIdentification *Max35Text `xml:"OrgnlStmtId,omitempty"`

	// Original unique identification, as assigned by the account servicer, to unambiguously identify the original entry.
	OriginalEntryIdentification *Max35Text `xml:"OrgnlNtryId,omitempty"`
}

func (u *UnderlyingStatementEntry1) AddOriginalGroupInformation() *OriginalGroupInformation3 {
	u.OriginalGroupInformation = new(OriginalGroupInformation3)
	return u.OriginalGroupInformation
}

func (u *UnderlyingStatementEntry1) SetOriginalStatementIdentification(value string) {
	u.OriginalStatementIdentification = (*Max35Text)(&value)
}

func (u *UnderlyingStatementEntry1) SetOriginalEntryIdentification(value string) {
	u.OriginalEntryIdentification = (*Max35Text)(&value)
}
