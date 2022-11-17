// package main
// // go doc method-name

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"regexp"
// 	"github.com/google/uuid"

// 	"github.com/gin-gonic/gin"
// 	"github.com/gin-gonic/gin/binding"
// 	"github.com/go-playground/validator/v10"
// )
// func verifyPassword(fl validator.FieldLevel) bool{
// 	var regex = regexp.MustCompile("\\w{8,}")
// 	var password=fl.Field().String()
// 	return regex.MatchString(password)
// }
// func main() {
// 	var router = gin.Default()
// 	var address = ":8080"

// 	// апишки в джине
// 	// router.GET
// 	// router.POST
// 	// router.DELETE
// 	// router.PUT
// 	// router.HEAD
// 	// router.OPTIONS

// 	// группирует похожие юрлки/префиксы в юрл
// 	var v1 = router.Group("/api/group-routes")
// 	// this method will response at endpoint-
// 	// http://localhost:8080/api/group-routes/get

// 	// controller- обработчик запросов
// 	v1.GET("/get",func(c *gin.Context){
// 		c.String(http.StatusOK, "get")
// 	})
// 	// this method will response at endpoint-
// 	// http://localhost:8080/hello
// 	router.GET("/hello",func(c *gin.Context){
// 		c.String(http.StatusOK, "hello")
// 	})
	
// 	// Accessing info from request
// 	// c.Param -returns a value of the url param
// 	router.GET("/products/:id",func(c *gin.Context){
// 		var id=c.Param("id")
// 		fmt.Println("Id:",id)
// 		c.String(http.StatusOK,"hello world")
// 	})
// 	// c.Query-returns the keyed url query value if it exists, otherwise empty string
// 	// var id=c.Query("id")

// 	// c.DefaultQuery
// 	// var id=c.Query("id",100) -разница в том что задаем дефолтный размер

// 	// c.PostFrom-returns the specified key from a POST urlencoded form or 
// 	// multipart form if it exists, otherwise it returns an empty string

// 	// c.DefaultPostFrom

// 	// c.GetHeader-returns value of request header
// 	// var id=c.GetHeader("id")


// 	// 	MODEL BINDING
// 	// when we want to map a specific information into our struct
// 	// c.ShouldBindUri-разница что дата поступает от url с методом ГЕТ
// 	type uriBinding struct{
// 		ID string `uri:"id"`

// 	}
// 	router.GET("/products/:id",func(c *gin.Context){
// 		var binding uriBinding
// 		if e:=c.ShouldBindUri(&binding);e!=nil{
// 			c.String(http.StatusBadRequest,e.Error())
// 			return
// 		}
// 		fmt.Println("Binding: ",binding)
// 		c.String(http.StatusOK,"hello world")
// 	})
// 	// c.ShouldBindJSON разница что дата поступает от json с методом ПОСТ
// 	type Product struct{
// 		ID string `json:"id"`
// 		Name string `json:"name"`
// 	}
// 	router.POST("/products",func(c *gin.Context){
// 		var product Product
		
// 		if e:=c.ShouldBindJSON(&product);e!=nil{
// 			c.String(http.StatusBadRequest,e.Error())
// 			return
// 		}
// 		fmt.Println("Product: ",product)
// 		c.String(http.StatusOK,"hello world")

// 	})
// 	// c.ShouldBind - разница что дата поступает от form-data с методом ПОСТ
// 	type Product2 struct{
// 		ID string `form:"id"`
// 		Name string `form:"name"`
// 	}
// 	router.POST("/products",func(c *gin.Context){
// 		var product Product2
		
// 		if e:=c.ShouldBind(&product);e!=nil{
// 			c.String(http.StatusBadRequest,e.Error())
// 			return
// 		}
// 		fmt.Println("Product: ",product)
// 		c.String(http.StatusOK,"hello world")

// 	})

// 	// c.ShouldBindHeader - разница что дата поступает от request header
// 	type headerBinding struct{
// 		RequestId string `header:"X-Request-ID"`
// 	}


// 	router.POST("/products",func(c *gin.Context){
// 		var binding headerBinding
		
// 		if e:=c.ShouldBind(&binding);e!=nil{
// 			c.String(http.StatusBadRequest,e.Error())
// 			return
// 		}
// 		fmt.Println("Header_binding: ",binding)
// 		c.String(http.StatusOK,"hello world")

// 	})

// 	// Validation

// 	type Customer struct{
// 		Email string `json:"email" binding:"required,email"`
// 		Password string `json:"password" binding:"required,password"`
// 		Role string `json:"role" binding:"required,oneof=BASIC ADMIN"`
// 		StreetAddress string `json:"streetAddress"`
// 		StreetNumber int `json:"streetNumber" binding:"required_with=StreetAddress"`
// 	}

	

