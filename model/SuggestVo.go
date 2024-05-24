package model

import (
	"Canteen/helper"
	"database/sql"
)

type Suggest struct {
	Id       int    `json:"id"`
	Category string `json:"category"`
	Location string `json:"location"`
	Details  string `json:"details"`
	User     string `json:"user"`
	Contract string `json:"contract"`
	Date     string `json:"date"`
}

func (s *Suggest) List() ([]*Suggest, error) {
	query := "SELECT " +
		"id, category, location, details, user, contract, date " +
		"FROM suggest " +
		"ORDER BY date DESC"
	mysqlEngine := helper.SqlContext
	rows, err := mysqlEngine.Query(query)
	if err != nil {
		return nil, err
	}
	data := make([]*Suggest, 0)
	for rows.Next() {
		obj := Suggest{}
		if err := rows.Scan(&obj.Id, &obj.Category, &obj.Location, &obj.Details, &obj.User, &obj.Contract, &obj.Date); err != nil {
			return nil, err
		} else {
			data = append(data, &obj)
		}
	}
	return data, nil
}

func (s *Suggest) Save() (sql.Result, error) {
	query := "INSERT " +
		"INTO suggest(category, location, details, user, contract, date) " +
		"VALUES (?, ?, ?, ?, ?, ?)"
	mysqlEngine := helper.SqlContext
	data, err := mysqlEngine.Exec(query, s.Category, s.Location, s.Details, s.User, s.Contract, s.Date)
	if err != nil {
		return nil, err
	}
	return data, nil
}
