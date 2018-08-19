package service

import (
	"context"
	"fmt"
	"github.com/aemengo/bosh-runc-cpi/pb"
	"github.com/aemengo/bosh-runc-cpi/utils"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)

func (s *Service) CreateVM(ctx context.Context, req *pb.CreateVMOpts) (*pb.IDParcel, error) {
	var (
		id                = uuid.NewV4().String()
		diskPath          = ""
		persistentDiskDir = "/persistent-disk"
		vmPath            = filepath.Join(s.config.VMDir, id)
		rootFsPath        = filepath.Join(vmPath, "rootfs")
		workDirPath       = filepath.Join(vmPath, "workdir")
		upperDirPath      = filepath.Join(vmPath, "upperdir")
		specPath          = filepath.Join(vmPath, "config.json")
		pidPath           = filepath.Join(vmPath, "pid")
		stemcellIDPath    = filepath.Join(vmPath, "stemcell-id")
		stemcellPath      = filepath.Join(s.config.StemcellDir, req.StemcellID)
		agentSettings     = attachVMID(req.AgentSettings, id)
		agentSettingsPath = filepath.Join(vmPath, "warden-cpi-agent-env.json")
	)

	err := utils.MkdirAll(rootFsPath, upperDirPath, workDirPath)
	if err != nil {
		return nil, err
	}

	err = utils.RunCommand("mount",
		"-t", "overlay",
		"-o", fmt.Sprintf("lowerdir=%s,upperdir=%s,workdir=%s", stemcellPath, upperDirPath, workDirPath),
		"overlay",
		rootFsPath,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to make rootfs: %s", err)
	}

	if req.DiskID != "" {
		agentSettings = attachPersistentDisk(agentSettings, req.DiskID, persistentDiskDir)
		diskPath = filepath.Join(s.config.DiskDir, req.DiskID)
	}

	err = ioutil.WriteFile(agentSettingsPath, agentSettings, 0666)
	if err != nil {
		return nil, fmt.Errorf("failed to write agent settings: %s", err)
	}

	f, err := os.Create(specPath)
	if err != nil {
		return nil, fmt.Errorf("failed to write container spec: %s", err)
	}
	defer f.Close()

	err = s.writeContainerSpec(f, containerOpts{
		ID:                id,
		Lowerdir:          stemcellPath,
		Upperdir:          upperDirPath,
		Workdir:           workDirPath,
		AgentSettingsPath: agentSettingsPath,
		DiskPath:          diskPath,
		PersistentDiskDir: persistentDiskDir,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to write container spec: %s", err)
	}

	err = s.startContainer(ctx, id, vmPath, pidPath, agentSettings, true)
	if err != nil {
		return nil, err
	}

	if req.DiskID != "" {
		saveDiskState(vmPath, req.DiskID)
	}

	err = ioutil.WriteFile(stemcellIDPath, []byte(req.StemcellID), 0666)
	if err != nil {
		return nil, err
	}

	return &pb.IDParcel{Value: id}, nil
}

type containerOpts struct {
	ID                string
	Lowerdir          string
	Upperdir          string
	Workdir           string
	AgentSettingsPath string
	DiskPath          string
	PersistentDiskDir string
}

func (s *Service) writeContainerSpec(f *os.File, opts containerOpts) error {
	t, err := template.New(opts.ID).Parse(containerSpec)
	if err != nil {
		return err
	}

	return t.Execute(f, opts)
}

var containerSpec = `{
	"ociVersion": "1.0.0",
	"process": {
		"terminal": false,
		"user": {
			"uid": 0,
			"gid": 0
		},
		"args": [
		    "bash",
		    "-c",
		    "umount /var/vcap/data/root_log; chmod 0777 /var/vcap/data; exec env -i /usr/sbin/runsvdir-start"
		],
		"env": [
			"PATH=/var/vcap/bosh/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
			"TERM=xterm"
		],
		"cwd": "/",
		"capabilities": {
			"bounding": [
              "CAP_AUDIT_CONTROL",
              "CAP_AUDIT_READ",
              "CAP_AUDIT_WRITE",
              "CAP_BLOCK_SUSPEND",
              "CAP_CHOWN",
              "CAP_DAC_OVERRIDE",
              "CAP_DAC_READ_SEARCH",
              "CAP_FOWNER",
              "CAP_FSETID",
              "CAP_IPC_LOCK",
              "CAP_IPC_OWNER",
              "CAP_KILL",
              "CAP_LEASE",
              "CAP_LINUX_IMMUTABLE",
              "CAP_MAC_ADMIN",
              "CAP_MAC_OVERRIDE",
              "CAP_MKNOD",
              "CAP_NET_ADMIN",
              "CAP_NET_BIND_SERVICE",
              "CAP_NET_BROADCAST",
              "CAP_NET_RAW",
              "CAP_SETGID",
              "CAP_SETFCAP",
              "CAP_SETPCAP",
              "CAP_SETUID",
              "CAP_SYS_ADMIN",
              "CAP_SYS_BOOT",
              "CAP_SYS_CHROOT",
              "CAP_SYS_MODULE",
              "CAP_SYS_NICE",
              "CAP_SYS_PACCT",
              "CAP_SYS_PTRACE",
              "CAP_SYS_RAWIO",
              "CAP_SYS_RESOURCE",
              "CAP_SYS_TIME",
              "CAP_SYS_TTY_CONFIG",
              "CAP_SYSLOG",
              "CAP_WAKE_ALARM"
			],
			"effective": [
              "CAP_AUDIT_CONTROL",
              "CAP_AUDIT_READ",
              "CAP_AUDIT_WRITE",
              "CAP_BLOCK_SUSPEND",
              "CAP_CHOWN",
              "CAP_DAC_OVERRIDE",
              "CAP_DAC_READ_SEARCH",
              "CAP_FOWNER",
              "CAP_FSETID",
              "CAP_IPC_LOCK",
              "CAP_IPC_OWNER",
              "CAP_KILL",
              "CAP_LEASE",
              "CAP_LINUX_IMMUTABLE",
              "CAP_MAC_ADMIN",
              "CAP_MAC_OVERRIDE",
              "CAP_MKNOD",
              "CAP_NET_ADMIN",
              "CAP_NET_BIND_SERVICE",
              "CAP_NET_BROADCAST",
              "CAP_NET_RAW",
              "CAP_SETGID",
              "CAP_SETFCAP",
              "CAP_SETPCAP",
              "CAP_SETUID",
              "CAP_SYS_ADMIN",
              "CAP_SYS_BOOT",
              "CAP_SYS_CHROOT",
              "CAP_SYS_MODULE",
              "CAP_SYS_NICE",
              "CAP_SYS_PACCT",
              "CAP_SYS_PTRACE",
              "CAP_SYS_RAWIO",
              "CAP_SYS_RESOURCE",
              "CAP_SYS_TIME",
              "CAP_SYS_TTY_CONFIG",
              "CAP_SYSLOG",
              "CAP_WAKE_ALARM"
			],
			"inheritable": [
              "CAP_AUDIT_CONTROL",
              "CAP_AUDIT_READ",
              "CAP_AUDIT_WRITE",
              "CAP_BLOCK_SUSPEND",
              "CAP_CHOWN",
              "CAP_DAC_OVERRIDE",
              "CAP_DAC_READ_SEARCH",
              "CAP_FOWNER",
              "CAP_FSETID",
              "CAP_IPC_LOCK",
              "CAP_IPC_OWNER",
              "CAP_KILL",
              "CAP_LEASE",
              "CAP_LINUX_IMMUTABLE",
              "CAP_MAC_ADMIN",
              "CAP_MAC_OVERRIDE",
              "CAP_MKNOD",
              "CAP_NET_ADMIN",
              "CAP_NET_BIND_SERVICE",
              "CAP_NET_BROADCAST",
              "CAP_NET_RAW",
              "CAP_SETGID",
              "CAP_SETFCAP",
              "CAP_SETPCAP",
              "CAP_SETUID",
              "CAP_SYS_ADMIN",
              "CAP_SYS_BOOT",
              "CAP_SYS_CHROOT",
              "CAP_SYS_MODULE",
              "CAP_SYS_NICE",
              "CAP_SYS_PACCT",
              "CAP_SYS_PTRACE",
              "CAP_SYS_RAWIO",
              "CAP_SYS_RESOURCE",
              "CAP_SYS_TIME",
              "CAP_SYS_TTY_CONFIG",
              "CAP_SYSLOG",
              "CAP_WAKE_ALARM"
			],
			"permitted": [
              "CAP_AUDIT_CONTROL",
              "CAP_AUDIT_READ",
              "CAP_AUDIT_WRITE",
              "CAP_BLOCK_SUSPEND",
              "CAP_CHOWN",
              "CAP_DAC_OVERRIDE",
              "CAP_DAC_READ_SEARCH",
              "CAP_FOWNER",
              "CAP_FSETID",
              "CAP_IPC_LOCK",
              "CAP_IPC_OWNER",
              "CAP_KILL",
              "CAP_LEASE",
              "CAP_LINUX_IMMUTABLE",
              "CAP_MAC_ADMIN",
              "CAP_MAC_OVERRIDE",
              "CAP_MKNOD",
              "CAP_NET_ADMIN",
              "CAP_NET_BIND_SERVICE",
              "CAP_NET_BROADCAST",
              "CAP_NET_RAW",
              "CAP_SETGID",
              "CAP_SETFCAP",
              "CAP_SETPCAP",
              "CAP_SETUID",
              "CAP_SYS_ADMIN",
              "CAP_SYS_BOOT",
              "CAP_SYS_CHROOT",
              "CAP_SYS_MODULE",
              "CAP_SYS_NICE",
              "CAP_SYS_PACCT",
              "CAP_SYS_PTRACE",
              "CAP_SYS_RAWIO",
              "CAP_SYS_RESOURCE",
              "CAP_SYS_TIME",
              "CAP_SYS_TTY_CONFIG",
              "CAP_SYSLOG",
              "CAP_WAKE_ALARM"
			],
			"ambient": [
              "CAP_AUDIT_CONTROL",
              "CAP_AUDIT_READ",
              "CAP_AUDIT_WRITE",
              "CAP_BLOCK_SUSPEND",
              "CAP_CHOWN",
              "CAP_DAC_OVERRIDE",
              "CAP_DAC_READ_SEARCH",
              "CAP_FOWNER",
              "CAP_FSETID",
              "CAP_IPC_LOCK",
              "CAP_IPC_OWNER",
              "CAP_KILL",
              "CAP_LEASE",
              "CAP_LINUX_IMMUTABLE",
              "CAP_MAC_ADMIN",
              "CAP_MAC_OVERRIDE",
              "CAP_MKNOD",
              "CAP_NET_ADMIN",
              "CAP_NET_BIND_SERVICE",
              "CAP_NET_BROADCAST",
              "CAP_NET_RAW",
              "CAP_SETGID",
              "CAP_SETFCAP",
              "CAP_SETPCAP",
              "CAP_SETUID",
              "CAP_SYS_ADMIN",
              "CAP_SYS_BOOT",
              "CAP_SYS_CHROOT",
              "CAP_SYS_MODULE",
              "CAP_SYS_NICE",
              "CAP_SYS_PACCT",
              "CAP_SYS_PTRACE",
              "CAP_SYS_RAWIO",
              "CAP_SYS_RESOURCE",
              "CAP_SYS_TIME",
              "CAP_SYS_TTY_CONFIG",
              "CAP_SYSLOG",
              "CAP_WAKE_ALARM"
			]
		},
		"rlimits": [
			{
				"type": "RLIMIT_NOFILE",
				"hard": 1024,
				"soft": 1024
			}
		],
		"noNewPrivileges": false
	},
	"root": {
		"path": "rootfs",
		"readonly": false
	},
	"hostname": "runc",
	"mounts": [
	    {
          "destination" : "/var/vcap/bosh/warden-cpi-agent-env.json",
          "source" : "{{ .AgentSettingsPath }}",
          "type" : "bind",
          "options" : [
            "mode=666",
			"ro",
			"rbind",
			"rprivate"
          ]
        },
{{if ne .DiskPath ""}}
        {
			"destination": "{{ .PersistentDiskDir }}",
			"source":      "{{ .DiskPath }}",
			"type":        "bind",
			"options": [
				"rw",
				"rbind",
				"rprivate"
			]
		},
{{end}}
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
				"relatime"
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
