package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type Plugin struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec PluginSpec `json:"spec"`
}

type PluginSpec struct {
	ConfigMapRef string `json:"configMapRef"`
}

type PluginList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Plugin `json:"items"`
}
