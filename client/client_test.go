package main

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("Client", func() {

	var (
		server     *ghttp.Server
		statusCode int
		body       []byte
		path       string
		addr       string
	)

	BeforeEach(func() {
		// start a test http server
		server = ghttp.NewServer()
	})

	AfterEach(func() {
		server.Close()
	})

	Context("When given empty url", func() {
		BeforeEach(func() {
			addr = ""
		})
		It("Returns the empty path", func() {
			_, err := getResponse(addr)
			Expect(err).Should(HaveOccurred())
		})
	})

	Context("When given unsupported protocol scheme", func() {
		BeforeEach(func() {
			addr = "tcp://localhost"
		})
		It("Returns the empty path", func() {
			_, err := getResponse(addr)
			Expect(err).Should(HaveOccurred())
		})
	})

	Context("When get request is sent to empty path", func() {
		BeforeEach(func() {
			statusCode = 200
			path = "/"
			body = []byte("Hi there, the end point is :!")
			addr = "http://" + server.Addr() + path
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", path),
					ghttp.RespondWithPtr(&statusCode, &body),
				))
		})
		It("Returns the empty path", func() {
			bdy, err := getResponse(addr)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(bdy).To(Equal(body))
		})
	})

	Context("When get request is sent to hello path", func() {
		BeforeEach(func() {
			statusCode = 200
			path = "/hello"
			body = []byte("Hi there, the end point is :hello!")
			addr = "http://" + server.Addr() + path
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", path),
					ghttp.RespondWithPtr(&statusCode, &body),
				))
		})
		It("Returns the hello path", func() {
			bdy, err := getResponse(addr)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(bdy).To(Equal(body))
		})
	})

	Context("When get request is sent to read path but there is no file", func() {
		BeforeEach(func() {
			statusCode = 500
			path = "/read"
			body = []byte("open data.txt: no such file or directory\r\n")
			addr = "http://" + server.Addr() + path
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", path),
					ghttp.RespondWithPtr(&statusCode, &body),
				))
		})
		It("Returns internal server error", func() {
			_, err := getResponse(addr)
			Expect(err).Should(HaveOccurred())
		})
	})

	Context("When get request is sent to read path but file exists", func() {
		BeforeEach(func() {
			file, err := os.Create("data.txt")
			Expect(err).NotTo(HaveOccurred())
			body = []byte("Hi there!")
			file.Write(body)
			statusCode = 200
			path = "/read"
			addr = "http://" + server.Addr() + path
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", path),
					ghttp.RespondWithPtr(&statusCode, &body),
				))
		})

		AfterEach(func() {
			err := os.Remove("data.txt")
			Expect(err).NotTo(HaveOccurred())
		})
		It("Reads data from file successfully", func() {
			bdy, err := getResponse(addr)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(bdy).To(Equal(body))
		})
	})
})
