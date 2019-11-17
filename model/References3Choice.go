package model

// Choice of reference.
type References3Choice struct {

	// Unambiguous identification of the confirmation to be cancelled.
	SecuritiesSettlementTransactionConfirmationIdentification *Max35Text `xml:"SctiesSttlmTxConfId"`

	// Unambiguous account servicer identification of the intra-position movement confirmation to be cancelled.
	IntraPositionMovementConfirmationIdentification *Max35Text `xml:"IntraPosMvmntConfId"`

	// Unambiguous identification of the report message to be cancelled.
	SecuritiesBalanceAccountingReportIdentification *Max35Text `xml:"SctiesBalAcctgRptId"`

	// Unambiguous identification of the report message to be cancelled.
	SecuritiesBalanceCustodyReportIdentification *Max35Text `xml:"SctiesBalCtdyRptId"`

	// Identification of the intra-position movement posting report to be cancelled.
	IntraPositionMovementPostingReportIdentification *Max35Text `xml:"IntraPosMvmntPstngRptId"`

	// Unambiguous identification of the confirmation to be cancelled.
	SecuritiesFinancingConfirmationIdentification *Max35Text `xml:"SctiesFincgConfId"`

	// Unambiguous identification of the report message to be cancelled.
	SecuritiesTransactionPendingReportIdentification *Max35Text `xml:"SctiesTxPdgRptId"`

	// Unambiguous identification of the report message to be cancelled.
	SecuritiesTransactionPostingReportIdentification *Max35Text `xml:"SctiesTxPstngRptId"`

	// Unambiguous identification of the report to be cancelled.
	SecuritiesSettlementTransactionAllegementReportIdentification *Max35Text `xml:"SctiesSttlmTxAllgmtRptId"`

	// Unambiguous identification of the allegement notification to be cancelled as know by the account servicer.
	SecuritiesSettlementTransactionAllegementNotificationTransactionIdentification *Max35Text `xml:"SctiesSttlmTxAllgmtNtfctnTxId"`

	// Identification of the portfolio transfer.
	PortfolioTransferNotificationIdentification *Max35Text `xml:"PrtflTrfNtfctnId"`

	// Unambiguous identification of the notification message to be cancelled.
	SecuritiesSettlementTransactionGenerationNotificationIdentification *Max35Text `xml:"SctiesSttlmTxGnrtnNtfctnId"`

	// Unambiguous identification of the message to be cancelled.
	OtherMessageIdentification *Max35Text `xml:"OthrMsgId"`
}

func (r *References3Choice) SetSecuritiesSettlementTransactionConfirmationIdentification(value string) {
	r.SecuritiesSettlementTransactionConfirmationIdentification = (*Max35Text)(&value)
}

func (r *References3Choice) SetIntraPositionMovementConfirmationIdentification(value string) {
	r.IntraPositionMovementConfirmationIdentification = (*Max35Text)(&value)
}

func (r *References3Choice) SetSecuritiesBalanceAccountingReportIdentification(value string) {
	r.SecuritiesBalanceAccountingReportIdentification = (*Max35Text)(&value)
}

func (r *References3Choice) SetSecuritiesBalanceCustodyReportIdentification(value string) {
	r.SecuritiesBalanceCustodyReportIdentification = (*Max35Text)(&value)
}

func (r *References3Choice) SetIntraPositionMovementPostingReportIdentification(value string) {
	r.IntraPositionMovementPostingReportIdentification = (*Max35Text)(&value)
}

func (r *References3Choice) SetSecuritiesFinancingConfirmationIdentification(value string) {
	r.SecuritiesFinancingConfirmationIdentification = (*Max35Text)(&value)
}

func (r *References3Choice) SetSecuritiesTransactionPendingReportIdentification(value string) {
	r.SecuritiesTransactionPendingReportIdentification = (*Max35Text)(&value)
}

func (r *References3Choice) SetSecuritiesTransactionPostingReportIdentification(value string) {
	r.SecuritiesTransactionPostingReportIdentification = (*Max35Text)(&value)
}

func (r *References3Choice) SetSecuritiesSettlementTransactionAllegementReportIdentification(value string) {
	r.SecuritiesSettlementTransactionAllegementReportIdentification = (*Max35Text)(&value)
}

func (r *References3Choice) SetSecuritiesSettlementTransactionAllegementNotificationTransactionIdentification(value string) {
	r.SecuritiesSettlementTransactionAllegementNotificationTransactionIdentification = (*Max35Text)(&value)
}

func (r *References3Choice) SetPortfolioTransferNotificationIdentification(value string) {
	r.PortfolioTransferNotificationIdentification = (*Max35Text)(&value)
}

func (r *References3Choice) SetSecuritiesSettlementTransactionGenerationNotificationIdentification(value string) {
	r.SecuritiesSettlementTransactionGenerationNotificationIdentification = (*Max35Text)(&value)
}

func (r *References3Choice) SetOtherMessageIdentification(value string) {
	r.OtherMessageIdentification = (*Max35Text)(&value)
}
