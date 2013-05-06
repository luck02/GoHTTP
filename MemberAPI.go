package main

import (
	"code.google.com/p/goweb/goweb"
	"fmt"
	"net/http"
)

type CreateMember struct {
	AccountNumber string
	FirstName     string
	LastName      string
	PhoneNumber   string
	Email         string
}

type TestEntity struct {
	Id   string
	Name string
	Age  int
}

type MemberController struct{}

func (cr *MemberController) Create(cx *goweb.Context) {
	cx.RespondWithData(TestEntity{"1", "Mat", 28})
}
func (cr *MemberController) Delete(id string, cx *goweb.Context) {
	cx.RespondWithOK()
}
func (cr *MemberController) DeleteMany(cx *goweb.Context) {
	cx.RespondWithStatus(http.StatusForbidden)
}
func (cr *MemberController) Read(id string, cx *goweb.Context) {

	if id == "1" {
		cx.RespondWithData(TestEntity{id, "Mat", 28})
	} else if id == "2" {
		cx.RespondWithData(TestEntity{id, "Laurie", 27})
	} else {
		cx.RespondWithNotFound()
	}

}
func (cr *MemberController) ReadMany(cx *goweb.Context) {
	cx.RespondWithData([]TestEntity{TestEntity{"1", "Mat", 28}, TestEntity{"2", "Laurie", 27}})
}
func (cr *MemberController) Update(id string, cx *goweb.Context) {
	cx.RespondWithData(TestEntity{id, "Mat", 28})
}
func (cr *MemberController) UpdateMany(cx *goweb.Context) {
	cx.RespondWithData(TestEntity{"1", "Mat", 28})
}

func main() {

	MemberController := new(MemberController)

	goweb.MapRest("/member", MemberController)

	goweb.MapFunc("/people/{name}/animals/{animal_name}", func(c *goweb.Context) {

		fmt.Fprintf(c.ResponseWriter, "Hey %s, your favourite animal is a %s", c.PathParams["name"], c.PathParams["animal_name"])

	})

	goweb.ListenAndServe(":8080")

}
