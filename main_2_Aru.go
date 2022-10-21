package main
import(
	"net/http"

    "github.com/gin-gonic/gin"
)

type Article struct{
	Id uint16
	Name,Anons,FullText,Image string
}
type User struct{
	Id uint16
	Name, Password string
}

var posts=[]Article{
	{
		Id: 1,
		Name: "AAAAAAAAa",
		Anons:"shdjdkldldfnbffkfkf",
		FullText: "gfxcjvkbknlnm;gchvkbjnlkbcfxhgh",
		Image: "https://yandex.kz/images/search?pos=2&from=tabbar&img_url=http%3A%2F%2Fimg-fotki.yandex.ru%2Fget%2F68556%2F328849705.1%2F0_225da4_ee23b47b_orig.jpg&text=photo+of+book&rpt=simage&lr=162",
	},{
		Id: 2,
		Name: "BBBBBBBBBB",
		Anons:"sbbsnsjddhdbdb",
		FullText: "gfxcjvkbknlnm;gchvkbjnlkbcfxhgh",
		Image: "https://yandex.kz/images/search?pos=2&from=tabbar&img_url=http%3A%2F%2Fimg-fotki.yandex.ru%2Fget%2F68556%2F328849705.1%2F0_225da4_ee23b47b_orig.jpg&text=photo+of+book&rpt=simage&lr=162",
	},
}
// var showPost=Article{}

var users=[]User{
	{
		Id: 1,
		Name: "Aruzhan",
		Password: "123456",
	},{
		Id: 2,
		Name: "Aidana",
		Password: "123456",
	},
}

// type Customer struct{
// 	Id uint16
// 	Name, Password, Surname string
// 	Age int
// 	Sex, CardNumber string
// }
// var customers=[]Customer{}

func getUsers(c *gin.Context){
	c.IndentedJSON(http.StatusOK, users)

}
func getArticles(c *gin.Context){
	c.IndentedJSON(http.StatusOK, posts)
}
func main(){
	router:=gin.Default()
	router.GET("/users",getUsers)
	router.GET("/articles",getArticles)
	router.Run("localhost:8080")
}