package model

// Securities financing transaction identification information, type (repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing) and other parameters.
type TransactionTypeAndAdditionalParameters18 struct {

	// Specifies the type of securities financing transaction, that is, repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing.
	SecuritiesFinancingTransactionType *SecuritiesFinancingTransactionType2Code `xml:"SctiesFincgTxTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *RestrictedFINXMax16Text `xml:"CmonId,omitempty"`

	// Settlement transaction has already been sent on the market. It is sent by an account owner to an account servicer for reconciliation purposes.
	ReconciliationIndicator *YesNoIndicator `xml:"RcncltnInd,omitempty"`
}

func (t *TransactionTypeAndAdditionalParameters18) SetSecuritiesFinancingTransactionType(value string) {
	t.SecuritiesFinancingTransactionType = (*SecuritiesFinancingTransactionType2Code)(&value)
}

func (t *TransactionTypeAndAdditionalParameters18) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionTypeAndAdditionalParameters18) SetCommonIdentification(value string) {
	t.CommonIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionTypeAndAdditionalParameters18) SetReconciliationIndicator(value string) {
	t.ReconciliationIndicator = (*YesNoIndicator)(&value)
}
