package data

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
)

type Lists struct {
	ListIndex map[string]*List
}

type List struct {
	ItemIndex map[string]*Item
	ApiKey    string
}

type Item struct {
	Ts       int64
	GUID     string
	Read     bool
	Bookmark bool
}

func (ls *Lists) SyncList(list *List) (*List, error) {
	l, err := ls.findList(list.ApiKey)
	if err != nil {
		return nil, err
	}

	if l.ItemIndex == nil {
		l.ItemIndex = make(map[string]*Item)
	}

	for guid, item := range list.ItemIndex {
		if existing, ok := l.ItemIndex[guid]; ok {
			if existing.Ts < item.Ts {
				slog.Info("Update", "i", existing)
				newItem := *item
				l.ItemIndex[guid] = &newItem
			}
		} else {
			slog.Info("Create", "i", item)
			l.ItemIndex[guid] = item
		}
	}

	return l, l.Save()
}

func (ls *Lists) findList(apiKey string) (*List, error) {
	if apiKey == "" {
		return nil, ErrNoApiKey
	}

	if ls.ListIndex == nil {
		ls.ListIndex = make(map[string]*List)
	}

	l := ls.ListIndex[apiKey]
	if l == nil {
		l = &List{
			ApiKey:    apiKey,
			ItemIndex: make(map[string]*Item),
		}
		ls.ListIndex[apiKey] = l
	}

	return l, nil
}

func (l *List) Save() error {
	dataFilePath, err := l.DataFilePath()
	if err != nil {
		return err
	}

	f, err := os.Create(dataFilePath)
	if err != nil {
		return err
	}
	defer f.Close()

	data, err := json.Marshal(l)
	if err != nil {
		return err
	}

	_, err = f.Write(data)
	if err == nil {
		slog.Info("Save", "l", dataFilePath)
	}

	return err
}

func Load() *Lists {
	// TODO: load from disk
	return &Lists{ListIndex: make(map[string]*List)}
}

func (l *List) DataFilePath() (string, error) {
	dir, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}

	appDir := filepath.Join(dir, "rssr-sync")
	if err := os.MkdirAll(appDir, 0755); err != nil {
		return "", err
	}

	filename := fmt.Sprintf("state-%s.json", l.ApiKey)
	return filepath.Join(appDir, filename), nil
}
