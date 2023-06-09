// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package rewards

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

type RewardStatus string

const (
	RewardStatusPending RewardStatus = "pending"
	RewardStatusSuccess RewardStatus = "success"
	RewardStatusFailed  RewardStatus = "failed"
)

func (e *RewardStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = RewardStatus(s)
	case string:
		*e = RewardStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for RewardStatus: %T", src)
	}
	return nil
}

type NullRewardStatus struct {
	RewardStatus RewardStatus
	Valid        bool // Valid is true if RewardStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullRewardStatus) Scan(value interface{}) error {
	if value == nil {
		ns.RewardStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.RewardStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullRewardStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.RewardStatus), nil
}

type RewardTypes string

const (
	RewardTypesR1 RewardTypes = "r1"
	RewardTypesR2 RewardTypes = "r2"
	RewardTypesR3 RewardTypes = "r3"
)

func (e *RewardTypes) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = RewardTypes(s)
	case string:
		*e = RewardTypes(s)
	default:
		return fmt.Errorf("unsupported scan type for RewardTypes: %T", src)
	}
	return nil
}

type NullRewardTypes struct {
	RewardTypes RewardTypes
	Valid       bool // Valid is true if RewardTypes is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullRewardTypes) Scan(value interface{}) error {
	if value == nil {
		ns.RewardTypes, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.RewardTypes.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullRewardTypes) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.RewardTypes), nil
}

type ScratchCard struct {
	ID              int64          `json:"id"`
	Schedule        sql.NullString `json:"schedule"`
	MaxCards        sql.NullInt32  `json:"max_cards"`
	MaxCardsPerUser sql.NullInt32  `json:"max_cards_per_user"`
	Weight          int32          `json:"weight"`
	RewardType      RewardTypes    `json:"reward_type"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
}

type ScratchCardsReward struct {
	ID            int64        `json:"id"`
	ScratchCardID int64        `json:"scratch_card_id"`
	UserID        int64        `json:"user_id"`
	OrderID       string       `json:"order_id"`
	Status        RewardStatus `json:"status"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
}

type User struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	ScratchCards int32     `json:"scratch_cards"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
