/*
 Copyright 2022 Juicedata Inc

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package driver

import (
	"context"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"k8s.io/klog"

	"github.com/juicedata/juicefs-csi-driver/pkg/config"
)

// GetPluginInfo returns plugin info
func (d *Driver) GetPluginInfo(ctx context.Context, req *csi.GetPluginInfoRequest) (*csi.GetPluginInfoResponse, error) {
	klog.V(6).Infof("GetPluginInfo: called with args %+v", *req)
	resp := &csi.GetPluginInfoResponse{
		Name:          config.DriverName,
		VendorVersion: driverVersion,
	}

	return resp, nil
}

// GetPluginCapabilities returns plugin capabilities
func (d *Driver) GetPluginCapabilities(ctx context.Context, req *csi.GetPluginCapabilitiesRequest) (*csi.GetPluginCapabilitiesResponse, error) {
	klog.V(6).Infof("GetPluginCapabilities: called with args %+v", *req)
	resp := &csi.GetPluginCapabilitiesResponse{
		Capabilities: []*csi.PluginCapability{
			{
				Type: &csi.PluginCapability_Service_{
					Service: &csi.PluginCapability_Service{
						Type: csi.PluginCapability_Service_CONTROLLER_SERVICE,
					},
				},
			},
		},
	}

	return resp, nil
}

// Probe probes driver
func (d *Driver) Probe(ctx context.Context, req *csi.ProbeRequest) (*csi.ProbeResponse, error) {
	klog.V(6).Infof("Probe: called with args %+v", *req)
	return &csi.ProbeResponse{}, nil
}
