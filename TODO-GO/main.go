package main

import{
	"net/http"
	"github.com/gin-gonic/gin"
}




type todo struct {
	ID			string 	`json:"id"`
	Item		string	`json:"title"`
	Completed	bool 	`json:"completed"`	
}


var todos = []{
{ID: "1" , Item: "clean Book" , Completed:false},
{ID: "2" , Item: "Study" , Completed:false},
{ID: "3" , Item: "Die" , Completed:false},
{ID: "4" , Item: "MAMA" , Completed:false},	
{ID: "5" , Item: "Just killed a man" , Completed:false},
}

func getTodos(context){
	
}



func main()  {
	router := gin.Default()
	router.GET("/todos")
	router.Run("localhost:9090")	
}