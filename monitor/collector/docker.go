package collector

import (
    "github.com/shirou/gopsutil/docker"
)

type Docker struct {
    DockersStat []docker.CgroupDockerStat `json:"dockers_stat"`
}

func (d Docker) Gather() *Docker {
    var err error
    if d.DockersStat, err = docker.GetDockerStat(); err != nil {
        d.DockersStat = []docker.CgroupDockerStat{}
    }
    return &d
}
