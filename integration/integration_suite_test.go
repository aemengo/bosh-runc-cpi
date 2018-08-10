package integration

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
	"github.com/onsi/gomega/gexec"
	"io/ioutil"
	"io"
	"fmt"
	"os/exec"
	"os"
	"time"
)

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration Suite")
}

const stemcellDownloadURL = "https://s3.amazonaws.com/bosh-core-stemcells/warden/bosh-stemcell-3586.25-warden-boshlite-ubuntu-trusty-go_agent.tgz"

var (
	cpiPath string
	configPath string
	serverSession *gexec.Session
	stemcellDir string
	diskDir string
	vmDir string
	config = func() string {
		return fmt.Sprintf(`
---
stemcell_dir: %s
disk_dir: %s
vm_dir: %s
network_type: unix
address: "/tmp/cpid.sock"
`, stemcellDir, diskDir, vmDir)
	}

)

var _ = BeforeSuite(func() {
	var err error
	cpiPath, err = gexec.Build("github.com/aemengo/bosh-runc-cpi/cmd/cpi")
	Expect(err).NotTo(HaveOccurred())

	cpidPath, err := gexec.Build("github.com/aemengo/bosh-runc-cpi/cmd/cpid")
	Expect(err).NotTo(HaveOccurred())

	stemcellDir, err = ioutil.TempDir("", "bosh-cpi-test-stemcell-")
	Expect(err).NotTo(HaveOccurred())

	diskDir, err = ioutil.TempDir("", "bosh-cpi-test-disk-")
	Expect(err).NotTo(HaveOccurred())

	vmDir, err = ioutil.TempDir("", "bosh-cpi-test-vm-")
	Expect(err).NotTo(HaveOccurred())

	configFile, err := ioutil.TempFile("", "bosh-cpi-test-config-")
	Expect(err).NotTo(HaveOccurred())

	_, err = io.WriteString(configFile, config())
	Expect(err).NotTo(HaveOccurred())

	configPath = configFile.Name()

	command := exec.Command(cpidPath, configPath)
	serverSession, err = gexec.Start(command, GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())

	// Wait for server to come up...
	time.Sleep(time.Second)
})

var _ = AfterSuite(func() {
	serverSession.Kill()
	gexec.CleanupBuildArtifacts()
	os.RemoveAll(configPath)
	os.RemoveAll(stemcellDir)
	os.RemoveAll(diskDir)
	os.RemoveAll(vmDir)
})

func must(session *gexec.Session, duration ...interface{}) *gexec.Session {
	if len(duration) == 0 {
		Eventually(session).Should(gexec.Exit(0))
	} else {
		Eventually(session, duration[0]).Should(gexec.Exit(0))
	}
	return session
}