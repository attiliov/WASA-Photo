package api

type Username struct {
	Value string `json:"value"`
}

type ResourceID struct {
	Value string `json:"value"`
}

type Date struct {
	Value string `json:"value"`
}

type Caption struct {
	Value string `json:"value"`
}

type Image struct {
	URI string `json:"uri"`
}

type Counter struct {
	Count int `json:"count"`
}

type UserCollection struct {
	Users []User `json:"users"`
}

type User struct {
	UserID        ResourceID `json:"userId"`
	Username      Username   `json:"username"`
	SignUpDate    Date       `json:"signUpDate"`
	LastSeenDate  Date       `json:"lastSeenDate"`
	Bio           Caption    `json:"bio"`
	ProfileImage  Image      `json:"profileImage"`
	Followers     Counter    `json:"followers"`
	Following     Counter    `json:"following"`
}

type UserPost struct {
	PostID         ResourceID `json:"postId"`
	AuthorUsername Username   `json:"authorUsername"`
	CreationDate   Date       `json:"creationDate"`
	Caption        Caption    `json:"caption"`
	Image          Image      `json:"image"`
	LikeCount      Counter    `json:"likeCount"`
	CommentCount   Counter    `json:"commentCount"`
}

type Comment struct {
	CommentID      ResourceID `json:"commentId"`
	AuthorUsername Username   `json:"authorUsername"`
	AuthorID       ResourceID `json:"authorId"`
	CreationDate   Date       `json:"creationDate"`
	Caption        Caption    `json:"caption"`
	LikeCount      Counter    `json:"likeCount"`
}

type PostStream struct {
	Posts []ResourceID `json:"posts"`
}

type CommentStream struct {
	Comments []Comment `json:"comments"`
}

type Error struct {
	Message string `json:"message"`
}

type Success struct {
	Message string      `json:"message"`
	Body    interface{} `json:"body"`
}

type Like struct {
	Resource ResourceID `json:"likeId"`		//The post or the comment
	UserID        ResourceID `json:"userId"`
	Username      Username   `json:"username"`
}

type LikeCollection struct {
	Likes []Like `json:"likes"`
}
