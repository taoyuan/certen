package caddy

import (
	"certen"
	"certen/utils"
	"errors"
	"fmt"
	"path/filepath"
	"strings"
)

func init() {
	certen.RegisterProvider("caddy", func(dir string) (certen.Provider, error) {
		caddy, err := New(dir)
		return caddy, err
	})
}

var POSSIBLES = []string{
	"{{.cwd}}/.caddy",
	"/opt/caddyserver/.caddy",
	"{{.home}}/.caddy",
}

type Caddy struct {
	certen.Provider
	dir string
}

func validate(dir string) bool {
	return utils.Exists(dir) && utils.Exists(filepath.Join(dir, "acme")) && utils.Exists(filepath.Join(dir, "ocsp"))
}

func New(dir string) (*Caddy, error) {
	var err error
	dir = strings.TrimSpace(dir)
	if dir == "" {
		dir, err = utils.FindDir(POSSIBLES, &utils.DefaultContext, validate)
		if err != nil {
			return nil, err
		}
	} else if !utils.Exists(dir) || !validate(dir) {
		err = errors.New(fmt.Sprintf("\"%s\" is not a valid caddy data directory", dir))
	}
	return &Caddy{dir: dir}, err
}

func (c *Caddy) FindDomains(pattern string) ([]*certen.Domain, error) {
	patterns := strings.Split(pattern, ",")
	answer := []*certen.Domain{}

	for _, p := range patterns {
		p = strings.TrimSpace(p)
		dirs, err := filepath.Glob(filepath.Join(c.dir, "acme/*/sites", p))
		if err != nil {
			continue
		}
		for _, dir := range dirs {
			answer = append(answer, &certen.Domain{
				Name: filepath.Base(dir),
				Dir:  dir,
			})
		}
	}

	return answer, nil
}

func (c *Caddy) ExportByName(name string, dest string, assemble bool) ([]string, error) {
	domains, err := c.FindDomains(name)
	if err != nil {
		return nil, err
	}
	return c.Export(domains, dest, assemble)
}

func (c *Caddy) Export(domains []*certen.Domain, dest string, assemble bool) ([]string, error) {
	answer := []string{}
	for _, domain := range domains {
		domainDir, err := filepath.Abs(domain.Dir)
		if err != nil {
			return nil, err
		}
		var files = []string{}
		for _, ext := range []string{"crt", "key"} {
			files = append(files, filepath.Join(domainDir, domain.Name+"."+ext))
		}

		if strings.TrimSpace(dest) == "" {
			dest = domainDir
		}

		if assemble {
			dest := filepath.Join(dest, domain.Name+".pem")
			err := utils.ConcatCerts(dest, files...)
			if err != nil {
				return nil, err
			}

		} else if dest != domainDir {
			for _, file := range files {
				err := utils.CopyFile(file, filepath.Join(dest, filepath.Base(file)))
				if err != nil {
					return nil, err
				}
			}
		}
		answer = append(answer, domain.Name)
	}

	return answer, nil
}
