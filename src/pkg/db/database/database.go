package database

type Config interface {
	GetDSN() string
}
type Driver string

const (
	PostgresDatabaseDriver Driver = "postgres"
	MysqlDatabaseDriver    Driver = "mysql"
)
