package uolo_center

import "artemis/component/avatar_keeper"

type (
	AvatarHandler struct {
	}
)

func NewAvatarHandler(conf OptionConfig) *AvatarHandler {
	keeper := avatar_keeper.NewAvatarKeeperWithFileStorage(conf.Path)
	return &AvatarHandler{
		config: conf,
		keeper: keeper,
	}
}