// 	if v, ok :=binding.Validator.Engine().(*validator.Validate); ok{
// 		v.RegisterValidation("password", verifyPassword)
// 	}

// 	router.POST("/customers",func(c *gin.Context){
// 		var customer Customer
		
// 		if e:=c.ShouldBindJSON(&customer);e!=nil{
// 			c.String(http.StatusBadRequest,e.Error())
// 			return
// 		}
// 		fmt.Println("Customer: ",customer)
// 		c.String(http.StatusOK,"hello world")

// 	})

// 	// SERVING STATIC FILES
// 	// router.Static("/assets","./assets/hello.txt")
// 	// router.StaticFS("/assets",http.Dir("./assets"))
// 	// router.StaticFile("/hello","./assets/hello.txt")

// 	// Serving HTML FILES
// 	// LoadHTMLGlob
// 	router.LoadHTMLGlob("./templates/*")
// 	router.GET("/home",func(c *gin.Context){
// 		c.HTML(http.StatusOK,"home.html",gin.H{
// 			"title":"Home",
// 			"desc":"description",
// 		})
// 	})
// 	router.GET("/about",func(c *gin.Context){
// 		c.HTML(http.StatusOK,"about.html",nil)
// 	})

// 	// LoadHTMLFiles- различается в том что загуржаемые хтмл файлы
// 	// нужно указывать отдельно
// 	router.LoadHTMLFiles(
// 		"./templates/home.html",
// 		"./templates/about.html",
// 	)
// 	router.GET("/home",func(c *gin.Context){
// 		c.HTML(http.StatusOK,"home.html",gin.H{
// 			"title":"Home",
// 			"desc":"description",
// 		})
// 	})
// 	router.GET("/about",func(c *gin.Context){
// 		c.HTML(http.StatusOK,"about.html",nil)
// 	})

// 	// REDIRECT 
// 	// здесь мы перенаправляем из /home на /about юрлку

// 	router.LoadHTMLFiles(
// 		"./templates/home.html",
// 		"./templates/about.html",
// 	)
// 	router.GET("/home",func(c *gin.Context){
// 		c.Redirect(http.StatusTemporaryRedirect,"/about")
// 	})
// 	router.GET("/about",func(c *gin.Context){
// 		c.HTML(http.StatusOK,"about.html",nil)
// 	})

// 	// MIDDLEWARES
// 	// можем писать функции чтобы обработать код
// 	// прежде чем request дойдет до controller

// 	// 1-common way
// 	// используем middleware  чтобы перехватить и назначить запросу айди если такого нету
// 	router.Use(func(c *gin.Context){
// 		var requestID = c.GetHeader("X-Request-Id")

// 		if len(requestID)==0{
// 			var id= uuid.New().String()
// 			c.Writer.Header().Add("X-Request-Id",id)
// 		}else{
// 			c.Writer.Header().Add("X-Request-Id",requestID)
// 		}
// 		fmt.Println("logging...")
// 	})
// 	router.GET("/ping",func(c *gin.Context){
// 		c.String(http.StatusOK,"pong")
// 	})
	

// 	// 2-common way
// 	// to authorize. we can pass user name and password, then specific gin method, that will return
// 	// a http middleware authentification, then we can access to pages/controllers

// 	var accounts = map[string]string{
// 		"john":"doe",
// 		"foo":"bar",
// 	}
// 	// так отдельно написав и сохранив middleware в переменной можем добавить/применить его только к определенным
// 	// контроллерам, если например хотим чтобы в некоторые страницы был доступ и без аутентификации
// 	// как видите только в /ping url добавлен миддлуэйр
// 	var authMiddleware = gin.BasicAuth(accounts)

// 	// router.Use(gin.BasicAuth(accounts))

// 	router.GET("/ping",authMiddleware,func(c *gin.Context){
// 		c.String(http.StatusOK,"pong")
// 	})

// 	router.GET("/hello",func(c *gin.Context){
// 		c.String(http.StatusOK, "hello-world")
// 	})

// 	// REQUESTS RESPONSE
// 	router.GET("/hello",func(c *gin.Context){
// 		var customer=Customer{
// 			Email: "test@gmail.com",
// 			Role: "BASIC",
// 			StreetAddress: "address",
// 			StreetNumber: 1,
// 		}
// 		c.JSON(http.StatusOK, customer) //превращает структуры в джейсон объект и передает как ответ
// 		// c.String(http.StatusOK, "hello-world") //передает ответ как стринг, но если передать структуру то ошибка будет
// 		// c.HTML() //в виде хтмл  файла, однако сначала надо парсить/рендерить хтмл файлы 
// 	})


	
// 	log.Fatalln(router.Run(address))
// }