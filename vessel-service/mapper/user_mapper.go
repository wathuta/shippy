package mapper

import (
	"vessel_service/model"
	"vessel_service/proto/vessel"
)

func Marshal_Vessel(vs *vessel.Vessel) *model.Vessel {
	return &model.Vessel{
		Id:        vs.Id,
		Capacity:  vs.Capacity,
		MaxWeight: vs.MaxWeight,
		Name:      vs.Name,
		OwnerId:   vs.OwnerId,
		Available: vs.Available,
	}
}
func Marshal_vessel_collection(vs_collection []*vessel.Vessel) []*model.Vessel {
	ret := make([]*model.Vessel, 0)
	for _, vs := range vs_collection {
		_vs := Marshal_Vessel(vs)
		ret = append(ret, _vs)
	}
	return ret
}
func Unmarshal_Vessel(_vs *model.Vessel) *vessel.Vessel {
	return &vessel.Vessel{
		Id:        _vs.Id,
		Capacity:  _vs.Capacity,
		MaxWeight: _vs.MaxWeight,
		Name:      _vs.Name,
		Available: _vs.Available,
		OwnerId:   _vs.OwnerId,
	}
}
func Unmarshal_vessel_collection(_vs_collection []*model.Vessel) []*vessel.Vessel {
	ret := make([]*vessel.Vessel, 0)
	for _, vs := range _vs_collection {
		_vs := Unmarshal_Vessel(vs)
		ret = append(ret, _vs)
	}
	return ret
}
func Marshal_Spec(sp *vessel.Specification) *model.Specification {
	return &model.Specification{
		Capacity:  sp.Capacity,
		MaxWeight: sp.MaxWeight,
	}
}
func Unmarshal_Spec(_sp *model.Specification) *vessel.Specification {
	return &vessel.Specification{
		Capacity:  _sp.Capacity,
		MaxWeight: _sp.MaxWeight,
	}
}
