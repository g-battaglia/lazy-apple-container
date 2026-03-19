package commands

import (
	"context"
	"fmt"
	"os/exec"
	"strings"

	"github.com/go-errors/errors"
	"github.com/jesseduffield/lazycontainer/pkg/i18n"
	"github.com/jesseduffield/lazycontainer/pkg/utils"
	"github.com/sasha-s/go-deadlock"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
)

type Container struct {
	Name            string
	ID              string
	AppleContainer  AppleContainer
	Client          *ContainerClient
	OSCommand       *OSCommand
	Log             *logrus.Entry
	StatHistory     []*RecordedStats
	MonitoringStats bool
	Tr              *i18n.TranslationSet

	StatsMutex deadlock.Mutex
}

func (c *Container) Remove(force bool) error {
	c.Log.Warn(fmt.Sprintf("removing container %s", c.Name))
	if err := c.Client.RemoveContainer(c.ID, force); err != nil {
		if strings.Contains(err.Error(), "Stop the container before attempting removal") {
			return ComplexError{
				Code:    MustStopContainer,
				Message: err.Error(),
				frame:   xerrors.Caller(1),
			}
		}
		return err
	}
	return nil
}

func (c *Container) Start() error {
	c.Log.Warn(fmt.Sprintf("starting container %s", c.Name))
	return c.Client.StartContainer(c.ID)
}

func (c *Container) Stop() error {
	c.Log.Warn(fmt.Sprintf("stopping container %s", c.Name))
	return c.Client.StopContainer(c.ID)
}

func (c *Container) Pause() error {
	c.Log.Warn(fmt.Sprintf("pausing container %s", c.Name))
	return errors.New("pause is not supported by Apple Container")
}

func (c *Container) Unpause() error {
	c.Log.Warn(fmt.Sprintf("unpausing container %s", c.Name))
	return errors.New("unpause is not supported by Apple Container")
}

func (c *Container) Restart() error {
	c.Log.Warn(fmt.Sprintf("restarting container %s", c.Name))
	return c.Client.RestartContainer(c.ID)
}

func (c *Container) Kill(signal string) error {
	c.Log.Warn(fmt.Sprintf("killing container %s with signal %s", c.Name, signal))
	return c.Client.KillContainer(c.ID, signal)
}

func (c *Container) Attach() (*exec.Cmd, error) {
	c.Log.Warn(fmt.Sprintf("attaching to container %s", c.Name))
	if c.AppleContainer.Status != "running" {
		return nil, errors.New(c.Tr.CannotAttachStoppedContainerError)
	}
	return c.Client.ExecContainer(c.ID, ExecOptions{
		Command:     []string{c.getShell()},
		Interactive: true,
		TTY:         true,
	}), nil
}

func (c *Container) getShell() string {
	for _, env := range c.AppleContainer.Configuration.InitProcess.Environment {
		if strings.HasPrefix(env, "PATH=") {
			path := strings.TrimPrefix(env, "PATH=")
			for _, dir := range strings.Split(path, ":") {
				for _, shell := range []string{"/bin/sh", "/bin/bash", "/usr/bin/bash"} {
					if dir == "/bin" || dir == "/usr/bin" {
						return shell
					}
				}
			}
		}
	}
	return "/bin/sh"
}

func (c *Container) Top(ctx context.Context) ([][]string, []string, error) {
	if c.AppleContainer.Status != "running" {
		return nil, nil, fmt.Errorf("container is not running")
	}

	stats, err := c.Client.StatsContainer(c.ID, true)
	if err != nil {
		return nil, nil, err
	}

	if len(stats) == 0 {
		return nil, nil, fmt.Errorf("no stats available")
	}

	s := stats[0]
	titles := []string{"METRIC", "VALUE"}
	processes := [][]string{
		{"CPU Usage", fmt.Sprintf("%d µs", s.CPUUsageUsec)},
		{"Memory Usage", utils.FormatBinaryBytes(int(s.MemoryUsageBytes))},
		{"Memory Limit", utils.FormatBinaryBytes(int(s.MemoryLimitBytes))},
		{"Network RX", utils.FormatBinaryBytes(int(s.NetworkRxBytes))},
		{"Network TX", utils.FormatBinaryBytes(int(s.NetworkTxBytes))},
		{"Block Read", utils.FormatBinaryBytes(int(s.BlockReadBytes))},
		{"Block Write", utils.FormatBinaryBytes(int(s.BlockWriteBytes))},
		{"Processes", fmt.Sprintf("%d", s.NumProcesses)},
	}

	return processes, titles, nil
}

func (c *Container) Inspect() (*AppleContainer, error) {
	return c.Client.InspectContainer(c.ID)
}

func (c *Container) RenderTop(ctx context.Context) (string, error) {
	processes, titles, err := c.Top(ctx)
	if err != nil {
		return "", err
	}

	return utils.RenderTable(append([][]string{titles}, processes...))
}

func (c *Container) DetailsLoaded() bool {
	return c.AppleContainer.Configuration.ID != ""
}

func (c *Container) GetStatus() string {
	return c.AppleContainer.Status
}

func (c *Container) GetImage() string {
	return c.AppleContainer.Configuration.Image.Reference
}

func (c *Container) GetIP() string {
	for _, net := range c.AppleContainer.Networks {
		if net.IPv4Address != "" {
			return strings.Split(net.IPv4Address, "/")[0]
		}
	}
	return ""
}

func (c *Container) GetPorts() []string {
	var ports []string
	for _, p := range c.AppleContainer.Configuration.PublishedPorts {
		ports = append(ports, fmt.Sprintf("%s:%d->%d/%s", p.HostIP, p.HostPort, p.ContainerPort, p.Protocol))
	}
	return ports
}

func (c *Container) GetEnv() []string {
	return c.AppleContainer.Configuration.InitProcess.Environment
}

func (c *Container) GetLabels() map[string]string {
	return c.AppleContainer.Configuration.Labels
}

func (c *Container) GetCPUs() int {
	return c.AppleContainer.Configuration.Resources.CPUS
}

func (c *Container) GetMemory() int64 {
	return c.AppleContainer.Configuration.Resources.MemoryInBytes
}
