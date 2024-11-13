package page

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/xr0-org/progstack-ssg/internal/assert"
	"github.com/xr0-org/progstack-ssg/internal/ast/area/sitefile"
)

type custompage struct {
	template string
	data     map[string]string
}

func CustomPage(template string, data map[string]string) *custompage {
	return &custompage{template, data}
}

func (pg *custompage) Link(path string, pi PageInfo) (string, error) {
	url, err := filepath.Rel(
		pi.Root(),
		rightextpath(path, pi.DynamicLinks()),
	)
	if err != nil {
		return "", fmt.Errorf("cannot get relative path: %w", err)
	}
	return "/" + url, nil
}

func (pg *custompage) GenerateIndex(
	w io.Writer, posts []Post, pi PageInfo,
) error {
	return fmt.Errorf("custom page cannot generate index")
}

func (pg *custompage) GenerateWithoutIndex(w io.Writer, pi PageInfo) error {
	return pi.Theme().ExecuteCustom(w, pg.template, pg.data)
}

func (pg *custompage) Generate(w io.Writer, pi PageInfo, index Page) error {
	assert.Assert(index != nil)
	indexppg, ok := index.(*parsedpage)
	assert.Assert(ok)

	return pi.Theme().ExecuteCustom(
		w,
		pg.template,
		pg.datawithmoremap(map[string]string{
			"SiteTitle": indexppg.title,
		}),
	)
}

func (pg *custompage) datawithmoremap(more map[string]string) map[string]string {
	m := map[string]string{}
	for k, v := range pg.data {
		m[k] = v
	}
	for k, v := range more {
		m[k] = v
	}
	return m
}

func (pg *custompage) IsPost() bool { return false }

func (pg *custompage) AsPost(_, _ string) *Post {
	assert.Assert(false)
	return nil
}

func (pg *custompage) ToFile(path string) sitefile.File {
	return sitefile.NonPostFile(path)
}