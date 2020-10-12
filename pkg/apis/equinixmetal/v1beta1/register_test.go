package v1beta1

import (
	"reflect"
	"testing"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

var (
	expectedProviderSpec = EquinixMetalMachineProviderSpec{
		MachineType: "n1-standard-1",
		Facility:    "us-east1",
		UserDataSecret: &corev1.LocalObjectReference{
			Name: "myUserData",
		},
	}
	expectedRawForProviderSpec = `{"metadata":{"creationTimestamp":null},"userDataSecret":{"name":"myUserData"},"machineType":"n1-standard-1","facility":"us-east1"}`

	instanceID             = "my-instance-id"
	instanceState          = "RUNNING"
	expectedProviderStatus = EquinixMetalMachineProviderStatus{
		InstanceID:    &instanceID,
		InstanceState: &instanceState,
	}
	expectedRawForProviderStatus = `{"metadata":{"creationTimestamp":null},"instanceId":"my-instance-id","instanceState":"RUNNING"}`
)

func TestRawExtensionFromProviderSpec(t *testing.T) {
	rawExtension, err := RawExtensionFromProviderSpec(&expectedProviderSpec)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if string(rawExtension.Raw) != expectedRawForProviderSpec {
		t.Errorf("Expected: %s, got: %s", expectedRawForProviderSpec, string(rawExtension.Raw))
	}
}

func TestProviderSpecFromRawExtension(t *testing.T) {
	rawExtension := runtime.RawExtension{
		Raw: []byte(expectedRawForProviderSpec),
	}

	providerSpec, err := ProviderSpecFromRawExtension(&rawExtension)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if reflect.DeepEqual(providerSpec, expectedProviderSpec) {
		t.Errorf("Expected: %v, got: %v", expectedProviderSpec, providerSpec)
	}
}

func TestRawExtensionFromProviderStatus(t *testing.T) {
	rawExtension, err := RawExtensionFromProviderStatus(&expectedProviderStatus)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if string(rawExtension.Raw) != expectedRawForProviderStatus {
		t.Errorf("Expected: %s, got: %s", expectedRawForProviderStatus, string(rawExtension.Raw))
	}
}

func TestProviderStatusFromRawExtension(t *testing.T) {
	rawExtension := runtime.RawExtension{
		Raw: []byte(expectedRawForProviderStatus),
	}

	providerStatus, err := ProviderSpecFromRawExtension(&rawExtension)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if reflect.DeepEqual(providerStatus, expectedProviderStatus) {
		t.Errorf("Expected: %v, got: %v", expectedProviderStatus, providerStatus)
	}
}
