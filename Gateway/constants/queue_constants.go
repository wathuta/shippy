package constants

const (
	CREATE_VESSEL_QUEUE   = "create_vessel"
	FIND_ALL_VESSEL_QUEUE = "find_all_vessel"
)

const (
	CREATE_VESSEL_CONSUMER_TAG   = "VESSEL_CREATION"
	FIND_ALL_VESSEL_CONSUMER_TAG = "VESSEL_FIND_ALL"
)

const (
	CREATE_VESSEL_ROUTING_KEY   = "one.create.new"
	FIND_ALL_VESSEL_ROUTING_KEY = "all.get.vessels"
)
