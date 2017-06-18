package apiutils

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/stupschwartz/qubit/applications/api/lib/pgutils"
)

// APIModel is an interface for objects passed to apiutils functions
type APIModel interface {
	GetCreateData() map[string]interface{}
	GetUpdateData() map[string]interface{}
	ValidateCreate() error
	ValidateUpdate(existingObject interface{}) error
}

type CreateConfig struct {
	DB     *sqlx.DB
	Object APIModel
	Table  string
}

type DeleteConfig struct {
	DB    *sqlx.DB
	Id    string
	Table string
}

type GetConfig struct {
	DB    *sqlx.DB
	Id    string
	Out   interface{}
	Table string
}

type ListConfig struct {
	DB    *sqlx.DB
	Out   interface{}
	Table string
}

type UpdateConfig struct {
	DB        *sqlx.DB
	Id        string
	NewObject APIModel
	OldObject APIModel
	Table     string
}

func Create(createConfig *CreateConfig) error {
	// TODO: Permissions
	err := createConfig.Object.ValidateCreate()
	if err != nil {
		return err
	}
	// Using reflection to pull the db struct field annotations
	columns := []string{}
	values := []interface{}{}
	//reflectedPointer := reflect.ValueOf(createConfig.Object)
	//reflectedStruct := reflect.Indirect(reflectedPointer)
	//for i := 0; i < reflectedStruct.Type().NumField(); i++ {
	//	columnName := reflectedStruct.Type().Field(i).Tag.Get("db")
	//	if columnName != "id" {
	//		columns = append(columns, columnName)
	//		values = append(values, reflectedStruct.Field(i).Interface())
	//	}
	//}
	for column, value := range createConfig.Object.GetCreateData() {
		columns = append(columns, column)
		values = append(values, value)
	}
	insertConfig := pgutils.InsertConfig{
		Columns: columns,
		DB:      createConfig.DB,
		Out:     createConfig.Object,
		Table:   createConfig.Table,
		Values: [][]interface{}{
			values,
		},
	}
	return pgutils.InsertOne(&insertConfig)
}

func Delete(deleteConfig *DeleteConfig) error {
	// TODO: Permissions
	// TODO: Delete dependent entities with service calls
	return pgutils.DeleteByID(&pgutils.DeleteConfig{
		DB:    deleteConfig.DB,
		Table: deleteConfig.Table,
		Id:    deleteConfig.Id,
	})
}

func Get(getConfig *GetConfig) error {
	// TODO: Permissions
	return pgutils.SelectByID(&pgutils.SelectConfig{
		DB:    getConfig.DB,
		Id:    getConfig.Id,
		Out:   getConfig.Out,
		Table: getConfig.Table,
	})
}

func List(listConfig *ListConfig) error {
	// TODO: Permissions
	return pgutils.Select(&pgutils.SelectConfig{
		DB:    listConfig.DB,
		Out:   listConfig.Out,
		Table: listConfig.Table,
	})
}

func Update(updateConfig *UpdateConfig) error {
	// TODO: Permissions
	tx, err := updateConfig.DB.Beginx()
	if err != nil {
		return errors.Wrap(err, "Failed to begin transaction")
	}
	err = pgutils.SelectByID(&pgutils.SelectConfig{
		ForClause: "FOR UPDATE",
		Id:        updateConfig.Id,
		Out:       updateConfig.OldObject,
		Table:     updateConfig.Table,
		Tx:        tx,
	})
	// TODO: 404 if not found
	if err != nil {
		return err
	}
	err = updateConfig.OldObject.ValidateUpdate(updateConfig.NewObject)
	if err != nil {
		return err
	}
	//updates := map[string]interface{}{}
	//reflectedPointer := reflect.ValueOf(updateConfig.NewObject)
	//reflectedStruct := reflect.Indirect(reflectedPointer)
	//for i := 0; i < reflectedStruct.Type().NumField(); i++ {
	//	columnName := reflectedStruct.Type().Field(i).Tag.Get("db")
	//	if columnName != "id" {
	//		updates[columnName] = reflectedStruct.Field(i).Interface()
	//	}
	//}
	err = pgutils.UpdateByID(&pgutils.UpdateConfig{
		Id:      updateConfig.Id,
		Table:   updateConfig.Table,
		Tx:      tx,
		Updates: updateConfig.NewObject.GetUpdateData(),
	})
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "Failed to commit transaction")
	}
	return nil
}
