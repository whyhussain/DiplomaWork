package hanlder

import (
	"DiplomaWork/internal/app/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
	"io"
	"log"
	"net/http"
)

func AddCategory(c echo.Context, conn *pgxpool.Pool) error {
	var category model.Category
	requestbody, err := io.ReadAll(c.Request().Body)
	if err != nil {
		fmt.Println("error reading request body: ", err)
		return err
	}
	if err := json.Unmarshal(requestbody, &category); err != nil {
		fmt.Errorf("error unmarshalling request body: %s", err)
	}

	_, err = conn.Query(context.Background(), "INSERT INTO category (name) VALUES ($1)", category.Name)
	if err != nil {
		fmt.Errorf("error inserting category: %s", err)
		return err
	}
	return c.JSON(http.StatusOK, category.Name)
}

func GetCategories(conn *pgxpool.Pool) ([]model.Category, error) {
	rows, err := conn.Query(context.Background(), "SELECT id, name FROM category")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Store the categories in a slice
	var categories []model.Category
	for rows.Next() {
		var category model.Category
		if err := rows.Scan(&category.Id, &category.Name); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func GetCategory(c echo.Context, conn *pgxpool.Pool) (model.Category, error) {
	id := c.Param("id")
	var category model.Category
	log.Println("getCategory called with id:", id)
	rows, err := conn.Query(context.Background(), "SELECT * FROM category WHERE id = $1", id)
	if err != nil {
		log.Println("error running query:", err)
		return model.Category{}, err
	}
	defer rows.Close()

	if !rows.Next() {
		if rows.Err() != nil {
			log.Println("error fetching row:", rows.Err())
			return model.Category{}, rows.Err()
		}
		log.Println("no rows returned")
		return model.Category{}, fmt.Errorf("no category found with id %s", id)
	}

	if err := rows.Scan(&category.Id, &category.Name); err != nil {
		log.Println("error scanning row:", err)
		return model.Category{}, err
	}

	return category, nil
}

func DeleteCategory(c echo.Context, conn *pgxpool.Pool) (string, error) {
	id := c.Param("id")
	log.Println("deleteCategory called with id:", id)
	_, err := conn.Query(context.Background(), "DELETE FROM category WHERE id = $1", id)
	if err != nil {
		log.Println("error running query:", err)
		return id, err
	}

	return id, nil
}

func UpdateCategory(c echo.Context, conn *pgxpool.Pool) (model.Category, error) {
	id := c.Param("id")
	ct := model.Category{}
	requestbody, err := io.ReadAll(c.Request().Body)
	if err != nil {
		fmt.Println("error reading request body: ", err)
		return model.Category{}, err
	}
	if err := json.Unmarshal(requestbody, &ct); err != nil {
		fmt.Errorf("error unmarshalling request body: %s", err)
	}

	_, err = conn.Query(context.Background(), "UPDATE category SET name = $1 WHERE id = $2", ct.Name, id)
	if err != nil {
		fmt.Errorf("error updating category: %s", err)
		return model.Category{}, err
	}
	updatedCategory, err := GetCategory(c, conn)
	if err != nil {
		return model.Category{}, err
	}

	return updatedCategory, nil
}
