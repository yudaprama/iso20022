package main

import (
	"encoding/xml"
	"github.com/yudaprama/iso20022/pacs"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	messages, err := ioutil.ReadFile("./example-message.xml")
	if err != nil {
		log.Fatalf("Unable to read file:  %v", err)
		os.Exit(1)
	}

	var messageParsed pacs.FIToFICustomerCreditTransferV06
	err = xml.Unmarshal(messages, &messageParsed)
	if err != nil {
		log.Fatalf("Unable to parse file:  %v", err)
		os.Exit(1)
	}

	log.Printf("Interbank settlement amount:  %v", messageParsed.CreditTransferTransactionInformation[0].InterbankSettlementAmount.Value)
	log.Printf("Interbank settlement currency:  %v", messageParsed.CreditTransferTransactionInformation[0].InterbankSettlementAmount.Currency)
}
