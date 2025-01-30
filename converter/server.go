package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmoteRequest struct {
	Link        string `json:"link"`
	Is2FrameGif bool   `json:"is_2_frame_gif"`
	DesiredName string `json:"desired_name"`
}

type Request struct {
	Links []EmoteRequest `json:"emotes"`
}
type Response struct {
}

func server() {
	r := gin.Default()
	r.POST("/api/emote", handleEmote)
	r.Run(":6999")
}
func handleEmote(c *gin.Context) {
	var request Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request format: " + err.Error(),
		})
		return
	}
	response := start_getting_emotes_or_sum_which_is_the_same_as_main_before_i_really_hate_naming_smh(request)
	fmt.Println(response)
	// fmt.Printf("%+v\n", request)
	c.JSON(http.StatusOK, gin.H{})
}
