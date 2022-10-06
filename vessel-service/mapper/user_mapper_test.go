package mapper

import (
	"reflect"
	"testing"
	"vessel_service/model"
	"vessel_service/proto/vessel"
)

func TestMarshal_Vessel(t *testing.T) {
	type args struct {
		vs *vessel.Vessel
	}
	tests := []struct {
		name string
		args args
		want *model.Vessel
	}{{
		name: "Marshal vessel",
		args: args{
			vs: &vessel.Vessel{
				Id:        "New Vessel",
				Capacity:  12,
				MaxWeight: 50,
				Name:      "Vessel",
				Available: true,
				OwnerId:   "Brian",
			},
		},
		want: &model.Vessel{
			Id:        "New Vessel",
			Capacity:  12,
			MaxWeight: 50,
			Name:      "Vessel",
			Available: true,
			OwnerId:   "Brian",
		},
	},
	{
		name: "Marshal vessel",
		args: args{
			vs: &vessel.Vessel{
				Id:        "New Vessel",
				Capacity:  12,
				MaxWeight: 50,
				Name:      "Vessel",
				Available: true,
				OwnerId:   "Brian",
			},
		},
		want: &model.Vessel{
			Id:        "New Vessel",
			Capacity:  12,
			MaxWeight: 50,
			Name:      "Vessel",
			Available: true,
			OwnerId:   "Brian",
		},
	},
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Marshal_Vessel(tt.args.vs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Marshal_Vessel() = %v, want %v", got, tt.want)
			}
		})

	}
}

func TestMarshal_vessel_collection(t *testing.T) {
	type args struct {
		vs_collection []*vessel.Vessel
	}
	tests := []struct {
		name string
		args args
		want []*model.Vessel
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Marshal_vessel_collection(tt.args.vs_collection); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Marshal_vessel_collection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnmarshal_Vessel(t *testing.T) {
	type args struct {
		_vs *model.Vessel
	}
	tests := []struct {
		name string
		args args
		want *vessel.Vessel
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Unmarshal_Vessel(tt.args._vs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unmarshal_Vessel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnmarshal_vessel_collection(t *testing.T) {
	type args struct {
		_vs_collection []*model.Vessel
	}
	tests := []struct {
		name string
		args args
		want []*vessel.Vessel
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Unmarshal_vessel_collection(tt.args._vs_collection); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unmarshal_vessel_collection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarshal_Spec(t *testing.T) {
	type args struct {
		sp *vessel.Specification
	}
	tests := []struct {
		name string
		args args
		want *model.Specification
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Marshal_Spec(tt.args.sp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Marshal_Spec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnmarshal_Spec(t *testing.T) {
	type args struct {
		_sp *model.Specification
	}
	tests := []struct {
		name string
		args args
		want *vessel.Specification
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Unmarshal_Spec(tt.args._sp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unmarshal_Spec() = %v, want %v", got, tt.want)
			}
		})
	}
}
