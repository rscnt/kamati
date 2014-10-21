package kamatis

import (
	"github.com/go-martini/martini"
	"labix.org/v2/mgo"
)

/*
 * different database,
 * *mgo.Session and database name.
 */
type DataBaseSession struct {
	*mgo.Session
	databasename string
}

/**
 * Creates a new connection to the local mongodb instance.
 */
func NewSession(name string) *DataBaseSession {
	session, err := mgo.Dial("mongodb://127.0.0.1")
	if err != nil {
		panic(err)
	}
	//addIndexToDescriptionActivity(session.DB(name))
	return &DataBaseSession{session, name}
}

/**
 * Add unique index to destripction field on activity collection.
 */
func addIndexToDescriptionActivity(db *mgo.Database) {
	index := mgo.Index{
		Key:      []string{"description"},
		Unique:   true,
		DropDups: true,
	}
	indexErrors := db.C("activities").EnsureIndex(index)
	if indexErrors != nil {
		panic(indexErrors)
	}
}

/**
 * Martini lets inject parameters on routing handlers, by using context.Map().
 */
func (session *DataBaseSession) DataBase() martini.Handler {
	return func(context martini.Context) {
		s := session.Clone()
		context.Map(s.DB(session.databasename))
		defer s.Close()
		context.Next()
	}
}
