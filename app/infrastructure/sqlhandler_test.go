package infrastructure

import (
	"reflect"
	"testing"
)

func Test_ConnectPostgres(t *testing.T) {
	_, err := ConnectPostgres()
	if err != nil {
		t.Fatal("DB connection error has occured.")
	}
}

func Test_Sqlhandler(t *testing.T) {
	type args struct {
		id int
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
	handler, err := ConnectPostgres()
	if err != nil {
		t.Fatal(err)
	}
	err = handler.Execute("DROP TABLE IF EXISTS foo;")
	if err != nil {
		t.Fatal(err)
	}
	handler.Execute("CREATE TABLE foo (id integer, name varchar(42));")
	insertQuery := "INSERT INTO foo (id, name) VALUES ($1, $2);"

	for _, tt := range tests {
		err = handler.Execute(insertQuery, tt.args.id, tt.args.name)
		if err != nil {
			t.Fatal(err)
		}
		var got args
		row := handler.Query("SELECT * FROM foo LIMIT 1;")
		row.Next()
		row.Scan(&got.id, &got.name)
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