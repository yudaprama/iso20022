package model

// Set of elements used to provide information on the statement entry for the resolution of the investigation.
type StatementResolutionEntry1 struct {

	// Set of elements used to provide information on the original messsage.
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
	Charges []*ChargesInformation6 `xml:"Chrgs,omitempty"`

	// Underlying reason for the payment transaction.
	// Usage: Purpose is used by the end-customers, that is initiating party, (ultimate) debtor, (ultimate) creditor to provide information concerning the nature of the payment. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
	Purpose *Purpose2Choice `xml:"Purp,omitempty"`
}

func (s *StatementResolutionEntry1) AddOriginalGroupInformation() *OriginalGroupInformation3 {
	s.OriginalGroupInformation = new(OriginalGroupInformation3)
	return s.OriginalGroupInformation
}

func (s *StatementResolutionEntry1) SetOriginalStatementIdentification(value string) {
	s.OriginalStatementIdentification = (*Max35Text)(&value)
}

func (s *StatementResolutionEntry1) SetAccountServicerReference(value string) {
	s.AccountServicerReference = (*Max35Text)(&value)
}

func (s *StatementResolutionEntry1) SetCorrectedAmount(value, currency string) {
	s.CorrectedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (s *StatementResolutionEntry1) AddCharges() *ChargesInformation6 {
	newValue := new(ChargesInformation6)
	s.Charges = append(s.Charges, newValue)
	return newValue
}

func (s *StatementResolutionEntry1) AddPurpose() *Purpose2Choice {
	s.Purpose = new(Purpose2Choice)
	return s.Purpose
}
