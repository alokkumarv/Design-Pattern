package main

import "log"

type Request struct {
	Header map[string]string
	Body   string
}

type Handler interface {
	SetNext(Handler)
	Handle(req *Request)
}

type AuthHandler struct {
	next Handler
}

func (a *AuthHandler) SetNext(next Handler) {
	a.next = next

}
func (a *AuthHandler) Handle(req *Request) {
	if req.Body == "auth" {
		log.Printf("Auth handler handle request ")
		return
	}
	a.next.Handle(req)

}

type LogMiddleware struct {
	next Handler
}

func (l *LogMiddleware) SetNext(next Handler) {
	l.next = next
}
func (l *LogMiddleware) Handle(req *Request) {
	if req.Body == "log" {
		log.Printf("Log handler handle reuest ")
		return
	}
	l.next.Handle(req)

}

type FinalHandler struct {
	next Handler
}

func (f *FinalHandler) SetNext(next Handler) {
	f.next = nil
}
func (f *FinalHandler) Handle(req *Request) {
	log.Println("No handler found to handle this request ")

}

func main() {

	auth := &AuthHandler{}
	log := &LogMiddleware{}
	final := &FinalHandler{}
	auth.SetNext(log)
	log.SetNext(final)
	req := Request{Header: nil, Body: "log"}
	auth.Handle(&req)
	req = Request{Header: nil, Body: "auth"}
	auth.Handle(&req)
	req = Request{Header: nil, Body: "final"}
	auth.Handle(&req)

}
