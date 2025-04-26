package event

import (
	"crypto/md5"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	ID       string `json:"user_id" gorm:"primaryKey"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Reply 响应
type Reply struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

// MsgJson 响应数据
func MsgJson(c *gin.Context, reply *Reply) {

	if reply.Code == 0 {
		reply.Code = 200
	}

	c.JSON(http.StatusOK, reply)
}

func InitUserDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("MUS.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	_ = db.AutoMigrate(&User{})
}

func loginUser(c *gin.Context) {
	var userInfo struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	c.GetHeader("")

	if err := c.ShouldBindJSON(&userInfo); err != nil {
		MsgJson(c, &Reply{
			Code: 400,
			Msg:  err.Error(),
		})
		return
	}

	// 查找用户
	var user User
	if db.Where("username = ? AND password = ?", userInfo.Username, userInfo.Password).First(&user).Error != nil {
		MsgJson(c, &Reply{
			Code: 400,
			Msg:  "登录失败",
		})
		return
	}

	token, _ := SubScriber.GenToken(user.Username)

	MsgJson(c, &Reply{
		Msg:  "登录成功",
		Data: gin.H{"token": token, "user_id": user.ID},
	})
}

func registerUser(c *gin.Context) {
	var newUser User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		MsgJson(c, &Reply{
			Code: 400,
			Msg:  err.Error(),
		})
		return
	}

	// 检查用户名是否已经存在
	var user User
	if db.Where("username = ?", newUser.Username).First(&user).Error == nil {
		MsgJson(c, &Reply{
			Code: 400,
			Msg:  "用户已存在",
		})
		return
	}

	genMd5 := func(s string) string {
		sum := md5.Sum([]byte(s))
		return fmt.Sprintf("%x", sum)
	}

	// 为新用户生成唯一的 ID
	newUser.ID = genMd5(newUser.Username)

	// 创建用户记录
	result := db.Create(&newUser)

	if result.Error != nil {
		MsgJson(c, &Reply{
			Code: 400,
			Msg:  "用户注册失败",
		})
		return
	}

	// 返回成功响应
	MsgJson(c, &Reply{
		Msg:  "用户注册成功",
		Data: newUser.ID,
	})
}

func deleteUser(c *gin.Context) {
	// 获取用户 ID 参数
	userID := c.Param("user_id")

	// 删除用户记录
	result := db.Delete(&User{}, userID)
	if result.Error != nil {
		MsgJson(c, &Reply{
			Code: 400,
			Msg:  "用户不存在",
		})
		return
	}

	// 返回成功响应
	MsgJson(c, &Reply{Msg: "用户删除成功"})
}

func updateUser(c *gin.Context) {
	// 获取用户 ID 参数
	userID := c.Param("user_id")

	// 解析请求中的 JSON 数据，更新用户信息
	var updatedUser User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		MsgJson(c, &Reply{
			Code: 400,
			Msg:  err.Error(),
		})
		return
	}

	// 更新用户记录
	result := db.Model(&User{}).Where("user_id = ?", userID).Updates(updatedUser)
	if result.Error != nil || result.RowsAffected == 0 {
		MsgJson(c, &Reply{
			Code: 400,
			Msg:  "用户不存在",
		})
		return
	}

	// 返回成功响应
	MsgJson(c, &Reply{Msg: "用户更新成功"})
}

// func generateToken(username string) string {
// 	hashed := sha256.New()
// 	timestamp := strconv.Itoa(int(time.Now().Unix()))
// 	hashed.Write([]byte(username + timestamp))
// 	hashedBytes := hashed.Sum(nil)
// 	return hex.EncodeToString(hashedBytes)
// }
