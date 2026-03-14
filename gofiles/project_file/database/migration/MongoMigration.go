package migration

import (
	"time"

	"github.com/kamva/mgm/v3"
)

type TestMongoMigration struct {
	mgm.DefaultModel `bson:",inline"`
	ID               string    `bson:"id"`
	CreatedAt        time.Time `bson:"created_at"`
}
