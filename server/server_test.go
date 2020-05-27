package main

import (
	"io/ioutil"
	"net/http"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("Server", func() {
	var server *ghttp.Server
	msg := "Hi there, the end point is :"

	BeforeEach(func() {
		// start a test http server
		server = ghttp.NewServer()
	})

	AfterEach(func() {
		server.Close()
	})

	Context("When get request is sent to empty path", func() {
		BeforeEach(func() {
			// Add your handler which has to be called for a given path
			// If there are multiple redirects append all the handlers
			server.AppendHandlers(
				Handler,
			)
		})
		It("Returns the empty path", func() {
			resp, err := http.Get(server.URL() + "/")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).Should(Equal(http.StatusOK))
			body, err := ioutil.ReadAll(resp.Body)
			resp.Body.Close()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(body)).To(Equal(msg + "!"))
		})
	})

	Context("When get request is sent to hello path", func() {
		BeforeEach(func() {
			server.AppendHandlers(
				Handler,
			)
		})
		It("Returns the empty path", func() {
			resp, err := http.Get(server.URL() + "/hello")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).Should(Equal(http.StatusOK))
			body, err := ioutil.ReadAll(resp.Body)
			resp.Body.Close()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(body)).To(Equal(msg + "hello!"))
		})
	})

	Context("When get request is sent to read path but there is no file", func() {
		BeforeEach(func() {
			server.AppendHandlers(
				ReadHandler,
			)
		})
		It("Returns internal server error", func() {
			resp, err := http.Get(server.URL() + "/read")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).Should(Equal(http.StatusInternalServerError))
			body, err := ioutil.ReadAll(resp.Body)
			resp.Body.Close()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(body)).To(Equal("open data.txt: no such file or directory\n"))
		})
	})

	Context("When get request is sent to read path but file exists", func() {
		BeforeEach(func() {
			file, err := os.Create("data.txt")
			Expect(err).NotTo(HaveOccurred())
			file.Write([]byte("Hi there!"))
			server.AppendHandlers(
				ReadHandler,
			)
		})

		AfterEach(func() {
			err := os.Remove("data.txt")
			Expect(err).NotTo(HaveOccurred())
		})
		It("Reads data from file successfully", func() {
			resp, err := http.Get(server.URL() + "/read")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).Should(Equal(http.StatusOK))
			body, err := ioutil.ReadAll(resp.Body)
			resp.Body.Close()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(body)).To(Equal("Content in file is...\r\nHi there!"))
		})
	})
})
