// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// McmsEmailLogColumns holds the columns for the "mcms_email_log" table.
	McmsEmailLogColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "target", Type: field.TypeString, Comment: "The target email address | 目标邮箱地址"},
		{Name: "content", Type: field.TypeString, Comment: "The content | 发送的内容"},
		{Name: "send_status", Type: field.TypeUint8, Comment: "The send status, 0 unknown 1 success 2 failed | 发送的状态, 0 未知， 1 成功， 2 失败"},
	}
	// McmsEmailLogTable holds the schema information for the "mcms_email_log" table.
	McmsEmailLogTable = &schema.Table{
		Name:       "mcms_email_log",
		Columns:    McmsEmailLogColumns,
		PrimaryKey: []*schema.Column{McmsEmailLogColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		McmsEmailLogTable,
	}
)

func init() {
	McmsEmailLogTable.Annotation = &entsql.Annotation{
		Table: "mcms_email_log",
	}
}
