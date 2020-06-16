package repository

import (
	"database/sql"

	"github.com/arham09/conn-db/supplier/models"
)

func AllSuppliers(db *sql.DB) ([]*models.Supplier, error) {
	rows, err := db.Query("SELECT id, code, name, address, status FROM suppliers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	supps := make([]*models.Supplier, 0)
	for rows.Next() {
		sup := new(models.Supplier)
		err := rows.Scan(&sup.ID, &sup.Code, &sup.Name, &sup.Address, &sup.Status)
		if err != nil {
			return nil, err
		}
		supps = append(supps, sup)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return supps, nil
}
