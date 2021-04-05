package main

import (
	"axe/dao/mysql"
	"axe/gl"
	"axe/v1/system"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("conf")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	pflag.String("app.cacheDir", "./cache/", "cache directory")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	//fmt.Println(viper.GetString("app.cacheDir"))

	mysql.InitMySQL()

	gl.App = iris.Default()
	gl.App.Post("/login", system.User.Login)
	gl.App.Post("/login_wx", system.User.Login)
	gl.App.Post("/login_wx", system.User.Logout)

	usersAPI := gl.App.Party("/users")
	{
		usersAPI.Get("/{uid:int64}", system.User.GetUser)
		usersAPI.Get("/your/groups", system.Group.GetGroupsByUserId)
		usersAPI.Get("/your/activities", system.Activity.GetActivitiesByUserId)
	}

	groupsAPI := gl.App.Party("/groups")
	{
		groupsAPI.Get("/{gid:int}", system.Group.GetGroupById)
		groupsAPI.Get("/", system.Group.GetGroups)
		groupsAPI.Get("/{gid:int}/pending", system.Group.GetApplyList)
		groupsAPI.Get("/{gid:int}/activities", system.Group.GetActivitiesByGroupId)

		groupsAPI.Post("/", system.Group.Create)
		groupsAPI.Post("/{gid:int}/gl.Apply", system.Group.Apply)
		groupsAPI.Post("/{gid:int}/gl.Approve", system.Group.Approve)
		groupsAPI.Post("/{gid:int}/promote/:mid", system.Group.Promote)
		groupsAPI.Post("/{gid:int}/transfer/:mid", system.Group.Transfer)
		groupsAPI.Post("/{gid:int}/remove/:mid", system.Group.Remove)
		groupsAPI.Post("/{gid:int}/quit", system.Group.Quit)

		groupsAPI.Put("/{gid:int}", system.Group.Update)
	}

	actsAPI := gl.App.Party("/activities")
	{
		actsAPI.Get("/{aid:int64}", system.Activity.GetActivityById)
		actsAPI.Get("/", system.Activity.GetActivities)

		actsAPI.Post("/", system.Activity.Create)
		actsAPI.Post("/{aid:int64}/end", system.Activity.End)
		actsAPI.Post("/{aid:int64}/gl.Apply", system.Activity.Apply)
		actsAPI.Post("/{aid:int64}/cancel", system.Activity.Cancel)
		actsAPI.Post("/{aid:int64}/remove/:index", system.Activity.Remove)

		actsAPI.Put("/{aid:int64}", system.Activity.Update)
	}

	booksAPI := gl.App.Party("/books")
	{
		booksAPI.Use(iris.Compression)

		// GET: http://localhost:8080/books
		booksAPI.Get("/", list)
		// POST: http://localhost:8080/books
		booksAPI.Post("/", create)
	}

	gl.App.Listen(viper.GetString("httpServer.addr"))
}

// Book example.
type Book struct {
	Title string `json:"title"`
}

func list(ctx iris.Context) {
	books := []Book{
		{"Mastering Concurrency in Go"},
		{"Go Design Patterns"},
		{"Black Hat Go"},
	}

	ctx.JSON(books)
	// TIP: negotiate the response between server's prioritizes
	// and client's requirements, instead of ctx.JSON:
	// ctx.Negotiation().JSON().MsgPack().Protobuf()
	// ctx.Negotiate(books)
}

func create(ctx iris.Context) {
	var b Book
	err := ctx.ReadJSON(&b)
	// TIP: use ctx.ReadBody(&b) to bind
	// any type of incoming data instead.
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("Book creation failure").DetailErr(err))
		// TIP: use ctx.StopWithError(code, err) when only
		// plain text responses are expected on errors.
		return
	}

	println("Received Book: " + b.Title)

	ctx.StatusCode(iris.StatusCreated)
}
