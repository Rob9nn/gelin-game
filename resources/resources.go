package resources

import (
	"embed"
	"io/fs"
)

//go:embed sql/*.sql
var sqlFiles embed.FS

func GetSqlFiles() ([]fs.DirEntry, error) {
	return sqlFiles.ReadDir("sql")
}
