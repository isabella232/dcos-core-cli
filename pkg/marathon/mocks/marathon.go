// Code generated by mockgen DO NOT EDIT.

package mocks

import (
	"net/url"
	"time"

	"github.com/gambol99/go-marathon"
)

type MarathonMock struct {
	AbdicateLeaderInvocations            int
	AbdicateLeaderFn                     func() (string, error)
	AddEventsListenerInvocations         int
	AddEventsListenerFn                  func(int) (marathon.EventsChannel, error)
	AllTasksInvocations                  int
	AllTasksFn                           func(*marathon.AllTasksOpts) (*marathon.Tasks, error)
	ApiPostInvocations                   int
	ApiPostFn                            func(string, interface{}, interface{}) error
	ApplicationInvocations               int
	ApplicationFn                        func(string) (*marathon.Application, error)
	ApplicationByInvocations             int
	ApplicationByFn                      func(string, *marathon.GetAppOpts) (*marathon.Application, error)
	ApplicationByVersionInvocations      int
	ApplicationByVersionFn               func(string, string) (*marathon.Application, error)
	ApplicationDeploymentsInvocations    int
	ApplicationDeploymentsFn             func(string) ([]*marathon.DeploymentID, error)
	ApplicationOKInvocations             int
	ApplicationOKFn                      func(string) (bool, error)
	ApplicationVersionsInvocations       int
	ApplicationVersionsFn                func(string) (*marathon.ApplicationVersions, error)
	ApplicationsInvocations              int
	ApplicationsFn                       func(url.Values) (*marathon.Applications, error)
	CreateApplicationInvocations         int
	CreateApplicationFn                  func(*marathon.Application) (*marathon.Application, error)
	CreateGroupInvocations               int
	CreateGroupFn                        func(*marathon.Group) error
	CreatePodInvocations                 int
	CreatePodFn                          func(*marathon.Pod) (*marathon.Pod, error)
	DeleteApplicationInvocations         int
	DeleteApplicationFn                  func(string, bool) (*marathon.DeploymentID, error)
	DeleteDeploymentInvocations          int
	DeleteDeploymentFn                   func(string, bool) (*marathon.DeploymentID, error)
	DeleteGroupInvocations               int
	DeleteGroupFn                        func(string, bool) (*marathon.DeploymentID, error)
	DeletePodInvocations                 int
	DeletePodFn                          func(string, bool) (*marathon.DeploymentID, error)
	DeletePodInstanceInvocations         int
	DeletePodInstanceFn                  func(string, string) (*marathon.PodInstance, error)
	DeletePodInstancesInvocations        int
	DeletePodInstancesFn                 func(string, []string) ([]*marathon.PodInstance, error)
	DeleteQueueDelayInvocations          int
	DeleteQueueDelayFn                   func(string) error
	DeploymentsInvocations               int
	DeploymentsFn                        func() ([]*marathon.Deployment, error)
	GetMarathonURLInvocations            int
	GetMarathonURLFn                     func() string
	GroupInvocations                     int
	GroupFn                              func(string) (*marathon.Group, error)
	GroupByInvocations                   int
	GroupByFn                            func(string, *marathon.GetGroupOpts) (*marathon.Group, error)
	GroupsInvocations                    int
	GroupsFn                             func() (*marathon.Groups, error)
	GroupsByInvocations                  int
	GroupsByFn                           func(*marathon.GetGroupOpts) (*marathon.Groups, error)
	HasApplicationVersionInvocations     int
	HasApplicationVersionFn              func(string, string) (bool, error)
	HasDeploymentInvocations             int
	HasDeploymentFn                      func(string) (bool, error)
	HasGroupInvocations                  int
	HasGroupFn                           func(string) (bool, error)
	InfoInvocations                      int
	InfoFn                               func() (*marathon.Info, error)
	KillApplicationTasksInvocations      int
	KillApplicationTasksFn               func(string, *marathon.KillApplicationTasksOpts) (*marathon.Tasks, error)
	KillTaskInvocations                  int
	KillTaskFn                           func(string, *marathon.KillTaskOpts) (*marathon.Task, error)
	KillTasksInvocations                 int
	KillTasksFn                          func([]string, *marathon.KillTaskOpts) error
	LeaderInvocations                    int
	LeaderFn                             func() (string, error)
	ListApplicationsInvocations          int
	ListApplicationsFn                   func(url.Values) ([]string, error)
	PingInvocations                      int
	PingFn                               func() (bool, error)
	PodInvocations                       int
	PodFn                                func(string) (*marathon.Pod, error)
	PodByVersionInvocations              int
	PodByVersionFn                       func(string, string) (*marathon.Pod, error)
	PodIsRunningInvocations              int
	PodIsRunningFn                       func(string) bool
	PodStatusInvocations                 int
	PodStatusFn                          func(string) (*marathon.PodStatus, error)
	PodStatusesInvocations               int
	PodStatusesFn                        func() ([]*marathon.PodStatus, error)
	PodVersionsInvocations               int
	PodVersionsFn                        func(string) ([]string, error)
	PodsInvocations                      int
	PodsFn                               func() ([]marathon.Pod, error)
	QueueInvocations                     int
	QueueFn                              func() (*marathon.Queue, error)
	RemoveEventsListenerInvocations      int
	RemoveEventsListenerFn               func(marathon.EventsChannel)
	RestartApplicationInvocations        int
	RestartApplicationFn                 func(string, bool) (*marathon.DeploymentID, error)
	ScaleApplicationInstancesInvocations int
	ScaleApplicationInstancesFn          func(string, int, bool) (*marathon.DeploymentID, error)
	SetApplicationVersionInvocations     int
	SetApplicationVersionFn              func(string, *marathon.ApplicationVersion) (*marathon.DeploymentID, error)
	SubscribeInvocations                 int
	SubscribeFn                          func(string) error
	SubscriptionsInvocations             int
	SubscriptionsFn                      func() (*marathon.Subscriptions, error)
	SupportsPodsInvocations              int
	SupportsPodsFn                       func() (bool, error)
	TaskEndpointsInvocations             int
	TaskEndpointsFn                      func(string, int, bool) ([]string, error)
	TasksInvocations                     int
	TasksFn                              func(string) (*marathon.Tasks, error)
	UnsubscribeInvocations               int
	UnsubscribeFn                        func(string) error
	UpdateApplicationInvocations         int
	UpdateApplicationFn                  func(*marathon.Application, bool) (*marathon.DeploymentID, error)
	UpdateGroupInvocations               int
	UpdateGroupFn                        func(string, *marathon.Group, bool) (*marathon.DeploymentID, error)
	UpdatePodInvocations                 int
	UpdatePodFn                          func(*marathon.Pod, bool) (*marathon.Pod, error)
	WaitOnApplicationInvocations         int
	WaitOnApplicationFn                  func(string, time.Duration) error
	WaitOnDeploymentInvocations          int
	WaitOnDeploymentFn                   func(string, time.Duration) error
	WaitOnGroupInvocations               int
	WaitOnGroupFn                        func(string, time.Duration) error
	WaitOnPodInvocations                 int
	WaitOnPodFn                          func(string, time.Duration) error
}

