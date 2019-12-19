package main

import (
	"devops-integral/basic"
	"devops-integral/basic/common/constants"
	"devops-integral/basic/config"
	"devops-integral/upm-srv/handler/privilege"
	"devops-integral/upm-srv/handler/project"
	"devops-integral/upm-srv/handler/role"
	"devops-integral/upm-srv/handler/user"
	pe "devops-integral/upm-srv/proto/privilege"
	pp "devops-integral/upm-srv/proto/project"
	pr "devops-integral/upm-srv/proto/role"
	pu "devops-integral/upm-srv/proto/user"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/util/log"
	"time"
)

func main() {
	log.Logf("开始启动服务，服务名称：%s", constants.ServiceNameUpmSrv)
	//加载基础配置信息
	basic.Init()
	// 使用etcd进行服务注册
	micReg := etcd.NewRegistry(registryOptions)
	// 创建service
	service := micro.NewService(
		micro.Name(constants.ServiceNameUpmSrv),
		micro.Registry(micReg),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)
	//初始化服务信息
	service.Init()
	// 注册handlers
	// 注册user服务handler
	_ = pu.RegisterUserServiceHandler(service.Server(), new(user.Handler))
	// 注册project服务handler
	_ = pp.RegisterProjectServiceHandler(service.Server(), new(project.Handler))
	_ = pe.RegisterPrivilegeServiceHandler(service.Server(), new(privilege.Handler))
	_ = pr.RegisterRoleServiceHandler(service.Server(), new(role.Handler))
	// 启动服务
	if err := service.Run(); err != nil {
		log.Fatal("服务启动失败：", err)
	}
}

func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}
