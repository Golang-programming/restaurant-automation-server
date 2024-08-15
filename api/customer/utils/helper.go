package utils

import (
	"fmt"
	"strconv"

	"github.co/golang-programming/restaurant/api/redis"
)

const activeUsersKey = "users:active"

func AddActiveUser(userID uint) error {
	return redis.SAdd(activeUsersKey, userID)
}

func RemoveActiveUser(userID uint) error {
	return redis.SRem(activeUsersKey, userID)
}

func GetActiveUsers() ([]uint, error) {
	members, err := redis.SMembers(activeUsersKey)
	if err != nil {
		return nil, err
	}

	var userIDs []uint
	for _, member := range members {
		id, err := strconv.ParseUint(member, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("failed to parse userID: %v", err)
		}
		userIDs = append(userIDs, uint(id))
	}

	return userIDs, nil
}
