package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"path/filepath"
	"time"
)

/*
	func main1() {
		app := iris.Default()
		app.Use(myMiddleware)

		app.Handle("GET", "/ping", func(ctx iris.Context) {
			ctx.JSON(iris.Map{"message": "pong"})
		})

		// Listens and serves incoming http requests
		// on http://localhost:8080.
		app.Run(iris.Addr(":8080"))
	}

	func myMiddleware(ctx iris.Context) {
		ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
		ctx.Next()
	}
*/
func main2() {
	app := iris.New()

	app.RegisterView(iris.HTML("./", ".html"))

	app.Get("/", func(ctx iris.Context) {
		// Bind: {{.message}} with "Hello world!"
		ctx.ViewData("message", "Hello iris!")
		// Render template file: ./views/hello.html
		ctx.View("iris.html")
	})

	// app.Get("/user/{id:string regexp(^[0-9]+$)}")
	app.Get("/user/{id:uint64}", func(ctx iris.Context) {
		userID, _ := ctx.Params().GetUint64("id")
		ctx.Writef("User ID: %d", userID)
	})

	//配置
	/*config := iris.WithConfiguration(iris.Configuration{
		DisableStartupLog: true,
		Charset:           "UTF-8",
	})

	app.Run(iris.Addr(":8080"), config)*/
	//函数式配置
	/*app.Run(iris.Addr(":8080"), iris.WithoutInterruptHandler,
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithoutBodyConsumptionOnUnmarshal,
		iris.WithoutAutoFireStatusCode,
		iris.WithOptimizations,
		iris.WithTimeFormat("Mon, 01 Jan 2006 15:04:05 GMT"),
	)*/
	//注册接受两个 int 参数的自定义的宏函数。
	//定义最短长度,最长长度
	app.Macros().Get("string").RegisterFunc("range",
		func(minLength, maxLength int) func(string) bool {
			return func(paramValue string) bool {
				return len(paramValue) >= minLength && len(paramValue) <= maxLength
			}
		})

	app.Get("/limitchar/{name:string range(10,200) else 400}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		ctx.Writef(`Hello %s | the name should be between 1 and 200 characters length
    otherwise this handler will not be executed`, name)
	})
	//注册接受一个 []string 参数的自定义的宏函数。
	app.Macros().Get("string").RegisterFunc("has",
		func(validNames []string) func(string) bool {
			return func(paramValue string) bool {
				for _, validName := range validNames {
					if validName == paramValue {
						return true
					}
				}
				return false
			}
		})

	app.Get("/static_validation/{name:string has([hello,world])}",
		func(ctx iris.Context) {
			name := ctx.Params().Get("name")
			ctx.Writef(`Hello %s | the name should be "hello" or "world"
    otherwise this handler will not be executed`, name)
		}).Name = "static_validation"

	//路由命名
	// define a function
	h := func(ctx iris.Context) {
		ctx.HTML("<b1>Hi</b1>")
	}

	// handler registration and naming
	home := app.Get("/home", h)
	home.Name = "home"
	// or
	app.Get("/about", h).Name = "about"
	app.Get("/page/{id}", h).Name = "page"

	app.Run(iris.Addr(":8080"))

}

/*// 编写中间件
func main3() {
	app := iris.New()
	app.Use(func(context iris.Context) {
		context.Record() //响应记录
		context.Next()
	})
	//app.Get("/middle", before, Handler, after)
	//全局范围
	app.Get("/middle", Handler)
	app.UseGlobal(before)
	app.DoneGlobal(after)

	//匹配所有除了被其他路由器处理的 GET 请求
	app.Get("{root:path}", func(context iris.Context) {

		context.Write([]byte("=================匹配所有============="))
		context.WriteString("响应器:" + string(context.Recorder().Body())) //响应器
		//Println(context.Recorder().Body())
	})

	//处理Http错误
	app.OnErrorCode(iris.StatusNotFound, notFound)
	app.OnErrorCode(iris.StatusInternalServerError, internalServerError)

	////子域名,需要修改hosts文件
	//admin := app.Subdomain("admin")
	//admin.Get("/", func(ctx iris.Context) {
	//	ctx.Writef("INDEX FROM admin.mydomain.com")
	//})
	//app.Run(iris.Addr("localhost:8080"))
	//包装路由器
	app.WrapRouter(func(w http.ResponseWriter, r *http.Request, router http.HandlerFunc) {
		path := r.URL.Path
		//localhost:8080/other....
		if strings.HasPrefix(path, "/other") {
			// acquire and release a context in order to use it to execute
			// our custom handler
			// remember: we use net/http.Handler because here
			// we are in the "low-level", before the router itself.
			ctx := app.ContextPool.Acquire(w, r)
			Handler(ctx)
			app.ContextPool.Release(ctx)
			return
		}

		// else continue serving routes as usual.
		router.ServeHTTP(w, r)
	})
	//API版本控制
	myCustomNotVersionFound := func(ctx iris.Context) {
		ctx.StatusCode(404)
		ctx.Writef("%s version not found", versioning.GetVersion(ctx))
	}

	userAPI := app.Party("/api/user")
	userAPI.Get("/", versioning.NewMatcher(versioning.Map{
		"1.0":               Handler,
		">= 2, < 3":         Handler,
		versioning.NotFound: myCustomNotVersionFound,
	}))
	//通过版本分组
	//http://localhost:8080/api/user/json
	userAPIV10 := versioning.NewGroup("4.0") //app.UseGlobal(before)设置版本
	userAPIV10.Get("/json", func(context iris.Context) {
		context.JSON(iris.StatusOK, iris.JSON{Indent: "json"})
	})
	//版本注册
	versioning.RegisterGroups(userAPI, versioning.NotFoundHandler, userAPIV10)
	//不同内用协商
	//json
	type testdata struct {
		Name string `json:"name" xml:"Name"`
		Age  int    `json:"age" xml:"Age"`
	}
	app.Get("/resource", func(ctx iris.Context) {
		data := testdata{
			Name: "test name",
			Age:  26,
		}

		ctx.Negotiation().JSON().XML().EncodingGzip()

		_, err := ctx.Negotiate(data)
		if err != nil {
			ctx.Writef("%v", err)
		}
	})
	//Http Referer
	app.Get("/refer", func(ctx iris.Context) {
		r := ctx.GetReferrer()
		switch r.Type {
		case iris.ReferrerSearch:
			ctx.Writef("=====Search %s: %s\n", r.Label, r.Query)
			ctx.Writef("Google=====: %s\n", r.GoogleType)
		case iris.ReferrerSocial:
			ctx.Writef("====Social=== %s\n", r.Label)
		case iris.ReferrerIndirect:
			ctx.Writef("====Indirect====: %s\n", r.URL)
		}
	})
	app.Run(iris.Addr(":8080"))
}
func before(ctx iris.Context) {
	//设置版本
	ctx.Values().Set(versioning.Key, ctx.URLParamDefault("version", "4.0"))

	value := "这是一个值"
	reqpath := ctx.Path()
	Println("before:" + reqpath)
	ctx.Values().Set("key", value)
	ctx.Next()
}
func after(ctx iris.Context) {
	Println("after:")
}
func Handler(ctx iris.Context) {
	value := ctx.Values().GetString("key")
	ctx.HTML("<h1>Response</h1>")
	ctx.HTML("<br/> value : " + value)

	ctx.Next() // execute the "after".
}
func notFound(ctx iris.Context) {

	//ctx.View("errors/404.html")
	ctx.View("iris.html")
}

func internalServerError(ctx iris.Context) {
	ctx.WriteString("Oups something went wrong, try again")
}*/

// jwt
func getTokenHandler(ctx iris.Context) {
	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
	}) //// 往jwt中写入了一对值

	tokenString, _ := token.SignedString([]byte("My Secret")) //生成给客户端的token

	ctx.HTML(`-------Token----: ` + tokenString + `<br/><br/>
    <a href="/secured?token=` + tokenString + `">点击验证:       secured?token=` + tokenString + `</a>`)
}

func myAuthenticatedHandler(ctx iris.Context) {
	//// 获取jwt里的值
	user := ctx.Values().Get("jwt").(*jwt.Token)

	ctx.Writef("This is an authenticated request\n")
	ctx.Writef("Claim content:\n")

	foobar := user.Claims.(jwt.MapClaims)
	for key, value := range foobar {
		ctx.Writef("%s = %s", key, value)
	}
}

func main4() {
	app := iris.New()

	j := jwt.New(jwt.Config{
		//// 从请求参数token中提取
		Extractor: jwt.FromParameter("token"),
		////设置一个函数返回秘钥
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("My Secret"), nil
		},
		////设置一个加密方法
		SigningMethod: jwt.SigningMethodHS256,
	})

	app.Get("/", getTokenHandler)                        //得到生成的令牌
	app.Get("/secured", j.Serve, myAuthenticatedHandler) //验证
	app.Run(iris.Addr(":8080"))
}

func main5() {
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.WriteString(ctx.URLParamDefault("name", "lhz"))
		ctx.Writef("%t", ctx.URLParamExists("name"))
		ok, err := ctx.URLParamBool("name")

		if err != nil {
			fmt.Println("err:", err)
		}
		ctx.WriteString(ctx.URLParamTrim("name"))
		ctx.Writef("%t", ok)
	})

	app.Run(iris.Addr(":8080"))

}

// 表单
func main6() {
	app := iris.Default()
	app.RegisterView(iris.HTML("./", ".html"))
	app.Post("/form_post", func(ctx iris.Context) {
		message := ctx.FormValue("message")
		nick := ctx.FormValueDefault("nick", "anonymous")

		ctx.JSON(iris.Map{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	app.Run(iris.Addr(":8080"))
}

const maxSize = 8 * iris.MB

func main9() {
	app := iris.New()
	app.Use(iris.StaticCache(-1)) //静态缓存

	app.RegisterView(iris.HTML("./", ".html"))
	app.Post("/upload", func(ctx iris.Context) {

		// Set a lower memory limit for multipart forms (default is 32 MiB)
		ctx.SetMaxRequestBodySize(maxSize)
		// OR
		// app.Use(iris.LimitRequestBodySize(maxSize))
		// OR
		// OR iris.WithPostMaxMemory(maxSize)

		// single file
		_, fileHeader, err := ctx.FormFile("file")
		if err != nil {
			ctx.StopWithError(iris.StatusBadRequest, err)
			return
		}

		// Upload the file to specific destination.

		dest := filepath.Join("./uploads", fileHeader.Filename)

		ctx.SaveFormFile(fileHeader, dest)

		ctx.Writef("File: %s uploaded!", fileHeader.Filename)
	}, iris.NoCache, iris.Cache304(time.Duration(time.Second))) //NOCache:在 HTML 路由上使用这个中间件；即使在浏览器的“后退”和“前进”箭头按钮上也可以刷新页面。

	app.Listen(":8000")
}

// 模型验证
type User struct {
	FirstName      string     `json:"fname"`
	LastName       string     `json:"lname"`
	Age            uint8      `json:"age" validate:"gte=0,lte=130"`
	Email          string     `json:"email" validate:"required,email"`
	FavouriteColor string     `json:"favColor" validate:"hexcolor|rgb|rgba"`
	Addresses      []*Address `json:"addresses" validate:"required,dive,required"`
}

// Address houses a users address information.
type Address struct {
	Street string `json:"street" validate:"required"`
	City   string `json:"city" validate:"required"`
	Planet string `json:"planet" validate:"required"`
	Phone  string `json:"phone" validate:"required"`
}

// Use a single instance of Validate, it caches struct info.
var validate *validator.Validate

func main8() {
	validate = validator.New()

	validate.RegisterStructValidation(UserStructLevelValidation, User{})

	app := iris.New()
	app.Post("/user", func(ctx iris.Context) {
		var user User
		if err := ctx.ReadJSON(&user); err != nil {
			// [handle error...]
		}
		// nil or ValidationErrors ( []FieldError )
		err := validate.Struct(user)
		if err != nil {
			// This check is only needed when your code could produce
			// an invalid value for validation such as interface with nil
			// value most including myself do not usually have code like this.
			if _, ok := err.(*validator.InvalidValidationError); ok {
				ctx.StatusCode(iris.StatusInternalServerError)
				ctx.WriteString(err.Error())
				return
			}
			ctx.StatusCode(iris.StatusBadRequest)
			for _, err := range err.(validator.ValidationErrors) {
				fmt.Println()
				fmt.Println(err.Namespace())
				fmt.Println(err.Field())
				fmt.Println(err.StructNamespace())
				fmt.Println(err.StructField())
				fmt.Println(err.Tag())
				fmt.Println(err.ActualTag())
				fmt.Println(err.Kind())
				fmt.Println(err.Type())
				fmt.Println(err.Value())
				fmt.Println(err.Param())
				fmt.Println()
			}
			return
		}
		// [save user to database...]
	})

	app.Run(iris.Addr(":8080"))
}

func UserStructLevelValidation(sl validator.StructLevel) {
	user := sl.Current().Interface().(User)

	if len(user.FirstName) == 0 && len(user.LastName) == 0 {
		sl.ReportError(user.FirstName, "FirstName", "fname", "fnameorlname", "")
		sl.ReportError(user.LastName, "LastName", "lname", "fnameorlname", "")
	}
}

// 视图
func main10() {
	app := iris.New()

	tmpl := iris.HTML("./", ".html")

	tmpl.Reload(true)

	app.RegisterView(tmpl)
	tmpl.AddFunc("greet", func(s string) string {
		return "Greetings " + s + "!"
	})
	app.Get("/", func(ctx iris.Context) {
		// Bind: {{.message}} with "Hello world!"
		ctx.ViewData("message", "Hello world!")
		// Render template file: ./views/hi.html
		ctx.View("iris.html")
	})

	app.Run(iris.Addr(":8080"))
}

// session
var (
	cookieNameForSessionID = "mycookiesessionnameid"
	sess                   = sessions.New(sessions.Config{Cookie: cookieNameForSessionID})
)

func secret(ctx iris.Context) {
	// Check if user is authenticated
	if auth, _ := sess.Start(ctx).GetBoolean("authenticated"); !auth {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}

	// Print secret message
	ctx.WriteString("The cake is a lie!")
}

func login(ctx iris.Context) {
	session := sess.Start(ctx)

	// Authentication goes here
	// ...

	// Set user as authenticated
	session.Set("authenticated", true)
}

func logout(ctx iris.Context) {
	session := sess.Start(ctx)

	// Revoke users authentication
	session.Set("authenticated", false)
	// Or to remove the variable:
	session.Delete("authenticated")
	// Or destroy the whole session:
	session.Destroy()
}

func main() {
	app := iris.New()

	app.Get("/secret", secret)
	app.Get("/login", login)
	//app.Get("/get", func(context iris.Context) {
	//	session := sess.Start(context)
	//	str := session.GetString("authenticated")
	//	context.WriteString(str)
	//})
	app.Get("/logout", logout)

	app.Run(iris.Addr(":8080"))
}
