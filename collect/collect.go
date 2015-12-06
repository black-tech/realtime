package collect

import (
	"github.com/black-tech/realtime/lib/db"
	"gopkg.in/mgo.v2"
	"time"
)

func checkExistingData() {
	db.Exce(db.MGODB_INIT, func(db *mgo.Database) error {

	})
}
