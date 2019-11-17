package model

// Choice of reference.
type References37Choice struct {

	// Unambiguous identification of the confirmation to be cancelled.
	SecuritiesSettlementTransactionConfirmationIdentification *SettlementTypeAndIdentification13 `xml:"SctiesSttlmTxConfId"`

	// Unambiguous account servicer identification of the intra-position movement confirmation to be cancelled.
	IntraPositionMovementConfirmationIdentification *Max35Text `xml:"IntraPosMvmntConfId"`

	// Unambiguous identification of the report message to be cancelled.
	SecuritiesBalanceAccountingReportIdentification *Max35Text `xml:"SctiesBalAcctgRptId"`

	// Unambiguous identification of the report message to be cancelled.
	SecuritiesBalanceCustodyReportIdentification *Max35Text `xml:"SctiesBalCtdyRptId"`

	// Identification of the intra-position movement posting report to be cancelled.
	IntraPositionMovementPostingReportIdentification *Max35Text `xml:"IntraPosMvmntPstngRptId"`

	// Unambiguous identification of the confirmation to be cancelled.
	SecuritiesFinancingConfirmationIdentification *SettlementTypeAndIdentification13 `xml:"SctiesFincgConfId"`

	// Unambiguous identification of the report message to be cancelled.
	SecuritiesTransactionPendingReportIdentification *Max35Text `xml:"SctiesTxPdgRptId"`

	// Unambiguous identification of the report message to be cancelled.
	SecuritiesTransactionPostingReportIdentification *Max35Text `xml:"SctiesTxPstngRptId"`

	// Unambiguous identification of the report to be cancelled.
	SecuritiesSettlementTransactionAllegementReportIdentification *Max35Text `xml:"SctiesSttlmTxAllgmtRptId"`

	// Unambiguous identification of the allegement notification to be cancelled as know by the account servicer.
	SecuritiesSettlementTransactionAllegementNotificationTransactionIdentification *SettlementTypeAndIdentification13 `xml:"SctiesSttlmTxAllgmtNtfctnTxId"`

	// Identification of the portfolio transfer.
	PortfolioTransferNotificationIdentification *Max35Text `xml:"PrtflTrfNtfctnId"`

	// Unambiguous identification of the notification message to be cancelled.
	SecuritiesSettlementTransactionGenerationNotificationIdentification *SettlementTypeAndIdentification13 `xml:"SctiesSttlmTxGnrtnNtfctnId"`

	// Unambiguous identification of the message to be cancelled.
	OtherMessageIdentification *Max35Text `xml:"OthrMsgId"`

	// Unique identification of the report.
	TotalPortfolioValuationReportIdentification *Max35Text `xml:"TtlPrtflValtnRptId"`
}

func (r *References37Choice) AddSecuritiesSettlementTransactionConfirmationIdentification() *SettlementTypeAndIdentification13 {
	r.SecuritiesSettlementTransactionConfirmationIdentification = new(SettlementTypeAndIdentification13)
	return r.SecuritiesSettlementTransactionConfirmationIdentification
}

func (r *References37Choice) SetIntraPositionMovementConfirmationIdentification(value string) {
	r.IntraPositionMovementConfirmationIdentification = (*Max35Text)(&value)
}

func (r *References37Choice) SetSecuritiesBalanceAccountingReportIdentification(value string) {
	r.SecuritiesBalanceAccountingReportIdentification = (*Max35Text)(&value)
}

func (r *References37Choice) SetSecuritiesBalanceCustodyReportIdentification(value string) {
	r.SecuritiesBalanceCustodyReportIdentification = (*Max35Text)(&value)
}

func (r *References37Choice) SetIntraPositionMovementPostingReportIdentification(value string) {
	r.IntraPositionMovementPostingReportIdentification = (*Max35Text)(&value)
}

func (r *References37Choice) AddSecuritiesFinancingConfirmationIdentification() *SettlementTypeAndIdentification13 {
	r.SecuritiesFinancingConfirmationIdentification = new(SettlementTypeAndIdentification13)
	return r.SecuritiesFinancingConfirmationIdentification
}

func (r *References37Choice) SetSecuritiesTransactionPendingReportIdentification(value string) {
	r.SecuritiesTransactionPendingReportIdentification = (*Max35Text)(&value)
}

func (r *References37Choice) SetSecuritiesTransactionPostingReportIdentification(value string) {
	r.SecuritiesTransactionPostingReportIdentification = (*Max35Text)(&value)
}

func (r *References37Choice) SetSecuritiesSettlementTransactionAllegementReportIdentification(value string) {
	r.SecuritiesSettlementTransactionAllegementReportIdentification = (*Max35Text)(&value)
}

func (r *References37Choice) AddSecuritiesSettlementTransactionAllegementNotificationTransactionIdentification() *SettlementTypeAndIdentification13 {
	r.SecuritiesSettlementTransactionAllegementNotificationTransactionIdentification = new(SettlementTypeAndIdentification13)
	return r.SecuritiesSettlementTransactionAllegementNotificationTransactionIdentification
}

func (r *References37Choice) SetPortfolioTransferNotificationIdentification(value string) {
	r.PortfolioTransferNotificationIdentification = (*Max35Text)(&value)
}

func (r *References37Choice) AddSecuritiesSettlementTransactionGenerationNotificationIdentification() *SettlementTypeAndIdentification13 {
	r.SecuritiesSettlementTransactionGenerationNotificationIdentification = new(SettlementTypeAndIdentification13)
	return r.SecuritiesSettlementTransactionGenerationNotificationIdentification
}

func (r *References37Choice) SetOtherMessageIdentification(value string) {
	r.OtherMessageIdentification = (*Max35Text)(&value)
}

func (r *References37Choice) SetTotalPortfolioValuationReportIdentification(value string) {
	r.TotalPortfolioValuationReportIdentification = (*Max35Text)(&value)
}
