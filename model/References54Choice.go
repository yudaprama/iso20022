package model

// Choice of reference.
type References54Choice struct {

	// Unambiguous identification of a securities settlement transaction as known by the account owner (or instructing party acting on its behalf).
	SecuritiesSettlementTransactionIdentification *RestrictedFINXMax16Text `xml:"SctiesSttlmTxId"`

	// Unambiguous identification of the intra-position movement transaction as known by the account owner (or instructing party acting on its behalf).
	IntraPositionMovementIdentification *RestrictedFINXMax16Text `xml:"IntraPosMvmntId"`

	// Unambiguous identification of the intra balance movement transaction as known by the account owner.
	IntraBalanceMovementIdentification *RestrictedFINXMax16Text `xml:"IntraBalMvmntId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctSvcrTxId"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *RestrictedFINXMax16Text `xml:"MktInfrstrctrTxId"`

	// Collective reference identifying a set of messages.
	PoolIdentification *RestrictedFINXMax16Text `xml:"PoolId"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *RestrictedFINXMax16Text `xml:"CmonId"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification *RestrictedFINXMax16Text `xml:"TradId"`

	// Identification of a transaction that cannot be identified using a standard reference element present in the message.
	OtherTransactionIdentification *RestrictedFINXMax16Text `xml:"OthrTxId"`
}

func (r *References54Choice) SetSecuritiesSettlementTransactionIdentification(value string) {
	r.SecuritiesSettlementTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (r *References54Choice) SetIntraPositionMovementIdentification(value string) {
	r.IntraPositionMovementIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (r *References54Choice) SetIntraBalanceMovementIdentification(value string) {
	r.IntraBalanceMovementIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (r *References54Choice) SetAccountServicerTransactionIdentification(value string) {
	r.AccountServicerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (r *References54Choice) SetMarketInfrastructureTransactionIdentification(value string) {
	r.MarketInfrastructureTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (r *References54Choice) SetPoolIdentification(value string) {
	r.PoolIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (r *References54Choice) SetCommonIdentification(value string) {
	r.CommonIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (r *References54Choice) SetTradeIdentification(value string) {
	r.TradeIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (r *References54Choice) SetOtherTransactionIdentification(value string) {
	r.OtherTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}
