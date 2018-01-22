package services

import "io/ioutil"

type Page struct {
	Title      string
	PageTitle  string
	ServerPath string
	Body       []byte
}

func (p *Page) loadPage() error {
	var err error
	filename := p.ServerPath + p.Title + ".html"
	p.Body, err = ioutil.ReadFile(filename)
	return err
}
func (p *Page) getFullPath() string {
	return p.ServerPath + p.Title + ".html"
}
