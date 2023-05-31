package main

import (
    "github.com/gin-gonic/gin"
    "github.com/ReedOlm/GoRestAPI/resource"
)

func main(){
    router := gin.Default()
    router.GET("/users", resource.GetUsers)
    router.GET("/users/:id", resource.GetUserByID)
    router.GET("/users/:id/friends", resource.GetFriendsOfUser)
    router.POST("/users", resource.AddUser)
    router.POST("/users/:id/friends", resource.AddFriendToUser)
    router.PUT("/users/:id", resource.UpdateUser)
    router.DELETE("/users/:id", resource.DeleteUser)
    router.DELETE("/users/:id/friends/:friendId", resource.DeleteFriendFromUser)
    router.Run("localhost:8080")
}
