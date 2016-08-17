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

  bus.AddHandler("sql",GetMachine)
  bus.AddHandler("sql",addMachine)
  bus.AddHandler("sql",RemoveOrgMachine)
}

func GetMachine(query *m.GetMachineQuery) error {

  query.Result = make([]*m.MachineDTO, 0)
  sess := x.Table("machine")
  sess.Where("machine.org_id=?", query.OrgId)
  sess.Cols("machine.org_id","machine.machine_name","machine.machine_id","machine.description","machine.updated_by","machine.vendor")

  err := sess.Find(&query.Result)
  return err
}


func addMachine(cmd *m.AddMachineCommand) error {
  return inTransaction(func(sess *xorm.Session) error {
    // check if user exists
    logger := log.New("main")
    logger.Info("AddProcessForCurrentOrg222",cmd.OrgId)




    entity := m.Machine{
      OrgId:   cmd.OrgId,
      MachineName:    cmd.MachineName,
      Description:cmd.Description,
      ProcessId:cmd.ProcessId,
      UpdatedBy:cmd.UpdatedBy,
      Created: time.Now(),
      Updated: time.Now(),
      Vendor:cmd.Vendor,
    }

    _, err := sess.Insert(&entity)
    return err
  })
}


func RemoveOrgMachine(cmd *m.DeleteMachineCommand) error {
  return inTransaction(func(sess *xorm.Session) error {
    var rawSql = "DELETE FROM machine WHERE org_id=? and machine_id=?"
    _, err := sess.Exec(rawSql, cmd.OrgId, cmd.MachineId)
    if err != nil {
      return err
    }

    return validateOneAdminLeftInOrg(cmd.OrgId, sess)
  })
}
