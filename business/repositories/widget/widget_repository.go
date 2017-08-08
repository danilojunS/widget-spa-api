package widgetRepo

import (
	"database/sql"
	widget "github.com/danilojunS/widgets-spa-api/business/entities/widget"
	database "github.com/danilojunS/widgets-spa-api/infra/database"
	_ "github.com/lib/pq"
)

// Create a widget
func Create(w widget.Widget) (int, error) {
	db, err := database.Connect()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	var lastInsertID int
	err = db.QueryRow("INSERT INTO widgets (name, color, price, inventory, melts) VALUES($1,$2,$3,$4,$5) returning id;", w.Name, w.Color, w.Price, w.Inventory, w.Melts).Scan(&lastInsertID)
	if err != nil {
		return 0, err
	}

	return lastInsertID, nil
}

// Read widgets
func Read(id int) ([]widget.Widget, error) {
	db, err := database.Connect()
	if err != nil {
		return []widget.Widget{}, err
	}
	defer db.Close()

	var rows *sql.Rows
	if id == 0 {
		rows, err = db.Query("SELECT * FROM widgets")
		if err != nil {
			return []widget.Widget{}, err
		}
	} else {
		rows, err = db.Query("SELECT * FROM widgets WHERE id=$1", id)
		if err != nil {
			return []widget.Widget{}, err
		}
	}

	var widgets []widget.Widget

	for rows.Next() {
		var id int
		var name string
		var color string
		var price string
		var inventory int
		var melts bool

		err = rows.Scan(&id, &name, &color, &price, &inventory, &melts)
		if err != nil {
			return []widget.Widget{}, err
		}

		widget := widget.Widget{
			ID:        id,
			Name:      name,
			Color:     color,
			Price:     price,
			Inventory: inventory,
			Melts:     melts,
		}
		widgets = append(widgets, widget)
	}

	return widgets, nil
}

// Update widget
func Update(w widget.Widget) (int, error) {
	db, err := database.Connect()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	stmt, err := db.Prepare("UPDATE widgets SET name=$1, color=$2, price=$3, inventory=$4, melts=$5 WHERE id=$6")
	if err != nil {
		return 0, err
	}

	_, err = stmt.Exec(w.Name, w.Color, w.Price, w.Inventory, w.Melts, w.ID)
	if err != nil {
		return 0, err
	}

	return w.ID, nil
}

// Delete widget
func Delete(w widget.Widget) (int, error) {
	db, err := database.Connect()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM widgets WHERE id=$1")
	if err != nil {
		return 0, err
	}

	_, err = stmt.Exec(w.ID)
	if err != nil {
		return 0, err
	}

	return w.ID, nil
}

// Clear all widgets
func Clear() error {
	db, err := database.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM widgets")
	if err != nil {
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	return nil
}
