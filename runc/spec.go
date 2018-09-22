package runc

import (
	"github.com/opencontainers/runtime-spec/specs-go"
	"strings"
)

func DefaultSpec() *specs.Spec {
	return &specs.Spec{
		Version: specs.Version,
		Root: &specs.Root{
			Path:     "rootfs",
			Readonly: false,
		},
		Process: &specs.Process{
			Terminal: false,
			User: specs.User{
				UID: 0,
				GID: 0,
			},
			Args: []string{
				"bash", "-c",
				strings.Join([]string{
					"umount /var/vcap/data/root_log",
					"umount /var/vcap/data/sys/run",
					"exec env -i /usr/sbin/runsvdir-start",
				}, "; "),
			},
			Env: []string{
				"PATH=/var/vcap/bosh/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
				"TERM=xterm",
			},
			Cwd: "/",
			Capabilities: &specs.LinuxCapabilities{
				Bounding:    capabilitiesAll,
				Inheritable: capabilitiesAll,
				Permitted:   capabilitiesAll,
			},
			NoNewPrivileges: false,
		},
		Linux: &specs.Linux{
			Resources: &specs.LinuxResources{
				Devices: []specs.LinuxDeviceCgroup{{
					Allow:  false,
					Access: "rwm",
				}},
			},
			Namespaces: []specs.LinuxNamespace{
				{Type: specs.NetworkNamespace},
				{Type: specs.PIDNamespace},
				{Type: specs.UTSNamespace},
				{Type: specs.IPCNamespace},
				{Type: specs.MountNamespace},
			},
		},
		Mounts: []specs.Mount{
			{
				Destination: "/proc",
				Source:      "proc",
				Type:        "proc",
				Options:     []string{"nosuid", "noexec", "nodev", "relatime"},
			},
			{
				Destination: "/dev",
				Source:      "tmpfs",
				Type:        "tmpfs",
				Options:     []string{"nosuid", "strictatime", "mode=755", "size=65536k"},
			},
			{
				Destination: "/dev/pts",
				Source:      "devpts",
				Type:        "devpts",
				Options:     []string{"nosuid", "noexec", "newinstance", "ptmxmode=0666", "mode=0620"},
			},
			{
				Destination: "/sys",
				Source:      "sysfs",
				Type:        "sysfs",
				Options:     []string{"nosuid", "noexec", "nodev"},
			},
			{
				Destination: "/sys/fs/cgroup",
				Source:      "cgroup",
				Type:        "cgroup",
				Options:     []string{"nosuid", "noexec", "nodev", "relatime"},
			},
		},
	}
}

type SpecOption func(*specs.Spec)

func Apply(spec *specs.Spec, opts ...SpecOption) {
	for _, opt := range opts {
		opt(spec)
	}
}

func PrependMounts(mounts []specs.Mount) SpecOption {
	return func(spec *specs.Spec) {
		spec.Mounts = append(mounts, spec.Mounts...)
	}
}

func AppendMounts(mounts []specs.Mount) SpecOption {
	return func(spec *specs.Spec) {
		spec.Mounts = append(spec.Mounts, mounts...)
	}
}

func RemoveMount(mountSrc string) SpecOption {
	return func(spec *specs.Spec) {
		var mts []specs.Mount

		for _, mount := range spec.Mounts {
			if mount.Source != mountSrc {
				mts = append(mts, mount)
			}
		}

		spec.Mounts = mts
	}
}
