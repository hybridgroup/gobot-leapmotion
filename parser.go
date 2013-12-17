package gobotLeap

import (
	"encoding/json"
	"regexp"
)

type LeapGesture struct {
	Direction     []float64
	Duration      int
	HandIDs       []int
	ID            int
	PointableIDs  []int
	Position      []float64
	Speed         float64
	StartPosition []float64
	State         string
	Type          string
}

type LeapHand struct {
	Direction              []float64
	ID                     int
	PalmNormal             []float64
	PalmPosition           []float64
	PalmVelocity           []float64
	R                      [][]float64
	S                      float64
	SphereCenter           []float64
	SphereRadius           float64
	StabilizedPalmPosition []float64
	T                      []float64
	TimeVisible            float64
}

type LeapPointable struct {
	Direction             []float64
	HandID                int
	ID                    int
	Length                float64
	StabilizedTipPosition []float64
	TimeVisible           float64
	TipPosition           []float64
	TipVelocity           []float64
	Tool                  bool
	TouchDistance         float64
	TouchZone             string
}

type LeapFrame struct {
	CurrentFrameRate float64
	Gestures         []LeapGesture
	Hands            []LeapHand
	ID               int
	InteractionBox   struct {
		Center []int
		Size   []float64
	}
	Pointables []LeapPointable
	R          [][]float64
	S          float64
	T          []float64
	Timestamp  int
}

func ParseLeapFrame(data []byte) LeapFrame {
	if !isAFrame(data) {
		return LeapFrame{}
	}
	var frame LeapFrame
	json.Unmarshal(data, &frame)
	return frame
}

func isAFrame(data []byte) bool {
	match, _ := regexp.Match("currentFrameRate", data)
	return match
}
