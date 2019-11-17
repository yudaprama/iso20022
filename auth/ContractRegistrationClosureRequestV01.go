package auth

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02000101 struct {
	XMLName xml.Name                               `xml:"urn:iso:std:iso:20022:tech:xsd:auth.020.001.01 Document"`
	Message *ContractRegistrationClosureRequestV01 `xml:"CtrctRegnClsrReq"`
}

func (d *Document02000101) AddMessage() *ContractRegistrationClosureRequestV01 {
	d.Message = new(ContractRegistrationClosureRequestV01)
	return d.Message
}

// The ContractRegistrationClosureRequest message is sent by the reporting party to the registration agent to close the registered contract subject to currency control.
type ContractRegistrationClosureRequestV01 struct {

	// Characteristics shared by all individual items included in the message.
	GroupHeader *model.CurrencyControlHeader1 `xml:"GrpHdr"`

	// Details on the closure of the registered contract.
	RegisteredContractClosure []*model.RegisteredContract2 `xml:"RegdCtrctClsr"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *ContractRegistrationClosureRequestV01) AddGroupHeader() *model.CurrencyControlHeader1 {
	c.GroupHeader = new(model.CurrencyControlHeader1)
	return c.GroupHeader
}

func (c *ContractRegistrationClosureRequestV01) AddRegisteredContractClosure() *model.RegisteredContract2 {
	newValue := new(model.RegisteredContract2)
	c.RegisteredContractClosure = append(c.RegisteredContractClosure, newValue)
	return newValue
}

func (c *ContractRegistrationClosureRequestV01) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
