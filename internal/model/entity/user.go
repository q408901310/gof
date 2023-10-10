// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id        uint        `json:"id"        ` //
	Passport  string      `json:"passport"  ` //
	Password  string      `json:"password"  ` //
	Name      string      `json:"name"      ` //
	Avatart   string      `json:"avatart"   ` //
	RoleIds   string      `json:"roleIds"   ` //
	Email     string      `json:"email"     ` //
	Phone     string      `json:"phone"     ` //
	Channel   string      `json:"channel"   ` //
	CreatedAt *gtime.Time `json:"createdAt" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
	DeletedAt *gtime.Time `json:"deletedAt" ` //
}
