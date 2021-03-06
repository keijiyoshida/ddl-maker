package mysql

import (
	"fmt"
	"testing"
)

func TestToSQL(t *testing.T) {
	m := MySQL{}
	args := make(map[string]string, 0)
	args["bool"] = "TINYINT(1)"
	args["int8"] = "TINYINT"
	args["int16"] = "SMALLINT"
	args["int32"] = "INTEGER"
	args["int64"] = "BIGINT"
	args["uint8"] = "TINYINT unsigned"
	args["uint16"] = "SMALLINT unsigned"
	args["uint32"] = "INTEGER unsigned"
	args["uint64"] = "BIGINT unsigned"
	args["float32"] = "FLOAT"
	args["float64"] = "DOUBLE"

	for k, v := range args {
		if m.ToSQL(k, 0) != v {
			t.Fatalf("error %s to sql %s. but result %s", k, v, m.ToSQL(k, 0))
		}
	}

	if m.ToSQL("string", 0) != fmt.Sprintf("VARCHAR(%d)", mysqlDefaultVarcharSize) {
		t.Fatalf("error %s to sql %s. but result %s", "string", fmt.Sprintf("VARCHAR(%d)", mysqlDefaultVarcharSize), m.ToSQL("string", 0))
	}

	if m.ToSQL("*string", 0) != fmt.Sprintf("VARCHAR(%d)", mysqlDefaultVarcharSize) {
		t.Fatalf("error %s to sql %s. but result %s", "*string", fmt.Sprintf("VARCHAR(%d)", mysqlDefaultVarcharSize), m.ToSQL("*string", 0))
	}

	if m.ToSQL("sql.NullString", 0) != fmt.Sprintf("VARCHAR(%d)", mysqlDefaultVarcharSize) {
		t.Fatalf("error %s to sql %s. but result %s", "sql.NullString", fmt.Sprintf("VARCHAR(%d)", mysqlDefaultVarcharSize), m.ToSQL("sql.NullString", 0))
	}

	size := uint64(10)
	if m.ToSQL("sql.NullString", size) != fmt.Sprintf("VARCHAR(%d)", size) {
		t.Fatalf("error %s to sql %s. but result %s", "sql.NullString", fmt.Sprintf("VARCHAR(%d)", size), m.ToSQL("sql.NullString", size))
	}

	if m.ToSQL("text", 0) != "TEXT" {
		t.Fatalf("error %s to sql %s. but result %s", "text", "TEXT", m.ToSQL("text", 0))
	}
}

func TestQuote(t *testing.T) {
	column := "id"

	if quote(column) != "`id`" {
		t.Fatalf("error %s quote. result:%s ", column, quote(column))
	}
}

func TestAuotIncrement(t *testing.T) {
	m := MySQL{}
	if m.AutoIncrement() != mysqlAutoIncreament {
		t.Fatalf("error auto increament: %s. result:%s", mysqlAutoIncreament, m.AutoIncrement())
	}
}

func TestAddIndex(t *testing.T) {
	index := AddIndex("player_id_idx", "player_id")
	if index.ToSQL() != "INDEX `player_id_idx` (`player_id`)" {
		t.Fatal("[error] parse player_id_idx. ", index.ToSQL())
	}

	index = AddIndex("player_entry_id_idx", "player_id", "entry_id")
	if index.ToSQL() != "INDEX `player_entry_id_idx` (`player_id`, `entry_id`)" {
		t.Fatal("[error] parse player_entry_id_idx", index.ToSQL())
	}
}

func TestAddUniqIndex(t *testing.T) {
	uniqIndex := AddUniqueIndex("player_id_idx", "player_id")
	if uniqIndex.ToSQL() != "UNIQUE `player_id_idx` (`player_id`)" {
		t.Fatal("[error] parse unique player_id_idx", uniqIndex.ToSQL())
	}

	uniqIndex = AddUniqueIndex("player_entry_id_idx", "player_id", "entry_id")
	if uniqIndex.ToSQL() != "UNIQUE `player_entry_id_idx` (`player_id`, `entry_id`)" {
		t.Fatal("[error] parse unique player_entry_id_idx", uniqIndex.ToSQL())
	}
}
func TestAddPrimaryKey(t *testing.T) {
	pk := AddPrimaryKey("id")
	if pk.ToSQL() != "PRIMARY KEY (`id`)" {
		t.Fatal("[error] parse primary key", pk.ToSQL())
	}

	pk = AddPrimaryKey("id", "created_at")
	if pk.ToSQL() != "PRIMARY KEY (`id`, `created_at`)" {
		t.Fatal("[error] parse primary key", pk.ToSQL())
	}

	pk = AddPrimaryKey("created_at", "id")
	if pk.ToSQL() != "PRIMARY KEY (`created_at`, `id`)" {
		t.Fatal("[error] parse primary key", pk.ToSQL())
	}
}
