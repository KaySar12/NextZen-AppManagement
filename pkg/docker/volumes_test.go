package docker_test

import (
	"fmt"
	"testing"

	"github.com/KaySar12/NextZen-AppManagement/pkg/docker"
)

func TestGetDir(t *testing.T) {
	fmt.Println(docker.GetDir("", "config"))
}
