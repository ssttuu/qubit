package apiutils

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

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
	Columns []string
	DB      *sqlx.DB
	Id      string
	Out     interface{}
	Table   string
}

type ListConfig struct {
	Columns []string
	DB      *sqlx.DB
	Out     interface{}
	Table   string
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
	columns := []string{}
	values := []interface{}{}
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
	// TODO: Not found vs. unknown vs. whatever
	return pgutils.DeleteByID(&pgutils.DeleteConfig{
		DB:    deleteConfig.DB,
		Table: deleteConfig.Table,
		Id:    deleteConfig.Id,
	})
}

func Get(getConfig *GetConfig) error {
	// TODO: Permissions
	err := pgutils.SelectByID(&pgutils.SelectConfig{
		Columns: getConfig.Columns,
		DB:      getConfig.DB,
		Id:      getConfig.Id,
		Out:     getConfig.Out,
		Table:   getConfig.Table,
	})
	if err != nil {
		log.Println(err)
		return grpc.Errorf(codes.NotFound, "Not found")
	}
	return nil
}

func List(listConfig *ListConfig) error {
	// TODO: Permissions
	return pgutils.Select(&pgutils.SelectConfig{
		Columns: listConfig.Columns,
		DB:      listConfig.DB,
		Out:     listConfig.Out,
		Table:   listConfig.Table,
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
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return grpc.Errorf(codes.NotFound, "Not found")
	}
	err = updateConfig.OldObject.ValidateUpdate(updateConfig.NewObject)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return grpc.Errorf(codes.InvalidArgument, "Invalid argument")
	}
	err = pgutils.UpdateByID(&pgutils.UpdateConfig{
		Id:      updateConfig.Id,
		Table:   updateConfig.Table,
		Tx:      tx,
		Updates: updateConfig.NewObject.GetUpdateData(),
	})
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return grpc.Errorf(codes.Internal, "Internal error")
	}
	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return grpc.Errorf(codes.Internal, "Internal error")
	}
	return nil
}
