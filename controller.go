package main

func createVideogame(videogame VideoGame) error {
	bd, err := getDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("INSERT INTO videogames (name,genre,year) values(?,?,?)", videogame.Name, videogame.Genre, videogame.Year)
	return err
}

func deleteVideogame(id int64) error {
	bd, err := getDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("DELETE FROM videogames WHERE id = ?", id)
	return err
}

func updateVideogame(videogame VideoGame) error {
	bd, err := getDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("UPDATE videogames SET name = ?, genre = ?, year = ? WHERE id = ?", videogame.Name, videogame.Genre, videogame.Year, videogame.Id)
	return err
}

func getVideogames() ([]VideoGame, error) {
	videogames := []VideoGame{}
	bd, err := getDB()
	if err != nil {
		return videogames, err
	}
	rows, err := bd.Query("SELECT id, name, genre, year FROM videogames")
	if err != nil {
		return videogames, err
	}
	for rows.Next() {
		var videogame VideoGame
		err = rows.Scan(&videogame.Id, &videogame.Name, &videogame.Genre, &videogame.Year)
		if err != nil {
			return videogames, err
		}
		videogames = append(videogames, videogame)
	}
	return videogames, nil
}

func getVideogameById(id int64) (VideoGame, error) {
	var videogame VideoGame
	bd, err := getDB()
	if err != nil {
		return videogame, err
	}
	row := bd.QueryRow("SELECT id, name, genre, year FROM videogames WHERE id = ?", id)
	err = row.Scan(&videogame.Id, &videogame.Name, &videogame.Genre, &videogame.Year)
	if err != nil {
		return videogame, err
	}
	return videogame, nil
}
