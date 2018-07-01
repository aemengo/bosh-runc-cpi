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

var (
	cpiPath string
	configPath string
	serverSession *gexec.Session
	stemcellDir string
	diskDir string
	config = func() string {
		return fmt.Sprintf(`
---
stemcell_dir: %s
disk_dir: %s
host: 127.0.0.1
port: 9999
`, stemcellDir, diskDir)
	}

)

var _ = BeforeSuite(func() {
	var err error
	cpiPath, err = gexec.Build("github.com/aemengo/bosh-containerd-cpi/cmd/cpi")
	Expect(err).NotTo(HaveOccurred())

	cpidPath, err := gexec.Build("github.com/aemengo/bosh-containerd-cpi/cmd/cpid")
	Expect(err).NotTo(HaveOccurred())

	stemcellDir, err = ioutil.TempDir("", "bosh-cpi-test-")
	Expect(err).NotTo(HaveOccurred())

	diskDir, err = ioutil.TempDir("", "bosh-cpi-test-")
	Expect(err).NotTo(HaveOccurred())

	configFile, err := ioutil.TempFile("", "bosh-cpi-test-")
	Expect(err).NotTo(HaveOccurred())

	_, err = io.WriteString(configFile, config())
	Expect(err).NotTo(HaveOccurred())

	configPath = configFile.Name()

	command := exec.Command(cpidPath, configPath)
	serverSession, err = gexec.Start(command, GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())

	fmt.Println("Waiting for server to come up...")
	time.Sleep(5 * time.Second)
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
	serverSession.Kill()
	os.RemoveAll(stemcellDir)
	os.RemoveAll(diskDir)
})

func must(session *gexec.Session) *gexec.Session {
	Eventually(session).Should(gexec.Exit(0))
	return session
}