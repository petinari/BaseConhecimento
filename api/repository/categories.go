package repository

import (
	"api/models"
	"database/sql"
)

type CategoriesRepository struct {
	db *sql.DB
}

func NewCategoriesRepository(db *sql.DB) *CategoriesRepository {
	return &CategoriesRepository{db}
}

func (c *CategoriesRepository) SaveCategories(category models.Categories) (uint64, error) {
	statement, err := c.db.Prepare("INSERT INTO categories(name) VALUES($1)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()
	result, err := statement.Exec(category.Name)
	if err != nil {
		return 0, err
	}
	lastinsert, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return uint64(lastinsert), nil
}

func (c *CategoriesRepository) GetAllCategories() ([]models.Categories, error) {
	rows, err := c.db.Query("select id, name from categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var categories []models.Categories
	for rows.Next() {
		var category models.Categories
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}
