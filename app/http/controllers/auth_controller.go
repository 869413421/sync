package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync/pkg/auth"
	"sync/pkg/model/user"
)

type AuthController struct {
}

// Login 登录
func (*AuthController) Login(context *gin.Context) {
	//1.获取表单数据
	email := context.Request.PostFormValue("email")
	password := context.Request.PostFormValue("password")

	//2.用户认证
	errors := auth.Attempt(email, password)
	if len(errors) > 0 {
		fmt.Println(errors)
		return
	}

	//3.登录成功
}

// Register 注册页面
func (*AuthController) Register(w http.ResponseWriter, r *http.Request) {
	view.RenderOnce(w, "register", view.D{}, "auth.register")
}

// DoRegister 注册
func (*AuthController) DoRegister(w http.ResponseWriter, r *http.Request) {
	//1.获取数据
	name := r.PostFormValue("name")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	passwordComfirm := r.PostFormValue("passwordComfirm")

	//2.构建数据
	_user := user.User{
		Name:            name,
		Email:           email,
		Password:        password,
		PasswordComfirm: passwordComfirm,
	}

	//3.验证数据
	errs := requests.VaildateRegisterationForm(_user)
	if len(errs) > 0 {
		flash.Danger("注册失败")
		view.RenderOnce(w, "register", view.D{
			"Errors": errs,
			"User":   _user,
		}, "auth.register")
		return
	}

	//4.验证成功入库
	_user.Create()
	if _user.ID <= 0 {
		flash.Danger("注册失败")
		view.RenderOnce(w, "register", view.D{
			"Errors": errs,
			"User":   _user,
		}, "auth.register")
		return
	}

	//5.返回成功信息
	flash.Success("注册成功")
	http.Redirect(w, r, "/test", http.StatusFound)
}
