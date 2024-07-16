package main

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
) 

func main()  {
	router:=gin.New()
	router.GET("/getData",getData)
	router.POST("/getPostData",getPostData)
	router.GET("/getQueryString",getQueryString)
	router.GET("/getUrlData/:name/:age",getUrlData)

	router.Run()

	
}
// hrrp://localhost:8080/getUrlData?name=Mark&age=30
func getUrlData(ctx *gin.Context) {
	name:=ctx.Param(("name"))
	age:=ctx.Param("age")
	ctx.JSON(200,gin.H{
		"data":"Hi I am gin framework url data",
		"name":name,
		"age":age,
	})
	
 }
 
// hrrp://localhost:8080/getQueryString?name=Mark&age=30
func getQueryString(ctx *gin.Context) {
	name:=ctx.Query(("name"))
	age:=ctx.Query("age")
	ctx.JSON(200,gin.H{
		"data":"Hi I am gin framework query string",
		"name":name,
		"age":age,
	})
	
 }
func getData(ctx *gin.Context) {
	ctx.JSON(200,gin.H{
		"data":"Hi I am gin framework get method", 
	})
	
 }
 func getPostData(ctx *gin.Context) {
	body:=ctx.Request.Body
	value,_:=ioutil.ReadAll(body)
	ctx.JSON(200,gin.H{
		"data":"I am post method",
		"bodyData":string(value),
	})
 }