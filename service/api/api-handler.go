package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.GET("/", rt.getHelloWorld)
	rt.router.GET("/context", rt.wrap(rt.getContextReply))


	rt.router.POST("/session", rt.getAuthToken)

	rt.router.GET("/users", rt.searchUser)
	
	rt.router.GET("/users/:userId", rt.getUserProfile)
	rt.router.PUT("/users/:userId", rt.updateUserProfile)
	rt.router.DELETE("/users/:userId", rt.deleteUserProfile)

	rt.router.GET("/users/:userId/posts", rt.getUserPosts)
	rt.router.POST("/users/:userId/posts", rt.createPost)

	rt.router.GET("/users/:userId/posts/postId", rt.getPost)
	rt.router.PUT("/users/:userId/posts/postId", rt.editPost)
	rt.router.DELETE("/users/:userId/posts/postId", rt.deletePost)

	rt.router.GET("/users/:userId/posts/postId/likes", rt.getPostLikes)

	rt.router.PUT("/users/:userId/posts/:postId/likes/:likeId", rt.likePost)
	rt.router.DELETE("/users/:userId/posts/:postId/likes/:likeId", rt.unlikePost)

	rt.router.GET("/users/:userId/posts/postId/comments", rt.getPostComments)
	rt.router.POST("/users/:userId/posts/postId/comments", rt.createComment)

	rt.router.GET("/users/:userId/posts/postId/comments/:commentId", rt.getComment)
	rt.router.PUT("/users/:userId/posts/postId/comments/commentId",rt.editComment)
	rt.router.DELETE("/users/:userId/posts/postId/comments/:commentId",rt.deleteComment)

	rt.router.GET("/users/:userId/posts/postId/comments/:commentId/likes", rt.getCommentLikes)

	rt.router.PUT("/users/:userId/posts/postId/comments/:commentId/likes/likeId", rt.likeComment)
	rt.router.DELETE("/users/:userId/posts/postId/comments/:commentId/likes/likeId", rt.unlikeComment)

	rt.router.GET("/users/:userId/follower", rt.getFollowersList)
	
	rt.router.GET("/user/:userId/following", rt.getFollowingsList)
	
	rt.router.PUT("/users/:userId/following/:followingId", rt.followUser)
	rt.router.DELETE("/users/:userId/following/followingId", rt.unfollowUser)

	rt.router.GET("/users/:userId/banned", rt.getUserBanList)

	rt.router.PUT("/users/:userId/banned/:bannedId", rt.banUser)
	rt.router.DELETE("/users/:userId/banned/:bannedId", rt.unbanUser)

	rt.router.POST("/users/:userId/photos", rt.uploadPhoto)

	rt.router.GET("/users/:userId/photos/photoId", rt.getPhoto)
	rt.router.DELETE("/users/:userId/photos/photoId", rt.deletePhoto)

	rt.router.GET("/users/:userId/feed", rt.getFeed)

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
