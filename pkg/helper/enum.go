package helper

const (
	// enum redis key
	REDIS_KEY_AUTH      = "auth:key"
	REDIS_KEY_USER_ID   = "user:%s"
	REDIS_KEY_USER_VIEW = "user:view:%s"

	// enum reaction
	REACTION_LIKE    int8 = 1
	REACTION_DISLIKE int8 = 2

	// Generate enum string
	ACTIVE_INT   int8 = 1
	INACTIVE_INT int8 = 0

	// enum Feature
	FEATURE_VERIFIED       int8 = 1
	FEATURE_NO_SWIPE_QUOTA int8 = 2

	FEATURE_VERIFIED_STRING       string = "VERIFIED"
	FEATURE_NO_SWIPE_QUOTA_STRING string = "NO_SWIPE_QUOTA"
)

var FEATURE_MAP = map[int8]string{
	FEATURE_VERIFIED:       FEATURE_VERIFIED_STRING,
	FEATURE_NO_SWIPE_QUOTA: FEATURE_NO_SWIPE_QUOTA_STRING,
}

func ValidateReaction(reaction int8) bool {
	var reactions = []interface{}{
		REACTION_LIKE,
		REACTION_DISLIKE,
	}

	return Contains(reactions, reaction)
}
