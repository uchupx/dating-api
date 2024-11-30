package helper

const (
	// enum redis key
	REDIS_KEY_AUTH = "auth:key"

	// enum reaction
	REACTION_LIKE    int8 = 1
	REACTION_DISLIKE int8 = 2
)

func ValidateReaction(reaction int8) bool {
	var reactions = []interface{}{
		REACTION_LIKE,
		REACTION_DISLIKE,
	}

	return Contains(reactions, reaction)
}
