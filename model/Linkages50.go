package model

// Information related to a linked transaction.
type Linkages50 struct {

	// When the transaction is to be executed relative to a linked transaction - for information only.
	ProcessingPosition *ProcessingPosition23Choice `xml:"PrcgPos,omitempty"`

	// Unambiguous identification of a securities settlement transaction as known by the account owner (or instructing party acting on its behalf).
	SecuritiesSettlementTransactionIdentification *RestrictedFINMax16Text `xml:"SctiesSttlmTxId"`
}

func (l *Linkages50) AddProcessingPosition() *ProcessingPosition23Choice {
	l.ProcessingPosition = new(ProcessingPosition23Choice)
	return l.ProcessingPosition
}

func (l *Linkages50) SetSecuritiesSettlementTransactionIdentification(value string) {
	l.SecuritiesSettlementTransactionIdentification = (*RestrictedFINMax16Text)(&value)
}
