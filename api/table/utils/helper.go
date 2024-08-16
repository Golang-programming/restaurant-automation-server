package utils

import (
	"fmt"
	"strconv"

	"github.co/golang-programming/restaurant/api/redis"
)

const availableTablesKey = "tables:availables"
const observedTablesKey = "tables:observed"

func MarkTableAvailable(tableID uint) error {
	redis.SRem(observedTablesKey, tableID)
	return redis.SAdd(availableTablesKey, tableID)
}

func MarkTableObserved(tableID uint) error {
	redis.SRem(availableTablesKey, tableID)
	return redis.SAdd(observedTablesKey, tableID)
}

func GetAvaiableTableIdz() ([]uint, error) {
	members, err := redis.SMembers(availableTablesKey)
	if err != nil {
		return nil, err
	}

	var tableIDz []uint
	for _, member := range members {
		id, err := strconv.ParseUint(member, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("failed to parse tableID: %v", err)
		}
		tableIDz = append(tableIDz, uint(id))
	}

	return tableIDz, nil
}
