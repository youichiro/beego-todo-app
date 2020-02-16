package controllers

import (
	"dev/bee-todo-app/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
)

type TodoController struct {
	beego.Controller
	model models.TodoModel
}

func (this *TodoController) OutputContext(status int, message string) {
	this.Ctx.Output.SetStatus(status)
	this.Ctx.Output.Body([]byte(message + "\n"))
}

func (this *TodoController) Index() {
	todos, err := this.model.All()
	if err != nil {
		fmt.Println(err)
	}
	this.Data["todos"] = todos
	this.TplName = "index.html"
}

func (this *TodoController) ListTodos() {
	todos, err := this.model.All()
	if err != nil {
		fmt.Println(err)
	}
	this.Data["json"] = todos
	this.ServeJSON()
}

func (this *TodoController) NewTodo() {
	req := struct{ Name string }{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if err != nil {
		this.OutputContext(400, "empty name"); return
	}
	err = this.model.Save(req.Name)
	if err != nil {
		this.OutputContext(400, err.Error()); return
	}
	this.OutputContext(200, "success new")
}

func (this *TodoController) GetTodo() {
	ids := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(ids)
	todo := this.model.Find(id)
	if todo == nil {
		this.OutputContext(400, "not found todo"); return
	}
	this.Data["json"] = todo
	this.ServeJSON()
}

func (this *TodoController) UpdateTodo() {
	// request from GET
	ids := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(ids)
	req := struct {
		Name string
		Done string
	}{}
	// request from POST
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if err != nil {
		this.OutputContext(400, err.Error()); return
	}
	// edit values
	todo := this.model.Find(id)
	if req.Name != "" {
		todo.Name = req.Name
	}
	if req.Done == "true" {
		todo.Done = true
	} else if req.Done == "false" {
		todo.Done = false
	}
	// update
	err = this.model.Update(id, todo.Name, todo.Done)
	if err != nil {
		this.OutputContext(400, err.Error()); return
	}
	this.OutputContext(200, "success update")
}

func (this *TodoController) DeleteTodo() {
	ids := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(ids)
	err := this.model.Delete(id)
	if err != nil {
		this.OutputContext(400, err.Error()); return
	}
	this.OutputContext(200, "success delete")
}
