//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.

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

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Account) DeepCopyInto(out *Account) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Account.
func (in *Account) DeepCopy() *Account {
	if in == nil {
		return nil
	}
	out := new(Account)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Account) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountList) DeepCopyInto(out *AccountList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Account, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountList.
func (in *AccountList) DeepCopy() *AccountList {
	if in == nil {
		return nil
	}
	out := new(AccountList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccountList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountSpec) DeepCopyInto(out *AccountSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountSpec.
func (in *AccountSpec) DeepCopy() *AccountSpec {
	if in == nil {
		return nil
	}
	out := new(AccountSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountStatus) DeepCopyInto(out *AccountStatus) {
	*out = *in
	if in.EncryptBalance != nil {
		in, out := &in.EncryptBalance, &out.EncryptBalance
		*out = new(string)
		**out = **in
	}
	if in.EncryptDeductionBalance != nil {
		in, out := &in.EncryptDeductionBalance, &out.EncryptDeductionBalance
		*out = new(string)
		**out = **in
	}
	if in.ChargeList != nil {
		in, out := &in.ChargeList, &out.ChargeList
		*out = make([]Charge, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountStatus.
func (in *AccountStatus) DeepCopy() *AccountStatus {
	if in == nil {
		return nil
	}
	out := new(AccountStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BillingInfoQuery) DeepCopyInto(out *BillingInfoQuery) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BillingInfoQuery.
func (in *BillingInfoQuery) DeepCopy() *BillingInfoQuery {
	if in == nil {
		return nil
	}
	out := new(BillingInfoQuery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BillingInfoQuery) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BillingInfoQueryList) DeepCopyInto(out *BillingInfoQueryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BillingInfoQuery, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BillingInfoQueryList.
func (in *BillingInfoQueryList) DeepCopy() *BillingInfoQueryList {
	if in == nil {
		return nil
	}
	out := new(BillingInfoQueryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BillingInfoQueryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BillingInfoQuerySpec) DeepCopyInto(out *BillingInfoQuerySpec) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BillingInfoQuerySpec.
func (in *BillingInfoQuerySpec) DeepCopy() *BillingInfoQuerySpec {
	if in == nil {
		return nil
	}
	out := new(BillingInfoQuerySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BillingInfoQueryStatus) DeepCopyInto(out *BillingInfoQueryStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BillingInfoQueryStatus.
func (in *BillingInfoQueryStatus) DeepCopy() *BillingInfoQueryStatus {
	if in == nil {
		return nil
	}
	out := new(BillingInfoQueryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BillingPropertiesForQuery) DeepCopyInto(out *BillingPropertiesForQuery) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BillingPropertiesForQuery.
func (in *BillingPropertiesForQuery) DeepCopy() *BillingPropertiesForQuery {
	if in == nil {
		return nil
	}
	out := new(BillingPropertiesForQuery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BillingRecord) DeepCopyInto(out *BillingRecord) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BillingRecord.
func (in *BillingRecord) DeepCopy() *BillingRecord {
	if in == nil {
		return nil
	}
	out := new(BillingRecord)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BillingRecordQuery) DeepCopyInto(out *BillingRecordQuery) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BillingRecordQuery.
func (in *BillingRecordQuery) DeepCopy() *BillingRecordQuery {
	if in == nil {
		return nil
	}
	out := new(BillingRecordQuery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BillingRecordQuery) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BillingRecordQueryItem) DeepCopyInto(out *BillingRecordQueryItem) {
	*out = *in
	in.Time.DeepCopyInto(&out.Time)
	in.BillingRecordQueryItemInline.DeepCopyInto(&out.BillingRecordQueryItemInline)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BillingRecordQueryItem.
func (in *BillingRecordQueryItem) DeepCopy() *BillingRecordQueryItem {
	if in == nil {
		return nil
	}
	out := new(BillingRecordQueryItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BillingRecordQueryItemInline) DeepCopyInto(out *BillingRecordQueryItemInline) {
	*out = *in
	if in.Costs != nil {
		in, out := &in.Costs, &out.Costs
		*out = make(Costs, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Payment != nil {
		in, out := &in.Payment, &out.Payment
		*out = new(PaymentForQuery)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BillingRecordQueryItemInline.
func (in *BillingRecordQueryItemInline) DeepCopy() *BillingRecordQueryItemInline {
	if in == nil {
		return nil
	}
	out := new(BillingRecordQueryItemInline)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BillingRecordQueryList) DeepCopyInto(out *BillingRecordQueryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BillingRecordQuery, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BillingRecordQueryList.
func (in *BillingRecordQueryList) DeepCopy() *BillingRecordQueryList {
	if in == nil {
		return nil
	}
	out := new(BillingRecordQueryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BillingRecordQueryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BillingRecordQuerySpec) DeepCopyInto(out *BillingRecordQuerySpec) {
	*out = *in
	in.StartTime.DeepCopyInto(&out.StartTime)
	in.EndTime.DeepCopyInto(&out.EndTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BillingRecordQuerySpec.
func (in *BillingRecordQuerySpec) DeepCopy() *BillingRecordQuerySpec {
	if in == nil {
		return nil
	}
	out := new(BillingRecordQuerySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BillingRecordQueryStatus) DeepCopyInto(out *BillingRecordQueryStatus) {
	*out = *in
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BillingRecordQueryItem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BillingRecordQueryStatus.
func (in *BillingRecordQueryStatus) DeepCopy() *BillingRecordQueryStatus {
	if in == nil {
		return nil
	}
	out := new(BillingRecordQueryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Charge) DeepCopyInto(out *Charge) {
	*out = *in
	in.Time.DeepCopyInto(&out.Time)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Charge.
func (in *Charge) DeepCopy() *Charge {
	if in == nil {
		return nil
	}
	out := new(Charge)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Costs) DeepCopyInto(out *Costs) {
	{
		in := &in
		*out = make(Costs, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Costs.
func (in Costs) DeepCopy() Costs {
	if in == nil {
		return nil
	}
	out := new(Costs)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Debt) DeepCopyInto(out *Debt) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Debt.
func (in *Debt) DeepCopy() *Debt {
	if in == nil {
		return nil
	}
	out := new(Debt)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Debt) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DebtList) DeepCopyInto(out *DebtList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Debt, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DebtList.
func (in *DebtList) DeepCopy() *DebtList {
	if in == nil {
		return nil
	}
	out := new(DebtList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DebtList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DebtSpec) DeepCopyInto(out *DebtSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DebtSpec.
func (in *DebtSpec) DeepCopy() *DebtSpec {
	if in == nil {
		return nil
	}
	out := new(DebtSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DebtStatus) DeepCopyInto(out *DebtStatus) {
	*out = *in
	if in.DebtStatusRecords != nil {
		in, out := &in.DebtStatusRecords, &out.DebtStatusRecords
		*out = make([]DebtStatusRecord, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DebtStatus.
func (in *DebtStatus) DeepCopy() *DebtStatus {
	if in == nil {
		return nil
	}
	out := new(DebtStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DebtStatusRecord) DeepCopyInto(out *DebtStatusRecord) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DebtStatusRecord.
func (in *DebtStatusRecord) DeepCopy() *DebtStatusRecord {
	if in == nil {
		return nil
	}
	out := new(DebtStatusRecord)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceBillingHistory) DeepCopyInto(out *NamespaceBillingHistory) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceBillingHistory.
func (in *NamespaceBillingHistory) DeepCopy() *NamespaceBillingHistory {
	if in == nil {
		return nil
	}
	out := new(NamespaceBillingHistory)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NamespaceBillingHistory) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceBillingHistoryList) DeepCopyInto(out *NamespaceBillingHistoryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NamespaceBillingHistory, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceBillingHistoryList.
func (in *NamespaceBillingHistoryList) DeepCopy() *NamespaceBillingHistoryList {
	if in == nil {
		return nil
	}
	out := new(NamespaceBillingHistoryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NamespaceBillingHistoryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceBillingHistorySpec) DeepCopyInto(out *NamespaceBillingHistorySpec) {
	*out = *in
	in.StartTime.DeepCopyInto(&out.StartTime)
	in.EndTime.DeepCopyInto(&out.EndTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceBillingHistorySpec.
func (in *NamespaceBillingHistorySpec) DeepCopy() *NamespaceBillingHistorySpec {
	if in == nil {
		return nil
	}
	out := new(NamespaceBillingHistorySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceBillingHistoryStatus) DeepCopyInto(out *NamespaceBillingHistoryStatus) {
	*out = *in
	if in.NamespaceList != nil {
		in, out := &in.NamespaceList, &out.NamespaceList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceBillingHistoryStatus.
func (in *NamespaceBillingHistoryStatus) DeepCopy() *NamespaceBillingHistoryStatus {
	if in == nil {
		return nil
	}
	out := new(NamespaceBillingHistoryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Payment) DeepCopyInto(out *Payment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Payment.
func (in *Payment) DeepCopy() *Payment {
	if in == nil {
		return nil
	}
	out := new(Payment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Payment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PaymentForQuery) DeepCopyInto(out *PaymentForQuery) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PaymentForQuery.
func (in *PaymentForQuery) DeepCopy() *PaymentForQuery {
	if in == nil {
		return nil
	}
	out := new(PaymentForQuery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PaymentList) DeepCopyInto(out *PaymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Payment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PaymentList.
func (in *PaymentList) DeepCopy() *PaymentList {
	if in == nil {
		return nil
	}
	out := new(PaymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PaymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PaymentSpec) DeepCopyInto(out *PaymentSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PaymentSpec.
func (in *PaymentSpec) DeepCopy() *PaymentSpec {
	if in == nil {
		return nil
	}
	out := new(PaymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PaymentStatus) DeepCopyInto(out *PaymentStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PaymentStatus.
func (in *PaymentStatus) DeepCopy() *PaymentStatus {
	if in == nil {
		return nil
	}
	out := new(PaymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PriceQuery) DeepCopyInto(out *PriceQuery) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PriceQuery.
func (in *PriceQuery) DeepCopy() *PriceQuery {
	if in == nil {
		return nil
	}
	out := new(PriceQuery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PriceQuery) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PriceQueryList) DeepCopyInto(out *PriceQueryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PriceQuery, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PriceQueryList.
func (in *PriceQueryList) DeepCopy() *PriceQueryList {
	if in == nil {
		return nil
	}
	out := new(PriceQueryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PriceQueryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PriceQuerySpec) DeepCopyInto(out *PriceQuerySpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PriceQuerySpec.
func (in *PriceQuerySpec) DeepCopy() *PriceQuerySpec {
	if in == nil {
		return nil
	}
	out := new(PriceQuerySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PriceQueryStatus) DeepCopyInto(out *PriceQueryStatus) {
	*out = *in
	if in.BillingRecords != nil {
		in, out := &in.BillingRecords, &out.BillingRecords
		*out = make([]BillingRecord, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PriceQueryStatus.
func (in *PriceQueryStatus) DeepCopy() *PriceQueryStatus {
	if in == nil {
		return nil
	}
	out := new(PriceQueryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PropertyQuery) DeepCopyInto(out *PropertyQuery) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PropertyQuery.
func (in *PropertyQuery) DeepCopy() *PropertyQuery {
	if in == nil {
		return nil
	}
	out := new(PropertyQuery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Transfer) DeepCopyInto(out *Transfer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Transfer.
func (in *Transfer) DeepCopy() *Transfer {
	if in == nil {
		return nil
	}
	out := new(Transfer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Transfer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransferList) DeepCopyInto(out *TransferList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Transfer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransferList.
func (in *TransferList) DeepCopy() *TransferList {
	if in == nil {
		return nil
	}
	out := new(TransferList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TransferList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransferSpec) DeepCopyInto(out *TransferSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransferSpec.
func (in *TransferSpec) DeepCopy() *TransferSpec {
	if in == nil {
		return nil
	}
	out := new(TransferSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransferStatus) DeepCopyInto(out *TransferStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransferStatus.
func (in *TransferStatus) DeepCopy() *TransferStatus {
	if in == nil {
		return nil
	}
	out := new(TransferStatus)
	in.DeepCopyInto(out)
	return out
}