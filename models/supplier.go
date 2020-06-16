package models

type Supplier struct {
	ID      int64  `form:"id" json:"id"`
	Code    string `form:"code" json:"code"`
	Name    string `form:"name" json:"name"`
	Address string `form:"address" json:"address"`
	Status  string `form:"status" json:"status"`
}

func (db *DB) AllSuppliers() ([]*Supplier, error) {
	rows, err := db.Query("SELECT id, code, name, address, status FROM suppliers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	supps := make([]*Supplier, 0)
	for rows.Next() {
		sup := new(Supplier)
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
