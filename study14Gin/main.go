package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"log"
	"net/http"
	"time"
)

import (
	Middle "src/study14Gin/session"
)

/*gin路由*/

func main1() {
	//1.创建路由
	r := gin.Default()
	//2.绑定路由规则，执行的函数
	//Restful风格
	//API参数
	r.GET("/user/:hz/*action", func(c *gin.Context) {
		hz, action := c.Param("hz"), c.Param("action")

		c.String(http.StatusOK, hz+"---------"+action)
	})
	r.POST("/post", func(c *gin.Context) {
		c.String(http.StatusOK, "hello Java")
	})
	r.PUT("/put", func(c *gin.Context) {
		c.String(http.StatusOK, "hello Golang")
	})
	r.DELETE("/delete", func(c *gin.Context) {
		c.String(http.StatusOK, "hello Python ")
	})

	//URL参数
	//DefaultQuery()
	r.GET("/hello", func(context *gin.Context) {
		//localhost:8080/hello?hz=xxx
		values := context.Query("hz")
		value := context.DefaultQuery("hz", "hello---") //获取查询参数，并给定默认值
		context.String(http.StatusOK, value, values)
	})
	//表单登录页面
	r.POST("/login", func(context *gin.Context) {
		types := context.DefaultPostForm("username", "alert")                            //获取Form表单数据，并给定默认值
		username, password := context.PostForm("username"), context.PostForm("password") //获取Form表单数据
		//多选框
		hobbys := context.PostFormArray("hobby")
		context.String(http.StatusOK, fmt.Sprintf("types is %s\nusername is %s\npassword is %s\nhobbys are %v\n",
			types, username, password, hobbys))
	})
	//上传文件
	r.POST("/upload", func(context *gin.Context) {
		//获取表单文件
		//设置最大8mb,默认32mb
		r.MaxMultipartMemory = 8 << 20
		/*file, err := context.FormFile("file")
		if err != nil {
			fmt.Println("upload failed...", err)
			return
		}
		log.Println(file.Filename)
		//传到项目根目录
		context.SaveUploadedFile(file, file.Filename)
		//打印此信息
		context.String(200, fmt.Sprintf("%s upload success!", file.Filename))*/
		form, err := context.MultipartForm()
		if err != nil {
			context.String(http.StatusBadRequest, fmt.Sprintf("err %s", err.Error()))
		}
		//获取所有图片
		files := form.File["file"]
		//遍历
		for _, file := range files {
			if err := context.SaveUploadedFile(file, file.Filename); err != nil {
				context.String(http.StatusOK, fmt.Sprintf("upload err %s", err.Error()))
				return
			}
		}
		context.String(http.StatusOK, fmt.Sprintf("upload ok %d file", len(files)))
	})
	//路由组
	g1, g2 := r.Group("g1"), r.Group("g2")
	//g1组处理get请求
	{
		g1.GET("/get1", login)
		g2.GET("/get2", submit)
	}
	//g2组处理post请求
	{
		g1.POST("/post1", login)
		g2.POST("/post2", submit)
	}
	r.Run(":8000")
}

func submit(context *gin.Context) {
	hz := context.DefaultQuery("hz", "lhz")
	context.String(200, hz)
}

func login(context *gin.Context) {
	name := context.DefaultQuery("name", "zhl")
	context.String(http.StatusOK, fmt.Sprintf("%s", name))
}

/*数据解析与绑定*/

