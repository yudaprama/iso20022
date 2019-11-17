package auth

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02300101 struct {
	XMLName xml.Name                                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.023.001.01 Document"`
	Message *ContractRegistrationStatementRequestV01 `xml:"CtrctRegnStmtReq"`
}

func (d *Document02300101) AddMessage() *ContractRegistrationStatementRequestV01 {
	d.Message = new(ContractRegistrationStatementRequestV01)
	return d.Message
}

// The ContractRegistrationStatementRequest message is sent by the reporting party to the registration agent to request for a statement of the operations related to the registered contract subject to currency control.
type ContractRegistrationStatementRequestV01 struct {

	// Characteristics shared by all individual items included in the message.
	GroupHeader *model.CurrencyControlHeader1 `xml:"GrpHdr"`

	// Details on the information requested for the contract registration statement.
	StatementRequest []*model.ContractRegistrationStatementRequest1 `xml:"StmtReq"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *ContractRegistrationStatementRequestV01) AddGroupHeader() *model.CurrencyControlHeader1 {
	c.GroupHeader = new(model.CurrencyControlHeader1)
	return c.GroupHeader
}

func (c *ContractRegistrationStatementRequestV01) AddStatementRequest() *model.ContractRegistrationStatementRequest1 {
	newValue := new(model.ContractRegistrationStatementRequest1)
	c.StatementRequest = append(c.StatementRequest, newValue)
	return newValue
}

func (c *ContractRegistrationStatementRequestV01) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
