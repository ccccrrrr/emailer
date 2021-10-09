package main

import (
	"github.com/ccccrrrr/config"
	"github.com/ccccrrrr/util"
	"errors"
	"fmt"
	"log"
	"net/smtp"
)

// content和to是一對一的關係
type Server struct {
	config *config.Config
	content []string
	to []string
	Subject string
	ContentType string
}

func CreateServer(config *config.Config) *Server {
	server := new(Server)
	server.config = config
	server.content = make([]string, 0)
	server.to = make([]string, 0)
	server.ContentType = "text/html; charset=UTF-8"
	return server
}


func (s *Server) Reset() {
	s.content = make([]string, 0)
	s.to = make([]string, 0)
	s.Subject = ""
}

func(s *Server) SetSubject(Subject string) {
	s.Subject = Subject
}

func (s *Server) Add(to, content string) {
	s.content = append(s.content, content)
	s.to = append(s.to, to)
}

func (s *Server) ListTo(print bool) string {
	var res string
	for k, v := range s.to {
		res += fmt.Sprintf("%d: %s\n", k+1, v)
	}
	if print {
		fmt.Print(res)
	}
	return res
}

func (s *Server) ListContent(num uint, print bool) string {
	var res string
	res += fmt.Sprintf("No.%d: to: %s\ncontent:\n%s\n", num+1, s.to[num], s.content[num])
	if print {
		fmt.Print(res)
	}
	return res
}

func (s *Server) Send() error {
	if s.Subject == "" {
		log.Print("Email has no Subject")
		return errors.New("email has no Subject")
	}
	if s.ContentType == "" {
		log.Print("Email has no content_type")
		return errors.New("email has no content_type")
	}
	if len(s.to) != len(s.content) {
		log.Print("receiver and body mismatch")
		return errors.New("receiver and body mismatch")
	}
	var __ error
	for k, v := range s.to {
		message := generateHeader(s, k) + "\r\n" + s.content[k]
		err := s.send(k, message)
		if err != nil {
			__ = err
			log.Printf("send to #%+v %+v error: %+v", k+1, v, err)
		}
	}
	return __
}

func (s *Server) send(k int, message string) error {
	auth := smtp.PlainAuth(
		"",
		s.config.Email,
		s.config.Password,
		s.config.Host,
	)
	err := util.SendMailUsingTLS(
		fmt.Sprintf("%s:%d", s.config.Host, s.config.Port),
		auth,
		s.config.Email,
		[]string{s.to[k]},
		[]byte(message),
	)
	return err
}

func generateHeader(s *Server, k int) string {
	header := make(map[string]string)
	header["From"] = "test" + "<" + s.config.Email + ">"
	header["To"] = s.to[k]
	header["Subject"] = s.Subject
	header["Content-Type"] = s.ContentType
	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	return message
}