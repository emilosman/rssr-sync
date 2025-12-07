package data

type Lists struct {
	ListIndex map[string]*List
}

type List struct {
	Ts     int64
	ApiKey string
	Data   []byte
}

func (ls *Lists) GetList(apiKey string, ts int64) ([]byte, error) {
	l, err := ls.findList(apiKey)
	if err != nil {
		return nil, err
	}

	if l.Ts > ts {
		return nil, ErrOldTimestamp
	}

	return l.Data, nil
}

func (ls *Lists) SetData(apiKey string, data []byte, ts int64) error {
	l, err := ls.findList(apiKey)
	if err != nil {
		return err
	}

	err = l.setTimestamp(ts)
	if err != nil {
		return err
	}

	err = l.setData(data)
	if err != nil {
		return err
	}

	return l.Save()
}

func (l *List) Save() error {
	// write to disk with API key filename
	return nil
}

func (ls *Lists) findList(apiKey string) (*List, error) {
	if apiKey == "" {
		return nil, ErrNoApiKey
	}

	l := ls.ListIndex[apiKey]
	if l == nil {
		l := &List{
			ApiKey: apiKey,
		}
		ls.ListIndex[apiKey] = l
	}
	return l, nil
}

func (l *List) setTimestamp(ts int64) error {
	if l.Ts < ts {
		l.Ts = ts
		return nil
	}
	return ErrOldTimestamp
}

func (l *List) setData(data []byte) error {
	l.Data = data
	return nil
}
