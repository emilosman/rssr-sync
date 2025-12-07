package web

type Data struct {
	ts     int64
	apiKey string
	body   []byte
}

func (d *Data) setTimestamp(ts int64) error {
	if d.ts < ts {
		d.ts = ts
		return nil
	}
	return ErrOldTimestamp
}

func (d *Data) setBody(data []byte) error {
	d.body = data
	return nil
}

func (d *Data) Save() error {
	// write to disk
	return nil
}

func (d *Data) UpdateData(data []byte, ts int64) error {
	err := d.setTimestamp(ts)
	if err != nil {
		return err
	}

	err = d.setBody(data)
	if err != nil {
		return err
	}

	return d.Save()
}
