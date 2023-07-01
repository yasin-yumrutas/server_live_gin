package main

import (
	"fmt"
	"net/http"
	"strconv"

	"net/mail"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() //default server router config

	router.GET("/", getTestOne)                             //localhost:9000/
	router.GET("/getYasinTest2", getTestTwo)                //localhost:9000/YasinTest2
	router.GET("/getYasinTest3/:text1", getTestThree)       //localhost:9000/YasinTest3/Yasin test değişken
	router.GET("/getYasinTest4/:text1/:text2", getTestFour) //localhost:9000/YasinTest4/Yasin test/One
	router.GET("/getYasinTest5", getTestFive)

	router.POST("/", posTestOne)                            //localhost:9000/
	router.POST("/postYasinTest2", filterTest, postTestTwo) //localhost:9000     body => {"name":"yasin", "id": 1}

	router.Run(":9000")
}

func getTestOne(d *gin.Context) {
	d.String(http.StatusOK, "Get  getTestOne")
}

func getTestTwo(d *gin.Context) {
	d.String(http.StatusOK, "Get getTestTwo")
}

func getTestThree(d *gin.Context) {

	txt1 := d.Param("text1")
	d.String(http.StatusOK, txt1)
}

func getTestFour(d *gin.Context) {
	txt1 := d.Param("text1")
	txt2 := d.Param("text2")
	d.String(http.StatusOK, txt1+txt2)
}

func getTestFive(d *gin.Context) {

	name5 := d.Query("name")
	yas5 := d.Query("yas")
	yas5Int, err := strconv.Atoi(yas5)
	if err != nil {
		fmt.Println("Hata: ", err)
		return
	}

	//d.String(http.StatusOK, name5 + yas5)

	d.JSON(http.StatusOK, struct {
		Name string
		Yas  int
	}{name5, yas5Int})

}

func posTestOne(d *gin.Context) {
	d.String(http.StatusOK, "Hello Post")
}

type kisi struct {
	Name string
	Mail string
}

type response_1 struct {
	KayitStatus bool
	Mail        string
	Msj         string
}

func postTestTwo(d *gin.Context) {

	var kisi_1 kisi

	filterGelenDeger, err := d.Get("aa")

	if !err {
		d.JSON(http.StatusBadRequest, "Bad request")

	}

	kisi_1 = filterGelenDeger.(kisi)

	d.JSON(http.StatusOK, response_1{true, kisi_1.Mail, "kayıt başarılı"})

}

func filterTest(d *gin.Context) {

	var kisi_1 kisi

	var responsefilt response_1

	d.BindJSON(&kisi_1)

	_, err := mail.ParseAddress(kisi_1.Mail)

	if err != nil {
		d.Abort() /*herhangi istenilmiyen koşullar oluşunca sonraki katmanlara geçişi durdurur ve default false
		fakat çalıştırılırsa true döner ve sonraki steplere geçmesi engellenir*/
		responsefilt.KayitStatus = false
		responsefilt.Mail = kisi_1.Name
		responsefilt.Msj = "lütfen geçerli bir mail giriniz"

		d.JSON(http.StatusBadRequest, responsefilt)
		return
	}

	if !d.IsAborted() {
		/* d.IsAborted() bu func ise d.Abort değişkenini okur ve false ise sonraki adımlara geçmesi için d.Next()
		   çalışır */
		d.Set("aa", kisi_1) /*herhangi bir sorun yoksa koşullar uygunsa sonraki func işlenmesi için
		  box içine kaydedilir  */
		d.Next()
	}
}
