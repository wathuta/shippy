package services

import (
	"context"
	"reflect"
	"testing"
	"vessel_service/proto/vessel"
	"vessel_service/repository"

	mocks "vessel_service/services/mocks"

	"github.com/sirupsen/logrus"
)

func TestNew_vessel_service(t *testing.T) {
	type args struct {
		repo  repository.Vessel_respository
		loger *logrus.Entry
	}
	tests := []struct {
		name string
		args args
		want VesselServiceHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New_vessel_service(tt.args.repo, tt.args.loger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New_vessel_service() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_vesselServiceHandler_FindAvailableVessel(t *testing.T) {
	type args struct {
		ctx context.Context
		req *vessel.Specification
		res *vessel.Response
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "FindAvailable vessel",
			args: args{
				ctx: context.Background(),
				req: &vessel.Specification{
					Capacity:  20,
					MaxWeight: 330,
				},
				res: &vessel.Response{},
			},
		},
	}
	for _, tt := range tests {
		// t.Run(tt.name, func(t *testing.T) {
		// 	v := &vesselServiceHandler{
		// 		repo:  tt.fields.repo,
		// 		loger: tt.fields.loger,
		// 	}
		// 	if err := v.FindAvailableVessel(tt.args.ctx, tt.args.req, tt.args.res); (err != nil) != tt.wantErr {
		// 		t.Errorf("vesselServiceHandler.FindAvailableVessel() error = %v, wantErr %v", err, tt.wantErr)
		// 	}
		// })
		t.Run(tt.name, func(t *testing.T) {

			repo := &mocks.VesselServiceHandler{}
			if !tt.wantErr{
				
			}
		})
	}
}

func Test_vesselServiceHandler_CreateVessel(t *testing.T) {
	type fields struct {
		repo  repository.Vessel_respository
		loger *logrus.Entry
	}
	type args struct {
		ctx context.Context
		in  *vessel.Vessel
		out *vessel.Response
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &vesselServiceHandler{
				repo:  tt.fields.repo,
				loger: tt.fields.loger,
			}
			if err := v.CreateVessel(tt.args.ctx, tt.args.in, tt.args.out); (err != nil) != tt.wantErr {
				t.Errorf("vesselServiceHandler.CreateVessel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
