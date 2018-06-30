package integration

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os/exec"
	"strings"
	"github.com/onsi/gomega/gexec"
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "gopkg.in/yaml.v2"
)

var _ = Describe("bosh-containerd-cpi", func() {

    var (
      configPath string
      config map[string]interface{}
    )

    AfterEach(func() {
        os.RemoveAll(configPath)
    })

    JustBeforeEach(func() {
        f, _ := ioutil.TempFile("", "bosh-cpi-test-")
        yaml.NewEncoder(f).Encode(&config)
        configPath = f.Name()
    })

	Describe("create_stemcell", func() {

        var (
          tempDir string
          stemcellSrcDir string
          stemcellDestDir string
        )

        BeforeEach(func() {
            var err error
            tempDir, err = ioutil.TempDir("", "bosh-cpi-test-")
            Expect(err).NotTo(HaveOccurred())

            stemcellSrcDir = filepath.Join(tempDir, "stemcell-src")
            stemcellDestDir = filepath.Join(tempDir, "stemcell-dest")

            os.MkdirAll(stemcellSrcDir, os.ModePerm)
            ioutil.WriteFile(
                filepath.Join(stemcellSrcDir, "some-stemcell.tgz"),
                []byte("some-content"),
                0600,
            )

            config = map[string]interface{}{
                "stemcell_dir" : stemcellDestDir,
            }
        })

        AfterEach(func() {
            os.RemoveAll(tempDir)
        })

		It("moves the stemcell to the stemcell dir", func() {
			var args = fmt.Sprintf(`{
              "method": "create_stemcell",
              "arguments": [
                "%s",
                {}
              ],
              "context": {
                "director_uuid": "e8c76164-7eda-405a-475a-cec0e51ee972"
              }
             }`, filepath.Join(stemcellSrcDir, "some-stemcell.tgz"))

			command := exec.Command(binaryPath, configPath)
			command.Stdin = strings.NewReader(args)
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session).Should(gexec.Exit(0))

            out := string(session.Out.Contents())
            Expect(out).To(MatchRegexp(`result":".*"`))
            Expect(out).To(MatchRegexp(`error":null`))

            files, _ := ioutil.ReadDir(stemcellDestDir)
            contents, _ := ioutil.ReadFile(filepath.Join(stemcellDestDir, files[0].Name()))
            Expect(string(contents)).To(Equal("some-content"))
		})
	})

    Describe("info", func() {
        It("returns info", func() {
            var args = `{
              "method": "info",
              "arguments": [],
              "context": {
                "director_uuid": "e8c76164-7eda-405a-475a-cec0e51ee972"
              }
             }`

            command := exec.Command(binaryPath, configPath)
            command.Stdin = strings.NewReader(args)
            session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
            Expect(err).NotTo(HaveOccurred())
            Eventually(session).Should(gexec.Exit(0))

            out := string(session.Out.Contents())
            Expect(out).To(MatchJSON(
                `{"result":{"stemcell_formats":["warden-tar","general-tar"]},"error":null,"log":""}`,
            ))
        })
    })

	Describe("delete_stemcell", func() {

        var (
          stemcellDir string
          stemcellPath string
        )

        BeforeEach(func() {
            var err error
            stemcellDir, err = ioutil.TempDir("", "bosh-cpi-test-")
            Expect(err).NotTo(HaveOccurred())

            stemcellPath = filepath.Join(stemcellDir, "abc-123-some-guid")

            ioutil.WriteFile(
                stemcellPath,
                []byte("some-content"),
                0600,
            )

            config = map[string]interface{}{
                "stemcell_dir" : stemcellDir,
            }
        })

        AfterEach(func() {
            os.RemoveAll(stemcellDir)
        })

		It("remove the stemcell from the stemcell dir", func() {
            var args = `{
              "method": "delete_stemcell",
              "arguments": [
                "abc-123-some-guid"
              ],
              "context": {
                "director_uuid": "e8c76164-7eda-405a-475a-cec0e51ee972"
              }
            }`

			command := exec.Command(binaryPath, configPath)
			command.Stdin = strings.NewReader(args)
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session).Should(gexec.Exit(0))

            out := string(session.Out.Contents())
            Expect(out).To(MatchJSON(`{"result":null,"error":null,"log":""}`))
            Expect(stemcellPath).NotTo(BeAnExistingFile())
		})
	})

    Describe("unimplemented cpi methods", func() {
        It("returns a 'NotImplemented' error", func() {
            var args = `{
              "method": "some-unimplemented-method",
              "arguments": [],
              "context": {
                "director_uuid": "e8c76164-7eda-405a-475a-cec0e51ee972"
              }
             }`

            command := exec.Command(binaryPath, configPath)
            command.Stdin = strings.NewReader(args)
            session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
            Expect(err).NotTo(HaveOccurred())

            <-session.Exited
            Expect(session.ExitCode()).NotTo(Equal(0))

            out := string(session.Out.Contents())
            Expect(out).To(MatchJSON(
                `{"result":null,"error":{"type":"Bosh::Clouds::NotImplemented","message":"\"some-unimplemented-method\" is not yet supported. Please call implemented method","ok_to_retry":false},"log":""}`,
            ))
        })
    })
})