func (m *MarathonMock) AbdicateLeader() (string, error) {
	m.AbdicateLeaderInvocations++
	return m.AbdicateLeaderFn()
}

func (m *MarathonMock) AddEventsListener(p0 int) (marathon.EventsChannel, error) {
	m.AddEventsListenerInvocations++
	return m.AddEventsListenerFn(p0)
}

func (m *MarathonMock) AllTasks(p0 *marathon.AllTasksOpts) (*marathon.Tasks, error) {
	m.AllTasksInvocations++
	return m.AllTasksFn(p0)
}

func (m *MarathonMock) ApiPost(p0 string, p1 interface{}, p2 interface{}) error {
	m.ApiPostInvocations++
	return m.ApiPostFn(p0, p1, p2)
}

func (m *MarathonMock) Application(p0 string) (*marathon.Application, error) {
	m.ApplicationInvocations++
	return m.ApplicationFn(p0)
}

func (m *MarathonMock) ApplicationBy(p0 string, p1 *marathon.GetAppOpts) (*marathon.Application, error) {
	m.ApplicationByInvocations++
	return m.ApplicationByFn(p0, p1)
}

func (m *MarathonMock) ApplicationByVersion(p0 string, p1 string) (*marathon.Application, error) {
	m.ApplicationByVersionInvocations++
	return m.ApplicationByVersionFn(p0, p1)
}

func (m *MarathonMock) ApplicationDeployments(p0 string) ([]*marathon.DeploymentID, error) {
	m.ApplicationDeploymentsInvocations++
	return m.ApplicationDeploymentsFn(p0)
}

func (m *MarathonMock) ApplicationOK(p0 string) (bool, error) {
	m.ApplicationOKInvocations++
	return m.ApplicationOKFn(p0)
}

func (m *MarathonMock) ApplicationVersions(p0 string) (*marathon.ApplicationVersions, error) {
	m.ApplicationVersionsInvocations++
	return m.ApplicationVersionsFn(p0)
}

func (m *MarathonMock) Applications(p0 url.Values) (*marathon.Applications, error) {
	m.ApplicationsInvocations++
	return m.ApplicationsFn(p0)
}

func (m *MarathonMock) CreateApplication(p0 *marathon.Application) (*marathon.Application, error) {
	m.CreateApplicationInvocations++
	return m.CreateApplicationFn(p0)
}

