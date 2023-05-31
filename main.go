package main

import (
    "github.com/gin-gonic/gin"
    "github.com/ReedOlm/GoRestAPI/data"
    "github.com/ReedOlm/GoRestAPI/model"
    "github.com/ReedOlm/GoRestAPI/service"
    "net/http"
)

// GET
func getUsers(c *gin.Context){
    c.IndentedJSON(http.StatusOK, data.Users)
}

func getUserByID(c *gin.Context){
    id := c.Param("id")
    user, err, _ := service.FindUserByID(id)

    if err != nil {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found."})
        return
    }

    c.IndentedJSON(http.StatusOK, user)
}

func getFriendsOfUser(c *gin.Context){
    id := c.Param("id")
    user, err, _ := service.FindUserByID(id)

    if err != nil {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found."})
        return
    }

    friends := service.FindFriends(user.Friends)
    c.IndentedJSON(http.StatusOK, friends)
}

// POST
func addUser(c *gin.Context){
    var newUser model.User

    if err := c.BindJSON(&newUser); err != nil {
        return
    }

    data.Users = append(data.Users, newUser)
    c.IndentedJSON(http.StatusCreated, newUser)
}

// TODO Post /users/:id/friends
func addFriendToUser(c *gin.Context){

}

// PUT
func updateUser(c *gin.Context){
    id := c.Param("id")
    user, err, _ := service.FindUserByID(id)

    if err != nil {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found."})
        return
    }
    var updatedUser model.User
    if er := c.BindJSON(&updatedUser); er != nil{
        return
    }

    user.Name = updatedUser.Name
    user.Email = updatedUser.Email
    user.Age = updatedUser.Age
    user.Friends = updatedUser.Friends

    c.IndentedJSON(http.StatusOK, gin.H{"message": "Updated user information"})
}

// DELETE
func deleteUser(c *gin.Context){
    id := c.Param("id")
    _, err, index := service.FindUserByID(id)

    if err != nil {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found."})
        return
    }

    data.Users = append(data.Users[:index], data.Users[index+1:]...)
    c.IndentedJSON(http.StatusNoContent, gin.H{"message": "Success"})
}

// TODO Delete /users/:id/friends/:friendId
func deleteFriendFromUser(c *gin.Context){
}

func main(){
    router := gin.Default()
    router.GET("/users", getUsers)
    router.GET("/users/:id", getUserByID)
    router.GET("/users/:id/friends", getFriendsOfUser)
    router.POST("/users", addUser)
    router.POST("/users/:id/friends", addFriendToUser)
    router.PUT("/users/:id", updateUser)
    router.DELETE("/users/:id", deleteUser)
    router.DELETE("/users/:id/friends/:friendId", deleteFriendFromUser)
    router.Run("localhost:8080")
}
