package api

import (
  m "github.com/grafana/grafana/pkg/models"
  "github.com/grafana/grafana/pkg/bus"
  "github.com/grafana/grafana/pkg/middleware"

  "github.com/grafana/grafana/pkg/log"

)

func addProcessHelper(cmd m.AddProcessCommand) Response {

  logger := log.New("main")
  logger.Info("Add ProcessForCurrentOrg111",cmd.OrgId)
  query:=m.AddProcessCommand{}

  query.OrgId=cmd.OrgId
  query.ProcessName=cmd.ProcessName
  query.ParentProcessName=cmd.ParentProcessName
  query.UpdatedBy=cmd.UpdatedBy

  if err := bus.Dispatch(&query); err != nil {
    return ApiError(500, "Could not add process to organization", err)
  }


  return ApiSuccess("Process Sucessfully added ")

}

// POST /api/process
func AddProcessToCurrentOrg(c *middleware.Context, cmd m.AddProcessCommand) Response {

  logger := log.New("main")
  logger.Info("Add ProcessForCurrentOrg",c.OrgId)
  cmd.OrgId = c.OrgId

  return addProcessHelper(cmd)
}

// POST /api/process/:orgId
func AddProcess(c *middleware.Context, cmd m.AddProcessCommand) Response {
  cmd.OrgId = c.ParamsInt64(":orgId")
  return addProcessHelper(cmd)
}

func getProcessHelper(OrgId int64) Response {

  query :=m.GetProcessQuery{OrgId:OrgId}
  if err := bus.Dispatch(&query); err != nil {
    return ApiError(500, "Failed to get Process1", err)
  }

  return Json(200, query.Result)

}
// GET /api/org/process
func GetProcessForCurrentOrg(c *middleware.Context) Response {
  logger := log.New("main")
  logger.Info("GetProcessForCurrentOrg12",c.OrgId)

  return getProcessHelper(c.OrgId)
}

// GET /api/orgs/:orgId/process
func GetOrgProcess(c *middleware.Context) Response {
  logger := log.New("main")
  logger.Info("GetProcess",c.ParamsInt64(":orgId"))
  processId:=c.ParamsInt64(":processId")
  return getProcessHelper1(c.OrgId, processId)

}
// DELETE /api/org/users/:userId
func RemoveProcessCurrentOrg(c *middleware.Context) Response {
  processId := c.ParamsInt64(":processId")
  return removeOrgProcessHelper(c.OrgId, processId)
}

// DELETE /api/orgs/:orgId/users/:userId
func RemoveOrgProcess(c *middleware.Context) Response {
  logger := log.New("main")
  logger.Info("GetProcess",c.ParamsInt64(":process_id"))
  processId := c.ParamsInt64(":process_id")
  orgId := c.ParamsInt64(":orgId")
  return removeOrgProcessHelper(orgId, processId)
}

func removeOrgProcessHelper(orgId int64, processId int64) Response {
  logger := log.New("main")
  logger.Info("GetProcess123 %s",orgId)
  cmd := m.DeleteProcessCommand{OrgId: orgId, ProcessId: processId}

  logger.Info("GetProcess456")
  if err := bus.Dispatch(&cmd); err != nil {
    if err == m.ErrLastOrgAdmin {
      return ApiError(400, "Cannot remove last organization admin", nil)
    }
    return ApiError(500, "Failed to remove user from organization", err)
  }

  return ApiSuccess("User removed from organization")
}

func getProcessHelper1(orgId int64, processId int64) Response {
  logger := log.New("main")
  logger.Info("GetProcess123 %s",orgId)
  cmd := m.GetProcessCommand{OrgId: orgId, ProcessId: processId}

  logger.Info("GetProcess456")
  if err := bus.Dispatch(&cmd); err != nil {
    if err == m.ErrLastOrgAdmin {
      return ApiError(400, "Cannot remove last organization admin", nil)
    }
    return ApiError(500, "Failed to remove user from organization", err)
  }

  return ApiSuccess("User removed from organization")
}
