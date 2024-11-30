package dating

import (
	"github.com/uchupx/dating-api/config"
	"github.com/uchupx/dating-api/src/repo"
)

type datingRepo struct {
	userRepo         *repo.UserRepo
	clientRepo       *repo.ClientRepo
	refreshTokenRepo *repo.RefreshTokenRepo
	reactionRepo     *repo.ReactionRepo
}

func (i *Dating) UserRepo(conf *config.Config) *repo.UserRepo {
	if i.userRepo == nil {
		i.userRepo = repo.NewUserRepo(i.DB(conf))
	}

	return i.userRepo
}

func (i *Dating) ClientRepo(conf *config.Config) *repo.ClientRepo {
	if i.clientRepo == nil {
		i.clientRepo = repo.NewClientRepo(i.DB(conf))
	}

	return i.clientRepo
}

func (i *Dating) RefreshTokenRepo(conf *config.Config) *repo.RefreshTokenRepo {
	if i.refreshTokenRepo == nil {
		i.refreshTokenRepo = repo.NewRefreshTokenRepo(i.DB(conf))
	}

	return i.refreshTokenRepo
}

func (i *Dating) ReactionRepo(conf *config.Config) *repo.ReactionRepo {
	if i.reactionRepo == nil {
		i.reactionRepo = repo.NewReactionRepo(i.DB(conf))
	}

	return i.reactionRepo
}
