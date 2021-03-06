/**
 * exec.go - Exec external process with timeout
 *
 * @author Yaroslav Pogrebnyak <yyyaroslav@gmail.com>
 * @author Ievgen Ponomarenko <kikomdev@gmail.com>
 */

package utils

import (
	"../logging"
	"os/exec"
	"time"
)

/**
 * Exec with timeout
 */
func ExecTimeout(timeout time.Duration, params ...string) (string, error) {

	log := logging.For("execTimeout")

	cmd := exec.Command(params[0], params[1:]...)

	timer := time.AfterFunc(timeout, func() {
		log.Info("Responce from exec ", params, " is timed out. Killing process...")
		cmd.Process.Kill()
	})

	out, err := cmd.Output()
	timer.Stop()

	if err != nil {
		return "", err
	}

	return string(out), nil
}
