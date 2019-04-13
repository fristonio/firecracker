package utils

import (
	"errors"
	"fmt"

	"github.com/sdslabs/beastv4/core"
	"github.com/sdslabs/beastv4/database"
	"github.com/sdslabs/beastv4/docker"
	tools "github.com/sdslabs/beastv4/utils"
)

func GetLogs(challname string, live bool) (*docker.Log, error) {
	chall, err := database.QueryFirstChallengeEntry("name", challname)
	if err != nil {
		return nil, fmt.Errorf("Error while database access : %s", err)
	}
	if chall.Format == core.STATIC_CHALLENGE_TYPE_NAME {
		return nil, fmt.Errorf("The challenge is a static challenge")
	}
	if !tools.IsContainerIdValid(chall.ContainerId) {
		return nil, fmt.Errorf("The container id is not valid")
	}
	containers, err := docker.SearchContainerByFilter(map[string]string{"id": chall.ContainerId})
	if err != nil {
		return nil, fmt.Errorf("Error while searching for container with id %s", chall.ContainerId)
	}

	if len(containers) > 1 {
		return nil, errors.New("Got more than one containers, something fishy here. Contact admin to check manually.")
	}

	if len(containers) == 0 {
		return nil, fmt.Errorf("The container does not exist")
	}

	if live {
		docker.ShowLiveDockerLogs(chall.ContainerId)
		return nil, nil
	}

	return docker.GetDockerStdLogs(chall.ContainerId)
}
