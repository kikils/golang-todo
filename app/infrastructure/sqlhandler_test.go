package infrastructure

import (
	"reflect"
	"testing"
)

func Test_Sqlhandler(t *testing.T) {
	type args struct {
		id   int
		name string
	}
	tests := []struct {
		name string
		args args
		want args
	}{
		{
			name: "normal",
			args: args{id: 0, name: "test"},
			want: args{id: 0, name: "test"},
		},
	}
	handler := NewSqlhandler()
	_, err := handler.Execute("DROP TABLE IF EXISTS foo;")
	if err != nil {
		t.Fatal(err)
	}
	_, err = handler.Execute("CREATE TABLE foo (id integer, name varchar(42));")
	if err != nil {
		t.Fatal(err)
	}
	insertQuery := "INSERT INTO foo (id, name) VALUES ($1, $2);"

	for _, tt := range tests {
		_, err = handler.Execute(insertQuery, tt.args.id, tt.args.name)
		if err != nil {
			t.Fatal(err)
		}
		var got args
		row, err := handler.Query("SELECT * FROM foo LIMIT 1;")
		if err != nil {
			t.Fatal(err)
		}
		row.Next()
		err = row.Scan(&got.id, &got.name)
		if err != nil {
			t.Fatal(err)
		}
		t.Run(tt.name, func(t *testing.T) {
			if !isEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func isEqual(got interface{}, want interface{}) bool {
	return reflect.DeepEqual(got, want)
}
