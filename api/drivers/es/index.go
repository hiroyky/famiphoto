package es

import "io"

type IndexBody interface {
	ID() string
	BodyReader() (io.Reader, error)
}

func (d *driver) Index(index string, body IndexBody) error {
	bodyReader, err := body.BodyReader()
	if err != nil {
		return err
	}

	res, err := d.client.Index(
		index, bodyReader,
		d.client.Index.WithDocumentID(body.ID()),
	)
	if err := d.handleError(res, err); err != nil {
		return err
	}

	return nil
}
