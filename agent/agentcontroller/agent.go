package agentcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/agent/agentdto"
	"github.com/noovertime7/gin-mysqlbak/agent/repository"
	"github.com/noovertime7/gin-mysqlbak/middleware"
	"github.com/noovertime7/gin-mysqlbak/public/globalError"
	"github.com/noovertime7/mysqlbak/pkg/log"
	"time"
)

type AgentController struct {
}

var AgentService *repository.AgentService

func AgentRegister(group *gin.RouterGroup) {
	agent := &AgentController{}
	group.GET("/agentlist", agent.GetAgentList)
	group.GET("/service_num_info", agent.GetServiceNumInfo)
	group.POST("/register", agent.Register)
	group.DELETE("/delete_agent", agent.Delete)
	group.PUT("/deregister", agent.DeRegister)
}

func (a *AgentController) GetAgentList(ctx *gin.Context) {
	params := &agentdto.AgentListInput{}
	if err := params.BindValidParams(ctx); err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	data, err := AgentService.GetAgentList(ctx, params)
	if err != nil {
		log.Logger.Error("查询列表出错", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.AgentGetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

func (a *AgentController) GetServiceNumInfo(ctx *gin.Context) {
	data, err := AgentService.GetServiceNumInfo(ctx)
	if err != nil {
		log.Logger.Error("获取服务任务数、完成任务数失败", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.AgentGetError, err))
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

func (a *AgentController) Register(ctx *gin.Context) {
	params := &agentdto.AgentRegisterInput{}
	if err := params.BindValidParams(ctx); err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	if err := AgentService.Register(ctx, params); err != nil {
		log.Logger.Error("注册失败", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.AgentDeRegisterError, err))
		return
	}
	listenChan := make(chan struct{})
	go ListenRegister(ctx, params.ServiceName, listenChan)
	middleware.ResponseSuccess(ctx, "注册成功")
	listenChan <- struct{}{}
}

func (a *AgentController) Delete(ctx *gin.Context) {
	params := &agentdto.AgentDeleteInput{}
	if err := params.BindValidParams(ctx); err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	if err := AgentService.Delete(ctx, params); err != nil {
		log.Logger.Error("删除服务失败", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.AgentDeleteError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "删除成功")
}

func (a *AgentController) DeRegister(ctx *gin.Context) {
	params := &agentdto.AgentDeRegisterInput{}
	if err := params.BindValidParams(ctx); err != nil {
		log.Logger.Error(err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.ParamBindError, err))
		return
	}
	if err := AgentService.DeRegister(ctx, params.ServiceName); err != nil {
		log.Logger.Error("注销失败", err)
		middleware.ResponseError(ctx, globalError.NewGlobalError(globalError.AgentDeRegisterError, err))
		return
	}
	middleware.ResponseSuccess(ctx, "注销成功")
}

func ListenRegister(ctx *gin.Context, serviceName string, ListenChan chan struct{}) {
	log.Logger.Infof("开始监听%s注册状态", serviceName)
	defer close(ListenChan)
	for {
		select {
		case <-ListenChan:
			log.Logger.Infof("收到注册消息,重置计时器,服务名:%s", serviceName)
		//	半小时内未收到注册消息，择认为这个服务离线了
		case <-time.After(time.Minute * 30):
			log.Logger.Warnf("%s注册超时", serviceName)
			//更改状态
			if err := AgentService.DeRegister(ctx, serviceName); err != nil {
				log.Logger.Error("监听器注销服务失败", err)
				return
			}
			log.Logger.Infof("监听器注销服务%s成功", serviceName)
			return
		}
	}
}
