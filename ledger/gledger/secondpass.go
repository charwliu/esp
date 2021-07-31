package main

import (
	"log"

	"go.vixal.xyz/esp/ledger/api"
)

func secondPass(db api.Datastorer) error {
	log.Printf("secondPass\n")
	if err := db.SecondPass(); err != nil {
		log.Printf("%v\n", err)
		return err
	}
	return nil
}
