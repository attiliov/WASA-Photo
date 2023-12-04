package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.GET("/", rt.getHelloWorld)
	rt.router.GET("/context", rt.wrap(rt.getContextReply))


	rt.router.POST("/session", rt.wrap(rt.createUser))

	rt.router.GET("/users", rt.wrap(rt.getUsersList))

	rt.router.GET("/users/:userId", rt.wrap(rt.getUserProfile))
	rt.router.PUT("/users/:userId", rt.wrap(rt.updateUserProfile))
	rt.router.DELETE("/users/:userId", rt.wrap(rt.deleteUserProfile))

	rt.router.GET("/users/:userId/posts", rt.wrap(rt.getUserPosts))
	rt.router.POST("/users/:userId/posts", rt.wrap(rt.createPost))

	rt.router.GET("/users/:userId/posts/postId", rt.wrap(rt.getPost))
	rt.router.PUT("/users/:userId/posts/postId", rt.wrap(rt.editPost))
	rt.router.DELETE("/users/:userId/posts/postId", rt.wrap(rt.deletePost))

	rt.router.GET("/users/:userId/posts/postId/likes", rt.wrap(rt.getPostLikes))

	rt.router.PUT("/users/:userId/posts/:postId/likes/:likeId", rt.wrap(rt.editPost))
	rt.router.DELETE("/users/:userId/posts/:postId/likes/:likeId", rt.wrap(rt.deletePost))

	rt.router.GET("/users/:userId/posts/postId/comments", rt.wrap(rt.getPostComments))
	rt.router.POST("/users/:userId/posts/postId/comments", rt.wrap(rt.createComment))

	rt.router.GET("/users/:userId/posts/postId/comments/:commentId", rt.wrap(rt.getComment))
	rt.router.PUT("/users/:userId/posts/postId/comments/commentId",rt.wrap(rt.editComment))
	rt.router.DELETE("/users/:userId/posts/postId/comments/:commentId",rt.wrap(rt.deleteComment))

	rt.router.GET("/users/:userId/posts/postId/comments/:commentId/likes", rt.wrap(rt.getCommentLikes))

	rt.router.PUT("/users/:userId/posts/postId/comments/:commentId/likes/likeId", rt.wrap(rt.likeComment))
	rt.router.DELETE("/users/:userId/posts/postId/comments/:commentId/likes/likeId", rt.wrap(rt.unlikeComment))

	rt.router.GET("/users/:userId/follower", rt.wrap(rt.getFollowersList))
	
	rt.router.GET("/user/:userId/following", rt.wrap(rt.getFollowingsList))
	
	rt.router.PUT("/users/:userId/following/:followingId", rt.wrap(rt.followUser))
	rt.router.DELETE("/users/:userId/following/followingId", rt.wrap(rt.unfollowUser))

	rt.router.GET("/users/:userId/banned", rt.wrap(rt.getUserBanList))

	rt.router.PUT("/users/:userId/banned/:bannedId", rt.wrap(rt.banUser))
	rt.router.DELETE("/users/:userId/banned/:bannedId", rt.wrap(rt.unbanUser))

	rt.router.POST("/users/:userId/photos", rt.wrap(rt.uploadPhoto))

	rt.router.GET("/users/:userId/photos/photoId", rt.wrap(rt.getPhoto))
	rt.router.DELETE("/users/:userId/photos/photoId", rt.wrap(rt.deletePhoto))

	rt.router.GET("/users/:userId/feed", rt.wrap(rt.getFeed))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
