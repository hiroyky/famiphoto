package es

import "io"

type Mapping interface {
	BodyReader() (io.Reader, error)
}

func (d *driver) CreateIndex(index string, mapping Mapping) error {
	reader, err := mapping.BodyReader()
	if err != nil {
		return err
	}
	res, err := d.client.Indices.Create(index, d.client.Indices.Create.WithBody(reader))
	return d.handleError(res, err)
}
