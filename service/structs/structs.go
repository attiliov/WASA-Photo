package structs

type Username struct {
	Username string `json:"username"`
}
type ResourceID struct {
	ResourceID string `json:"resourceId"`
}
type Date struct {
	Date string `json:"date"`
}
type Caption struct {
	Caption string `json:"caption"`
}
type Image struct {
	Id string `json:"image"`
}
type Counter struct {
	Value int `json:"counter"`
}

type UserCollection struct {
	Users []User `json:"users"`
}

type User struct {
	UserID       string `json:"userId"`
	Username     string `json:"username"`
	SignUpDate   string `json:"signUpDate"`
	LastSeenDate string `json:"lastSeenDate"`
	Bio          string `json:"bio"`
	ProfileImage string `json:"profileImage"`
	Followers    int    `json:"followers"`
	Following    int    `json:"following"`
}

type UserPost struct {
	PostID         string `json:"postId"`
	AuthorUsername string `json:"authorUsername"`
	AuthorID       string `json:"authorId"`
	CreationDate   string `json:"creationDate"`
	Caption        string `json:"caption"`
	Image          string `json:"image"`
	LikeCount      int    `json:"likeCount"`
	CommentCount   int    `json:"commentCount"`
}

type Comment struct {
	CommentID      string `json:"commentId"`
	AuthorUsername string `json:"authorUsername"`
	AuthorID       string `json:"authorId"`
	CreationDate   string `json:"creationDate"`
	Caption        string `json:"caption"`
	LikeCount      int    `json:"likeCount"`
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
	Resource string `json:"resourceId"` // The post or the comment
	UserID   string `json:"userId"`
	Username string `json:"username"`
}

type LikeCollection struct {
	Likes []Like `json:"likes"`
}
