package route

import (
	"example.com/go-gin-blog-api/auth"
	"example.com/go-gin-blog-api/comment"
	"example.com/go-gin-blog-api/post"
	"example.com/go-gin-blog-api/reaction"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	post.RegisterPostRoutes(r)
	auth.RegisterAuthRoutes(r)
	comment.RegisterCommentRoutes(r)
	reaction.RegisterReactionRoutes(r)
}
