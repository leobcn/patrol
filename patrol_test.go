package main

import (
	"log"
	"sabey.co/unittest"
	"testing"
)

func TestPatrol(t *testing.T) {
	log.Println("TestPatrol")

	patrol := &Patrol{}

	unittest.Equals(t, patrol.validate(), ERR_APPS_EMPTY)
	// Apps must be initialized, the creation of the patrol object will not do this for you
	patrol.Apps = make(map[string]*PatrolApp)
	unittest.Equals(t, patrol.validate(), ERR_APPS_EMPTY)

	patrol.Apps[""] = &PatrolApp{}
	// check that key exists
	_, exists := patrol.Apps[""]
	unittest.Equals(t, exists, true)
	unittest.Equals(t, patrol.validate(), ERR_APP_KEY_EMPTY)

	// delete empty key
	delete(patrol.Apps, "")
	_, exists = patrol.Apps[""]
	unittest.Equals(t, exists, false)

	app := &PatrolApp{
		// empty object
	}
	patrol.Apps["http"] = app

	unittest.Equals(t, patrol.validate(), ERR_APP_NAME_EMPTY)
	app.Name = "name"

	unittest.Equals(t, patrol.validate(), ERR_APP_DIRECTORY_EMPTY)
	app.Directory = "directory"

	unittest.Equals(t, patrol.validate(), ERR_APP_FILE_EMPTY)
	app.File = "file"

	unittest.Equals(t, patrol.validate(), ERR_APP_LOG_DIRECTORY_EMPTY)
	app.LogDirectory = "log-directory"

	unittest.Equals(t, patrol.validate(), ERR_APP_PID_EMPTY)
	app.PID = "pid"

	unittest.IsNil(t, app.validate())
}
func TestPatrolApp(t *testing.T) {
	log.Println("TestPatrolApp")

	app := &PatrolApp{}
	unittest.Equals(t, app.validate(), ERR_APP_NAME_EMPTY)
	app.Name = "name"

	unittest.Equals(t, app.validate(), ERR_APP_DIRECTORY_EMPTY)
	app.Directory = "directory"

	unittest.Equals(t, app.validate(), ERR_APP_FILE_EMPTY)
	app.File = "file"

	unittest.Equals(t, app.validate(), ERR_APP_LOG_DIRECTORY_EMPTY)
	app.LogDirectory = "log-directory"

	unittest.Equals(t, app.validate(), ERR_APP_PID_EMPTY)
	app.PID = "pid"

	unittest.IsNil(t, app.validate())
}