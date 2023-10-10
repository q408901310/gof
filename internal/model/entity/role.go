// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Role is the golang structure for table role.
type Role struct {
	Id        uint        `json:"id"        ` //
	Uid       uint        `json:"uid"       ` //
	Sid       uint        `json:"sid"       ` //
	Name      string      `json:"name"      ` //
	Avatart   string      `json:"avatart"   ` //
	CreatedAt *gtime.Time `json:"createdAt" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
	DeletedAt *gtime.Time `json:"deletedAt" ` //
}
