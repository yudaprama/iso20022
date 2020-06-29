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

	log.Printf("nama pengirim:  %v", *messageParsed.CreditTransferTransactionInformation[0].Creditor.Name)
	log.Printf("alamat pengirim:  %v", *messageParsed.CreditTransferTransactionInformation[0].Creditor.PostalAddress.StreetName)
	log.Printf("negara pengirim:  %v", *messageParsed.CreditTransferTransactionInformation[0].Creditor.PostalAddress.Country)
	log.Printf("nama penerima:  %v", *messageParsed.CreditTransferTransactionInformation[0].Debtor.Name)
	log.Printf("alamat penerima:  %v", *messageParsed.CreditTransferTransactionInformation[0].Debtor.PostalAddress.StreetName)
	log.Printf("negara penerima:  %v", *messageParsed.CreditTransferTransactionInformation[0].Debtor.PostalAddress.Country)
	log.Printf("amount:  %v", messageParsed.CreditTransferTransactionInformation[0].InterbankSettlementAmount.Value)
	log.Printf("currency:  USD")
}
