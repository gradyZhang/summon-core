/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 *
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package core

//服务中心，处理服务的调度和查询
import (
	"github.com/samuel/go-zookeeper/zk"
	"time"
	"github.com/name5566/leaf/log"
	"encoding/json"
	"sync"
)

const SERVICE_NODE_NAME string = "service"


type ServiceCenter struct {
	sync.RWMutex
	zkCon       		*zk.Conn
	zkName      		string
	serviceRoot 		string
	serviceContainer    	map[string]*ServiceCategory  //服务容器 key 服务名
	serviceFactory		IServiceFactory
}

//启动服务中心客户端
func (this *ServiceCenter) Connect(address string, timeout int, zkName string, serviceFactory IServiceFactory) {
	this.zkName = zkName
	this.serviceFactory = serviceFactory
	c, _, err := zk.Connect([]string{address}, time.Second)
	if err != nil {
		panic(err)
	}
	this.serviceContainer = make(map[string]*ServiceCategory)
	this.serviceRoot = NODE_SPLIT + this.zkName + NODE_SPLIT + SERVICE_NODE_NAME
	this.zkCon = c
	this.confirmNode(NODE_SPLIT + this.zkName)
	this.confirmNode(this.serviceRoot)
}

func (this *ServiceCenter) assert() {
	if (this.zkCon == nil) {
		panic("mast start service center first")
	}
}

//关闭服务中心
func (this *ServiceCenter) Close() {
	if (this.zkCon != nil) {
		this.zkCon.Close()
	}
}

//更新服务
func (this *ServiceCenter) UpdateService(service IService) {
	this.Lock()
	defer this.Unlock()
	if (this.serviceContainer[service.GetType()] == nil) {
		this.serviceContainer[service.GetType()] =
			&ServiceCategory{Category:service.GetType(), services:make(map[string]IService)}
	}
	this.serviceContainer[service.GetType()].updateService(service)
}

//根据服务类型获取一个空闲的服务节点
func (this *ServiceCenter) AllocService(serviceType string) map[string]IService {
	this.RLock()
	defer this.RUnlock()
	return this.serviceContainer[serviceType].services
}

//订阅服务  能实时更新服务信息
func (this *ServiceCenter) SubscribeServices(serviceTypes ...string) {
	this.assert()
	for _, serviceType := range serviceTypes {
		this.SubscribeService(serviceType)
	}
}

func (this *ServiceCenter) SubscribeService(serviceType string) {
	path := this.serviceRoot + NODE_SPLIT + serviceType
	this.confirmNode(path)
	serviceIDs, _, ch, err := this.zkCon.ChildrenW(path)
	if (err != nil) {
		log.Release("subscribe service %v error: %v",path, err)
		return
	}
	this.Lock()
	defer this.Unlock()
	oldContainer := this.serviceContainer[serviceType]
	serviceCategory := &ServiceCategory{Category:serviceType, services:make(map[string]IService)}
	for _, serviceID := range serviceIDs {
		data, _, err := this.zkCon.Get(path + NODE_SPLIT + serviceID)
		service := this.serviceFactory.CreateService(data)
		//&Service{ID:serviceID, Type:serviceType}
		//err = json.Unmarshal(data, service)
		if (service == nil) {
			log.Release("%v unmarshal json error : %v",path, err)
			continue
		}
		service.SetID(serviceID)
		service.SetType(serviceType)
		if (oldContainer != nil) {
			oldService := oldContainer.takeoutService(service)
			if (oldService != nil) {
				oldService.SetID(service.GetID())
				serviceCategory.updateService(oldService)
				continue
			}
		}
		//新服务需要连接上才能更新
		if (service.Connect()) {
			serviceCategory.updateService(service)
		}
	}
	this.serviceContainer[serviceType] = serviceCategory
	go this.openListener(serviceType, path , ch)
}

func (this *ServiceCenter) openListener(serviceType string, path string, ch <- chan zk.Event) {
	event, _ := <- ch
	//更新服务节点信息
	if (event.Type == zk.EventNodeChildrenChanged) {
		this.SubscribeService(serviceType)
	}
}

//
func (this *ServiceCenter) confirmNode(path string, flags ...int32) bool {
	_, err := this.zkCon.Create(path, nil, 0, zk.WorldACL(zk.PermAll))
	return err == nil
}

//发布服务
func (this *ServiceCenter) PublicService(service IService) bool {
	this.assert()
	if (!service.IsLocal()) {
		log.Release("service info is invalid")
		return false
	}
	//path string, data []byte, version int32
	data, err := json.Marshal(service)
	if (err != nil) {
		log.Release("marshal json service data error : %v", err)
		return false
	}
	servicePath := this.serviceRoot + NODE_SPLIT + service.GetType()
	this.confirmNode(servicePath)
	id, err := this.zkCon.Create(servicePath + NODE_SPLIT + service.GetType(), data,
		zk.FlagEphemeral|zk.FlagSequence, zk.WorldACL(zk.PermAll))
	if (err != nil) {
		log.Release("create service error : %v", err)
		return false
	}
	log.Release("public service success : %v", id)
	service.SetID(id)
	//服务注册在容器
	this.UpdateService(service)
	return true
}

