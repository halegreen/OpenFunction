//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 The OpenFunction Authors.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	apiv1alpha1 "github.com/kedacore/keda/v2/api/v1alpha1"
	"github.com/openfunction/apis/core/v1beta1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterEventBus) DeepCopyInto(out *ClusterEventBus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterEventBus.
func (in *ClusterEventBus) DeepCopy() *ClusterEventBus {
	if in == nil {
		return nil
	}
	out := new(ClusterEventBus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterEventBus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterEventBusList) DeepCopyInto(out *ClusterEventBusList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterEventBus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterEventBusList.
func (in *ClusterEventBusList) DeepCopy() *ClusterEventBusList {
	if in == nil {
		return nil
	}
	out := new(ClusterEventBusList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterEventBusList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CronSpec) DeepCopyInto(out *CronSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CronSpec.
func (in *CronSpec) DeepCopy() *CronSpec {
	if in == nil {
		return nil
	}
	out := new(CronSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventBus) DeepCopyInto(out *EventBus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventBus.
func (in *EventBus) DeepCopy() *EventBus {
	if in == nil {
		return nil
	}
	out := new(EventBus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventBus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventBusList) DeepCopyInto(out *EventBusList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EventBus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventBusList.
func (in *EventBusList) DeepCopy() *EventBusList {
	if in == nil {
		return nil
	}
	out := new(EventBusList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventBusList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventBusSpec) DeepCopyInto(out *EventBusSpec) {
	*out = *in
	if in.NatsStreaming != nil {
		in, out := &in.NatsStreaming, &out.NatsStreaming
		*out = new(NatsStreamingSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventBusSpec.
func (in *EventBusSpec) DeepCopy() *EventBusSpec {
	if in == nil {
		return nil
	}
	out := new(EventBusSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventSource) DeepCopyInto(out *EventSource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventSource.
func (in *EventSource) DeepCopy() *EventSource {
	if in == nil {
		return nil
	}
	out := new(EventSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventSource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventSourceList) DeepCopyInto(out *EventSourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EventSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventSourceList.
func (in *EventSourceList) DeepCopy() *EventSourceList {
	if in == nil {
		return nil
	}
	out := new(EventSourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventSourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventSourceSpec) DeepCopyInto(out *EventSourceSpec) {
	*out = *in
	if in.Redis != nil {
		in, out := &in.Redis, &out.Redis
		*out = make(map[string]*RedisSpec, len(*in))
		for key, val := range *in {
			var outVal *RedisSpec
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(RedisSpec)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.Kafka != nil {
		in, out := &in.Kafka, &out.Kafka
		*out = make(map[string]*KafkaSpec, len(*in))
		for key, val := range *in {
			var outVal *KafkaSpec
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(KafkaSpec)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.Cron != nil {
		in, out := &in.Cron, &out.Cron
		*out = make(map[string]*CronSpec, len(*in))
		for key, val := range *in {
			var outVal *CronSpec
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(CronSpec)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Mqtt != nil {
		in, out := &in.Mqtt, &out.Mqtt
		*out = make(map[string]*MQTTSpec, len(*in))
		for key, val := range *in {
			var outVal *MQTTSpec
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(MQTTSpec)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.Sink != nil {
		in, out := &in.Sink, &out.Sink
		*out = new(SinkSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Build != nil {
		in, out := &in.Build, &out.Build
		*out = new(v1beta1.BuildImpl)
		(*in).DeepCopyInto(*out)
	}
	if in.LogLevel != nil {
		in, out := &in.LogLevel, &out.LogLevel
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventSourceSpec.
func (in *EventSourceSpec) DeepCopy() *EventSourceSpec {
	if in == nil {
		return nil
	}
	out := new(EventSourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventSourceStatus) DeepCopyInto(out *EventSourceStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventSourceStatus.
func (in *EventSourceStatus) DeepCopy() *EventSourceStatus {
	if in == nil {
		return nil
	}
	out := new(EventSourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericScaleOption) DeepCopyInto(out *GenericScaleOption) {
	*out = *in
	if in.PollingInterval != nil {
		in, out := &in.PollingInterval, &out.PollingInterval
		*out = new(int32)
		**out = **in
	}
	if in.CooldownPeriod != nil {
		in, out := &in.CooldownPeriod, &out.CooldownPeriod
		*out = new(int32)
		**out = **in
	}
	if in.MinReplicaCount != nil {
		in, out := &in.MinReplicaCount, &out.MinReplicaCount
		*out = new(int32)
		**out = **in
	}
	if in.MaxReplicaCount != nil {
		in, out := &in.MaxReplicaCount, &out.MaxReplicaCount
		*out = new(int32)
		**out = **in
	}
	if in.Advanced != nil {
		in, out := &in.Advanced, &out.Advanced
		*out = new(apiv1alpha1.AdvancedConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AuthRef != nil {
		in, out := &in.AuthRef, &out.AuthRef
		*out = new(apiv1alpha1.ScaledObjectAuthRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericScaleOption.
func (in *GenericScaleOption) DeepCopy() *GenericScaleOption {
	if in == nil {
		return nil
	}
	out := new(GenericScaleOption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Input) DeepCopyInto(out *Input) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Input.
func (in *Input) DeepCopy() *Input {
	if in == nil {
		return nil
	}
	out := new(Input)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaScaleOption) DeepCopyInto(out *KafkaScaleOption) {
	*out = *in
	if in.GenericScaleOption != nil {
		in, out := &in.GenericScaleOption, &out.GenericScaleOption
		*out = new(GenericScaleOption)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaScaleOption.
func (in *KafkaScaleOption) DeepCopy() *KafkaScaleOption {
	if in == nil {
		return nil
	}
	out := new(KafkaScaleOption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaSpec) DeepCopyInto(out *KafkaSpec) {
	*out = *in
	if in.SaslUsername != nil {
		in, out := &in.SaslUsername, &out.SaslUsername
		*out = new(string)
		**out = **in
	}
	if in.SaslPassword != nil {
		in, out := &in.SaslPassword, &out.SaslPassword
		*out = new(string)
		**out = **in
	}
	if in.MaxMessageBytes != nil {
		in, out := &in.MaxMessageBytes, &out.MaxMessageBytes
		*out = new(int64)
		**out = **in
	}
	if in.ScaleOption != nil {
		in, out := &in.ScaleOption, &out.ScaleOption
		*out = new(KafkaScaleOption)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaSpec.
func (in *KafkaSpec) DeepCopy() *KafkaSpec {
	if in == nil {
		return nil
	}
	out := new(KafkaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MQTTSpec) DeepCopyInto(out *MQTTSpec) {
	*out = *in
	if in.ConsumerID != nil {
		in, out := &in.ConsumerID, &out.ConsumerID
		*out = new(string)
		**out = **in
	}
	if in.Qos != nil {
		in, out := &in.Qos, &out.Qos
		*out = new(int64)
		**out = **in
	}
	if in.Retain != nil {
		in, out := &in.Retain, &out.Retain
		*out = new(bool)
		**out = **in
	}
	if in.CleanSession != nil {
		in, out := &in.CleanSession, &out.CleanSession
		*out = new(bool)
		**out = **in
	}
	if in.CaCert != nil {
		in, out := &in.CaCert, &out.CaCert
		*out = new(string)
		**out = **in
	}
	if in.ClientCert != nil {
		in, out := &in.ClientCert, &out.ClientCert
		*out = new(string)
		**out = **in
	}
	if in.ClientKey != nil {
		in, out := &in.ClientKey, &out.ClientKey
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MQTTSpec.
func (in *MQTTSpec) DeepCopy() *MQTTSpec {
	if in == nil {
		return nil
	}
	out := new(MQTTSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NatsStreamingScaleOption) DeepCopyInto(out *NatsStreamingScaleOption) {
	*out = *in
	if in.GenericScaleOption != nil {
		in, out := &in.GenericScaleOption, &out.GenericScaleOption
		*out = new(GenericScaleOption)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NatsStreamingScaleOption.
func (in *NatsStreamingScaleOption) DeepCopy() *NatsStreamingScaleOption {
	if in == nil {
		return nil
	}
	out := new(NatsStreamingScaleOption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NatsStreamingSpec) DeepCopyInto(out *NatsStreamingSpec) {
	*out = *in
	if in.AckWaitTime != nil {
		in, out := &in.AckWaitTime, &out.AckWaitTime
		*out = new(string)
		**out = **in
	}
	if in.MaxInFlight != nil {
		in, out := &in.MaxInFlight, &out.MaxInFlight
		*out = new(int64)
		**out = **in
	}
	if in.DeliverNew != nil {
		in, out := &in.DeliverNew, &out.DeliverNew
		*out = new(bool)
		**out = **in
	}
	if in.StartAtSequence != nil {
		in, out := &in.StartAtSequence, &out.StartAtSequence
		*out = new(int64)
		**out = **in
	}
	if in.StartWithLastReceived != nil {
		in, out := &in.StartWithLastReceived, &out.StartWithLastReceived
		*out = new(bool)
		**out = **in
	}
	if in.DeliverAll != nil {
		in, out := &in.DeliverAll, &out.DeliverAll
		*out = new(bool)
		**out = **in
	}
	if in.StartAtTimeDelta != nil {
		in, out := &in.StartAtTimeDelta, &out.StartAtTimeDelta
		*out = new(string)
		**out = **in
	}
	if in.StartAtTime != nil {
		in, out := &in.StartAtTime, &out.StartAtTime
		*out = new(string)
		**out = **in
	}
	if in.StartAtTimeFormat != nil {
		in, out := &in.StartAtTimeFormat, &out.StartAtTimeFormat
		*out = new(string)
		**out = **in
	}
	if in.ConsumerID != nil {
		in, out := &in.ConsumerID, &out.ConsumerID
		*out = new(string)
		**out = **in
	}
	if in.ScaleOption != nil {
		in, out := &in.ScaleOption, &out.ScaleOption
		*out = new(NatsStreamingScaleOption)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NatsStreamingSpec.
func (in *NatsStreamingSpec) DeepCopy() *NatsStreamingSpec {
	if in == nil {
		return nil
	}
	out := new(NatsStreamingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisSpec) DeepCopyInto(out *RedisSpec) {
	*out = *in
	if in.EnableTLS != nil {
		in, out := &in.EnableTLS, &out.EnableTLS
		*out = new(bool)
		**out = **in
	}
	if in.Failover != nil {
		in, out := &in.Failover, &out.Failover
		*out = new(bool)
		**out = **in
	}
	if in.SentinelMasterName != nil {
		in, out := &in.SentinelMasterName, &out.SentinelMasterName
		*out = new(string)
		**out = **in
	}
	if in.RedeliverInterval != nil {
		in, out := &in.RedeliverInterval, &out.RedeliverInterval
		*out = new(string)
		**out = **in
	}
	if in.ProcessingTimeout != nil {
		in, out := &in.ProcessingTimeout, &out.ProcessingTimeout
		*out = new(string)
		**out = **in
	}
	if in.RedisType != nil {
		in, out := &in.RedisType, &out.RedisType
		*out = new(string)
		**out = **in
	}
	if in.RedisDB != nil {
		in, out := &in.RedisDB, &out.RedisDB
		*out = new(int64)
		**out = **in
	}
	if in.RedisMaxRetries != nil {
		in, out := &in.RedisMaxRetries, &out.RedisMaxRetries
		*out = new(int64)
		**out = **in
	}
	if in.RedisMinRetryInterval != nil {
		in, out := &in.RedisMinRetryInterval, &out.RedisMinRetryInterval
		*out = new(string)
		**out = **in
	}
	if in.RedisMaxRetryInterval != nil {
		in, out := &in.RedisMaxRetryInterval, &out.RedisMaxRetryInterval
		*out = new(string)
		**out = **in
	}
	if in.DialTimeout != nil {
		in, out := &in.DialTimeout, &out.DialTimeout
		*out = new(string)
		**out = **in
	}
	if in.ReadTimeout != nil {
		in, out := &in.ReadTimeout, &out.ReadTimeout
		*out = new(string)
		**out = **in
	}
	if in.WriteTimeout != nil {
		in, out := &in.WriteTimeout, &out.WriteTimeout
		*out = new(string)
		**out = **in
	}
	if in.PoolSize != nil {
		in, out := &in.PoolSize, &out.PoolSize
		*out = new(int64)
		**out = **in
	}
	if in.PoolTimeout != nil {
		in, out := &in.PoolTimeout, &out.PoolTimeout
		*out = new(string)
		**out = **in
	}
	if in.MaxConnAge != nil {
		in, out := &in.MaxConnAge, &out.MaxConnAge
		*out = new(string)
		**out = **in
	}
	if in.MinIdleConns != nil {
		in, out := &in.MinIdleConns, &out.MinIdleConns
		*out = new(int64)
		**out = **in
	}
	if in.IdleCheckFrequency != nil {
		in, out := &in.IdleCheckFrequency, &out.IdleCheckFrequency
		*out = new(string)
		**out = **in
	}
	if in.IdleTimeout != nil {
		in, out := &in.IdleTimeout, &out.IdleTimeout
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisSpec.
func (in *RedisSpec) DeepCopy() *RedisSpec {
	if in == nil {
		return nil
	}
	out := new(RedisSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Reference) DeepCopyInto(out *Reference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Reference.
func (in *Reference) DeepCopy() *Reference {
	if in == nil {
		return nil
	}
	out := new(Reference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SinkSpec) DeepCopyInto(out *SinkSpec) {
	*out = *in
	if in.Ref != nil {
		in, out := &in.Ref, &out.Ref
		*out = new(Reference)
		**out = **in
	}
	if in.Uri != nil {
		in, out := &in.Uri, &out.Uri
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SinkSpec.
func (in *SinkSpec) DeepCopy() *SinkSpec {
	if in == nil {
		return nil
	}
	out := new(SinkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Subscriber) DeepCopyInto(out *Subscriber) {
	*out = *in
	if in.Sink != nil {
		in, out := &in.Sink, &out.Sink
		*out = new(SinkSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.DeadLetterSink != nil {
		in, out := &in.DeadLetterSink, &out.DeadLetterSink
		*out = new(SinkSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Subscriber.
func (in *Subscriber) DeepCopy() *Subscriber {
	if in == nil {
		return nil
	}
	out := new(Subscriber)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Trigger) DeepCopyInto(out *Trigger) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Trigger.
func (in *Trigger) DeepCopy() *Trigger {
	if in == nil {
		return nil
	}
	out := new(Trigger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Trigger) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerList) DeepCopyInto(out *TriggerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Trigger, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerList.
func (in *TriggerList) DeepCopy() *TriggerList {
	if in == nil {
		return nil
	}
	out := new(TriggerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TriggerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerSpec) DeepCopyInto(out *TriggerSpec) {
	*out = *in
	if in.Inputs != nil {
		in, out := &in.Inputs, &out.Inputs
		*out = make(map[string]*Input, len(*in))
		for key, val := range *in {
			var outVal *Input
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(Input)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Subscribers != nil {
		in, out := &in.Subscribers, &out.Subscribers
		*out = make([]*Subscriber, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Subscriber)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.LogLevel != nil {
		in, out := &in.LogLevel, &out.LogLevel
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerSpec.
func (in *TriggerSpec) DeepCopy() *TriggerSpec {
	if in == nil {
		return nil
	}
	out := new(TriggerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerStatus) DeepCopyInto(out *TriggerStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerStatus.
func (in *TriggerStatus) DeepCopy() *TriggerStatus {
	if in == nil {
		return nil
	}
	out := new(TriggerStatus)
	in.DeepCopyInto(out)
	return out
}
