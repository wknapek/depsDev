package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	_ "modernc.org/sqlite"
)

var ErrNoFound = errors.New("not found")

type DBHandler struct {
	db *sql.DB
}

func (han *DBHandler) InitDB() error {
	var err error
	han.db, err = sql.Open("sqlite", "./depsDev.db") // Open a connection to the SQLite database file named app.db
	if err != nil {
		return err
	}
	_, err = han.db.Exec("CREATE TABLE IF NOT EXISTS depsDev(packageName TEXT PRIMARY KEY, json TEXT, packages TEXT);")
	if err != nil {
		return err
	}

	return nil
}

func (han *DBHandler) Close() {
	err := han.db.Close()
	if err != nil {
		return
	}
}

func (han *DBHandler) Insert(packageName string, packagesValue Check) error {
	packages, err := han.GetPackages(packageName)
	if err != nil {
		return err
	}
	packages = append(packages, packagesValue)
	stmt, err := han.db.Prepare("UPDATE depsDev SET packages = ? WHERE packageName = ?;")
	dataByte, err := json.Marshal(packages)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(string(dataByte), packageName)
	if err != nil {
		return err
	}
	return nil
}

func (han *DBHandler) InitPackage(packageName string, json string, packagesValue string) error {
	stmt, err := han.db.Prepare("INSERT INTO depsDev(packageName, json, packages) VALUES (?, ?, ?);")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(packageName, json, packagesValue)
	if err != nil {
		return err
	}
	return nil
}

func (han *DBHandler) GetPackages(name string) ([]Check, error) {
	rows := han.db.QueryRow("SELECT * FROM depsDev WHERE packageName = ?", name)
	row := Row{}
	err := rows.Scan(&row.Name, &row.Json, &row.Packages)
	if err != nil {
		return nil, err
	}
	var ret []Check
	err = json.Unmarshal([]byte(row.Packages), &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (han *DBHandler) DeletePackage(name string) error {
	packages, err := han.GetPackages("github.com/cli/cli")
	if err != nil {
		return err
	}
	for idx, pkg := range packages {
		if pkg.Name == name {
			stmt, err := han.db.Prepare("UPDATE depsDev SET packages = ? WHERE packageName = ?;")
			if err != nil {
				return err
			}
			packages = remove(packages, idx)
			dataByte, err := json.Marshal(packages)
			if err != nil {
				return err
			}
			_, err = stmt.Exec(string(dataByte), "github.com/cli/cli")
			if err != nil {
				return err
			}
			return nil
		}
	}
	return ErrNoFound
}

func remove(s []Check, i int) []Check {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
