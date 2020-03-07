package mongodao

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/imdigo/DolimGoLangStudy/task"
)

// MongoAccessor is an Accessor for MongoDB
type MongoAccessor struct {
	session *mgo.Session
	collection *mgo.Collection
}

// New returns a new MongoAccessor.
func New(path, db, c string) task.Accessor {
	session, err := mgo.Dial(path)
	if err != nil {
		return nil
	}
	collection := session.DB(db).C(c)
	return &MongoAccessor {
		session: session,
		collection: collection,
	}
}

// Close closes MongoAccessor session
func (m *MongoAccessor) Close() error {
	m.session.Close()
	return nil
}


// idToObjectId returns bson.ObjectID converted from id.
func idToObjectId(id task.ID) bson.ObjectId {
	return bson.ObjectIdHex(string(id))
}

// objectIdToID returns task.Id converted from objID.
func objectIdToID(objID bson.ObjectId) task.ID {
	return task.ID(objID.Hex())
}

// Get returns a task with a given ID.
func (m *MongoAccessor) Get(id task.ID) (task.Task, error) {
	t := task.Task{}
	err := m.collection.FindId(idToObjectId(id)).One(&t)
	return t, err
}

// Put updates a task with a given ID with it.
func (m *MongoAccessor) Put(id task.ID, t task.Task) error {
	return m.collection.UpdateId(idToObjectId(id), t)
}

// Post addes a new task.
func (m *MongoAccessor) Post(t task.Task) (task.ID, error) {
	objID := bson.NewObjectId()
	_, err := m.collection.UpsertId(objID, &t)
	return objectIdToID(objID), err
}

// Delete removes the task with a given ID
func (m *MongoAccessor) Delete(id task.ID) error {
	return m.collection.RemoveId(idToObjectId(id))
}
