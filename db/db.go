package db

import (
	"gopkg.in/mgo.v2"
	"log"
	"gopkg.in/mgo.v2/bson"
)

var dbstring = "heroku_lnc7sl64"
type DB struct {
	Session *mgo.Session
}


func Connect() *DB{

	s, err := mgo.Dial("mongodb://test:test@ds123371.mlab.com:23371/" + dbstring)
	if err != nil {
		panic(err)
	}

	return &DB{Session : s}
}

func (DB *DB)LoadCollection(c string){
	if DB.Session == nil { panic("Session not instanciated") }

	//coll := Session.DB("test").C(c)

}

func (DB *DB)UpdateCollection(c string, data interface{}){
	if DB.Session == nil { panic("Session not instanciated") }
	coll := DB.Session.DB(dbstring).C(c)
	err := coll.Insert(bson.M{"id" : 1, "data" : &data}) ; if err != nil {
		log.Fatal(err)
	}

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

func (DB *DB)UpsertTypes(t string, icon string){
	if DB.Session == nil { panic("Session not instanciated") }
	coll := DB.Session.DB(dbstring).C("TypesIcons")
	_, err:= coll.Upsert(bson.M{"value" : t}, bson.M{"value" : t, "icon" : icon}) ; if err != nil{
		log.Fatal(err)
	} else {

	}
}

func (DB *DB)UpsertIcon(icon string, t string){
	if DB.Session == nil { panic("Session not instanciated") }
	coll := DB.Session.DB(dbstring).C("TypesIcons")
	_, err:= coll.Upsert(bson.M{"icon" : icon}, bson.M{"icon" : icon, "type" : t}) ; if err != nil{
		log.Fatal(err)
	} else {
		log.Println("Adding " + icon + " to type + " + t + " to Icons")
	}
}
func (DB *DB)UpsertMods(mod string, t string){
	if DB.Session == nil { panic("Session not instanciated") }
	coll := DB.Session.DB(dbstring).C("Mods")
	_, err:= coll.Upsert(bson.M{"mod" : mod}, bson.M{"mod" : mod, "type" : t}) ; if err != nil{
		log.Fatal(err)
	} else {
		log.Println("Adding " + mod + " to type + " + t + " to Mods")
	}
}

func (DB *DB)InitializeStringMap(c string) map[string]bool{
	if DB.Session == nil { panic("Session not instanciated") }
	coll := DB.Session.DB(dbstring).C(c)
	var d bson.D
	err := coll.Find(nil).All(&d) ; if err != nil{
		log.Fatal(err)
	}
	m := make(map[string]bool)
	for _, name := range d {
		m[name.Value.(string)] = true
	}
	log.Printf("Loaded map with %v %v\n", len(m), c)
	return m
}

func (DB *DB)InitializeIconsMap() map[string]string{
	if DB.Session == nil { panic("Session not instanciated") }
	coll := DB.Session.DB(dbstring).C("TypesIcons")
	var d []struct{
		ID     bson.ObjectId `bson:"_id,omitempty" json:"-"`
		Icon string `bson:"icon"`
		Type string `bson:"type"`
	}
	err := coll.Find(nil).All(&d) ; if err != nil{
		log.Fatal(err)
	}

	m := make(map[string]string)
	for _, entry := range d {

		m[entry.Icon] = entry.Type
	}
	log.Printf("Loaded map with %v %v\n", len(m), "TypesIcons")
	return m
}

func (DB *DB)InitializeModsMap() map[string]string{
	if DB.Session == nil { panic("Session not instanciated") }
	coll := DB.Session.DB(dbstring).C("Mods")
	var d []struct{
		ID     bson.ObjectId `bson:"_id,omitempty" json:"-"`
		Mod string `bson:"mod"`
		Type string `bson:"type"`
	}
	err := coll.Find(nil).All(&d) ; if err != nil{
		log.Fatal(err)
	}

	m := make(map[string]string)
	for _, entry := range d {

		m[entry.Mod] = entry.Type
	}
	log.Printf("Loaded map with %v %v\n", len(m), "Mods")
	return m
}


