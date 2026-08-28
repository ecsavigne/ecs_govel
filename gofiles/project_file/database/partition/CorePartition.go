package partition

import (
	"ecs_govel/pkg/pkglog"
	"reflect"

	"gorm.io/gorm"
)

type corePartition struct {
	db                     *gorm.DB
	ifViewImportMigrations bool
}

// create new corePartition
func New(db *gorm.DB) *corePartition {
	return &corePartition{db: db}
}

func (self corePartition) Run() error {
	typ := reflect.TypeFor[*corePartition]() // information of struct
	value := reflect.ValueOf(&self)          // value of struct

	for method := range typ.Methods() {
		methodValue := value.MethodByName(method.Name)

		if method.Name == "Run" {
			continue
		}

		if method.Func.IsValid() {
			pkglog.Log.Infof("executing partition Method: %+v\n", method.Name)
			sliceValue := methodValue.Call([]reflect.Value{})

			if len(sliceValue) > 0 {
				v := sliceValue[0].Interface()
				if v != nil {
					if err, ok := v.(error); ok {
						pkglog.Log.Errorf("executed partition Method: %v\n", err.Error())
					}
				}
			}
			pkglog.Log.Infof("executed partition Method: %+v\n", method.Name)
		}
	}

	return nil
}
