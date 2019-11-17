package model

// Choice of reference.
type References43Choice struct {

	// Unambiguous identification of the confirmation to be cancelled.
	SecuritiesSettlementTransactionConfirmationIdentification *SettlementTypeAndIdentification18 `xml:"SctiesSttlmTxConfId"`

	// Unambiguous account servicer identification of the intra-position movement confirmation to be cancelled.
	IntraPositionMovementConfirmationIdentification *Max35Text `xml:"IntraPosMvmntConfId"`

	// Unambiguous identification of the report message to be cancelled.
	SecuritiesBalanceAccountingReportIdentification *Max35Text `xml:"SctiesBalAcctgRptId"`

	// Unambiguous identification of the report message to be cancelled.
	SecuritiesBalanceCustodyReportIdentification *Max35Text `xml:"SctiesBalCtdyRptId"`

	// Identification of the intra-position movement posting report to be cancelled.
	IntraPositionMovementPostingReportIdentification *Max35Text `xml:"IntraPosMvmntPstngRptId"`

	// Unambiguous identification of the confirmation to be cancelled.
	SecuritiesFinancingConfirmationIdentification *SettlementTypeAndIdentification18 `xml:"SctiesFincgConfId"`

	// Unambiguous identification of the report message to be cancelled.
	SecuritiesTransactionPendingReportIdentification *Max35Text `xml:"SctiesTxPdgRptId"`

	// Unambiguous identification of the report message to be cancelled.
	SecuritiesTransactionPostingReportIdentification *Max35Text `xml:"SctiesTxPstngRptId"`

	// Unambiguous identification of the report to be cancelled.
	SecuritiesSettlementTransactionAllegementReportIdentification *Max35Text `xml:"SctiesSttlmTxAllgmtRptId"`

	// Unambiguous identification of the allegement notification to be cancelled as know by the account servicer.
	SecuritiesSettlementTransactionAllegementNotificationTransactionIdentification *SettlementTypeAndIdentification18 `xml:"SctiesSttlmTxAllgmtNtfctnTxId"`

	// Identification of the portfolio transfer.
	PortfolioTransferNotificationIdentification *Max35Text `xml:"PrtflTrfNtfctnId"`

	// Unambiguous identification of the notification message to be cancelled.
	SecuritiesSettlementTransactionGenerationNotificationIdentification *SettlementTypeAndIdentification18 `xml:"SctiesSttlmTxGnrtnNtfctnId"`

	// Unambiguous identification of the message to be cancelled.
	OtherMessageIdentification *Max35Text `xml:"OthrMsgId"`

	// Unique identification of the report.
	TotalPortfolioValuationReportIdentification *Max35Text `xml:"TtlPrtflValtnRptId"`
}

func (r *References43Choice) AddSecuritiesSettlementTransactionConfirmationIdentification() *SettlementTypeAndIdentification18 {
	r.SecuritiesSettlementTransactionConfirmationIdentification = new(SettlementTypeAndIdentification18)
	return r.SecuritiesSettlementTransactionConfirmationIdentification
}

func (r *References43Choice) SetIntraPositionMovementConfirmationIdentification(value string) {
	r.IntraPositionMovementConfirmationIdentification = (*Max35Text)(&value)
}

func (r *References43Choice) SetSecuritiesBalanceAccountingReportIdentification(value string) {
	r.SecuritiesBalanceAccountingReportIdentification = (*Max35Text)(&value)
}

func (r *References43Choice) SetSecuritiesBalanceCustodyReportIdentification(value string) {
	r.SecuritiesBalanceCustodyReportIdentification = (*Max35Text)(&value)
}

func (r *References43Choice) SetIntraPositionMovementPostingReportIdentification(value string) {
	r.IntraPositionMovementPostingReportIdentification = (*Max35Text)(&value)
}

func (r *References43Choice) AddSecuritiesFinancingConfirmationIdentification() *SettlementTypeAndIdentification18 {
	r.SecuritiesFinancingConfirmationIdentification = new(SettlementTypeAndIdentification18)
	return r.SecuritiesFinancingConfirmationIdentification
}

func (r *References43Choice) SetSecuritiesTransactionPendingReportIdentification(value string) {
	r.SecuritiesTransactionPendingReportIdentification = (*Max35Text)(&value)
}

func (r *References43Choice) SetSecuritiesTransactionPostingReportIdentification(value string) {
	r.SecuritiesTransactionPostingReportIdentification = (*Max35Text)(&value)
}

func (r *References43Choice) SetSecuritiesSettlementTransactionAllegementReportIdentification(value string) {
	r.SecuritiesSettlementTransactionAllegementReportIdentification = (*Max35Text)(&value)
}

func (r *References43Choice) AddSecuritiesSettlementTransactionAllegementNotificationTransactionIdentification() *SettlementTypeAndIdentification18 {
	r.SecuritiesSettlementTransactionAllegementNotificationTransactionIdentification = new(SettlementTypeAndIdentification18)
	return r.SecuritiesSettlementTransactionAllegementNotificationTransactionIdentification
}

func (r *References43Choice) SetPortfolioTransferNotificationIdentification(value string) {
	r.PortfolioTransferNotificationIdentification = (*Max35Text)(&value)
}

func (r *References43Choice) AddSecuritiesSettlementTransactionGenerationNotificationIdentification() *SettlementTypeAndIdentification18 {
	r.SecuritiesSettlementTransactionGenerationNotificationIdentification = new(SettlementTypeAndIdentification18)
	return r.SecuritiesSettlementTransactionGenerationNotificationIdentification
}

func (r *References43Choice) SetOtherMessageIdentification(value string) {
	r.OtherMessageIdentification = (*Max35Text)(&value)
}

func (r *References43Choice) SetTotalPortfolioValuationReportIdentification(value string) {
	r.TotalPortfolioValuationReportIdentification = (*Max35Text)(&value)
}
