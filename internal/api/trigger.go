package api

import (
	"encoding/json"
	"github.com/kataras/iris"
	"github.com/lishimeng/iot-link/internal/db/repo"
	"github.com/lishimeng/iot-link/internal/model"
)

func SetupTrigger(app *iris.Application) {
	routTrigger(app)
}

func routTrigger(app *iris.Application) {

	p := app.Party("api/{owner}/application/{appId}/trigger")
	{
		p.Get("", listTriggers)
		p.Get("/{id}", getTrigger)
		//p.Get("/tpl", triggerTemplate)
		p.Post("/", createTrigger)
		p.Post("/del/{id}", delTrigger)
	}
}

type Trigger struct {
	Id         string         `json:"id"`
	AppId      string         `json:"appId"`
	Content    *model.Trigger `json:"content,omitempty"`
	CreateTime int64          `json:"createTime,omitempty"`
	UpdateTime int64          `json:"updateTime,omitempty"`
	Status     int            `json:"status,omitempty"`
}

func listTriggers(ctx iris.Context) {
	res := NewBean()
	appId := ctx.Params().Get("appId")
	triggers, err := repo.GetTriggerByApp(appId)
	if err != nil {
		res.Code = -1
	} else {
		if len(triggers) > 0 {
			list := make([]Trigger, len(triggers))
			for index, item := range triggers {
				t := Trigger{
					Id:         item.Id,
					AppId:      item.AppId,
					CreateTime: item.CreateTime,
					UpdateTime: item.UpdateTime,
					Status:     item.Status,
				}
				list[index] = t
			}
			res.Item = &list
		}
	}

	_, _ = ctx.JSON(&res)
}

func getTrigger(ctx iris.Context) {
	res := NewBean()
	appId := ctx.Params().Get("appId")
	triggerId := ctx.Params().Get("id")
	triggerConfig, err := repo.GetTrigger(triggerId)
	if err != nil {
		res.Code = -1
	} else {
		if triggerConfig.AppId == appId {
			t := Trigger{
				Id:     triggerConfig.Id,
				AppId:  triggerConfig.AppId,
				Status: triggerConfig.Status,
			}
			content := model.Trigger{}
			err = json.Unmarshal([]byte(triggerConfig.Content), &content)
			if err == nil {
				t.Content = &content
			}
			res.Item = &t
		} else {
			res.Code = -1
		}
	}

	_, _ = ctx.JSON(&res)
}

type TriggerForm struct {
	Content string
}

// update or create
func createTrigger(ctx iris.Context) {

	res := NewBean()
	appId := ctx.Params().Get("appId")
	form := TriggerForm{}
	err := ctx.ReadForm(&form)
	if err == nil {

	}
	var triggerConfig repo.TriggerConfig
	triggerConfig, err = repo.CreateTrigger(appId, form.Content)
	if err == nil {
		res.Item = &triggerConfig
	}

	if err != nil {
		res.Code = -1
	}
	_, _ = ctx.JSON(&res)
}

func delTrigger(ctx iris.Context) {

	res := NewBean()
	//appId := ctx.Params().Get("appId")
	// TODO check app_id

	_, _ = ctx.JSON(&res)
}
