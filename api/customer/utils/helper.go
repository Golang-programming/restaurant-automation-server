/* package utils

import (
	"fmt"


	"github.co/golang-programming/restaurant/api/redis"
)

const activeUsersKey = "active_users"

func AddActiveUser(userID uint) error {
	// Convert the userID to a string
	member := fmt.Sprintf("%d", userID)

	// Add the userID to the Redis set
	err := redis.SAdd(activeUsersKey, member)
	if err != nil {
		return fmt.Errorf("failed to add active user: %v", err)
	}
	return nil
}

func RemoveActiveUser(userID uint) error {
	// Convert the userID to a string
	member := fmt.Sprintf("%d", userID)

	// Remove the userID from the Redis set
	err := redis.SRem(activeUsersKey, member)
	if err != nil {
		return fmt.Errorf("failed to remove active user: %v", err)
	}
	return nil
}

func GetActiveUsers() ([]uint, error) {
	// Get all active userIDs from the Redis set
	members, err := redis.GetRedisClient().SMembers(activeUsersKey).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get active users: %v", err)
	}

	// Convert the result from strings to uints
	var userIDs []uint
	for _, member := range members {
		var userID uint
		_, err := fmt.Sscanf(member, "%d", &userID)
		if err != nil {
			return nil, fmt.Errorf("failed to parse userID: %v", err)
		}
		userIDs = append(userIDs, userID)
	}

	return userIDs, nil
}
*/

package utils

import (
	"fmt"
	"github.co/golang-programming/restaurant/api/redis"
	"strconv"
)

const activeUsersKey = "active_users"

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
