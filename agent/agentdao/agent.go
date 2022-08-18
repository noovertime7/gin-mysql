package agentdao

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/gin-mysqlbak/agent/agentdto"
	"gorm.io/gorm"
	"time"
)

type AgentDB struct {
	Id          int       `json:"id" gorm:"primary_key" description:"自增主键"`
	ServiceName string    `json:"service_name"  gorm:"column:service_name"`
	Content     string    `json:"content"  gorm:"column:content"`
	Address     string    `json:"address"  gorm:"column:address"`
	AgentStatus int       `json:"agent_status"  gorm:"column:agent_status"`
	LastTime    time.Time `json:"last_time"  gorm:"column:last_time"`
	CreateAt    time.Time `json:"create_at"  gorm:"column:create_at"`
	IsDeleted   int       `json:"is_deleted"  gorm:"column:is_deleted"`
}

func (a *AgentDB) TableName() string {
	return "t_agent"
}

func (a *AgentDB) Find(c *gin.Context, tx *gorm.DB, search *AgentDB) (*AgentDB, error) {
	out := &AgentDB{}
	err := tx.WithContext(c).Where(search).Find(out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (a *AgentDB) Save(c *gin.Context, tx *gorm.DB) error {
	return tx.WithContext(c).Save(a).Error
}

func (a *AgentDB) Update(c *gin.Context, tx *gorm.DB) error {
	return tx.WithContext(c).Table(a.TableName()).Where("id = ?", a.Id).Updates(map[string]interface{}{
		"service_name": a.ServiceName,
		"content":      a.Content,
		"address":      a.Address,
		"agent_status": a.AgentStatus,
		"last_time":    time.Now(),
		"is_deleted":   a.IsDeleted,
	}).Error
}

func (a *AgentDB) UpdateStatus(c *gin.Context, tx *gorm.DB) error {
	if a.Id == 0 {
		return errors.New("ID 不存在")
	}
	return tx.WithContext(c).Table(a.TableName()).Where("id = ?", a.Id).Updates(map[string]interface{}{
		"agent_status": a.AgentStatus,
	}).Error
}

func (s *AgentDB) PageList(c *gin.Context, tx *gorm.DB, params *agentdto.AgentListInput) ([]AgentDB, int, error) {
	var total int64 = 0
	var list []AgentDB
	offset := (params.PageNo - 1) * params.PageSize
	query := tx.WithContext(c)
	query = query.Table(s.TableName()).Where("is_deleted=0")
	if params.Info != "" {
		query = query.Where("( service_name like ?)", "%"+params.Info+"%")
	}
	if err := query.Limit(params.PageSize).Offset(offset).Order("id desc").Find(&list).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	query.Find(&list).Count(&total)
	return list, int(total), nil
}
