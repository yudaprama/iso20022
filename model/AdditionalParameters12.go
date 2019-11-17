package model

// Specifies additional parameters to the message or transaction.
type AdditionalParameters12 struct {

	// Specifies whether there exists a pre-confirmation.
	PreConfirmation *PreConfirmation1Code `xml:"PreConf,omitempty"`

	// Specifies partial settlement information.
	PartialSettlement *PartialSettlement2Code `xml:"PrtlSttlm,omitempty"`

	// Identification of the confirmation previously sent to confirm the partial settlement of a transaction.
	PreviousPartialConfirmationIdentification *Max35Text `xml:"PrvsPrtlConfId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the triparty agent's point of view.
	TripartyAgentCollateralTransactionIdentification *Max35Text `xml:"TrptyAgtCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *Max35Text `xml:"ClntTrptyCollTxId,omitempty"`
}

func (a *AdditionalParameters12) SetPreConfirmation(value string) {
	a.PreConfirmation = (*PreConfirmation1Code)(&value)
}

func (a *AdditionalParameters12) SetPartialSettlement(value string) {
	a.PartialSettlement = (*PartialSettlement2Code)(&value)
}

func (a *AdditionalParameters12) SetPreviousPartialConfirmationIdentification(value string) {
	a.PreviousPartialConfirmationIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters12) SetTripartyAgentCollateralTransactionIdentification(value string) {
	a.TripartyAgentCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters12) SetClientTripartyCollateralTransactionIdentification(value string) {
	a.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}
