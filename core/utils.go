package core

import (
	"errors"

	"github.com/sdslabs/beastv4/core/database"
	"github.com/sdslabs/beastv4/docker"
	log "github.com/sirupsen/logrus"
)

func CleanupContainerByFilter(filter, filterVal string) error {
	if filter != "id" && filter != "name" {
		return fmt.Errorf("Not a valid filter %s", filter)
	}

	containers, err := docker.SearchContainerByFilter(map[string]string{filter: filterVal})
	if err != nil {
		log.Error("Error while searching for container with %s : ", filter, filterVal)
		return err
	}

	var erroredContainers []string
	if len(containers) != 0 {
		log.Infof("Cleaning up container with %s %s", filter, filterVal)
		for _, container := range containers {
			err = docker.StopAndRemoveContainer(container.ID)
			if err != nil {
				erroredContainers.append(container.ID)
				log.Errorf("Error while cleaning up container %s : %s", container.ID, err)
			}
		}
	}

	if len(erroredContainers) != 0 {
		return fmt.Errorf("Error wihle cleaning up container : %s", erroredContainers)
	}

	return nil
}

func CleanupChallengeContainers(chall database.Challenge, config BeastConfig) error {
	if chall.ContainerId != "" {
		err = CleanupContainerByFilter("id", chall.ContainerId)
		if err != nil {
			return err
		}

		chall.ContainerId = ""
		database.Db.Save(&chall)
	}

	err = CleanupContainerByFilter("name", config.Challenge.Name)
	return err
}

func CleanupChallengeImage(chall database.Challenge) error {
	err := docker.RemoveImage(chall.ImageId)
	if err != nil {
		log.Error("Error while cleaning up image with id ", chall.ImageId)
		return err
	}

	chall.ImageId = ""
	database.Db.Save(&chall)

	return nil
}

func CleanupChallengeIfExist(config BeastConfig) error {
	chall, err := database.QueryFirstChallengeEntry("challenge_id", config.Challenge.ChallengeId)
	if err != nil {
		log.Errorf("Error while database query for challenge Id %s", config.Challenge.ChallengeId)
		return err
	}

	if len(chall) == 0 {
		log.Info("No such challenge exist in the database")
		return nil
	}

	err := CleanupChallengeContainers(chall, config)
	if err != nil {
		return err
	}

	if chall.ImageId == "" {
		log.Warn("Looks like we don't have the image ID in database for challenge, Nothing to remove")
		return nil
	}
	err = CleanupChallengeImage(chall)
	return err
}
