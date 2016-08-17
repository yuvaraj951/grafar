package api

import (
  m "github.com/grafana/grafana/pkg/models"
  "github.com/grafana/grafana/pkg/bus"
  "github.com/grafana/grafana/pkg/middleware"

  "github.com/grafana/grafana/pkg/log"

)
// get machines
func getMachineHelper(OrgId int64) Response {

  query :=m.GetMachineQuery{OrgId:OrgId}
  if err := bus.Dispatch(&query); err != nil {
    return ApiError(500, "Failed to get Machines", err)
  }

  return Json(200, query.Result)

}
// GET /api/org/machine
func GetMachineForCurrentOrg(c *middleware.Context) Response {
  logger := log.New("main")
  logger.Info("GetProcessForCurrentOrg12",c.OrgId)

  return getMachineHelper(c.OrgId)
}

// add a machine
func addMachineHelper(cmd m.AddMachineCommand) Response {

  logger := log.New("main")
  logger.Info("Add ProcessForCurrentOrg111",cmd.OrgId)
  query:=m.AddMachineCommand{}

  query.OrgId=cmd.OrgId
  query.ProcessId=cmd.ProcessId
  query.MachineName=cmd.MachineName
  query.Description=cmd.Description
  query.UpdatedBy=cmd.UpdatedBy
  query.Vendor=cmd.Vendor

  if err := bus.Dispatch(&query); err != nil {
    return ApiError(500, "Could not add process to organization", err)
  }


  return ApiSuccess("Process Sucessfully added ")

}

// POST /api/process
func AddMachineToCurrentOrg(c *middleware.Context, cmd m.AddMachineCommand) Response {

  logger := log.New("main")
  logger.Info("Add ProcessForCurrentOrg",c.OrgId)
  cmd.OrgId = c.OrgId
  return addMachineHelper(cmd)
}


//remove the machine
func RemoveMachineCurrentOrg(c *middleware.Context) Response {
  machineId := c.ParamsInt64(":machineId")
  return removeOrgMachineHelper(c.OrgId, machineId)
}
func removeOrgMachineHelper(orgId int64, machineId int64) Response {
  logger := log.New("main")
  logger.Info("GetProcess123 %s",orgId)
  cmd := m.DeleteMachineCommand{OrgId: orgId, MachineId: machineId}

  logger.Info("GetProcess456")
  if err := bus.Dispatch(&cmd); err != nil {
    if err == m.ErrLastOrgAdmin {
      return ApiError(400, "Cannot remove last organization admin", nil)
    }
    return ApiError(500, "Failed to remove user from organization", err)
  }

  return ApiSuccess("User removed from organization")
}
