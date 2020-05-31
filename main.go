package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	"github.com/candrea/backend/model"
	"github.com/gin-gonic/gin"
	//"gopkg.in/mgo.v2/bson"
	"github.com/candrea/backend/connection"
	"github.com/gin-contrib/cors"
)

var prefixPath = "/api/reciclaje"

func InsertUserController (c *gin.Context){
	
	var usuario model.Usuario
	err := c.BindJSON(&usuario)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	fmt.Println(usuario)
	//usuario.ID = bson.NewObjectId() 
	if err := connection.InsertUser(usuario);
	 err != nil {
	
		return
	}
	
}




func main() {
	r := gin.Default()

	r.Use(cors.Default())


	r.POST("/api/register", InsertUserController)
	r.Run()
	
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}


