package configs_Package_Test

import (
	"BiTaksi_Backend_Driver/configs"
	"github.com/stretchr/testify/assert"
	"testing"
)

var db = configs.ConnectDB()
var collection = configs.GetCollection(db, "driver")

func TestConnectDB(t *testing.T) {

	assert.NotNil(t, db)
}

func TestGetCollection(t *testing.T) {
	assert.NotNil(t, collection)
	assert.NotEmpty(t, collection)
}

func TestGetCollection2(t *testing.T) {
	want := configs.GetCollection(db, "driveasdsadsdassdasadrs")
	assert.NotNil(t, want)
	assert.NotEqual(t, want, collection)
}

func TestCreateLocationData(t *testing.T) {
	tests := []struct {
		name string
		want error
	}{
		{"Example", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := configs.CreateLocationData(); err != nil {
				t.Errorf("Fatal: %v ", err)
			}
		})
	}
}

func TestEnvMongoURI(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"ENV TEST", "mongodb.ahmetikrdg:1532sifre@cluster0.aoxmb.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := configs.EnvMongoURI(); got != tt.want {
				assert.NotEqual(t, got, tt.want)
			}
		})
	}
}

func TestEnvMongoURI2(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"ENV TEST", "mongodb+srv://ahmetikrdg:1532sifre@cluster0.aoxmb.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := configs.EnvMongoURI(); got != tt.want {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestEnvMongoURI1(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{

		{"Get ENV file", "mongodb+srv"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := configs.EnvMongoURI(); got[:11] != tt.want {
				t.Errorf("EnvMongoURI() = %v, want %v", got, tt.want)
			}
		})
	}
}
