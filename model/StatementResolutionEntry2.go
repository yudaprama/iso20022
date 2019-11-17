package model

// Provides further details on the statement entry for the resolution of the investigation.
type StatementResolutionEntry2 struct {

	// Provides information on the original message.
	OriginalGroupInformation *OriginalGroupInformation3 `xml:"OrgnlGrpInf,omitempty"`

	// Unique identification, as assigned by the account servicer, to unambiguously identify the original statement.
	OriginalStatementIdentification *Max35Text `xml:"OrgnlStmtId,omitempty"`

	// Unique reference, as assigned by the account servicing institution, to unambiguously identify the entry.
	AccountServicerReference *Max35Text `xml:"AcctSvcrRef,omitempty"`

	// Corrected debit or credit amount, compared to the original entry where the amount is incorrect.
	//
	// Usage: This amount may only be present if an incorrect statement entry has been reported.
	CorrectedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"CrrctdAmt,omitempty"`

	// Provides information on the charges included in the original entry amount.
	Charges []*Charges3 `xml:"Chrgs,omitempty"`

	// Underlying reason for the payment transaction.
	// Usage: Purpose is used by the end-customers, that is initiating party, (ultimate) debtor, (ultimate) creditor to provide information concerning the nature of the payment. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
	Purpose *Purpose2Choice `xml:"Purp,omitempty"`
}

func (s *StatementResolutionEntry2) AddOriginalGroupInformation() *OriginalGroupInformation3 {
	s.OriginalGroupInformation = new(OriginalGroupInformation3)
	return s.OriginalGroupInformation
}

func (s *StatementResolutionEntry2) SetOriginalStatementIdentification(value string) {
	s.OriginalStatementIdentification = (*Max35Text)(&value)
}

func (s *StatementResolutionEntry2) SetAccountServicerReference(value string) {
	s.AccountServicerReference = (*Max35Text)(&value)
}

func (s *StatementResolutionEntry2) SetCorrectedAmount(value, currency string) {
	s.CorrectedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (s *StatementResolutionEntry2) AddCharges() *Charges3 {
	newValue := new(Charges3)
	s.Charges = append(s.Charges, newValue)
	return newValue
}

func (s *StatementResolutionEntry2) AddPurpose() *Purpose2Choice {
	s.Purpose = new(Purpose2Choice)
	return s.Purpose
}
