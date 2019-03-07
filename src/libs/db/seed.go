package db

import (
	"bufio"
	"database/sql"
	"io"
	"os"
	"strings"
)

func SeedFile(path string, db *sql.DB) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	f := bufio.NewReader(file)
	var line string
	for read_line, err := f.ReadString('\n'); err != io.EOF; read_line, err = f.ReadString('\n') {
		if strings.HasPrefix(read_line, `--`) {
			continue
		}
		line += strings.TrimSuffix(read_line, "\n")
		if strings.HasSuffix(line, ";") {
			db.Exec(strings.TrimRight(line, ";\n\t "))
			if err != nil {
				return err
			}
			line = ``
		}
	}
	return nil
}
