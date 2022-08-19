package controllers

import (
	"fmt"
	"log"

	"go.etcd.io/bbolt"
)

var (
	Db  *bbolt.DB
	err error
)

func main() {
	Db, err = bbolt.Open("./data.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = Db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("gitlabhook"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
}
