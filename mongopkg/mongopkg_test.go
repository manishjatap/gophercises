package mongopkg

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2"
)

func TestGetSessionActualImplemenation(t *testing.T) {
	session, err := getSession()
	closeSession(session)
	assert.NoError(t, err, "No error required")
}

func TestInsertActualImplemenation(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			assert.Equal(t, err, "Session already closed", "Expected : Session already closed Error")
		}
	}()

	mockGetSession()

	dummyTask := Task{"fake-task", "fake-date", "Incomplete", "fake-date"}

	_ = dummyTask.Insert()

}

func TestGetActualImplemenation(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			assert.Equal(t, err, "Session already closed", "Expected : Session already closed Error")
		}
	}()

	mockGetSession()

	dummyTask := Task{"fake-task", "fake-date", "Incomplete", "fake-date"}

	_, _ = dummyTask.Get()

}

func TestDeleteActualImplemenation(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			assert.Equal(t, err, "Session already closed", "Expected : Session already closed Error")
		}
	}()

	mockGetSession()

	dummyTask := Task{"fake-task", "fake-date", "Incomplete", "fake-date"}

	_ = dummyTask.Delete()

}

func TestUpdateActualImplemenation(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			assert.Equal(t, err, "Session already closed", "Expected : Session already closed Error")
		}
	}()

	mockGetSession()

	dummyTask := Task{"fake-task", "fake-date", "Incomplete", "fake-date"}

	_ = dummyTask.Update()

}

func TestInsertSuccess(t *testing.T) {
	mockGetSession()
	mockInsertDoc()
	dummyTask := Task{"fake-task", "fake-date", "Incomplete", "fake-date"}

	err := dummyTask.Insert()

	assert.NoError(t, err, "Expected : No Error")
	resetInsertDoc()

}

func TestGetSuccess(t *testing.T) {
	mockGetSession()
	mockGetDoc()
	dummyTask := new(Task)

	_, err := dummyTask.Get()

	assert.NoError(t, err, "Expected : No Error")
	resetGetDoc()

}

func TestDeleteSuccess(t *testing.T) {
	mockGetSession()
	mockRemoveDoc()
	dummyTask := Task{"fake-task", "fake-date", "Incomplete", "fake-date"}

	err := dummyTask.Delete()

	assert.NoError(t, err, "Expected : No Error")
	resetRemoveDoc()

}

func TestUpdateSuccess(t *testing.T) {
	mockGetSession()
	mockUpdateDoc()
	dummyTask := Task{"fake-task", "fake-date", "Incomplete", "fake-date"}

	err := dummyTask.Update()

	assert.NoError(t, err, "Expected : No Error")
	resetUpdateDoc()

}

func TestInsertError(t *testing.T) {
	expectedErr := "Error while inserting docs"
	mockGetSession()
	errorInsertDoc(expectedErr)
	dummyTask := Task{"fake-task", "fake-date", "Incomplete", "fake-date"}

	err := dummyTask.Insert()

	assert.Equal(t, err.Error(), expectedErr, "Expected : Error")
	resetInsertDoc()

}

func TestGetError(t *testing.T) {
	expectedErr := "Error while fetching docs"
	mockGetSession()
	errorGetDoc(expectedErr)
	dummyTask := new(Task)

	_, err := dummyTask.Get()

	assert.Equal(t, err.Error(), expectedErr, "Expected : Error")
	resetGetDoc()

}

func TestDeleteError(t *testing.T) {
	expectedErr := "Error while deleting docs"
	mockGetSession()
	errorRemoveDoc(expectedErr)
	dummyTask := Task{"fake-task", "fake-date", "Incomplete", "fake-date"}

	err := dummyTask.Delete()

	assert.Equal(t, err.Error(), expectedErr, "Expected : Error")
	resetRemoveDoc()

}

