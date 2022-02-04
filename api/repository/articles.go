package repository

import (
	"api/models"
	"database/sql"
)

type ArticlesRepository struct {
	db *sql.DB
}

func NewArticlesRepository(db *sql.DB) *ArticlesRepository {
	return &ArticlesRepository{db}
}

func (c *ArticlesRepository) SaveArticles(articles models.Articles) (uint64, error) {

	statement, err := c.db.Prepare("INSERT INTO article(name, description, image_url, content, user_id, categorie_id ) VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()
	result, err := statement.Exec(articles.Name, articles.Description, articles.ImageUrl, articles.Content, articles.UserId, articles.CategorieId)
	if err != nil {
		return 0, err
	}
	lastinsert, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return uint64(lastinsert), nil
}

func (c *ArticlesRepository) GetAllArticles() ([]models.Articles, error) {
	rows, err := c.db.Query("SELECT id, name, description, image_url, content, user_id, categorie_id FROM public.article")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var articles []models.Articles
	for rows.Next() {
		var article models.Articles
		if err := rows.Scan(&article.ID, &article.Name, &article.Description, &article.ImageUrl, &article.Content, &article.UserId, &article.CategorieId); err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}
	return articles, nil
}

func (c *ArticlesRepository) GetArticlesById(id uint64) (models.Articles, error) {
	rows, err := c.db.Query("SELECT id, name, description, image_url, content, user_id, categorie_id FROM public.article where id = $1", id)
	if err != nil {
		return models.Articles{}, err
	}
	defer rows.Close()
	var article models.Articles
	for rows.Next() {
		if err := rows.Scan(&article.ID, &article.Name, &article.Description, &article.ImageUrl, &article.Content, &article.UserId, &article.CategorieId); err != nil {
			return models.Articles{}, err
		}

	}
	return article, nil
}

func (c *ArticlesRepository) GetAllArticlesByCategory(category string) ([]models.Articles, error) {
	rows, err := c.db.Query("select id from categories where name = $1", category)
	if err != nil {
		return []models.Articles{}, err
	}
	defer rows.Close()
	var idCategorie uint64
	for rows.Next() {
		if err := rows.Scan(&idCategorie); err != nil {
			return []models.Articles{}, err
		}
	}
	rows, err = c.db.Query("select * from article where categorie_id = $1", idCategorie)
	if err != nil {
		return []models.Articles{}, err
	}
	var articles []models.Articles
	for rows.Next() {
		var article models.Articles
		if err := rows.Scan(&article.ID, &article.Name, &article.Description, &article.ImageUrl, &article.Content, &article.UserId, &article.CategorieId); err != nil {
			return []models.Articles{}, err
		}
		articles = append(articles, article)
	}
	return articles, nil
}
