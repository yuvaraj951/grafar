package migrations

import  ."github.com/grafana/grafana/pkg/services/sqlstore/migrator"

func addProcessMigrations(mg *Migrator)  {

  processV1:= Table{
    Name: "process",
    Columns: []*Column{
      {Name: "proces_id", Type: DB_BigInt, IsPrimaryKey: true, IsAutoIncrement: true},
      {Name: "process_name", Type: DB_NVarchar, Nullable: false},
      {Name: "org_id", Type: DB_Int,  Nullable: false},
      {Name: "parent_process_name", Type: DB_NVarchar,  Nullable: false},
      {Name: "created", Type: DB_DateTime},
      {Name: "updated", Type: DB_DateTime},
      {Name: "updated_by", Type: DB_NVarchar, Length: 255, Nullable: true},
    },
    Indices: []*Index{
      {Cols: []string{"proces_id"}, Type: IndexType},
      {Cols: []string{"org_id"}, Type: IndexType},
      {Cols: []string{"parent_process_name"}, Type: IndexType},
      {Cols: []string{"updated_by"}, Type: IndexType},
    },

  }
  mg.AddMigration("create process  table v1-7", NewAddTableMigration(processV1))
  addTableIndicesMigrations(mg, "v1-7", processV1)



}