func TestUpdateError(t *testing.T) {
	expectedErr := "Error while updating docs"
	mockGetSession()
	errorUpdateDoc(expectedErr)
	dummyTask := Task{"fake-task", "fake-date", "Incomplete", "fake-date"}

	err := dummyTask.Update()

	assert.Equal(t, err.Error(), expectedErr, "Expected : Error")
	resetUpdateDoc()

}

func TestInsertSessionError(t *testing.T) {
	expectedErr := "Error while fetching session"
	errorGetSession(expectedErr)
	dummyTask := Task{"fake-task", "fake-date", "Incomplete", "fake-date"}

	err := dummyTask.Insert()

	assert.Equal(t, err.Error(), expectedErr, "Expected : Error")

}

func TestGetSessionError(t *testing.T) {
	expectedErr := "Error while fetching session"
	errorGetSession(expectedErr)
	dummyTask := new(Task)

	_, err := dummyTask.Get()

	assert.Equal(t, err.Error(), expectedErr, "Expected : Error")

}

func TestDeleteSessionError(t *testing.T) {
	expectedErr := "Error while fetching session"
	errorGetSession(expectedErr)
	dummyTask := Task{"fake-task", "fake-date", "Incomplete", "fake-date"}

	err := dummyTask.Delete()

	assert.Equal(t, err.Error(), expectedErr, "Expected : Error")

}

func TestUpdateSessionError(t *testing.T) {
	expectedErr := "Error while fetching session"
	errorGetSession(expectedErr)
	dummyTask := Task{"fake-task", "fake-date", "Incomplete", "fake-date"}

	err := dummyTask.Update()

	assert.Equal(t, err.Error(), expectedErr, "Expected : Error")

}

func mockGetSession() {
	getSession = func() (*mgo.Session, error) {
		return new(mgo.Session), nil
	}
}

func errorGetSession(errMsg string) {
	getSession = func() (*mgo.Session, error) {
		return new(mgo.Session), errors.New(errMsg)
	}
}

func resetGetSession() {
	getSession = func() (*mgo.Session, error) {
		return mgo.Dial("mongodb://localhost:27017")
	}
}

func mockInsertDoc() {
	insertDoc = func(collection *mgo.Collection, t Task) error {
		return nil
	}
}

func errorInsertDoc(errMsg string) {
	insertDoc = func(collection *mgo.Collection, t Task) error {
		return errors.New(errMsg)
	}
}

func resetInsertDoc() {
	insertDoc = func(collection *mgo.Collection, t Task) error {
		return collection.Insert(t)
	}
}

func mockGetDoc() {
	getDoc = func(collection *mgo.Collection, filter interface{}, tlist *[]Task) error {
		return nil
	}
}

func errorGetDoc(errMsg string) {
	getDoc = func(collection *mgo.Collection, filter interface{}, tlist *[]Task) error {
		return errors.New(errMsg)
	}
}

func resetGetDoc() {
	getDoc = func(collection *mgo.Collection, filter interface{}, tlist *[]Task) error {
		iter := collection.Find(filter).Iter()
		return iter.All(tlist)
	}
}

func mockUpdateDoc() {
	updateDoc = func(collection *mgo.Collection, filter interface{}, update interface{}) error {
		return nil
	}
}

func errorUpdateDoc(errMsg string) {
	updateDoc = func(collection *mgo.Collection, filter interface{}, update interface{}) error {
		return errors.New(errMsg)
	}
}

func resetUpdateDoc() {
	updateDoc = func(collection *mgo.Collection, filter interface{}, update interface{}) error {
		return collection.Update(filter, update)
	}
}

func mockRemoveDoc() {
	removeDoc = func(collection *mgo.Collection, filter interface{}) error {
		return nil
	}
}

func errorRemoveDoc(errMsg string) {
	removeDoc = func(collection *mgo.Collection, filter interface{}) error {
		return errors.New(errMsg)
	}
}

func resetRemoveDoc() {
	removeDoc = func(collection *mgo.Collection, filter interface{}) error {
		return collection.Remove(filter)
	}
}
