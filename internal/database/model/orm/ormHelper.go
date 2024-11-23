package orm

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/paulborgen/goLangArb/internal/database"
)

type Scannable interface {
	Scan(dest ...interface{}) error
}

func GetColumnNames(myStruct any) string {
	temp := GetColumnNamesAsArray(myStruct)
	return strings.Join(temp[:], ",")
}

func GetColumnNamesAsArray(myStruct any) []string {
	var columnNames []string

	st := reflect.TypeOf(myStruct)
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)

		if field.Tag.Get("postgres.Table") != "" {
			columnNames = append(columnNames, field.Tag.Get("postgres.Table"))
		}
	}

	return columnNames
}

func GetColumnNamesNoPrimaryKey(myStruct any, primaryKeyColumnName string) string {

	columnNames := GetColumnNames(myStruct)

	// Remove everything up until and including the first comma
	columnNames = columnNames[strings.Index(columnNames, ",")+1:]

	return columnNames
}

func CreateInsertStatement(myStruct any, tableName string, primaryKeyColumnName string) string {
	pairColumnNamesNoPrimaryKey := GetColumnNamesNoPrimaryKey(myStruct, primaryKeyColumnName)

	numberOfColumns := strings.Count(pairColumnNamesNoPrimaryKey, ",") + 1
	log.Print(numberOfColumns)

	valuesStatement := "VALUES ("
	for i := 1; i <= numberOfColumns; i++ {
		valuesStatement = valuesStatement + "$" + strconv.Itoa(i) + ","
	}

	valuesStatement = strings.TrimSuffix(valuesStatement, ",")

	valuesStatement = valuesStatement + ")"

	sqlStatementTemp := `
		INSERT INTO %s (%s)
		%s
		RETURNING %s`

	sqlStatement := fmt.Sprintf(sqlStatementTemp, tableName, pairColumnNamesNoPrimaryKey, valuesStatement, primaryKeyColumnName)

	return sqlStatement
}

func CreateUpdateStatement(myStruct any, tableName string, primaryKeyColumnName string, primaryKey int) string {
	pairColumnNamesNoPrimaryKey := GetColumnNamesNoPrimaryKey(myStruct, primaryKeyColumnName)

	numberOfColumns := strings.Count(pairColumnNamesNoPrimaryKey, ",") + 1
	log.Print(numberOfColumns)

	//updateStatement := " SET "
	//for i, columnName := range pairColumnNamesNoPrimaryKey {
	//	updateStatement = updateStatement + " " + columnName + " = $" + strconv.Itoa(i) + ","
	//}

	sqlStatementTemp := `
		UPDATE %s 
		%s
		WHERE %s = %s`

	sqlStatement := fmt.Sprintf(sqlStatementTemp, tableName, pairColumnNamesNoPrimaryKey, primaryKeyColumnName, primaryKey)

	return sqlStatement
}

func MyExists(tableName string, columnName string, value any) bool {
	db := database.GetDBConnection()

	var exists bool = false
	var count sql.NullInt64

	row := db.QueryRow("SELECT Count(1) FROM "+tableName+" WHERE "+columnName+"=$1", value)
	err := row.Scan(&count)

	if err != nil {
		panic("error")
	}

	if count.Valid {
		if int(count.Int64) == 1 {
			exists = true
		}
	}

	return exists

}
