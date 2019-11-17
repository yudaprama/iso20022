package model

// Structured information related to the invoice to be financed.
type ReferredDocumentInformation2 struct {

	// Specifies the type of the document, for example commercial invoice.
	Type *ReferredDocumentType1 `xml:"Tp,omitempty"`

	// Unique and unambiguous identification number of the referred document.
	DocumentNumber *Max35Text `xml:"DocNb,omitempty"`

	// Date associated with the referred document, eg, date of issue.
	RelatedDate *ISODate `xml:"RltdDt,omitempty"`

	// Amount of money and currency of a document referred to invoice to be financed.
	DocumentAmount *ActiveCurrencyAndAmount `xml:"DocAmt,omitempty"`
}

func (r *ReferredDocumentInformation2) AddType() *ReferredDocumentType1 {
	r.Type = new(ReferredDocumentType1)
	return r.Type
}

func (r *ReferredDocumentInformation2) SetDocumentNumber(value string) {
	r.DocumentNumber = (*Max35Text)(&value)
}

func (r *ReferredDocumentInformation2) SetRelatedDate(value string) {
	r.RelatedDate = (*ISODate)(&value)
}

func (r *ReferredDocumentInformation2) SetDocumentAmount(value, currency string) {
	r.DocumentAmount = NewActiveCurrencyAndAmount(value, currency)
}
