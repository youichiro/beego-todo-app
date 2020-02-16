package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Todo struct {
	Id 		int
	Name 	string
	Done	bool
}

type TodoModel struct {
}

func init() {
	orm.RegisterModel(new(Todo))
}

func (this *TodoModel) All() ([]*Todo, error) {
	o := orm.NewOrm()
	var todos []*Todo
	_, err := o.QueryTable("todo").All(&todos)
	return todos, err
}

func (this *TodoModel) Save(name string) error {
	if name == "" {
		return fmt.Errorf("empty name\n")
	}
	o := orm.NewOrm()
	todo := new(Todo)
	todo.Name = name
	_, err := o.Insert(todo)
	return err
}

func (this *TodoModel) Find(id int) *Todo {
	o := orm.NewOrm()
	todo := Todo{Id: id}
	err := o.Read(&todo)
	if err != nil {
		return nil
	}
	return &todo
}

func (this *TodoModel) Update(id int, name string, done bool) error {
	o := orm.NewOrm()
	todo := this.Find(id)
	if todo == nil {
		return fmt.Errorf("not found todo")
	}
	todo.Name = name
	todo.Done = done
	_, err := o.Update(todo)
	if err != nil { return err }
	return nil
}

func (this *TodoModel) Delete(id int) error {
	o := orm.NewOrm()
	todo := this.Find(id)
	if todo == nil {
		return fmt.Errorf("not found todo")
	}
	_, err := o.Delete(todo)
	return err
}
