package api

import (
	"fmt"
	"os"
	"os/user"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/khulnasoft/shipyard/pkg/logging"
	"github.com/khulnasoft/shipyard/pkg/logging/logfields"
)

var log = logging.DefaultLogger.WithField(logfields.LogSubsys, "api")

// getGroupIDByName returns the group ID for the given grpName.
func getGroupIDByName(grpName string) (int, error) {
	group, err := user.LookupGroup(grpName)
	if err != nil {
		return -1, err
	}
	return strconv.Atoi(group.Gid)
}

// SetDefaultPermissions sets the given socket's group to `ShipyardGroupName` and
// mode to `SocketFileMode`.
func SetDefaultPermissions(socketPath string) error {
	gid, err := getGroupIDByName(ShipyardGroupName)
	if err != nil {
		log.WithError(err).WithFields(logrus.Fields{
			logfields.Path: socketPath,
			"group":        ShipyardGroupName,
		}).Debug("Group not found")
	} else {
		if err := os.Chown(socketPath, 0, gid); err != nil {
			return fmt.Errorf("failed while setting up %s's group ID"+
				" in %q: %s", ShipyardGroupName, socketPath, err)
		}
	}
	if err := os.Chmod(socketPath, SocketFileMode); err != nil {
		return fmt.Errorf("failed while setting up file permissions in %q: %w",
			socketPath, err)
	}
	return nil
}
