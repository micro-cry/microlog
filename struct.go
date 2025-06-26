package microlog

import "time"

// // // // // // // // // // // // // // // // // //

type LineObj struct {
	UID  [16]byte
	Text string
}

type InstanceObj struct {
	UID  [8]byte
	Text string
}

type PathObj struct {
	UID  [12]byte
	Text string
}

type StatusObj struct {
	UID  [4]byte
	Text string
}

// //

type StreamObj struct {
	UID      [24]byte
	Start    time.Time
	Stop     time.Time
	Instance *InstanceObj
	Line     *LineObj
}

type TimelineObj struct {
	UID       [32]byte
	TimeStamp time.Time
	Instance  *InstanceObj
	Path      *PathObj
	Status    *StatusObj
	Line      *LineObj
}
