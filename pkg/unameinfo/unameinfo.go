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

// Creates the UnameInfo struct
func NewUnameInfo() (*UnameInfo, error) {
	var uts syscall.Utsname

	err := syscall.Uname(&uts)
	if err != nil {
		return nil, err
	}

	return &UnameInfo{
		Sysname:    util.ByteToString(uts.Sysname),
		Nodename:   util.ByteToString(uts.Nodename),
		Release:    util.ByteToString(uts.Release),
		Version:    util.ByteToString(uts.Version),
		Machine:    util.ByteToString(uts.Machine),
		Domainname: util.ByteToString(uts.Domainname),
	}, nil
}
