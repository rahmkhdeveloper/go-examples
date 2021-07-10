package model

type IUser interface {
	GetMyInitials() (byte, byte)
	WriteMyInitials(byte, byte)
	WriteMyNameAndInitials() string
}