func (m *MarathonMock) CreateGroup(p0 *marathon.Group) error {
	m.CreateGroupInvocations++
	return m.CreateGroupFn(p0)
}

func (m *MarathonMock) CreatePod(p0 *marathon.Pod) (*marathon.Pod, error) {
	m.CreatePodInvocations++
	return m.CreatePodFn(p0)
}

func (m *MarathonMock) DeleteApplication(p0 string, p1 bool) (*marathon.DeploymentID, error) {
	m.DeleteApplicationInvocations++
	return m.DeleteApplicationFn(p0, p1)
}

func (m *MarathonMock) DeleteDeployment(p0 string, p1 bool) (*marathon.DeploymentID, error) {
	m.DeleteDeploymentInvocations++
	return m.DeleteDeploymentFn(p0, p1)
}

func (m *MarathonMock) DeleteGroup(p0 string, p1 bool) (*marathon.DeploymentID, error) {
	m.DeleteGroupInvocations++
	return m.DeleteGroupFn(p0, p1)
}

func (m *MarathonMock) DeletePod(p0 string, p1 bool) (*marathon.DeploymentID, error) {
	m.DeletePodInvocations++
	return m.DeletePodFn(p0, p1)
}

func (m *MarathonMock) DeletePodInstance(p0 string, p1 string) (*marathon.PodInstance, error) {
	m.DeletePodInstanceInvocations++
	return m.DeletePodInstanceFn(p0, p1)
}

func (m *MarathonMock) DeletePodInstances(p0 string, p1 []string) ([]*marathon.PodInstance, error) {
	m.DeletePodInstancesInvocations++
	return m.DeletePodInstancesFn(p0, p1)
}

func (m *MarathonMock) DeleteQueueDelay(p0 string) error {
	m.DeleteQueueDelayInvocations++
	return m.DeleteQueueDelayFn(p0)
}

func (m *MarathonMock) Deployments() ([]*marathon.Deployment, error) {
	m.DeploymentsInvocations++
	return m.DeploymentsFn()
}

func (m *MarathonMock) GetMarathonURL() string {
	m.GetMarathonURLInvocations++
	return m.GetMarathonURLFn()
}

func (m *MarathonMock) Group(p0 string) (*marathon.Group, error) {
	m.GroupInvocations++
	return m.GroupFn(p0)
}

func (m *MarathonMock) GroupBy(p0 string, p1 *marathon.GetGroupOpts) (*marathon.Group, error) {
	m.GroupByInvocations++
	return m.GroupByFn(p0, p1)
}

func (m *MarathonMock) Groups() (*marathon.Groups, error) {
	m.GroupsInvocations++
	return m.GroupsFn()
}

func (m *MarathonMock) GroupsBy(p0 *marathon.GetGroupOpts) (*marathon.Groups, error) {
	m.GroupsByInvocations++
	return m.GroupsByFn(p0)
}

func (m *MarathonMock) HasApplicationVersion(p0 string, p1 string) (bool, error) {
	m.HasApplicationVersionInvocations++
	return m.HasApplicationVersionFn(p0, p1)
}

func (m *MarathonMock) HasDeployment(p0 string) (bool, error) {
	m.HasDeploymentInvocations++
	return m.HasDeploymentFn(p0)
}

func (m *MarathonMock) HasGroup(p0 string) (bool, error) {
	m.HasGroupInvocations++
	return m.HasGroupFn(p0)
}

func (m *MarathonMock) Info() (*marathon.Info, error) {
	m.InfoInvocations++
	return m.InfoFn()
}

func (m *MarathonMock) KillApplicationTasks(p0 string, p1 *marathon.KillApplicationTasksOpts) (*marathon.Tasks, error) {
	m.KillApplicationTasksInvocations++
	return m.KillApplicationTasksFn(p0, p1)
}

func (m *MarathonMock) KillTask(p0 string, p1 *marathon.KillTaskOpts) (*marathon.Task, error) {
	m.KillTaskInvocations++
	return m.KillTaskFn(p0, p1)
}

func (m *MarathonMock) KillTasks(p0 []string, p1 *marathon.KillTaskOpts) error {
	m.KillTasksInvocations++
	return m.KillTasksFn(p0, p1)
}

func (m *MarathonMock) Leader() (string, error) {
	m.LeaderInvocations++
	return m.LeaderFn()
}

func (m *MarathonMock) ListApplications(p0 url.Values) ([]string, error) {
	m.ListApplicationsInvocations++
	return m.ListApplicationsFn(p0)
}

func (m *MarathonMock) Ping() (bool, error) {
	m.PingInvocations++
	return m.PingFn()
}

