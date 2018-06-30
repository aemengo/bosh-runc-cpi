package integration

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

var _ = Describe("bosh-containerd-cpi", func() {

	var (
		configPath string
		config     map[string]interface{}
        executeCPI  = func(args string) *gexec.Session {
            command := exec.Command(binaryPath, configPath)
            command.Stdin = strings.NewReader(args)
            session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
            Expect(err).NotTo(HaveOccurred())
            return session
        }
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
			tempDir         string
			stemcellSrcDir  string
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
				"stemcell_dir": stemcellDestDir,
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

            session := must(executeCPI(args))

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

            session := must(executeCPI(args))

			out := string(session.Out.Contents())
			Expect(out).To(MatchJSON(
				`{"result":{"stemcell_formats":["warden-tar","general-tar"]},"error":null,"log":""}`,
			))
		})
	})

	Describe("delete_stemcell", func() {

		var (
			stemcellDir  string
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
				"stemcell_dir": stemcellDir,
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

            session := must(executeCPI(args))

			out := string(session.Out.Contents())
			Expect(out).To(MatchJSON(`{"result":null,"error":null,"log":""}`))
			Expect(stemcellPath).NotTo(BeAnExistingFile())
		})
	})

	Describe("create_disk", func() {

		var (
			diskDir string
		)

		BeforeEach(func() {
			if runtime.GOOS != "linux" {
				Skip("the following test require a linux environment")
			}

			var err error
			diskDir, err = ioutil.TempDir("", "bosh-cpi-test-")
			Expect(err).NotTo(HaveOccurred())

			config = map[string]interface{}{
				"disk_dir": diskDir,
			}
		})

		AfterEach(func() {
			os.RemoveAll(diskDir)
		})

		It("creates a disk in the disk_directory", func() {
			var args = `  {
              "method": "create_disk",
              "arguments": [
                100,
                {},
                "vm-870c3e28-a4a7-4d2f-5272-18f2a136cb58"
              ],
              "context": {
                "director_uuid": "e8c76164-7eda-405a-475a-cec0e51ee972"
              }
            }`

            session := must(executeCPI(args))

			out := string(session.Out.Contents())
			Expect(out).To(MatchRegexp(`result":".*"`))
			Expect(out).To(MatchRegexp(`error":null`))

			files, _ := ioutil.ReadDir(diskDir)
			Expect(files).To(HaveLen(1))
		})
	})

	Describe("delete_disk", func() {

		var (
			tempDir string
			diskDir string
		)

		BeforeEach(func() {
			var err error
			tempDir, err = ioutil.TempDir("", "bosh-cpi-test-")
			Expect(err).NotTo(HaveOccurred())

			diskDir = filepath.Join(tempDir, "abc-123-some-guid")
			os.MkdirAll(diskDir, os.ModePerm)

			config = map[string]interface{}{
				"disk_dir": tempDir,
			}
		})

		AfterEach(func() {
			os.RemoveAll(tempDir)
		})

		It("remove the disk from the disk dir", func() {
			var args = `{
              "method": "delete_disk",
              "arguments": [
                "abc-123-some-guid"
              ],
              "context": {
                "director_uuid": "e8c76164-7eda-405a-475a-cec0e51ee972"
              }
            }`

			Expect(diskDir).To(BeADirectory())

            session := must(executeCPI(args))

			out := string(session.Out.Contents())
			Expect(out).To(MatchJSON(`{"result":null,"error":null,"log":""}`))
			Expect(diskDir).NotTo(BeAnExistingFile())
		})
	})

	Describe("has_disk", func() {

		var (
			tempDir string
			diskDir string
		)

		BeforeEach(func() {
			var err error
			tempDir, err = ioutil.TempDir("", "bosh-cpi-test-")
			Expect(err).NotTo(HaveOccurred())

			diskDir = filepath.Join(tempDir, "abc-123-some-guid")
			os.MkdirAll(diskDir, os.ModePerm)

			config = map[string]interface{}{
				"disk_dir": tempDir,
			}
		})

		AfterEach(func() {
			os.RemoveAll(tempDir)
		})

		Context("when the disk exists", func() {
			It("returns true", func() {
                var args = `{
                  "method": "has_disk",
                  "arguments": [
                    "abc-123-some-guid"
                  ],
                  "context": {
                    "director_uuid": "e8c76164-7eda-405a-475a-cec0e51ee972"
                  }
                }`

                session := must(executeCPI(args))

				out := string(session.Out.Contents())
				Expect(out).To(MatchJSON(`{"result":true,"error":null,"log":""}`))
			})
		})

		Context("when the disk does not exist", func() {
			It("returns false", func() {
				var args = `{
                  "method": "has_disk",
                  "arguments": [
                    "def-456-some-non-existent-guid"
                  ],
                  "context": {
                    "director_uuid": "e8c76164-7eda-405a-475a-cec0e51ee972"
                  }
                }`

                session := must(executeCPI(args))

				out := string(session.Out.Contents())
				Expect(out).To(MatchJSON(`{"result":false,"error":null,"log":""}`))
			})
		})
	})

    Describe("set_vm_metadata", func() {
        It("is a no-op action that returns successfully", func() {
            var args = `{
              "method": "set_vm_metadata",
              "arguments": [
                "23b20ab4-7d08-4cb6-5136-a31486af2139",
                {
                  "director": "bosh-lite",
                  "deployment": "zookeeper",
                  "id": "fef527e2-cbd8-4d07-aa8a-39aea1fbb420",
                  "job": "compilation-35a43e79-e770-41fd-81d5-0816306db687",
                  "instance_group": "compilation-35a43e79-e770-41fd-81d5-0816306db687",
                  "index": "0",
                  "name": "compilation-35a43e79-e770-41fd-81d5-0816306db687/fef527e2-cbd8-4d07-aa8a-39aea1fbb420",
                  "created_at": "2018-06-24T22:45:44Z"
                }
              ],
              "context": {
                "director_uuid": "eb896bb7-d54b-4e5b-97cd-dc4e1ab1204f",
                "request_id": "cpi-685135"
              },
              "api_version": 1
            }`

            session := must(executeCPI(args))

            out := string(session.Out.Contents())
            Expect(out).To(MatchJSON(`{"result":null,"error":null,"log":""}`))
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

            session := executeCPI(args)

			<-session.Exited
			Expect(session.ExitCode()).NotTo(Equal(0))

			out := string(session.Out.Contents())
			Expect(out).To(MatchJSON(
				`{"result":null,"error":{"type":"Bosh::Clouds::NotImplemented","message":"'some-unimplemented-method' is not yet supported. Please call implemented method","ok_to_retry":false},"log":""}`,
			))
		})
	})
})
