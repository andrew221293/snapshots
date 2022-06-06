package entity

import "time"

type Services struct {
	Kind string `json:"kind"`
	Spec struct {
		Type  string `json:"type"`
		Ports []struct {
			Name       string `json:"name"`
			Port       int    `json:"port"`
			Protocol   string `json:"protocol"`
			TargetPort string `json:"targetPort"`
		} `json:"ports"`
		Selector struct {
			AppKubernetesIoName     string `json:"app.kubernetes.io/name"`
			AppKubernetesIoInstance string `json:"app.kubernetes.io/instance"`
		} `json:"selector"`
		ClusterIP       string `json:"clusterIP"`
		SessionAffinity string `json:"sessionAffinity"`
	} `json:"spec"`
	Status struct {
		LoadBalancer struct {
		} `json:"loadBalancer"`
	} `json:"status"`
	Metadata struct {
		UID    string `json:"uid"`
		Name   string `json:"name"`
		Labels struct {
			Product                  string `json:"product"`
			Service                  string `json:"service"`
			HelmShChart              string `json:"helm.sh/chart"`
			AppKubernetesIoName      string `json:"app.kubernetes.io/name"`
			AppKubernetesIoPartOf    string `json:"app.kubernetes.io/part-of"`
			AppKubernetesIoInstance  string `json:"app.kubernetes.io/instance"`
			AppKubernetesIoManagedBy string `json:"app.kubernetes.io/managed-by"`
		} `json:"labels"`
		SelfLink    string `json:"selfLink"`
		Namespace   string `json:"namespace"`
		Annotations struct {
			MetaHelmShReleaseName      string `json:"meta.helm.sh/release-name"`
			MetaHelmShReleaseNamespace string `json:"meta.helm.sh/release-namespace"`
			A8RIoIgnore                string `json:"a8r.io/ignore"`
		} `json:"annotations"`
		ResourceVersion   string    `json:"resourceVersion"`
		CreationTimestamp time.Time `json:"creationTimestamp"`
	} `json:"metadata"`
	APIVersion string `json:"apiVersion"`
}
