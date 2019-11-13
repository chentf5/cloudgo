package service
import (
   "github.com/codegangsta/martini" 

   "github.com/martini-contrib/render"
   "github.com/martini-contrib/binding"
	"time"
	//"github.com/urfave/negroni"
)

type Post struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
func NewServer(port string) {   
	m := martini.Classic()
	m.Use(martini.Static("assets"))
	m.Use(render.Renderer())
    //提交请求的处理
    m.Get("/", func (r render.Render)  {
		curTime:=time.Now().Format("2006-01-02 15:04:05")
		r.HTML(200,"homePage",map[string]interface{}{"Time":curTime})
		
	})
	m.Post("/",binding.Bind(Post{}),func (post Post,r render.Render)  {
		curTime:= time.Now().Format("2006-01-02 15:04:05")
		p := Post{Username: post.Username, Password: post.Password}
		r.HTML(200,"info",map[string]interface{}{"post":p,"Time":curTime})		
	})
    m.RunOnAddr(":"+port)   
}
