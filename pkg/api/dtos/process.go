package dtos

type AddProcessForm struct {
  ProcessName         string     `json:"processName"`
  ParentProcessName     string     `json:"parentProcessName" binding:"Required"`
  UpdatedBy   string       `json:"updatedBy"`
}
