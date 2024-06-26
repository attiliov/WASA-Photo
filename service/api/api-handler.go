package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {

	// Register routes
	rt.router.GET("/context", rt.wrap(rt.getContextReply))

	rt.router.POST("/session", rt.getAuthToken) // TESTED, TESTED ON FRONTEND TODO: add last seen update and chech why dates are not working properly in db

	rt.router.GET("/users", rt.searchUser) // TESTED

	rt.router.GET("/users/:userId", rt.getUserProfile)       // TESTED, on frontend
	rt.router.PUT("/users/:userId", rt.updateUserProfile)    // TESTED, ON FRONTEND
	rt.router.DELETE("/users/:userId", rt.deleteUserProfile) // TESTED, ON FRONTEND

	rt.router.GET("/users/:userId/posts", rt.getUserPosts) // TESTED, on frontend
	rt.router.POST("/users/:userId/posts", rt.createPost)  // TESTED, ON FRONTEND TODO: add chcek that if the photo is not null, the photo is saved in the db

	rt.router.GET("/users/:userId/posts/:postId", rt.getPost)       // TESTED, ON FRONTEND
	rt.router.PUT("/users/:userId/posts/:postId", rt.editPost)      // TESTED, ON FRONTEND
	rt.router.DELETE("/users/:userId/posts/:postId", rt.deletePost) // TESTED, ON FRONTEND

	rt.router.GET("/users/:userId/posts/:postId/likes", rt.getPostLikes) // TESTED, ON FRONTEND

	rt.router.PUT("/users/:userId/posts/:postId/likes/:likeId", rt.likePost)      // TESTED, ON FRONTEND
	rt.router.DELETE("/users/:userId/posts/:postId/likes/:likeId", rt.unlikePost) // TESTED, ON FRONTEND

	rt.router.GET("/users/:userId/posts/:postId/comments", rt.getPostComments) // TESTED, ON FRONTEND
	rt.router.POST("/users/:userId/posts/:postId/comments", rt.createComment)  // TESTED, ON FRONTEND

	rt.router.GET("/users/:userId/posts/:postId/comments/:commentId", rt.getComment)       // TESTED, ON FRONTEND
	rt.router.PUT("/users/:userId/posts/:postId/comments/:commentId", rt.editComment)      // TESTED, ON FRONTEND
	rt.router.DELETE("/users/:userId/posts/:postId/comments/:commentId", rt.deleteComment) // TESTED, ON FRONTEND

	rt.router.GET("/users/:userId/posts/:postId/comments/:commentId/likes", rt.getCommentLikes) // TESTED, ON FRONTEND

	rt.router.PUT("/users/:userId/posts/:postId/comments/:commentId/likes/:likeId", rt.likeComment)      // TESTED, ON FRONTEND
	rt.router.DELETE("/users/:userId/posts/:postId/comments/:commentId/likes/:likeId", rt.unlikeComment) // TESTED, ON FRONTEND

	rt.router.GET("/users/:userId/followers", rt.getFollowersList) // TESTED

	rt.router.GET("/users/:userId/following", rt.getFollowingsList) // TESTED

	rt.router.PUT("/users/:userId/following/:followingId", rt.followUser)      // TESTED
	rt.router.DELETE("/users/:userId/following/:followingId", rt.unfollowUser) // TESTED

	rt.router.GET("/users/:userId/banned", rt.getUserBanList) // TESTED

	rt.router.PUT("/users/:userId/banned/:bannedId", rt.banUser)      // TESTED
	rt.router.DELETE("/users/:userId/banned/:bannedId", rt.unbanUser) // TESTED

	rt.router.POST("/users/:userId/photos", rt.savePhoto) // TESTED on frontend

	rt.router.GET("/users/:userId/photos/:photoId", rt.getPhoto)       // TESTED on frontend
	rt.router.DELETE("/users/:userId/photos/:photoId", rt.deletePhoto) // TESTED on frontend

	rt.router.GET("/users/:userId/feed", rt.getFeed) // TESTED

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
