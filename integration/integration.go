package integration

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"io"
)

var _ = Describe("bosh-containerd-cpi", func() {

	var (
        executeCPI = func(args string) *gexec.Session {
            command := exec.Command(cpiPath, configPath)
            command.Stdin = strings.NewReader(args)
            session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
            Expect(err).NotTo(HaveOccurred())
            return session
        }
	)

	AfterEach(func() {
		os.RemoveAll(stemcellDir)
		os.RemoveAll(diskDir)
		os.MkdirAll(stemcellDir, os.ModePerm)
		os.MkdirAll(diskDir, os.ModePerm)
	})

	Describe("create_stemcell", func() {

		var (
			stemcellSrcPath         string
		)

		BeforeEach(func() {
			stemcellSrcFile, err := ioutil.TempFile("", "bosh-cpi-test-")
			Expect(err).NotTo(HaveOccurred())

			_, err = io.WriteString(stemcellSrcFile, "some-stemcell-content")
			Expect(err).NotTo(HaveOccurred())

			stemcellSrcPath = stemcellSrcFile.Name()
		})

		AfterEach(func() {
			os.RemoveAll(stemcellSrcPath)
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
             }`, stemcellSrcPath)

            session := must(executeCPI(args))

            out := string(session.Out.Contents())
            Expect(out).To(MatchRegexp(`result":".*"`))
            Expect(out).To(MatchRegexp(`error":null`))

            files, _ := ioutil.ReadDir(stemcellDir)
            contents, _ := ioutil.ReadFile(filepath.Join(stemcellDir, files[0].Name()))
            Expect(string(contents)).To(Equal("some-stemcell-content"))
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
			stemcellPath string
		)

		BeforeEach(func() {
			stemcellPath = filepath.Join(stemcellDir, "abc-123-some-guid")

			ioutil.WriteFile(
				stemcellPath,
				[]byte("some-content"),
				0600,
			)
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

		BeforeEach(func() {
			if runtime.GOOS != "linux" {
				Skip("the following test require a linux environment")
			}
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
			diskPath string
		)

		BeforeEach(func() {
			diskPath = filepath.Join(diskDir, "abc-123-some-guid")
			os.MkdirAll(diskPath, os.ModePerm)
		})

		AfterEach(func() {
			os.RemoveAll(diskPath)
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

			Expect(diskPath).To(BeADirectory())

            session := must(executeCPI(args))

			out := string(session.Out.Contents())
			Expect(out).To(MatchJSON(`{"result":null,"error":null,"log":""}`))
			Expect(diskPath).NotTo(BeAnExistingFile())
		})
	})

	Describe("has_disk", func() {
		Context("when the disk exists", func() {
			var (
				diskPath string
			)

			BeforeEach(func() {
				diskPath = filepath.Join(diskDir, "abc-123-some-guid")
				os.MkdirAll(diskPath, os.ModePerm)
			})

			AfterEach(func() {
				os.RemoveAll(diskPath)
			})

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
