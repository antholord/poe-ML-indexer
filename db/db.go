package db

import (
	"gopkg.in/mgo.v2"
	"log"
	"os"
	"github.com/antholord/poe-ML-indexer/api"
)

var dbString = os.Getenv("db")
type DB struct {
	Session *mgo.Session
	ScTempColl *mgo.Collection
}


func Connect() *DB{
	if (dbString == ""){//dbString = "mongodb://test:test@ds123371.mlab.com:23371/heroku_lnc7sl64"
		dbString = "mongodb://woned:w0n3dp455 @34.209.73.153:27017/test"
	}
	s, err := mgo.Dial(dbString)
	if err != nil {
		panic(err)
	}
	coll := s.DB("heroku_lnc7sl64").C("Legacy")

	return &DB{Session : s, ScTempColl : coll}
}


func (DB *DB)BulkInsert(itemSlice []*api.Item){
	if (len(itemSlice) == 0) {return}
	bulk := DB.ScTempColl.Bulk()
	bulk.Unordered()
	for _, v := range itemSlice {
		bulk.Insert(v)
	}
	_, err := bulk.Run()
	if err != nil {
		log.Fatal(err)
	}
}
