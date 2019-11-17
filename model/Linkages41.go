package model

// Information related to a linked transaction.
type Linkages41 struct {

	// When the transaction is to be executed relative to a linked transaction - for information only.
	ProcessingPosition *ProcessingPosition9Choice `xml:"PrcgPos,omitempty"`

	// Unambiguous identification of a securities settlement transaction as known by the account owner (or instructing party acting on its behalf).
	SecuritiesSettlementTransactionIdentification *Max35Text `xml:"SctiesSttlmTxId"`
}

func (l *Linkages41) AddProcessingPosition() *ProcessingPosition9Choice {
	l.ProcessingPosition = new(ProcessingPosition9Choice)
	return l.ProcessingPosition
}

func (l *Linkages41) SetSecuritiesSettlementTransactionIdentification(value string) {
	l.SecuritiesSettlementTransactionIdentification = (*Max35Text)(&value)
}
