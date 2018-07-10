package patrol

func (self *App) apiRequest(
	ping bool,
	request *API_Request,
) {
	if ping && request.PID > 0 &&
		(self.config.KeepAlive == APP_KEEPALIVE_HTTP ||
			self.config.KeepAlive == APP_KEEPALIVE_UDP) {
		// only HTTP and UDP can have their PIDs managed
		self.pid = request.PID
	}
	// keyvalue
	if request.KeyValueReplace {
		// replace
		// dereference
		kv := make(map[string]interface{})
		for k, v := range request.KeyValue {
			kv[k] = v
		}
		self.keyvalue = kv
	} else {
		// merge
		for k, v := range request.KeyValue {
			self.keyvalue[k] = v
		}
	}
	// toggle
	if request.Toggle == API_TOGGLE_ENABLE {
		self.disabled = false
	} else if request.Toggle == API_TOGGLE_DISABLE {
		self.disabled = true
	} else if request.Toggle == API_TOGGLE_RUNONCE {
		// TODO
	}
}
func (self *App) apiResponse(
	status bool,
) *API_Response {
	result := &API_Response{
		PID:      self.pid,
		Disabled: self.disabled,
		Shutdown: self.patrol.shutdown,
		KeyValue: self.getKeyValue(),
	}
	if !status {
		// we don't need these values for individual Apps
		result.ID = self.id
		result.Group = "app"
		// we need to read lock patrol
		self.patrol.mu.RLock()
		result.Shutdown = self.patrol.shutdown
		self.patrol.mu.RUnlock()
	}
	if !self.started.IsZero() {
		result.Started = &Timestamp{
			Time: self.started,
			f:    self.patrol.config.Timestamp,
		}
	}
	if !self.lastseen.IsZero() {
		result.LastSeen = &Timestamp{
			Time: self.lastseen,
			f:    self.patrol.config.Timestamp,
		}
	}
	return result
}
