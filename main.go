package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	wallet "BTC-tracer/wallet"
)

func main() {

	router := gin.Default()

	router.POST("/form_post", func(c *gin.Context) {
		fmt.Println("Analyzing " + c.PostForm("message"))
		resp, err := http.Get("https://blockchain.info/rawaddr/" + c.PostForm("message"))
		if err != nil {
			panic(err)
		} else {
			body, _ := ioutil.ReadAll(resp.Body)
			var data wallet.DResp
			json.Unmarshal(body, &data)
			fmt.Println("===================== TOTAL RECIEVED ===================")
			fmt.Println("Hash160: ", data.Hash160)
		}
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	router.Run(":8080")
}