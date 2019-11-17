package model

// Specifies additional parameters to the message or transaction.
type AdditionalParameters4 struct {

	// Specifies whether there exists a pre-confirmation.
	PreConfirmation *PreConfirmation1Code `xml:"PreConf,omitempty"`

	// Specifies partial settlement information.
	PartialSettlement *PartialSettlement1Code `xml:"PrtlSttlm,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the triparty agent's point of view.
	TripartyAgentCollateralTransactionIdentification *Max35Text `xml:"TrptyAgtCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *Max35Text `xml:"ClntTrptyCollTxId,omitempty"`
}

func (a *AdditionalParameters4) SetPreConfirmation(value string) {
	a.PreConfirmation = (*PreConfirmation1Code)(&value)
}

func (a *AdditionalParameters4) SetPartialSettlement(value string) {
	a.PartialSettlement = (*PartialSettlement1Code)(&value)
}

func (a *AdditionalParameters4) SetTripartyAgentCollateralTransactionIdentification(value string) {
	a.TripartyAgentCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters4) SetClientTripartyCollateralTransactionIdentification(value string) {
	a.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}
