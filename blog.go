package main

import (
	"blog/controller/admin"
	"blog/controller/page"
	"blog/dao"
	"blog/middleware"
	"encoding/gob"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"html/template"
)

/*
 url  --> controller -->  logic   -->   model
请求         控制层       业务逻辑层    模型层的增删改查
*/

func main() {

	err := dao.InitMysql()
	if err != nil {
		panic(err)
	}
	defer dao.Close()

	//把user这个接头体注册进来，后面跨路由才可以获取到user数据
	gob.Register(middleware.U{})
	r := gin.New()
	r.SetFuncMap(template.FuncMap{
		"add":                    middleware.Add,
		"notAdd":                 middleware.NotAdd,
		"year":                   dao.BlogYear,
		"month":                  dao.ArchMonth,
		"selectTypeNameByBlog":   dao.SelectTypeNameByBlog,
		"clearComments":          dao.ClearComments,
		"selectCommentByComment": dao.SelectCommentByComment,
		"SUMBlog": dao.SUMBlog,
		"SUMBlogNUM": dao.SUMBlogNUM,
		"NewBlog": dao.NewBlog,
		"SelectBlog": dao.SelectBlog,
		"SelectNewComment": dao.SelectNewComment,
		"SelectOneType": dao.SelectOneType,
		"safe": dao.Safe,
		"notes": dao.SelectNewNotes,
	})
	store := cookie.NewStore([]byte("secret"))
	//设置session中间件
	r.Use(sessions.Sessions("mysession", store))

	r.Static("/static", "static")
	r.LoadHTMLGlob("templates/**/*")

	//首页
	r.Any("/", page.IndexHandler)
	r.Any("/notes", page.NotesController)
	r.Any("/research", page.ResearchHandler)

	r.Any("/blogs/page/:page", page.PagesHandler)
	r.Any("/comments", page.CommentHandler)
	//分类
	r.GET("/types", page.TypesHandler)
	r.Any("/types/:name", page.TypesNameHandler)
	r.Any("/types/:name/:id", page.BlogHandler)
	//标签
	r.GET("/tags", page.TagsHandler)
	r.Any("/tags/:name", page.TagsNameHandler)
	//归档
	r.GET("/archives", page.ArchivesHandler)
	//关于
	r.GET("/about", page.AboutHandler)

	r.Any("/admin/login", admin.LoginHandler)
	r.GET("/admin/logout", admin.LogoutHandler)

	adminGroup := r.Group("/admin")
	adminGroup.Use(middleware.Authorization)
	{
		adminGroup.GET("/", admin.AdminHandler)
		adminGroup.GET("/index", admin.AdminIndexHandler)

		typeGroup := adminGroup.Group("/type")
		{
			typeGroup.GET("/", admin.TypeHandler)
			typeGroup.Any("/edit", admin.TypeEditHandler)
			typeGroup.Any("/add", admin.TypeAddHandler)
		}
		tagGroup := adminGroup.Group("/tag")
		{
			tagGroup.Any("/", admin.TagHandler)
			tagGroup.Any("/edit", admin.TagEditHandler)
			tagGroup.Any("/add", admin.TagAddHandler)
		}
		BlogGroup := adminGroup.Group("/blog")
		{
			BlogGroup.Any("/", admin.BlogHandler)
			BlogGroup.Any("/edit", admin.BlogEditHandler)
			BlogGroup.Any("/add", admin.BlogAddHandler)
		}
		UserGroup := adminGroup.Group("/user")
		{
			UserGroup.Any("/", admin.UserEditHandler)
			//UserGroup.Any("/edit", controller.UserEditHandler)
			//UserGroup.Any("/add", controller.BlogAddHandler)
		}

	}

	//404
	r.NoRoute(page.NotFoundHandler)

	//启动服务
	r.Run(":81")

	//api优雅重启接口初始化
	//if err := endless.ListenAndServe(":81", r); err!=nil{
	//	log.Fatalf("listen: %s\n", err)
	//}

}
