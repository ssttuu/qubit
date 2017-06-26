package pgutils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/stupschwartz/qubit/applications/lib/stringutils"
)

// DeleteConfig is a configuration for deleting rows
type DeleteConfig struct {
	DB    *sqlx.DB
	Id    string
	Table string
	Tx    *sqlx.Tx
}

// InsertConfig is a configuration for inserting rows
type InsertConfig struct {
	Columns []string
	DB      *sqlx.DB
	Out     interface{}
	Values  [][]interface{}
	Table   string
	Tx      *sqlx.Tx
}

// SelectConfig is a configuration for selecting rows
type SelectConfig struct {
	Args          []interface{}
	Columns       []string
	DB            *sqlx.DB
	ForClause     string
	Id            string
	JoinClause    string
	Limit         int
	Offset        int
	OrderByClause string
	Out           interface{}
	Table         string
	Tx            *sqlx.Tx
	WhereClause   string
}

// UpdateConfig is a configuration for updating rows
type UpdateConfig struct {
	DB      *sqlx.DB
	Id      string
	Table   string
	Tx      *sqlx.Tx
	Updates map[string]interface{}
}

func getColumnString(columns []string) string {
	if len(columns) == 0 {
		return "*"
	}
	return strings.Join(columns, ", ")
}

// DeleteByID deletes a record by ID
func DeleteByID(deleteConfig *DeleteConfig) error {
	objId, err := strconv.ParseInt(deleteConfig.Id, 10, 64)
	if err != nil {
		return errors.Wrapf(err, "Could not convert %v to integer", deleteConfig.Id)
	}
	query := fmt.Sprintf("DELETE FROM %v WHERE id=$1", deleteConfig.Table)
	if deleteConfig.Tx == nil {
		_, err = deleteConfig.DB.Queryx(query, objId)
	} else {
		_, err = deleteConfig.Tx.Queryx(query, objId)
	}
	if err != nil {
		return errors.Wrapf(err, "Could not delete row from table %v with ID %v", deleteConfig.Table, objId)
	}
	return nil
}

// Insert inserts new records into a table
func Insert(insertConfig *InsertConfig) error {
	valuesMap := map[string]interface{}{}
	var params []string
	for i, rowData := range insertConfig.Values {
		rowParams := []string{}
		for j, value := range rowData {
			param := stringutils.Concat("param_", strconv.Itoa(i), "_", strconv.Itoa(j))
			valuesMap[param] = value
			rowParams = append(rowParams, stringutils.Concat(":", param))
		}
		params = append(params, stringutils.Concat("(", strings.Join(rowParams, ", "), ")"))
	}
	query := fmt.Sprintf(
		`INSERT INTO %v (%v) VALUES (%v) RETURNING id`,
		insertConfig.Table,
		strings.Join(insertConfig.Columns, ", "),
		strings.Join(params, ", "),
	)
	stmt, err := insertConfig.DB.PrepareNamed(query)
	if err != nil {
		return errors.Wrapf(err, "Failed to prepare statement, %s", query)
	}
	err = stmt.Select(insertConfig.Out, valuesMap)
	if err != nil {
		return errors.Wrapf(err, "Failed to create rows in %v", insertConfig.Table)
	}
	return nil
}

// InsertOne inserts a single row
func InsertOne(insertConfig *InsertConfig) error {
	if len(insertConfig.Values) != 1 {
		return errors.New("InsertOne expects only one array of items in InsertConfig.Values")
	}
	valuesMap := map[string]interface{}{}
	var params []string
	for i, value := range insertConfig.Values[0] {
		param := stringutils.Concat("param_", strconv.Itoa(i))
		valuesMap[param] = value
		params = append(params, stringutils.Concat(":", param))
	}
	query := fmt.Sprintf(
		`INSERT INTO %v (%v) VALUES (%v) RETURNING id`,
		insertConfig.Table,
		strings.Join(insertConfig.Columns, ", "),
		strings.Join(params, ", "),
	)
	stmt, err := insertConfig.DB.PrepareNamed(query)
	if err != nil {
		return errors.Wrapf(err, "Failed to prepare statement, %s", query)
	}
	err = stmt.Get(insertConfig.Out, valuesMap)
	if err != nil {
		return errors.Wrapf(err, "Failed to create row in %v", insertConfig.Table)
	}
	return nil
}

// Select selects records from a table
func Select(selectConfig *SelectConfig) error {
	columnString := getColumnString(selectConfig.Columns)
	var limitClause string
	if selectConfig.Limit != 0 {
		limitClause = fmt.Sprintf("LIMIT %v", selectConfig.Limit)
	}
	var offsetClause string
	if selectConfig.Offset != 0 {
		offsetClause = fmt.Sprintf("OFFSET %v", selectConfig.Offset)
	}
	query := fmt.Sprintf(
		`SELECT %v
		FROM %v
		%v
		%v
		%v
		%v
		%v
		%v`,
		columnString,
		selectConfig.Table,
		selectConfig.JoinClause,
		selectConfig.WhereClause,
		selectConfig.OrderByClause,
		limitClause,
		offsetClause,
		selectConfig.ForClause,
	)
	var err error
	if selectConfig.Tx == nil {
		err = selectConfig.DB.Select(selectConfig.Out, query, selectConfig.Args...)
	} else {
		err = selectConfig.Tx.Select(selectConfig.Out, query, selectConfig.Args...)
	}
	if err != nil {
		return errors.Wrapf(err, "Could not select from table %v", selectConfig.Table)
	}
	return nil
}

// SelectByID selects a record by ID
func SelectByID(selectConfig *SelectConfig) error {
	columnString := getColumnString(selectConfig.Columns)
	objId, err := strconv.ParseInt(selectConfig.Id, 10, 64)
	if err != nil {
		return errors.Wrapf(err, "Could not convert %v to integer", selectConfig.Id)
	}
	query := fmt.Sprintf("SELECT %v FROM %v WHERE id=$1 %v", columnString, selectConfig.Table, selectConfig.ForClause)
	if selectConfig.Tx == nil {
		err = selectConfig.DB.Get(selectConfig.Out, query, objId)
	} else {
		err = selectConfig.Tx.Get(selectConfig.Out, query, objId)
	}
	if err != nil {
		return errors.Wrapf(err, "Could not get row from table %v with ID %v", selectConfig.Table, objId)
	}
	return nil
}

// UpdateByID updates a record by ID
func UpdateByID(updateConfig *UpdateConfig) error {
	updateFields := []string{}
	updateArgs := []interface{}{}
	i := 1
	for column, value := range updateConfig.Updates {
		updateArgs = append(updateArgs, value)
		updateFields = append(updateFields, stringutils.Concat(column, "=$", strconv.Itoa(i)))
		i++
	}
	query := fmt.Sprintf(`UPDATE %v SET %v WHERE id=$%v`, updateConfig.Table, strings.Join(updateFields, ", "), i)
	updateArgs = append(updateArgs, updateConfig.Id)
	var err error
	if updateConfig.Tx == nil {
		_, err = updateConfig.DB.Exec(query, updateArgs...)
	} else {
		_, err = updateConfig.Tx.Exec(query, updateArgs...)
	}
	if err != nil {
		return errors.Wrapf(err, "Failed to update record on table %v with ID %v", updateConfig.Table, updateConfig.Id)
	}
	return nil
}
