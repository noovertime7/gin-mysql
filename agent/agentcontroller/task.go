package agentcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/agent/agentdto"
	"github.com/noovertime7/gin-mysqlbak/agent/agentservice"
	"github.com/noovertime7/gin-mysqlbak/middleware"
	"github.com/noovertime7/gin-mysqlbak/public/globalError"
	"github.com/noovertime7/mysqlbak/pkg/log"
)

type AgentTaskController struct {
	service *agentservice.TaskService
}

func AgentTaskRegister(group *gin.RouterGroup) {
	task := &AgentTaskController{service: agentservice.GetClusterTaskService()}
	group.POST("/taskadd", task.TaskAdd)
	group.POST("/task_auto_add", task.TaskAutoAdd)
	group.GET("/tasklist", task.TaskList)
	group.GET("/taskdetail", task.TaskDetail)
	group.DELETE("/taskdelete", task.TaskDelete)
	group.GET("/task_restore", task.TaskRestore)
	group.PUT("/taskupdate", task.TaskUpdate)
}

func (a *AgentTaskController) TaskAdd(ctx *gin.Context) {
	params := &agentdto.TaskAddInput{}
	if err := params.BindValidParams(ctx); err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	if err := a.service.TaskAdd(ctx, params); err != nil {
		log.Logger.Error("agent添加任务失败", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.TaskAddError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "添加成功")
	log.Logger.Info("agent添加任务成功")
}

func (a *AgentTaskController) TaskAutoAdd(ctx *gin.Context) {
	params := &agentdto.TaskAutoAddInput{}
	if err := params.BindValidParams(ctx); err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	if err := a.service.TaskAutoAdd(ctx, params); err != nil {
		log.Logger.Error("agent自动创建任务成功失败", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.TaskAddError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "自动创建任务成功")
	log.Logger.Info("agent自动创建任务成功成功")
}

func (a *AgentTaskController) TaskList(ctx *gin.Context) {
	params := &agentdto.TaskListInput{}
	if err := params.BindValidParams(ctx); err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := a.service.TaskList(ctx, params)
	if err != nil {
		log.Logger.Error("agent查询任务列表失败", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
	log.Logger.Info("agent查询任务列表成功")
}

func (a *AgentTaskController) TaskDetail(ctx *gin.Context) {
	params := &agentdto.TaskIDInput{}
	if err := params.BindValidParams(ctx); err != nil {
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := a.service.TaskDetail(ctx, params)
	if err != nil {
		log.Logger.Error("agent查询任务详情失败", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
	log.Logger.Info("agent查询任务详情成功")
}

func (a *AgentTaskController) TaskDelete(ctx *gin.Context) {
	params := &agentdto.TaskDeleteInput{}
	if err := params.BindValidParams(ctx); err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	if err := a.service.TaskDelete(ctx, params); err != nil {
		log.Logger.Error("删除失败", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.TaskDeleteError, err))
	}
	middleware.ResponseSuccess(ctx, "删除成功")
	log.Logger.Info("agent删除任务成功")
}

func (a *AgentTaskController) TaskRestore(ctx *gin.Context) {
	params := &agentdto.TaskDeleteInput{}
	if err := params.BindValidParams(ctx); err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	if err := a.service.TaskRestore(ctx, params); err != nil {
		log.Logger.Error("还原失败", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.TaskDeleteError, err))
	}
	middleware.ResponseSuccess(ctx, "还原成功")
	log.Logger.Info("agent还原删除成功")
}

func (a *AgentTaskController) TaskUpdate(ctx *gin.Context) {
	params := &agentdto.TaskUpdateInput{}
	if err := params.BindValidParams(ctx); err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	if err := a.service.TaskUpdate(ctx, params); err != nil {
		log.Logger.Error("更新失败")
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.TaskUpdateError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "更新成功")
	log.Logger.Info("agent更新任务成功")
}
