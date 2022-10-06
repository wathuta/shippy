package configs

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestCreateConn(t *testing.T) {
	type args struct {
		ctx     context.Context
		uri     string
		retries int32
	}
	tests := []struct {
		name    string
		args    args
		want    *mongo.Client
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "connections Test",
			args: args{
				ctx:     context.Background(),
				uri:     "mongodb+srv://brayo:brayo@myblogcluster.976g4.mongodb.net/shippy?retryWrites=true&w=majority",
				retries: 3,
			},
			want:    &mongo.Client{},
			wantErr: false,
		}, {
			name: "wrong uri Test",
			args: args{
				ctx:     context.Background(),
				uri:     "",
				retries: 3,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// got, err := CreateConn(tt.args.ctx, tt.args.uri, tt.args.retries)
			// if (err != nil) != tt.wantErr {
			// 	t.Errorf("CreateConn() error = %v, wantErr %v", err, tt.wantErr)
			// 	return
			// }
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("CreateConn() = %v, want %v", got, tt.want)
			// }
			got, _ := CreateConn(tt.args.ctx, tt.args.uri, tt.args.retries)
			assert.Equal(t, tt.want, got, tt.name)

		})

	}
}

func TestNewRabbitConfigs(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name string
		want *Config
	}{
		{
			name: "config",
			want: &Config{
				Host:        "localhost",
				Port:        "5672",
				User:        "guest",
				Password:    "guest",
				Exchange:    "Vessel",
				Queue:       "VesselService",
				RoutingKey:  "vessel.create",
				ConsumerTag: "vessel",
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		// t.Run(tt.name, func(t *testing.T) {
		// 	if got := NewRabbitConfigs(); !reflect.DeepEqual(got, tt.want) {
		// 		t.Errorf("NewRabbitConfigs() = %v, want %v", got, tt.want)
		// 	}
		// })
		got := NewRabbitConfigs()
		assert.Equal(tt.want, got, tt.name)
	}
}
