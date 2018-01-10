package main

//import (
//	"github.com/astaxie/beego"
//)
//
//type HomeController struct {
//	beego.Controller
//}
//func (this *HomeController) Get() {
//	this.Ctx.WriteString("Hello world!")
//}
//func main()  {
//	beego.Router("/",&HomeController{})
//	beego.Run()
//}

import (
	//"io"
	//"log"

)

func main() {

}




//func main()  {
//	http.HandleFunc("/",sayHello)
//
//	err := http.ListenAndServe(":8080",nil)
//	if err !=nil{
//		log.Fatal(err)
//	}
//}
//func sayHello(w http.ResponseWriter, r *http.Request) {
//	io.WriteString(w,"Hello World,One")
//}




//func main(){
//	routers := gin.Default()
//
//	routers.GET("/user/:name/*action", func(c *gin.Context) {
//		name := c.Param("name")
//		action := c.Param("action")
//		message := name + " is " + action
//		c.String(http.StatusOK, message)
//	})
//}


//func main()  {
//	max := http.NewServeMux()
//	max.Handle("/",&myhandler{})
//	max.HandleFunc("/hello",sayHello)
//		err := http.ListenAndServe(":8080",max)
//		if err !=nil{
//			log.Fatal(err)
//		}
//}

//type myhandler struct {}
//
//func (*myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	io.WriteString(w,"URL:"+r.URL.String())
//}
//func sayHello(w http.ResponseWriter, r *http.Request) {
//	io.WriteString(w,"Hello World,Two")
//}