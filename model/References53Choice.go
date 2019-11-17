package model

// Choice of reference.
type References53Choice struct {

	// Unambiguous identification of the confirmation to be cancelled.
	SecuritiesSettlementTransactionConfirmationIdentification *SettlementTypeAndIdentification22 `xml:"SctiesSttlmTxConfId"`

	// Unambiguous account servicer identification of the intra-position movement confirmation to be cancelled.
	IntraPositionMovementConfirmationIdentification *RestrictedFINXMax16Text `xml:"IntraPosMvmntConfId"`

	// Unambiguous identification of the report message to be cancelled.
	SecuritiesBalanceAccountingReportIdentification *RestrictedFINXMax16Text `xml:"SctiesBalAcctgRptId"`

	// Unambiguous identification of the report message to be cancelled.
	SecuritiesBalanceCustodyReportIdentification *RestrictedFINXMax16Text `xml:"SctiesBalCtdyRptId"`

	// Identification of the intra-position movement posting report to be cancelled.
	IntraPositionMovementPostingReportIdentification *RestrictedFINXMax16Text `xml:"IntraPosMvmntPstngRptId"`

	// Unambiguous identification of the confirmation to be cancelled.
	SecuritiesFinancingConfirmationIdentification *SettlementTypeAndIdentification22 `xml:"SctiesFincgConfId"`

	// Unambiguous identification of the report message to be cancelled.
	SecuritiesTransactionPendingReportIdentification *RestrictedFINXMax16Text `xml:"SctiesTxPdgRptId"`

	// Unambiguous identification of the report message to be cancelled.
	SecuritiesTransactionPostingReportIdentification *RestrictedFINXMax16Text `xml:"SctiesTxPstngRptId"`

	// Unambiguous identification of the report to be cancelled.
	SecuritiesSettlementTransactionAllegementReportIdentification *RestrictedFINXMax16Text `xml:"SctiesSttlmTxAllgmtRptId"`

	// Unambiguous identification of the allegement notification to be cancelled as know by the account servicer.
	SecuritiesSettlementTransactionAllegementNotificationTransactionIdentification *SettlementTypeAndIdentification22 `xml:"SctiesSttlmTxAllgmtNtfctnTxId"`

	// Identification of the portfolio transfer.
	PortfolioTransferNotificationIdentification *RestrictedFINXMax16Text `xml:"PrtflTrfNtfctnId"`

	// Unambiguous identification of the notification message to be cancelled.
	SecuritiesSettlementTransactionGenerationNotificationIdentification *SettlementTypeAndIdentification22 `xml:"SctiesSttlmTxGnrtnNtfctnId"`

	// Unambiguous identification of the message to be cancelled.
	OtherMessageIdentification *RestrictedFINXMax16Text `xml:"OthrMsgId"`

	// Unique identification of the report.
	TotalPortfolioValuationReportIdentification *RestrictedFINXMax16Text `xml:"TtlPrtflValtnRptId"`
}

func (r *References53Choice) AddSecuritiesSettlementTransactionConfirmationIdentification() *SettlementTypeAndIdentification22 {
	r.SecuritiesSettlementTransactionConfirmationIdentification = new(SettlementTypeAndIdentification22)
	return r.SecuritiesSettlementTransactionConfirmationIdentification
}

func (r *References53Choice) SetIntraPositionMovementConfirmationIdentification(value string) {
	r.IntraPositionMovementConfirmationIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (r *References53Choice) SetSecuritiesBalanceAccountingReportIdentification(value string) {
	r.SecuritiesBalanceAccountingReportIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (r *References53Choice) SetSecuritiesBalanceCustodyReportIdentification(value string) {
	r.SecuritiesBalanceCustodyReportIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (r *References53Choice) SetIntraPositionMovementPostingReportIdentification(value string) {
	r.IntraPositionMovementPostingReportIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (r *References53Choice) AddSecuritiesFinancingConfirmationIdentification() *SettlementTypeAndIdentification22 {
	r.SecuritiesFinancingConfirmationIdentification = new(SettlementTypeAndIdentification22)
	return r.SecuritiesFinancingConfirmationIdentification
}

func (r *References53Choice) SetSecuritiesTransactionPendingReportIdentification(value string) {
	r.SecuritiesTransactionPendingReportIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (r *References53Choice) SetSecuritiesTransactionPostingReportIdentification(value string) {
	r.SecuritiesTransactionPostingReportIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (r *References53Choice) SetSecuritiesSettlementTransactionAllegementReportIdentification(value string) {
	r.SecuritiesSettlementTransactionAllegementReportIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (r *References53Choice) AddSecuritiesSettlementTransactionAllegementNotificationTransactionIdentification() *SettlementTypeAndIdentification22 {
	r.SecuritiesSettlementTransactionAllegementNotificationTransactionIdentification = new(SettlementTypeAndIdentification22)
	return r.SecuritiesSettlementTransactionAllegementNotificationTransactionIdentification
}

func (r *References53Choice) SetPortfolioTransferNotificationIdentification(value string) {
	r.PortfolioTransferNotificationIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (r *References53Choice) AddSecuritiesSettlementTransactionGenerationNotificationIdentification() *SettlementTypeAndIdentification22 {
	r.SecuritiesSettlementTransactionGenerationNotificationIdentification = new(SettlementTypeAndIdentification22)
	return r.SecuritiesSettlementTransactionGenerationNotificationIdentification
}

func (r *References53Choice) SetOtherMessageIdentification(value string) {
	r.OtherMessageIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (r *References53Choice) SetTotalPortfolioValuationReportIdentification(value string) {
	r.TotalPortfolioValuationReportIdentification = (*RestrictedFINXMax16Text)(&value)
}
