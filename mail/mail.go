package main

import (
	"bytes"
	"context"
	"html/template"
	"log"
	"path/filepath"

	"github.com/aitorfernandez/roshambo/pkg/projectpath"
	pb "github.com/aitorfernandez/roshambo/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

// mailServer is the API for mailServer service.
type mailServer struct{}

// SendToken sends a send_token mail
func (s mailServer) SendToken(ctx context.Context, req *pb.SendTokenReq) (*emptypb.Empty, error) {
	var (
		body string
		err  error
	)
	data := struct {
		ID    string
		Token string
	}{
		ID:    req.ID,
		Token: req.Token,
	}
	if body, err = getContent("send_auth_token.html", data); err != nil {
		log.Println(err)
		return &emptypb.Empty{}, nil
	}
	if err = send([]string{req.Receiver}, "Get started link for roshambo", body); err != nil {
		log.Println(err)
	}
	return &emptypb.Empty{}, nil
}

func parse(filename string, data interface{}) (string, error) {
	var (
		buf bytes.Buffer
		err error
		t   *template.Template
	)
	if t, err = template.ParseFiles(filename); err != nil {
		return "", err
	}
	if err = t.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func getContent(filename string, data interface{}) (string, error) {
	path := filepath.Join(projectpath.Base(), "mail", filename)
	return parse(path, data)
}
