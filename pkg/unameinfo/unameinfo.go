package unameinfo

import (
	"syscall"

	"github.com/pirxthepilot/go-uname/pkg/util"
)

type UnameInfo struct {
	Sysname    string
	Nodename   string
	Release    string
	Version    string
	Machine    string
	Domainname string
}

// Populates the UnameInfo struct
func (u *UnameInfo) Create() error {
	var uts syscall.Utsname

	err := syscall.Uname(&uts)
	if err != nil {
		return err
	}

	u.Sysname = util.ByteToString(uts.Sysname)
	u.Nodename = util.ByteToString(uts.Nodename)
	u.Release = util.ByteToString(uts.Release)
	u.Version = util.ByteToString(uts.Version)
	u.Machine = util.ByteToString(uts.Machine)
	u.Domainname = util.ByteToString(uts.Domainname)

	return nil
}
