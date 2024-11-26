package handlers

import (
	"menu-search/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Menu struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Category    string  `json:"category"`
	Ingredients string  `json:"ingredients"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

func GetMenus(c *gin.Context) {
	query := c.Query("query")
	menus := []Menu{}

	rows, err := db.DB.Query("SELECT * FROM menus WHERE name ILIKE $1 OR category ILIKE $1", "%"+query+"%")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var menu Menu
		rows.Scan(&menu.ID, &menu.Name, &menu.Category, &menu.Ingredients, &menu.Price, &menu.Description)
		menus = append(menus, menu)
	}
	c.JSON(http.StatusOK, menus)
}
