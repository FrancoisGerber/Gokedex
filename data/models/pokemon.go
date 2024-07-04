package models

import (
	"Gokedex/data"
)

// Implements IModel
type Pokemon struct {
	Id        int64
	Name      string
	Category  string
	Height    float64
	Weight    float64
	Type      string
	Weakness  string
	Abilities string
}

func (u *Pokemon) GetAll() (pokemons []Pokemon, err error) {
	selectAllQuery := `SELECT Id, Name, Category, Height, Weight, Type, Weakness, Abilities FROM Pokemon`
	rows, err := data.DB.Query(selectAllQuery)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var pokemon Pokemon
		err := rows.Scan(&pokemon.Id, &pokemon.Name, &pokemon.Category, &pokemon.Height, &pokemon.Weight, &pokemon.Type, &pokemon.Weakness, &pokemon.Abilities)

		if err != nil {
			return nil, err
		}

		pokemons = append(pokemons, pokemon)
	}

	return pokemons, nil
}

func (pokemon *Pokemon) Get(id int64) error {
	selectSingleQuery := `SELECT Id, Name, Category, Height, Weight, Type, Weakness, Abilities FROM Pokemon WHERE Id = ?`
	row := data.DB.QueryRow(selectSingleQuery, id)

	err := row.Scan(&pokemon.Id, &pokemon.Name, &pokemon.Category, &pokemon.Height, &pokemon.Weight, &pokemon.Type, &pokemon.Weakness, &pokemon.Abilities)

	if err != nil {
		return err
	}

	return nil
}

func (pokemon *Pokemon) Create() error {
	insertQuery := `
	INSERT INTO Pokemon(Name, Category, Height, Weight, Type, Weakness, Abilities)
	VALUES(?,?,?,?,?,?,?)
	`
	result, err := data.DB.Exec(insertQuery, &pokemon.Name, &pokemon.Category, &pokemon.Height, &pokemon.Weight, &pokemon.Type, &pokemon.Weakness, &pokemon.Abilities)

	if err != nil {
		return err
	}

	pokemon.Id, err = result.LastInsertId()

	if err != nil {
		return err
	}

	return nil
}

func (pokemon *Pokemon) Update(id int64) (int64, error) {
	updateQuery := `
	UPDATE Pokemon
	SET Name=?, Category=?, Height=?, Weight=?, Type=?, Weakness=?, Abilities=?
	WHERE Id = ?`

	result, err := data.DB.Exec(updateQuery, &pokemon.Name, &pokemon.Category, &pokemon.Height, &pokemon.Weight, &pokemon.Type, &pokemon.Weakness, &pokemon.Abilities, id)

	if err != nil {
		return 0, err
	}

	rowsChanged, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsChanged, nil
}

func (u *Pokemon) Delete(id int64) (bool, error) {
	deleteQuery := `
	DELETE FROM Pokemon
	WHERE Id = ?`

	result, err := data.DB.Exec(deleteQuery, id)

	if err != nil {
		return false, err
	}

	rowsChanged, err := result.RowsAffected()

	if err != nil {
		return false, err
	}

	return rowsChanged > 0, nil
}
