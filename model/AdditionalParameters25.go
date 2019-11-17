package model

// Specifies additional parameters to the message or transaction.
type AdditionalParameters25 struct {

	// Specifies partial settlement information.
	PartialSettlement *PartialSettlement2Code `xml:"PrtlSttlm,omitempty"`

	// Identification of the confirmation previously sent to confirm the partial settlement of a transaction.
	PreviousPartialConfirmationIdentification *RestrictedFINXMax16Text `xml:"PrvsPrtlConfId,omitempty"`

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctOwnrTxId,omitempty"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctSvcrTxId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *RestrictedFINXMax16Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"CorpActnEvtId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *RestrictedFINXMax16Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *RestrictedFINXMax16Text `xml:"PrcrTxId,omitempty"`
}

func (a *AdditionalParameters25) SetPartialSettlement(value string) {
	a.PartialSettlement = (*PartialSettlement2Code)(&value)
}

func (a *AdditionalParameters25) SetPreviousPartialConfirmationIdentification(value string) {
	a.PreviousPartialConfirmationIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (a *AdditionalParameters25) SetAccountOwnerTransactionIdentification(value string) {
	a.AccountOwnerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (a *AdditionalParameters25) SetAccountServicerTransactionIdentification(value string) {
	a.AccountServicerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (a *AdditionalParameters25) SetPoolIdentification(value string) {
	a.PoolIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (a *AdditionalParameters25) SetCorporateActionEventIdentification(value string) {
	a.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (a *AdditionalParameters25) SetMarketInfrastructureTransactionIdentification(value string) {
	a.MarketInfrastructureTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (a *AdditionalParameters25) SetProcessorTransactionIdentification(value string) {
	a.ProcessorTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}
