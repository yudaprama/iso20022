package model

// Unique identifier of a document, message or transaction.
type Identification24 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *RestrictedFINXMax16Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *RestrictedFINXMax16Text `xml:"PrcrTxId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *RestrictedFINXMax16Text `xml:"CmonId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*RestrictedFINXMax16Text `xml:"TradId,omitempty"`

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterIdentification *RestrictedFINXMax16Text `xml:"MstrId,omitempty"`

	// Identification of a basket trade.
	BasketIdentification *RestrictedFINXMax16Text `xml:"BsktId,omitempty"`

	// Reference identifying a index trade.
	IndexIdentification *RestrictedFINXMax16Text `xml:"IndxId,omitempty"`

	// Unique identifier for a list, as assigned by the trading party. The identifier must be unique within a single trading day.
	ListIdentification *RestrictedFINXMax16Text `xml:"ListId,omitempty"`

	// Program reference which identifies a program trade.
	ProgramIdentification *RestrictedFINXMax16Text `xml:"PrgmId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *RestrictedFINXMax16Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"CorpActnEvtId,omitempty"`
}

func (i *Identification24) SetAccountOwnerTransactionIdentification(value string) {
	i.AccountOwnerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (i *Identification24) SetAccountServicerTransactionIdentification(value string) {
	i.AccountServicerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (i *Identification24) SetMarketInfrastructureTransactionIdentification(value string) {
	i.MarketInfrastructureTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (i *Identification24) SetProcessorTransactionIdentification(value string) {
	i.ProcessorTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (i *Identification24) SetCommonIdentification(value string) {
	i.CommonIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (i *Identification24) AddTradeIdentification(value string) {
	i.TradeIdentification = append(i.TradeIdentification, (*RestrictedFINXMax16Text)(&value))
}

func (i *Identification24) SetMasterIdentification(value string) {
	i.MasterIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (i *Identification24) SetBasketIdentification(value string) {
	i.BasketIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (i *Identification24) SetIndexIdentification(value string) {
	i.IndexIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (i *Identification24) SetListIdentification(value string) {
	i.ListIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (i *Identification24) SetProgramIdentification(value string) {
	i.ProgramIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (i *Identification24) SetPoolIdentification(value string) {
	i.PoolIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (i *Identification24) SetCorporateActionEventIdentification(value string) {
	i.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}
