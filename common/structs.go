package common

type Empty struct {
	// Success bool
	// Error   error // else Error string
}

type MapInput struct {
	NoOfReducers int
	DataFileCid  string
	MasterPeerId string
	Round        int
}

type MapOutput struct {
	DataFileCid string
	KvFileCids  []string
}

type KeyValue struct {
	Key   string
	Value string
}

type ReduceInput struct {
	KvFileCids   []string
	MasterPeerId string
	ReducerNo    int
	Round        int
}

type ReduceOutput struct {
	ReducerNo     int
	OutputFileCid string
}
