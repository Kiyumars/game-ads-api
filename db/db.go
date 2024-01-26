package db

import "github.com/hashicorp/go-memdb"

func Create() (*memdb.MemDB, error) {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"ad": &memdb.TableSchema{
				Name: "ad",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.IntFieldIndex{Field: "id"},
					},
					"filetype": &memdb.IndexSchema{
						Name:    "filetype",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "filetype"},
					},
					"ref": &memdb.IndexSchema{
						Name:    "ref",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ref"},
					},
					"size": &memdb.IndexSchema{
						Name:    "size",
						Unique:  false,
						Indexer: &memdb.IntFieldIndex{Field: "size"},
					},
					"adult": &memdb.IndexSchema{
						Name:    "adult",
						Unique:  false,
						Indexer: &memdb.BoolFieldIndex{Field: "adult"},
					},
					"children": &memdb.IndexSchema{
						Name: "children",
						Unique: false,
						Indexer: &memdb.BoolFieldIndex{Field: "children"},
					},
				},
			},
		},
	}
	return memdb.NewMemDB(schema)
}
