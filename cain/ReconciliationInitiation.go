package cain

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00700101 struct {
	XMLName xml.Name                  `xml:"urn:iso:std:iso:20022:tech:xsd:cain.007.001.01 Document"`
	Message *ReconciliationInitiation `xml:"RcncltnInitn"`
}

func (d *Document00700101) AddMessage() *ReconciliationInitiation {
	d.Message = new(ReconciliationInitiation)
	return d.Message
}

// The ReconciliationInitiation message is sent by an acquirer or an agent to an issuer or an agent, to initiate an exchange of totals to be reconciled for debits, credits, chargebacks and other transactions.
type ReconciliationInitiation struct {

	// Information related to the protocol management.
	Header *model.Header17 `xml:"Hdr"`

	// Information related to the reconciliation.
	ReconciliationInitiation *model.AcquirerReconciliationInitiation1 `xml:"RcncltnInitn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *model.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (r *ReconciliationInitiation) AddHeader() *model.Header17 {
	r.Header = new(model.Header17)
	return r.Header
}

func (r *ReconciliationInitiation) AddReconciliationInitiation() *model.AcquirerReconciliationInitiation1 {
	r.ReconciliationInitiation = new(model.AcquirerReconciliationInitiation1)
	return r.ReconciliationInitiation
}

func (r *ReconciliationInitiation) AddSecurityTrailer() *model.ContentInformationType15 {
	r.SecurityTrailer = new(model.ContentInformationType15)
	return r.SecurityTrailer
}
