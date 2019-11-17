package model

// Specifies the conditions under which the order/trade is to be settled.
type TransactionTypeAndAdditionalParameters20 struct {

	// Unambiguous identification of the transaction (unique per piece of collateral) as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction (unique per piece of collateral) as known by the account servicer.
	AccountServicerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctSvcrTxId,omitempty"`

	// Specifies the type of securities financing transaction, that is, repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing.
	SecuritiesFinancingTransactionType *SecuritiesFinancingTransactionType2Code `xml:"SctiesFincgTxTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Identifies the type of securities financing modification requested.
	ModificationType *RepurchaseType31Choice `xml:"ModTp,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *RestrictedFINXMax16Text `xml:"CmonId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *RestrictedFINXMax16Text `xml:"PoolId,omitempty"`
}

func (t *TransactionTypeAndAdditionalParameters20) SetAccountOwnerTransactionIdentification(value string) {
	t.AccountOwnerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionTypeAndAdditionalParameters20) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionTypeAndAdditionalParameters20) SetSecuritiesFinancingTransactionType(value string) {
	t.SecuritiesFinancingTransactionType = (*SecuritiesFinancingTransactionType2Code)(&value)
}

func (t *TransactionTypeAndAdditionalParameters20) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionTypeAndAdditionalParameters20) AddModificationType() *RepurchaseType31Choice {
	t.ModificationType = new(RepurchaseType31Choice)
	return t.ModificationType
}

func (t *TransactionTypeAndAdditionalParameters20) SetCommonIdentification(value string) {
	t.CommonIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionTypeAndAdditionalParameters20) SetPoolIdentification(value string) {
	t.PoolIdentification = (*RestrictedFINXMax16Text)(&value)
}
