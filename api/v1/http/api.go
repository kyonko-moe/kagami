// Package http provides http api
package http

import (
	"context"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

// Register a new node
func Register(ctx context.Context, urlstr string, name string) (err error) {
	var (
		form = url.Values{}
	)
	form.Set("name", name)
	if _, err := http.PostForm(urlstr, form); err != nil {
		err = errors.WithStack(err)
	}
	return
}

// Locate query a node by name
func Locate(ctx context.Context) {

}

// Connect to a spcified node , when connected it will be return
func Connect(ctx context.Context) {

}