type Login struct {
	User     string `form:"username" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}
type Person struct {
	ID   string `uri:"id" binding:"required,uuid"` //必须是uuid	格式
	Name string `uri:"name" binding:"required"`
}

func main2() {
	router := gin.Default()
	// Binding from JSON
	// curl -v -X POST http://localhost:8000/loginJSON  -H 'content-type: application/json' -d '{"user": "lhz","password": "123" }'
	// Example for binding JSON ({"user": "lhz", "password": "123"})
	router.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		//将request的body中的数据，自动按照json格式解析到结构体
		if err := c.ShouldBindJSON(&json); err != nil {
			//gin.H封装了生成json数据的工具
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if json.User != "lhz" || json.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// Example for binding XML (
	//		<?xml version="1.0" encoding="UTF-8"?>
	//		<root>
	//			<user>lhz</user>
	//			<password>123</password>
	//		</root>)
	router.POST("/loginXML", func(c *gin.Context) {
		var xml Login
		if err := c.ShouldBindXML(&xml); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if xml.User != "lhz" || xml.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// Example for binding a HTML form (user=lhz&password=123)
	router.POST("/loginForm", func(c *gin.Context) {
		var form Login
		// This will infer what binder to use depending on the content-type header.
		if err := c.Bind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if form.User != "lhz" || form.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})
	router.GET("/:name/:id", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBindUri(&person); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"name": person.Name, "uuid": person.ID})
	})

	router.Run(":8000")
}

/*gin 渲染*/
func main3() {
	r := gin.Default()
	//1.json
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hello", "status": "200"})
	})
	//2.结构体
	r.GET("someSTRUCT", func(c *gin.Context) {
		var msg struct {
			name    string `json:"user"`
			message string
			number  int
		}
		msg.name, msg.message, msg.number = "lhz", "hello", 666
		fmt.Println(msg)
		c.JSON(http.StatusOK, msg)
	})
	//3.XML
	r.GET("someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "xml", "status": "200"})
	})
	//4.YAML
	r.GET("someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "YAML", "status": "200"})
	})
	//5,ProtoBuf
	r.GET("somePROTOBUF", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "lhz"
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		c.ProtoBuf(http.StatusOK, data)
	})

	//Html模版渲染
	r.LoadHTMLGlob("src/templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{"title": "模版渲染"})
	})

	//重定向
	r.GET("study2_test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://leetcode.cn")
	})
	r.GET("test1", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/someJSON")
	})
	r.GET("/test2", func(c *gin.Context) {
		c.Request.URL.Path = "/test3"
		r.HandleContext(c)
	})
	r.GET("/test3", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hello": "world"})
	})
	//异步
	r.GET("test4", func(c *gin.Context) {
		//异步不应该适用原始上下文，必须适用他的可读副本
		copyContext := c.Copy()
		go func() {
			time.Sleep(time.Second)
			log.Println("异步执行" + copyContext.Request.URL.Path)
		}()
	})
	r.Run(":8000")
}

/*中间件*/
func main4() {
	r := gin.New()
	r.Use(Middleware())
	//代码规范
	{
		r.GET("/middle1", func(c *gin.Context) {
			name := c.MustGet("lhz1")
			log.Println(name)
		})

		r.GET("/middle2", Middleware(), func(c *gin.Context) {
			name := c.MustGet("lhz2")
			log.Println(name)
		})
	}

	r.Run(":8000")
}

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t1 := time.Now()
		c.Set("lhz", "hello lhz")
		c.Next()
		t2 := time.Since(t1)
		log.Printf("执行用时：%dns", t2)
		log.Printf("执行用时：%fs", t2.Seconds())
		status := c.Writer.Status()
		log.Println(status)
	}
}

// BasicAuth
var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func main5() {
	r := gin.Default()

	//路由组使用中间件
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))

	// /admin/secrets endpoint
	// hit "localhost:8080/admin/secrets
	authorized.GET("/secrets", func(c *gin.Context) {
		// 获取中间件设置的key
		user := c.MustGet(gin.AuthUserKey).(string)

		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})
	r.Run(":8000")

}

// cookie验证练习
func main6() {
	r := gin.Default()

	//先访问login
	r.GET("/login", func(c *gin.Context) {
		//设置cookie信息
		c.SetCookie("lhz", "123", 60, "/", "localhost", false, true)
		//返回信息
		c.String(http.StatusOK, "login success")
	})
	//再访问home
	r.GET("/home", AuthMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "home"})
	})
	r.Run(":9000")
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取客户端cookie校验
		if str, err := c.Cookie("lhz"); err == nil {
			if str == "123" {
				c.Next()
				return
			}
		}
		//返回错误
		c.JSON(http.StatusUnauthorized, gin.H{"error": "err"})
		//不再向下执行
		c.Abort()
	}
}

// 自定义session中间件
func main() {
	memory, err := Middle.Init("memory", "127.0.0.1", "1234")
	if err != nil {
		fmt.Println("初始化错误:", err)
		return
	}
	memory.Init("127.0.0.1", "1234")
	session, err := memory.CreateSession()

	err = session.Set("lhz", "12345")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	value, err := session.Get("lhz")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	fmt.Println(value)
	err = session.Del("lhz")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	value, err = session.Get("lhz")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	fmt.Println(value)
}
