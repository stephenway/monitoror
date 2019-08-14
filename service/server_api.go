//+build !faker

package service

import (
	"fmt"

	"github.com/monitoror/monitoror/monitorable/config"
	_configDelivery "github.com/monitoror/monitoror/monitorable/config/delivery/http"
	_configRepository "github.com/monitoror/monitoror/monitorable/config/repository"
	_configUsecase "github.com/monitoror/monitoror/monitorable/config/usecase"

	"github.com/monitoror/monitoror/monitorable/jenkins"
	_jenkinsDelivery "github.com/monitoror/monitoror/monitorable/jenkins/delivery/http"
	_jenkinsModels "github.com/monitoror/monitoror/monitorable/jenkins/models"
	_jenkinsRepository "github.com/monitoror/monitoror/monitorable/jenkins/repository"
	_jenkinsUsecase "github.com/monitoror/monitoror/monitorable/jenkins/usecase"
	"github.com/monitoror/monitoror/monitorable/ping"
	_pingDelivery "github.com/monitoror/monitoror/monitorable/ping/delivery/http"
	_pingModels "github.com/monitoror/monitoror/monitorable/ping/models"
	_pingRepository "github.com/monitoror/monitoror/monitorable/ping/repository"
	_pingUsecase "github.com/monitoror/monitoror/monitorable/ping/usecase"
	"github.com/monitoror/monitoror/monitorable/port"
	_portDelivery "github.com/monitoror/monitoror/monitorable/port/delivery/http"
	_portModels "github.com/monitoror/monitoror/monitorable/port/models"
	_portRepository "github.com/monitoror/monitoror/monitorable/port/repository"
	_portUsecase "github.com/monitoror/monitoror/monitorable/port/usecase"
	"github.com/monitoror/monitoror/monitorable/travisci"
	_travisciDelivery "github.com/monitoror/monitoror/monitorable/travisci/delivery/http"
	_travisciModels "github.com/monitoror/monitoror/monitorable/travisci/models"
	_travisciRepository "github.com/monitoror/monitoror/monitorable/travisci/repository"
	_travisciUsecase "github.com/monitoror/monitoror/monitorable/travisci/usecase"
)

func (s *Server) registerConfig() config.Helper {
	repository := _configRepository.NewConfigRepository()
	usecase := _configUsecase.NewConfigUsecase(repository)
	delivery := _configDelivery.NewHttpConfigDelivery(usecase)

	s.v1.GET("/config", s.cm.UpstreamCacheHandler(delivery.GetConfig))

	return usecase
}

func (s *Server) registerPing(configHelper config.Helper) {
	defer logStatus(ping.PingTileType, true)

	repository := _pingRepository.NewPingRepository(s.config)
	usecase := _pingUsecase.NewPingUsecase(repository)
	delivery := _pingDelivery.NewHttpPingDelivery(usecase)

	// Register route to echo
	route := s.v1.GET("/ping", s.cm.UpstreamCacheHandler(delivery.GetPing))

	// Register param and path to config usecase
	configHelper.RegisterTile(ping.PingTileType, route.Path, &_pingModels.PingParams{})
}

func (s *Server) registerPort(configHelper config.Helper) {
	defer logStatus(port.PortTileType, true)

	repository := _portRepository.NewPortRepository(s.config)
	usecase := _portUsecase.NewPortUsecase(repository)
	delivery := _portDelivery.NewHttpPortDelivery(usecase)

	// Register route to echo
	route := s.v1.GET("/port", s.cm.UpstreamCacheHandler(delivery.GetPort))

	// Register param and path to config usecase
	configHelper.RegisterTile(port.PortTileType, route.Path, &_portModels.PortParams{})
}

func (s *Server) registerTravisCI(configHelper config.Helper) {
	for label, travisCIConf := range s.config.Monitorable.TravisCI {
		// Associate github config
		githubConf := s.config.Monitorable.Github[label]

		defer logStatusWithConfigVariant(travisci.TravisCIBuildTileType, label, travisCIConf.IsValid())
		if !travisCIConf.IsValid() {
			continue
		}

		repository := _travisciRepository.NewTravisCIRepository(travisCIConf, githubConf)
		usecase := _travisciUsecase.NewTravisCIUsecase(repository)
		delivery := _travisciDelivery.NewHttpTravisCIDelivery(usecase)

		// Register route to echo
		travisCIGroup := s.v1.Group(fmt.Sprintf("/travisci/%s", label))
		route := travisCIGroup.GET("/build", s.cm.UpstreamCacheHandler(delivery.GetBuild))

		// Register param and path to config usecase
		configHelper.RegisterTileWithConfigVariant(travisci.TravisCIBuildTileType, label, route.Path, &_travisciModels.BuildParams{})
	}
}

func (s *Server) registerJenkins(configHelper config.Helper) {
	for label, jenkinsConf := range s.config.Monitorable.Jenkins {
		defer logStatusWithConfigVariant(jenkins.JenkinsBuildTileType, label, jenkinsConf.IsValid())
		if !jenkinsConf.IsValid() {
			continue
		}

		repository := _jenkinsRepository.NewJenkinsRepository(jenkinsConf)
		usecase := _jenkinsUsecase.NewJenkinsUsecase(repository)
		delivery := _jenkinsDelivery.NewHttpJenkinsDelivery(usecase)

		// Register route to echo
		jenkinsGroup := s.v1.Group(fmt.Sprintf("/jenkins/%s", label))
		route := jenkinsGroup.GET("/build", s.cm.UpstreamCacheHandler(delivery.GetBuild))

		// Register param and path to config usecase
		configHelper.RegisterTileWithConfigVariant(jenkins.JenkinsBuildTileType, label, route.Path, &_jenkinsModels.BuildParams{})
	}
}