func (m *MarathonMock) Pod(p0 string) (*marathon.Pod, error) {
	m.PodInvocations++
	return m.PodFn(p0)
}

func (m *MarathonMock) PodByVersion(p0 string, p1 string) (*marathon.Pod, error) {
	m.PodByVersionInvocations++
	return m.PodByVersionFn(p0, p1)
}

func (m *MarathonMock) PodIsRunning(p0 string) bool {
	m.PodIsRunningInvocations++
	return m.PodIsRunningFn(p0)
}

func (m *MarathonMock) PodStatus(p0 string) (*marathon.PodStatus, error) {
	m.PodStatusInvocations++
	return m.PodStatusFn(p0)
}

func (m *MarathonMock) PodStatuses() ([]*marathon.PodStatus, error) {
	m.PodStatusesInvocations++
	return m.PodStatusesFn()
}

func (m *MarathonMock) PodVersions(p0 string) ([]string, error) {
	m.PodVersionsInvocations++
	return m.PodVersionsFn(p0)
}

func (m *MarathonMock) Pods() ([]marathon.Pod, error) {
	m.PodsInvocations++
	return m.PodsFn()
}

func (m *MarathonMock) Queue() (*marathon.Queue, error) {
	m.QueueInvocations++
	return m.QueueFn()
}

func (m *MarathonMock) RemoveEventsListener(p0 marathon.EventsChannel) {
	m.RemoveEventsListenerInvocations++
	m.RemoveEventsListenerFn(p0)
}

func (m *MarathonMock) RestartApplication(p0 string, p1 bool) (*marathon.DeploymentID, error) {
	m.RestartApplicationInvocations++
	return m.RestartApplicationFn(p0, p1)
}

func (m *MarathonMock) ScaleApplicationInstances(p0 string, p1 int, p2 bool) (*marathon.DeploymentID, error) {
	m.ScaleApplicationInstancesInvocations++
	return m.ScaleApplicationInstancesFn(p0, p1, p2)
}

func (m *MarathonMock) SetApplicationVersion(p0 string, p1 *marathon.ApplicationVersion) (*marathon.DeploymentID, error) {
	m.SetApplicationVersionInvocations++
	return m.SetApplicationVersionFn(p0, p1)
}

func (m *MarathonMock) Subscribe(p0 string) error {
	m.SubscribeInvocations++
	return m.SubscribeFn(p0)
}

func (m *MarathonMock) Subscriptions() (*marathon.Subscriptions, error) {
	m.SubscriptionsInvocations++
	return m.SubscriptionsFn()
}

func (m *MarathonMock) SupportsPods() (bool, error) {
	m.SupportsPodsInvocations++
	return m.SupportsPodsFn()
}

func (m *MarathonMock) TaskEndpoints(p0 string, p1 int, p2 bool) ([]string, error) {
	m.TaskEndpointsInvocations++
	return m.TaskEndpointsFn(p0, p1, p2)
}

func (m *MarathonMock) Tasks(p0 string) (*marathon.Tasks, error) {
	m.TasksInvocations++
	return m.TasksFn(p0)
}

func (m *MarathonMock) Unsubscribe(p0 string) error {
	m.UnsubscribeInvocations++
	return m.UnsubscribeFn(p0)
}

func (m *MarathonMock) UpdateApplication(p0 *marathon.Application, p1 bool) (*marathon.DeploymentID, error) {
	m.UpdateApplicationInvocations++
	return m.UpdateApplicationFn(p0, p1)
}

func (m *MarathonMock) UpdateGroup(p0 string, p1 *marathon.Group, p2 bool) (*marathon.DeploymentID, error) {
	m.UpdateGroupInvocations++
	return m.UpdateGroupFn(p0, p1, p2)
}

func (m *MarathonMock) UpdatePod(p0 *marathon.Pod, p1 bool) (*marathon.Pod, error) {
	m.UpdatePodInvocations++
	return m.UpdatePodFn(p0, p1)
}

func (m *MarathonMock) WaitOnApplication(p0 string, p1 time.Duration) error {
	m.WaitOnApplicationInvocations++
	return m.WaitOnApplicationFn(p0, p1)
}

func (m *MarathonMock) WaitOnDeployment(p0 string, p1 time.Duration) error {
	m.WaitOnDeploymentInvocations++
	return m.WaitOnDeploymentFn(p0, p1)
}

func (m *MarathonMock) WaitOnGroup(p0 string, p1 time.Duration) error {
	m.WaitOnGroupInvocations++
	return m.WaitOnGroupFn(p0, p1)
}

func (m *MarathonMock) WaitOnPod(p0 string, p1 time.Duration) error {
	m.WaitOnPodInvocations++
	return m.WaitOnPodFn(p0, p1)
}
