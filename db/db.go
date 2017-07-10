package db

import (
	"gopkg.in/mgo.v2"
	"log"
	"gopkg.in/mgo.v2/bson"
	"os"
)

var dbString = os.Getenv("db")
var dbDefault = "heroku_lnc7sl64"
type DB struct {
	Session *mgo.Session
	ScTempColl *mgo.Collection
}


func (DB *DB) Connect() *DB{
	if (dbString == ""){
		dbString = "mongodb://test:test@ds123371.mlab.com:23371/heroku_lnc7sl64"
	}
	s, err := mgo.Dial(dbString)
	if err != nil {
		panic(err)
	}
	DB.ScTempColl = DB.Session.DB(dbString).C("Legacy")
	return &DB{Session : s}
}



func (DB *DB)UpsertString(c string, s string){
	if DB.Session == nil { panic("Session not instanciated") }

	coll := DB.Session.DB(dbstring).C(c)
	log.Println("Adding " + s + " to " + c)
	_, err:= coll.Upsert(bson.M{"value" : s}, bson.M{"value" : s}) ; if err != nil{
		log.Fatal(err)
	} else {

	}
}
