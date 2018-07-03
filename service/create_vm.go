package service

import (
	"context"
	"github.com/aemengo/bosh-containerd-cpi/pb"
	"os"
	"text/template"
	"github.com/satori/go.uuid"
	"path/filepath"
	"fmt"
	"github.com/aemengo/bosh-containerd-cpi/utils"
)

type containerOpts struct {
	ID string
}

func (s *Service) CreateVM(ctx context.Context, req *pb.CreateVMOpts) (*pb.IDParcel, error) {
	var (
		id           = uuid.NewV4().String()
		vmPath       = filepath.Join(s.config.VMDir, id)
		rootFsPath   = filepath.Join(vmPath, "rootfs")
		specPath     = filepath.Join(vmPath, "config.json")
		stemcellPath = filepath.Join(s.config.StemcellDir, req.StemcellID)
	)

	err := os.MkdirAll(rootFsPath, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("failed to make rootfs: %s", err)
	}

	err = utils.RunCommand("/bin/sh", "-c", fmt.Sprintf("tar -O -xzf %s image | tar -xzf - -C %s", stemcellPath, rootFsPath))
	if err != nil {
		return nil, fmt.Errorf("failed to make rootfs: %s", err)
	}

	f, err := os.Create(specPath)
	if err != nil {
		return nil, fmt.Errorf("failed to write container spec: %s", err)
	}
	defer f.Close()

	err = s.writeContainerSpec(f, containerOpts{ID: id})
	if err != nil {
		return nil, fmt.Errorf("failed to write container spec: %s", err)
	}

	err = s.runc.Run(id, vmPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create container: %s", err)
	}

	return &pb.IDParcel{Value: id}, nil
}

func (s *Service) writeContainerSpec(f *os.File, opts containerOpts) error {
	//TODO read config from json
	//t, err := template.New(opts.ID).ParseFiles("./config.json")
	t, err := template.New(opts.ID).Parse(thing)
	if err != nil {
		return err
	}

	return t.Execute(f, opts)
}

var thing = `{
	"ociVersion": "1.0.0",
	"process": {
		"terminal": false,
		"user": {
			"uid": 0,
			"gid": 0
		},
		"args": [
			"sleep",
		    "infinity"
		],
		"env": [
			"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
			"TERM=xterm"
		],
		"cwd": "/",
		"capabilities": {
			"bounding": [
				"CAP_AUDIT_WRITE",
				"CAP_KILL",
				"CAP_NET_BIND_SERVICE"
			],
			"effective": [
				"CAP_AUDIT_WRITE",
				"CAP_KILL",
				"CAP_NET_BIND_SERVICE"
			],
			"inheritable": [
				"CAP_AUDIT_WRITE",
				"CAP_KILL",
				"CAP_NET_BIND_SERVICE"
			],
			"permitted": [
				"CAP_AUDIT_WRITE",
				"CAP_KILL",
				"CAP_NET_BIND_SERVICE"
			],
			"ambient": [
				"CAP_AUDIT_WRITE",
				"CAP_KILL",
				"CAP_NET_BIND_SERVICE"
			]
		},
		"rlimits": [
			{
				"type": "RLIMIT_NOFILE",
				"hard": 1024,
				"soft": 1024
			}
		],
		"noNewPrivileges": true
	},
	"root": {
		"path": "rootfs",
		"readonly": true
	},
	"hostname": "runc",
	"mounts": [
		{
			"destination": "/proc",
			"type": "proc",
			"source": "proc"
		},
		{
			"destination": "/dev",
			"type": "tmpfs",
			"source": "tmpfs",
			"options": [
				"nosuid",
				"strictatime",
				"mode=755",
				"size=65536k"
			]
		},
		{
			"destination": "/dev/pts",
			"type": "devpts",
			"source": "devpts",
			"options": [
				"nosuid",
				"noexec",
				"newinstance",
				"ptmxmode=0666",
				"mode=0620",
				"gid=5"
			]
		},
		{
			"destination": "/dev/shm",
			"type": "tmpfs",
			"source": "shm",
			"options": [
				"nosuid",
				"noexec",
				"nodev",
				"mode=1777",
				"size=65536k"
			]
		},
		{
			"destination": "/dev/mqueue",
			"type": "mqueue",
			"source": "mqueue",
			"options": [
				"nosuid",
				"noexec",
				"nodev"
			]
		},
		{
			"destination": "/sys",
			"type": "sysfs",
			"source": "sysfs",
			"options": [
				"nosuid",
				"noexec",
				"nodev",
				"ro"
			]
		},
		{
			"destination": "/sys/fs/cgroup",
			"type": "cgroup",
			"source": "cgroup",
			"options": [
				"nosuid",
				"noexec",
				"nodev",
				"relatime",
				"ro"
			]
		}
	],
	"linux": {
		"resources": {
			"devices": [
				{
					"allow": false,
					"access": "rwm"
				}
			]
		},
		"namespaces": [
			{
				"type": "pid"
			},
			{
				"type": "network"
			},
			{
				"type": "ipc"
			},
			{
				"type": "uts"
			},
			{
				"type": "mount"
			}
		],
		"maskedPaths": [
			"/proc/kcore",
			"/proc/latency_stats",
			"/proc/timer_list",
			"/proc/timer_stats",
			"/proc/sched_debug",
			"/sys/firmware",
			"/proc/scsi"
		],
		"readonlyPaths": [
			"/proc/asound",
			"/proc/bus",
			"/proc/fs",
			"/proc/irq",
			"/proc/sys",
			"/proc/sysrq-trigger"
		]
	}
}`