package req

type UserLoginReq struct {
	Username string `json:"username" binding:"required,min=2,max=20"` // 账号
	Password string `json:"password" binding:"required,min=6,max=32"` // 密码
}

type UserRegister struct {
	Username string `json:"username" binding:"required,min=2,max=20"` // 账号
	Password string `json:"password" binding:"required,min=6,max=32"` // 密码
	Sex      string `json:"sex" binding:"required"`
	Age      string `json:"age" binding:"required"`
	Phone    string `json:"phone" binding:"required,min=9,max=11"`
}
