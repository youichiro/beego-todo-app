package routers

import (
	"dev/bee-todo-app/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.TodoController{}, "get:Index")
    beego.Router("/api/todo", &controllers.TodoController{}, "get:ListTodos;post:NewTodo")
    beego.Router("/api/todo/:id:int", &controllers.TodoController{}, "get:GetTodo;post:UpdateTodo")
    beego.Router("/api/todo/delete/:id", &controllers.TodoController{}, "get:DeleteTodo")
}
