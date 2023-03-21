package routes

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"go_nosql/app/databases"
	"go_nosql/app/model"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var DB = databases.ConnectDB()
	var db = databases.GetCollection(DB, "Routes")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	route := new(model.Create)
	defer cancel()

	if err := c.BindJSON(&route); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		log.Fatal(err)
		return
	}

	if route.TotalRoutes <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": errors.New("invalid total routes")})
		return
	}

	routes := make([]interface{}, 0)
	for i := 0; i <= route.TotalRoutes; i++ {
		payload := model.Route{
			TravelID: rand.Int(),
			SiteID:   createRandomString(),
			Status:   createRandomString(),
			CreateAt: time.Now().UTC(),
		}
		routes = append(routes, payload)

	}
	result, err := db.InsertMany(ctx, routes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusCreated,
		gin.H{
			"message": "routes created successfully",
			"data":    result.InsertedIDs},
	)
}

func createRandomUInt() uint64 {
	return rand.Uint64()
}

func createRandomString() string {
	return strconv.FormatUint(createRandomUInt(), 10)
}
