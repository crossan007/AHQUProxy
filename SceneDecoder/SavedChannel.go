package main

type SavedChannel struct {
	Index           int
	Offset          int
	Id              int
	Type            int
	PhysicalName    string
	DisplayName     string
	GainValue       KnobValue
	FaderValue      FaderValue
	RawValue        string
	EQ              ChannelEQ
	Gate            ChannelGate
	Compression     ChannelCompression
	SendToMainFader FaderValue
}

type ChannelEQ struct {
	Enabled        string
	HighPassFilter FaderValue
	Band1          EQBand
	Band2          EQBand
	Band3          EQBand
	Band4          EQBand
	RawBytes       string
}

type ChannelGate struct {
	Enabled   string
	Threshold KnobValue
	Depth     KnobValue
	Attack    FaderValue
	Release   KnobValue
	Hold      KnobValue
	RawBytes  string
}

type ChannelCompression struct {
	Enabled  string
	RawBytes string
}
