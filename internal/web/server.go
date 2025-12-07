package web

type Lists struct {
	listIndex map[string]*List
}

type List struct {
	Ts     int64
	ApiKey string
	Data   []byte
}

func (ls *Lists) GetList(apiKey string, ts int64) ([]byte, error) {
	l := ls.findList(apiKey)

	if l.Ts < ts {
		return nil, ErrOldTimestamp
	}

	return l.Data, nil
}

func (ls *Lists) SetData(apiKey string, data []byte, ts int64) error {
	l := ls.findList(apiKey)

	err := l.setTimestamp(ts)
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
	// write to disk
	return nil
}

func (ls *Lists) findList(apiKey string) *List {
	l := ls.listIndex[apiKey]
	if l == nil {
		l := &List{
			ApiKey: apiKey,
		}
		ls.listIndex[apiKey] = l
	}
	return l
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
