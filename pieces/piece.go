package pieces

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"tetris/component"
)

const (
	BAR = iota
	BOX
	LEFT_L
	LEFT_Z
	RIGHT_L
	RIGHT_Z
	TEE
)

type Piece interface {
	Position() (int, int)
	Pixels() []component.Pixel
	NewPosition(int, int)
	Rotate()
	RotateBack()
}

type SerializablePiece struct {
	PieceType int
	Piece
}

func (p *SerializablePiece) UnmarshalJSON(data []byte) error {
	var jsonObject map[string]interface{}
	err := json.Unmarshal(data, &jsonObject)
	if err != nil {
		fmt.Println("Error deserializing piece")
	}
	pieceType := int(jsonObject["PieceType"].(float64))
	// TODO: There shouldn't be a need to marshal it again
	pieceData, _ := json.Marshal(jsonObject["Piece"])
	switch pieceType {
	case BAR:
		var piece Bar
		json.Unmarshal(pieceData, &piece)
		*p = SerializablePiece{pieceType, &piece}
	case BOX:
		var piece Box
		json.Unmarshal(pieceData, &piece)
		*p = SerializablePiece{pieceType, &piece}
	case LEFT_L:
		var piece LeftL
		json.Unmarshal(pieceData, &piece)
		*p = SerializablePiece{pieceType, &piece}
	case LEFT_Z:
		var piece LeftZ
		json.Unmarshal(pieceData, &piece)
		*p = SerializablePiece{pieceType, &piece}
	case RIGHT_L:
		var piece RightL
		json.Unmarshal(pieceData, &piece)
		*p = SerializablePiece{pieceType, &piece}
	case RIGHT_Z:
		var piece RightZ
		json.Unmarshal(pieceData, &piece)
		*p = SerializablePiece{pieceType, &piece}
	case TEE:
		var piece Tee
		json.Unmarshal(pieceData, &piece)
		*p = SerializablePiece{pieceType, &piece}
	}
	return nil
}

func NewRandomPiece() *SerializablePiece {
	c := rand.IntN(7)
	switch c {
	case 0:
		return &SerializablePiece{BAR, NewBar()}
	case 1:
		return &SerializablePiece{BOX, NewBox()}
	case 2:
		return &SerializablePiece{LEFT_L, NewLeftL()}
	case 3:
		return &SerializablePiece{LEFT_Z, NewLeftZ()}
	case 4:
		return &SerializablePiece{RIGHT_L, NewRightL()}
	case 5:
		return &SerializablePiece{RIGHT_Z, NewRightZ()}
	default:
		return &SerializablePiece{TEE, NewTee()}
	}
}
