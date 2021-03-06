/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kubeform. DO NOT EDIT.

package main

import (
	jsoniter "github.com/json-iterator/go"
	"k8s.io/apimachinery/pkg/runtime/schema"
	cloudv1alpha1 "kubeform.dev/provider-ovh-api/apis/cloud/v1alpha1"
	dbaasv1alpha1 "kubeform.dev/provider-ovh-api/apis/dbaas/v1alpha1"
	dedicatedv1alpha1 "kubeform.dev/provider-ovh-api/apis/dedicated/v1alpha1"
	domainv1alpha1 "kubeform.dev/provider-ovh-api/apis/domain/v1alpha1"
	ipv1alpha1 "kubeform.dev/provider-ovh-api/apis/ip/v1alpha1"
	iploadbalancingv1alpha1 "kubeform.dev/provider-ovh-api/apis/iploadbalancing/v1alpha1"
	mev1alpha1 "kubeform.dev/provider-ovh-api/apis/me/v1alpha1"
	vrackv1alpha1 "kubeform.dev/provider-ovh-api/apis/vrack/v1alpha1"
	"kubeform.dev/provider-ovh-controller/controllers"
)

type Data struct {
	JsonIt       jsoniter.API
	ResourceType string
}

var allJsonIt = map[schema.GroupVersionResource]Data{
	{
		Group:    "cloud.ovh.kubeform.com",
		Version:  "v1alpha1",
		Resource: "projects",
	}: {
		JsonIt:       controllers.GetJSONItr(cloudv1alpha1.GetEncoder(), cloudv1alpha1.GetDecoder()),
		ResourceType: "ovh_cloud_project",
	},
	{
		Group:    "cloud.ovh.kubeform.com",
		Version:  "v1alpha1",
		Resource: "projectcontainerregistries",
	}: {
		JsonIt:       controllers.GetJSONItr(cloudv1alpha1.GetEncoder(), cloudv1alpha1.GetDecoder()),
		ResourceType: "ovh_cloud_project_containerregistry",
	},
	{
		Group:    "cloud.ovh.kubeform.com",
		Version:  "v1alpha1",
		Resource: "projectcontainerregistryusers",
	}: {
		JsonIt:       controllers.GetJSONItr(cloudv1alpha1.GetEncoder(), cloudv1alpha1.GetDecoder()),
		ResourceType: "ovh_cloud_project_containerregistry_user",
	},
	{
		Group:    "cloud.ovh.kubeform.com",
		Version:  "v1alpha1",
		Resource: "projectkubes",
	}: {
		JsonIt:       controllers.GetJSONItr(cloudv1alpha1.GetEncoder(), cloudv1alpha1.GetDecoder()),
		ResourceType: "ovh_cloud_project_kube",
	},
	{
		Group:    "cloud.ovh.kubeform.com",
		Version:  "v1alpha1",
		Resource: "projectkubenodepools",
	}: {
		JsonIt:       controllers.GetJSONItr(cloudv1alpha1.GetEncoder(), cloudv1alpha1.GetDecoder()),
		ResourceType: "ovh_cloud_project_kube_nodepool",
	},
	{
		Group:    "cloud.ovh.kubeform.com",
		Version:  "v1alpha1",
		Resource: "projectnetworkprivates",
	}: {
		JsonIt:       controllers.GetJSONItr(cloudv1alpha1.GetEncoder(), cloudv1alpha1.GetDecoder()),
		ResourceType: "ovh_cloud_project_network_private",
	},
	{
		Group:    "cloud.ovh.kubeform.com",
		Version:  "v1alpha1",
		Resource: "projectnetworkprivatesubnets",
	}: {
		JsonIt:       controllers.GetJSONItr(cloudv1alpha1.GetEncoder(), cloudv1alpha1.GetDecoder()),
		ResourceType: "ovh_cloud_project_network_private_subnet",
	},
	{
		Group:    "cloud.ovh.kubeform.com",
		Version:  "v1alpha1",
		Resource: "projectusers",
	}: {
		JsonIt:       controllers.GetJSONItr(cloudv1alpha1.GetEncoder(), cloudv1alpha1.GetDecoder()),
		ResourceType: "ovh_cloud_project_user",
	},
	{
		Group:    "dbaas.ovh.kubeform.com",
		Version:  "v1alpha1",
		Resource: "logsinputs",
	}: {
		JsonIt:       controllers.GetJSONItr(dbaasv1alpha1.GetEncoder(), dbaasv1alpha1.GetDecoder()),
		ResourceType: "ovh_dbaas_logs_input",
	},
	{
		Group:    "dbaas.ovh.kubeform.com",
		Version:  "v1alpha1",
		Resource: "logsoutputgraylogstreams",
	}: {
		JsonIt:       controllers.GetJSONItr(dbaasv1alpha1.GetEncoder(), dbaasv1alpha1.GetDecoder()),
		ResourceType: "ovh_dbaas_logs_output_graylog_stream",
	},
	{
		Group:    "dedicated.ovh.kubeform.com",
		Version:  "v1alpha1",
		Resource: "cephacls",
	}: {
		JsonIt:       controllers.GetJSONItr(dedicatedv1alpha1.GetEncoder(), dedicatedv1alpha1.GetDecoder()),
		ResourceType: "ovh_dedicated_ceph_acl",
	},
	{
		Group:    "dedicated.ovh.kubeform.com",
		Version:  "v1alpha1",
		Resource: "serverinstalltasks",
	}: {
		JsonIt:       controllers.GetJSONItr(dedicatedv1alpha1.GetEncoder(), dedicatedv1alpha1.GetDecoder()),
		ResourceType: "ovh_dedicated_server_install_task",
	},
	{
		Group:    "dedicated.ovh.kubeform.com",
		Version:  "v1alpha1",
		Resource: "serverreboottasks",
	}: {
		JsonIt:       controllers.GetJSONItr(dedicatedv1alpha1.GetEncoder(), dedicatedv1alpha1.GetDecoder()),
		ResourceType: "ovh_dedicated_server_reboot_task",
	},
	{
		Group:    "dedicated.ovh.kubeform.com",
		Version:  "v1alpha1",
		Resource: "serverupdates",
	}: {
		JsonIt:       controllers.GetJSONItr(dedicatedv1alpha1.GetEncoder(), dedicatedv1alpha1.GetDecoder()),
		ResourceType: "ovh_dedicated_server_update",
	},
	{
		Group:    "domain.ovh.kubeform.com",
		Version:  "v1alpha1",
		Resource: "zones",
	}: {
		JsonIt:       controllers.GetJSONItr(domainv1alpha1.GetEncoder(), domainv1alpha1.GetDecoder()),
		ResourceType: "ovh_domain_zone",
	},
	{
		Group:    "domain.ovh.kubeform.com",
		Version:  "v1alpha1",
		Resource: "zonerecords",
	}: {
		JsonIt:       controllers.GetJSONItr(domainv1alpha1.GetEncoder(), domainv1alpha1.GetDecoder()),
		ResourceType: "ovh_domain_zone_record",
	},
	{
		Group:    "domain.ovh.kubeform.com",
		Version:  "v1alpha1",
		Resource: "zoneredirections",
	}: {
		JsonIt:       controllers.GetJSONItr(domainv1alpha1.GetEncoder(), domainv1alpha1.GetDecoder()),
		ResourceType: "ovh_domain_zone_redirection",
	},
	{
		Group:    "ip.ovh.kubeform.com",
		Version:  "v1alpha1",
		Resource: "reverses",
	}: {
		JsonIt:       controllers.GetJSONItr(ipv1alpha1.GetEncoder(), ipv1alpha1.GetDecoder()),
		ResourceType: "ovh_ip_reverse",
	},
	{
		Group:    "ip.ovh.kubeform.com",
		Version:  "v1alpha1",
		Resource: "services",
	}: {
		JsonIt:       controllers.GetJSONItr(ipv1alpha1.GetEncoder(), ipv1alpha1.GetDecoder()),
		ResourceType: "ovh_ip_service",
	},
	{
		Group:    "iploadbalancing.ovh.kubeform.com",
		Version:  "v1alpha1",
		Resource: "iploadbalancings",
	}: {
		JsonIt:       controllers.GetJSONItr(iploadbalancingv1alpha1.GetEncoder(), iploadbalancingv1alpha1.GetDecoder()),
		ResourceType: "ovh_iploadbalancing",
	},
	{
		Group:    "iploadbalancing.ovh.kubeform.com",
		Version:  "v1alpha1",
		Resource: "httpfarms",
	}: {
		JsonIt:       controllers.GetJSONItr(iploadbalancingv1alpha1.GetEncoder(), iploadbalancingv1alpha1.GetDecoder()),
		ResourceType: "ovh_iploadbalancing_http_farm",
	},
	{
		Group:    "iploadbalancing.ovh.kubeform.com",
		Version:  "v1alpha1",
		Resource: "httpfarmservers",
	}: {
		JsonIt:       controllers.GetJSONItr(iploadbalancingv1alpha1.GetEncoder(), iploadbalancingv1alpha1.GetDecoder()),
		ResourceType: "ovh_iploadbalancing_http_farm_server",
	},
	{
		Group:    "iploadbalancing.ovh.kubeform.com",
		Version:  "v1alpha1",
		Resource: "httpfrontends",
	}: {
		JsonIt:       controllers.GetJSONItr(iploadbalancingv1alpha1.GetEncoder(), iploadbalancingv1alpha1.GetDecoder()),
		ResourceType: "ovh_iploadbalancing_http_frontend",
	},
	{
		Group:    "iploadbalancing.ovh.kubeform.com",
		Version:  "v1alpha1",
		Resource: "httproutes",
	}: {
		JsonIt:       controllers.GetJSONItr(iploadbalancingv1alpha1.GetEncoder(), iploadbalancingv1alpha1.GetDecoder()),
		ResourceType: "ovh_iploadbalancing_http_route",
	},
	{
		Group:    "iploadbalancing.ovh.kubeform.com",
		Version:  "v1alpha1",
		Resource: "httprouterules",
	}: {
		JsonIt:       controllers.GetJSONItr(iploadbalancingv1alpha1.GetEncoder(), iploadbalancingv1alpha1.GetDecoder()),
		ResourceType: "ovh_iploadbalancing_http_route_rule",
	},
	{
		Group:    "iploadbalancing.ovh.kubeform.com",
		Version:  "v1alpha1",
		Resource: "refreshes",
	}: {
		JsonIt:       controllers.GetJSONItr(iploadbalancingv1alpha1.GetEncoder(), iploadbalancingv1alpha1.GetDecoder()),
		ResourceType: "ovh_iploadbalancing_refresh",
	},
	{
		Group:    "iploadbalancing.ovh.kubeform.com",
		Version:  "v1alpha1",
		Resource: "tcpfarms",
	}: {
		JsonIt:       controllers.GetJSONItr(iploadbalancingv1alpha1.GetEncoder(), iploadbalancingv1alpha1.GetDecoder()),
		ResourceType: "ovh_iploadbalancing_tcp_farm",
	},
	{
		Group:    "iploadbalancing.ovh.kubeform.com",
		Version:  "v1alpha1",
		Resource: "tcpfarmservers",
	}: {
		JsonIt:       controllers.GetJSONItr(iploadbalancingv1alpha1.GetEncoder(), iploadbalancingv1alpha1.GetDecoder()),
		ResourceType: "ovh_iploadbalancing_tcp_farm_server",
	},
	{
		Group:    "iploadbalancing.ovh.kubeform.com",
		Version:  "v1alpha1",
		Resource: "tcpfrontends",
	}: {
		JsonIt:       controllers.GetJSONItr(iploadbalancingv1alpha1.GetEncoder(), iploadbalancingv1alpha1.GetDecoder()),
		ResourceType: "ovh_iploadbalancing_tcp_frontend",
	},
	{
		Group:    "iploadbalancing.ovh.kubeform.com",
		Version:  "v1alpha1",
		Resource: "tcproutes",
	}: {
		JsonIt:       controllers.GetJSONItr(iploadbalancingv1alpha1.GetEncoder(), iploadbalancingv1alpha1.GetDecoder()),
		ResourceType: "ovh_iploadbalancing_tcp_route",
	},
	{
		Group:    "iploadbalancing.ovh.kubeform.com",
		Version:  "v1alpha1",
		Resource: "tcprouterules",
	}: {
		JsonIt:       controllers.GetJSONItr(iploadbalancingv1alpha1.GetEncoder(), iploadbalancingv1alpha1.GetDecoder()),
		ResourceType: "ovh_iploadbalancing_tcp_route_rule",
	},
	{
		Group:    "iploadbalancing.ovh.kubeform.com",
		Version:  "v1alpha1",
		Resource: "vracknetworks",
	}: {
		JsonIt:       controllers.GetJSONItr(iploadbalancingv1alpha1.GetEncoder(), iploadbalancingv1alpha1.GetDecoder()),
		ResourceType: "ovh_iploadbalancing_vrack_network",
	},
	{
		Group:    "me.ovh.kubeform.com",
		Version:  "v1alpha1",
		Resource: "identityusers",
	}: {
		JsonIt:       controllers.GetJSONItr(mev1alpha1.GetEncoder(), mev1alpha1.GetDecoder()),
		ResourceType: "ovh_me_identity_user",
	},
	{
		Group:    "me.ovh.kubeform.com",
		Version:  "v1alpha1",
		Resource: "installationtemplates",
	}: {
		JsonIt:       controllers.GetJSONItr(mev1alpha1.GetEncoder(), mev1alpha1.GetDecoder()),
		ResourceType: "ovh_me_installation_template",
	},
	{
		Group:    "me.ovh.kubeform.com",
		Version:  "v1alpha1",
		Resource: "installationtemplatepartitionschemes",
	}: {
		JsonIt:       controllers.GetJSONItr(mev1alpha1.GetEncoder(), mev1alpha1.GetDecoder()),
		ResourceType: "ovh_me_installation_template_partition_scheme",
	},
	{
		Group:    "me.ovh.kubeform.com",
		Version:  "v1alpha1",
		Resource: "installationtemplatepartitionschemehardwareraids",
	}: {
		JsonIt:       controllers.GetJSONItr(mev1alpha1.GetEncoder(), mev1alpha1.GetDecoder()),
		ResourceType: "ovh_me_installation_template_partition_scheme_hardware_raid",
	},
	{
		Group:    "me.ovh.kubeform.com",
		Version:  "v1alpha1",
		Resource: "installationtemplatepartitionschemepartitions",
	}: {
		JsonIt:       controllers.GetJSONItr(mev1alpha1.GetEncoder(), mev1alpha1.GetDecoder()),
		ResourceType: "ovh_me_installation_template_partition_scheme_partition",
	},
	{
		Group:    "me.ovh.kubeform.com",
		Version:  "v1alpha1",
		Resource: "ipxescripts",
	}: {
		JsonIt:       controllers.GetJSONItr(mev1alpha1.GetEncoder(), mev1alpha1.GetDecoder()),
		ResourceType: "ovh_me_ipxe_script",
	},
	{
		Group:    "me.ovh.kubeform.com",
		Version:  "v1alpha1",
		Resource: "sshkeys",
	}: {
		JsonIt:       controllers.GetJSONItr(mev1alpha1.GetEncoder(), mev1alpha1.GetDecoder()),
		ResourceType: "ovh_me_ssh_key",
	},
	{
		Group:    "vrack.ovh.kubeform.com",
		Version:  "v1alpha1",
		Resource: "vracks",
	}: {
		JsonIt:       controllers.GetJSONItr(vrackv1alpha1.GetEncoder(), vrackv1alpha1.GetDecoder()),
		ResourceType: "ovh_vrack",
	},
	{
		Group:    "vrack.ovh.kubeform.com",
		Version:  "v1alpha1",
		Resource: "cloudprojects",
	}: {
		JsonIt:       controllers.GetJSONItr(vrackv1alpha1.GetEncoder(), vrackv1alpha1.GetDecoder()),
		ResourceType: "ovh_vrack_cloudproject",
	},
	{
		Group:    "vrack.ovh.kubeform.com",
		Version:  "v1alpha1",
		Resource: "dedicatedservers",
	}: {
		JsonIt:       controllers.GetJSONItr(vrackv1alpha1.GetEncoder(), vrackv1alpha1.GetDecoder()),
		ResourceType: "ovh_vrack_dedicated_server",
	},
	{
		Group:    "vrack.ovh.kubeform.com",
		Version:  "v1alpha1",
		Resource: "dedicatedserverinterfaces",
	}: {
		JsonIt:       controllers.GetJSONItr(vrackv1alpha1.GetEncoder(), vrackv1alpha1.GetDecoder()),
		ResourceType: "ovh_vrack_dedicated_server_interface",
	},
	{
		Group:    "vrack.ovh.kubeform.com",
		Version:  "v1alpha1",
		Resource: "ips",
	}: {
		JsonIt:       controllers.GetJSONItr(vrackv1alpha1.GetEncoder(), vrackv1alpha1.GetDecoder()),
		ResourceType: "ovh_vrack_ip",
	},
	{
		Group:    "vrack.ovh.kubeform.com",
		Version:  "v1alpha1",
		Resource: "iploadbalancings",
	}: {
		JsonIt:       controllers.GetJSONItr(vrackv1alpha1.GetEncoder(), vrackv1alpha1.GetDecoder()),
		ResourceType: "ovh_vrack_iploadbalancing",
	},
}

func getJsonItAndResType(gvr schema.GroupVersionResource) Data {
	return allJsonIt[gvr]
}
