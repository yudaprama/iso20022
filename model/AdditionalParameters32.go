package model

// Specifies additional parameters to the message or transaction.
type AdditionalParameters32 struct {

	// Specifies whether there exists a pre-confirmation.
	PreConfirmation *PreConfirmation1Code `xml:"PreConf,omitempty"`

	// Specifies partial settlement information.
	PartialSettlement *PartialSettlement2Code `xml:"PrtlSttlm,omitempty"`

	// Unique identification identifying the triparty collateral management transaction from the triparty-agent's/service-provider's point of view.
	TripartyAgentServiceProviderCollateralTransactionIdentification *RestrictedFINXMax16Text `xml:"TrptyAgtSvcPrvdrCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *RestrictedFINXMax16Text `xml:"ClntTrptyCollTxId,omitempty"`
}

func (a *AdditionalParameters32) SetPreConfirmation(value string) {
	a.PreConfirmation = (*PreConfirmation1Code)(&value)
}

func (a *AdditionalParameters32) SetPartialSettlement(value string) {
	a.PartialSettlement = (*PartialSettlement2Code)(&value)
}

func (a *AdditionalParameters32) SetTripartyAgentServiceProviderCollateralTransactionIdentification(value string) {
	a.TripartyAgentServiceProviderCollateralTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (a *AdditionalParameters32) SetClientTripartyCollateralTransactionIdentification(value string) {
	a.ClientTripartyCollateralTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}
