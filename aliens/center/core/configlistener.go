/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2017/4/13
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package core

type ConfigListener struct {
	Handlers map[string]func(data []byte)
}

func (this *ConfigListener)ConfigChange(configType string, content []byte) {
	handler := this.Handlers[configType]
	if (handler != nil) {
		handler(content)
	}
}
