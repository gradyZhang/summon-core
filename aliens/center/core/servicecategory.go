/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2017/4/28
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package core

type ServiceCategory struct {
	Category 	string
	services 	map[string]IService
}

//更新服务
func (this *ServiceCategory) updateService(service IService) {
	this.services[service.GetID()] = service
}

//取出相同的服务
func (this *ServiceCategory) takeoutService(serviceConfig IService) IService {
	//服务地址信息没有变，不需要连接
	for key, service := range this.services {
		if (service.Equals(serviceConfig)) {
			delete(this.services, key)
			return service
		}
	}
	return nil
}
