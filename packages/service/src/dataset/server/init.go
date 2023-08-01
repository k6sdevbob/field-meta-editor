package main

import (
	"fmt"
)

func (a *api) InitPG() error {
	// Prepare the sqlmock for an INSERT operation
	tianic := []struct {
		passengerId int
		survived    int
		pclass      int
		name        string
		sex         string
		age         int
		ticket      string
		parch       int
	}{
		{1, 0, 3, "Braund, Mr. Owen Harris", "male", 22, "A/5 21171", 0},
		{2, 1, 1, "Cumings, Mrs. John Bradley (Florence Briggs Thayer)", "female", 38, "PC 17599", 0},
		{3, 0, 3, "Heikkinen, Miss. Laina", "female", 26, "STON/O2. 3101282", 0},
		{4, 1, 1, "Futrelle, Mrs. Jacques Heath (Lily May Peel)", "female", 35, "113803", 0},
		{5, 0, 3, "Allen, Mr. William Henry", "male", 35, "373450", 0},
		{6, 0, 1, "McCarthy, Mr. Timothy J", "female", 54, "17463", 0},
		{7, 0, 3, "Palsson, Master. Gosta Leonard", "male", 2, "34990", 1},
		{8, 1, 3, "Johnson, Mrs. Oscar W (Elisabeth Vilhelmina Berg)", "female", 27, "34990", 2},
		{9, 0, 1, "Nasser, Mrs. Nicholas (Adele Achem)", "female", 14, "237736", 0},
		{10, 1, 2, "Sandstrom, Miss. Marguerite Rut", "female", 4, "PP 9549", 0},
		{11, 0, 3, "Bonnell, Miss. Elizabeth", "female", 58, "113783", 5},
		{12, 1, 1, "Saundercock, Mr. William Henry", "male", 20, "A/5. 2151", 0},
	}

	// Create table
	_, err := a.db.Exec("CREATE TABLE tianic (PassengerId int, Survived int, Pclass int, Name text, Sex text, Age text, Ticket int, Parch int)")
	if err != nil {
		return fmt.Errorf("An error '%s' was not expected while creating table", err)
	}

	for _, row := range tianic {
		_, err := a.db.Exec("INSERT INTO tianic (PassengerId, Survived, Pclass, Name, Sex, Age, Ticket, Parch) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", row.passengerId, row.survived, row.pclass, row.name, row.sex, row.age, row.ticket, row.parch)
		if err != nil {
			return fmt.Errorf("An error '%s' was not expected while inserting a row", err)
		}
	}
	return nil
}

func (a *api) InitDataset() error {
	datasets := []*Dataset{
		{
			DataID: 1,
			Name:   "tianic",
			Type:   "table",
			Meta: []Meta{
				{
					Fid:          "col_0",
					Name:         "passengerId",
					Desc:         "",
					SemanticType: "quantitative",
					GeoRole:      "none",
					DataType:     "int",
					Forder:        0,
				},
				{
					Fid:          "col_1",
					Name:         "survived",
					Desc:         "",
					SemanticType: "quantitative",
					GeoRole:      "none",
					DataType:     "int",
					Forder:        1,
				},
				{
					Fid:          "col_2",
					Name:         "pclass",
					Desc:         "",
					SemanticType: "quantitative",
					GeoRole:      "none",
					DataType:     "int",
					Forder:        2,
				},
				{
					Fid:          "col_3",
					Name:         "name",
					Desc:         "",
					SemanticType: "quantitative",
					GeoRole:      "none",
					DataType:     "string",
					Forder:        3,
				},
				{
					Fid:          "col_4",
					Name:         "sex",
					Desc:         "",
					SemanticType: "quantitative",
					GeoRole:      "none",
					DataType:     "string",
					Forder:        4,
				},
				{
					Fid:          "col_5",
					Name:         "age",
					Desc:         "",
					SemanticType: "quantitative",
					GeoRole:      "none",
					DataType:     "int",
					Forder:        6,
				},
				{
					Fid:          "col_6",
					Name:         "ticket",
					Desc:         "",
					SemanticType: "quantitative",
					GeoRole:      "none",
					DataType:     "string",
					Forder:        6,
				},
				{
					Fid:          "col_7",
					Name:         "parch",
					Desc:         "",
					SemanticType: "quantitative",
					GeoRole:      "none",
					DataType:     "int",
					Forder:        7,
				},
			},
		},
	}

	// Create tables
	_, err := a.db.Exec("CREATE TABLE dataset (DataID int, Name text, Type text)")
	if err != nil {
		return fmt.Errorf("An error '%s' was not expected while creating dataset table", err)
	}
	_, err = a.db.Exec("CREATE TABLE meta (Fid text, Name text, Desc text, SemanticType text, GeoRole text, DataType text, Forder int, DatasetID int)")
	if err != nil {
		return fmt.Errorf("An error '%s' was not expected while creating meta table", err)
	}

	// Begin transaction
	tx, err := a.db.Begin()
	if err != nil {
		return fmt.Errorf("An error '%s' was not expected when starting a transaction", err)
	}

	for _, dataset := range datasets {
		res, err := tx.Exec("INSERT INTO dataset (DataID, Name, Type) VALUES (?, ?, ?)", dataset.DataID, dataset.Name, dataset.Type)
		if err != nil {
			return fmt.Errorf("An error '%s' was not expected while inserting a row into dataset", err)
		}
		dataID, err := res.LastInsertId()
		if err != nil {
			return fmt.Errorf("An error '%s' was not expected while getting the last inserted ID", err)
		}

		for _, meta := range dataset.Meta {
			_, err = tx.Exec("INSERT INTO meta (Fid, Name, Desc, SemanticType, GeoRole, DataType, Forder, DatasetID) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", meta.Fid, meta.Name, meta.Desc, meta.SemanticType, meta.GeoRole, meta.DataType, meta.Forder, dataID)
			if err != nil {
				return fmt.Errorf("An error '%s' was not expected while inserting a row into meta", err)
			}
		}
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("An error '%s' was not expected when committing the transaction", err)
	}

	return nil
}

func (a *api) InitDB() error {
	if err := a.InitPG(); err != nil {
		return err
	}
	if err := a.InitDataset(); err != nil {
		return err
	}
	return nil
}
