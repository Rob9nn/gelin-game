package resources

import (
	"embed"
	"io/fs"
)

//go:embed sql/*.sql
var sqlFiles embed.FS

func ReadSqlFile(file string) ([]byte, error) {
	return sqlFiles.ReadFile("sql/" + file)
}
func GetSqlFiles() ([]fs.DirEntry, error) {
	return sqlFiles.ReadDir("sql")
}
