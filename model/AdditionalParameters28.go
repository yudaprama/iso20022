package model

// Specifies additional parameters to the message or transaction.
type AdditionalParameters28 struct {

	// Specifies whether there exists a pre-confirmation.
	PreConfirmation *PreConfirmation1Code `xml:"PreConf,omitempty"`

	// Specifies partial settlement information.
	PartialSettlement *PartialSettlement2Code `xml:"PrtlSttlm,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the triparty agent's point of view.
	TripartyAgentCollateralTransactionIdentification *RestrictedFINXMax16Text `xml:"TrptyAgtCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *RestrictedFINXMax16Text `xml:"ClntTrptyCollTxId,omitempty"`
}

func (a *AdditionalParameters28) SetPreConfirmation(value string) {
	a.PreConfirmation = (*PreConfirmation1Code)(&value)
}

func (a *AdditionalParameters28) SetPartialSettlement(value string) {
	a.PartialSettlement = (*PartialSettlement2Code)(&value)
}

func (a *AdditionalParameters28) SetTripartyAgentCollateralTransactionIdentification(value string) {
	a.TripartyAgentCollateralTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (a *AdditionalParameters28) SetClientTripartyCollateralTransactionIdentification(value string) {
	a.ClientTripartyCollateralTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}
