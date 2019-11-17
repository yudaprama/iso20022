package model

// Currency control document entry supporting the contract registration.
type SupportingDocumentEntry1 struct {

	// Unique and unambiguous identification of the supporting document entry.
	EntryIdentification *Max35Text `xml:"NtryId"`

	// Identification of the original document for which the supporting documents are provided.
	OriginalDocument *DocumentIdentification22 `xml:"OrgnlDoc"`

	// Document type in a coded form.
	//
	// TBC: Data must support "_".
	DocumentType *Exact4AlphaNumericText `xml:"DocTp"`

	// Total amount of the supporting document entry.
	TotalAmount *ActiveCurrencyAndAmount `xml:"TtlAmt,omitempty"`

	// Total amount after shipment of the supporting document entry.
	TotalAmountAfterShipment *ActiveCurrencyAndAmount `xml:"TtlAmtAftrShipmnt,omitempty"`

	// Total amount of the supporting document entry in the currency of the contract.
	TotalAmountInContractCurrency *ActiveCurrencyAndAmount `xml:"TtlAmtInCtrctCcy,omitempty"`

	// Total amount after shipment of the supporting document entry in the currency of the contract.
	TotalAmountAfterShipmentInContractCurrency *ActiveCurrencyAndAmount `xml:"TtlAmtAftrShipmntInCtrctCcy,omitempty"`

	// Conditions and attributes related to the shipment.
	ShipmentAttributes *ShipmentAttribute1 `xml:"ShipmntAttrbts"`

	// Further details on the supporting document entry.
	AdditionalInformation *Max500Text `xml:"AddtlInf,omitempty"`

	// Documents provided as attachments to the supporting document entry.
	Attachment []*DocumentGeneralInformation3 `xml:"Attchmnt,omitempty"`
}

func (s *SupportingDocumentEntry1) SetEntryIdentification(value string) {
	s.EntryIdentification = (*Max35Text)(&value)
}

func (s *SupportingDocumentEntry1) AddOriginalDocument() *DocumentIdentification22 {
	s.OriginalDocument = new(DocumentIdentification22)
	return s.OriginalDocument
}

func (s *SupportingDocumentEntry1) SetDocumentType(value string) {
	s.DocumentType = (*Exact4AlphaNumericText)(&value)
}

func (s *SupportingDocumentEntry1) SetTotalAmount(value, currency string) {
	s.TotalAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SupportingDocumentEntry1) SetTotalAmountAfterShipment(value, currency string) {
	s.TotalAmountAfterShipment = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SupportingDocumentEntry1) SetTotalAmountInContractCurrency(value, currency string) {
	s.TotalAmountInContractCurrency = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SupportingDocumentEntry1) SetTotalAmountAfterShipmentInContractCurrency(value, currency string) {
	s.TotalAmountAfterShipmentInContractCurrency = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SupportingDocumentEntry1) AddShipmentAttributes() *ShipmentAttribute1 {
	s.ShipmentAttributes = new(ShipmentAttribute1)
	return s.ShipmentAttributes
}

func (s *SupportingDocumentEntry1) SetAdditionalInformation(value string) {
	s.AdditionalInformation = (*Max500Text)(&value)
}

func (s *SupportingDocumentEntry1) AddAttachment() *DocumentGeneralInformation3 {
	newValue := new(DocumentGeneralInformation3)
	s.Attachment = append(s.Attachment, newValue)
	return newValue
}
