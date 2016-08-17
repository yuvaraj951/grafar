package sqlstore

import (
  "github.com/grafana/grafana/pkg/bus"
  m "github.com/grafana/grafana/pkg/models"
  "github.com/go-xorm/xorm"
  "time"
  "github.com/grafana/grafana/pkg/log"
 // "fmt"
  //"github.com/Unknwon/bra/cmd"
)

func init()  {

  bus.AddHandler("sql",addProcess)
  bus.AddHandler("sql",GetProcess)
  bus.AddHandler("sql",RemoveOrgProcess)
  bus.AddHandler("sql",GetOrgProcess)

}

func addProcess(cmd *m.AddProcessCommand) error {
  return inTransaction(func(sess *xorm.Session) error {
    // check if user exists
    logger := log.New("main")
    logger.Info("AddProcessForCurrentOrg")

    entity := m.Process{
      OrgId:   cmd.OrgId,

      ProcessName:    cmd.ProcessName,
      ParentProcessName:cmd.ParentProcessName,
      UpdatedBy:cmd.UpdatedBy,
      Created: time.Now(),
      Updated: time.Now(),
    }

    _, err := sess.Insert(&entity)
    return err
  })
}

func GetProcess(query *m.GetProcessQuery) error {

  query.Result = make([]*m.ProcessDTO, 0)
  sess := x.Table("process")
  sess.Where("process.org_id=?", query.OrgId)
  sess.Cols("process.org_id","process.process_name","process.parent_process_name","process.updated_by","process.process_id")

  err := sess.Find(&query.Result)
  return err
}


func RemoveOrgProcess(cmd *m.DeleteProcessCommand) error {
  return inTransaction(func(sess *xorm.Session) error {
    var rawSql = "DELETE FROM process WHERE org_id=? and process_id=?"
    _, err := sess.Exec(rawSql, cmd.OrgId, cmd.ProcessId)
    if err != nil {
      return err
    }

    return validateOneAdminLeftInOrg(cmd.OrgId, sess)
  })
}
func GetOrgProcess(cmd *m.GetProcessCommand) error {
  return inTransaction(func(sess *xorm.Session) error {
    var rawSql = "SELECT FROM process WHERE org_id=? and process_id=?"
    _, err := sess.Exec(rawSql, cmd.OrgId, cmd.ProcessId)
    if err != nil {
      return err
    }

    return validateOneAdminLeftInOrg(cmd.OrgId, sess)
  })
}


