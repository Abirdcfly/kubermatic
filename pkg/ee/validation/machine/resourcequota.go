//go:build ee

/*
                  Kubermatic Enterprise Read-Only License
                         Version 1.0 ("KERO-1.0”)
                     Copyright © 2022 Kubermatic GmbH

   1.	You may only view, read and display for studying purposes the source
      code of the software licensed under this license, and, to the extent
      explicitly provided under this license, the binary code.
   2.	Any use of the software which exceeds the foregoing right, including,
      without limitation, its execution, compilation, copying, modification
      and distribution, is expressly prohibited.
   3.	THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND,
      EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
      MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
      IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
      CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
      TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
      SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

   END OF TERMS AND CONDITIONS
*/

package machine

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	clusterv1alpha1 "github.com/kubermatic/machine-controller/pkg/apis/cluster/v1alpha1"
	"k8c.io/kubermatic/v2/pkg/resources/certificates"

	"k8s.io/apimachinery/pkg/api/resource"
	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

// ValidateQuota validates if the requested Machine resource consumption fits in the quota of the clusters project.
func ValidateQuota(ctx context.Context, log *zap.SugaredLogger, seedClient, userClient ctrlruntimeclient.Client,
	machine *clusterv1alpha1.Machine, caBundle *certificates.CABundle) error {
	machineResourceUsage, err := GetMachineResourceUsage(ctx, userClient, machine, caBundle)
	if err != nil {
		return fmt.Errorf("error getting machine resource request: %w", err)
	}

	// TODO Get quota and usage from ResourceQuota CRD when its implemented
	quota, currentUsage, err := getResourceQuota()
	if err != nil {
		return fmt.Errorf("failed to get resource quota: %w", err)
	}

	// add requested resources to current usage and compare
	combinedUsage := NewResourceDetails(currentUsage.cpu, currentUsage.mem, currentUsage.storage)
	combinedUsage.Cpu().Add(*machineResourceUsage.Cpu())
	combinedUsage.Memory().Add(*machineResourceUsage.Memory())
	combinedUsage.Storage().Add(*machineResourceUsage.Storage())

	if quota.Cpu().Cmp(*combinedUsage.Cpu()) < 0 {
		log.Debugw("requested CPU would exceed current quota", "request",
			machineResourceUsage.Cpu(), "quota", quota.Cpu(), "used", currentUsage.Cpu())
		return fmt.Errorf("requested CPU %q would exceed current quota (quota/used %q/%q)",
			machineResourceUsage.Cpu(), quota.Cpu(), currentUsage.Cpu())
	}

	if quota.Memory().Cmp(*combinedUsage.Memory()) < 0 {
		log.Debugw("requested Memory would exceed current quota", "request",
			machineResourceUsage.Memory(), "quota", quota.Memory(), "used", currentUsage.Memory())
		return fmt.Errorf("requested Memory %q would exceed current quota (quota/used %q/%q)",
			machineResourceUsage.Memory(), quota.Memory(), currentUsage.Memory())
	}

	if quota.Storage().Cmp(*combinedUsage.Storage()) < 0 {
		log.Debugw("requested disk size would exceed current quota", "request",
			machineResourceUsage.Storage(), "quota", quota.Storage(), "used", currentUsage.Storage())
		return fmt.Errorf("requested disk size %q would exceed current quota (quota/used %q/%q)",
			machineResourceUsage.Storage(), quota.Storage(), currentUsage.Storage())
	}

	return nil
}

// TODO we should get it from the ResourceQuota CRD for the project, for now just some hardcoded values for tests.
func getResourceQuota() (*ResourceDetails, *ResourceDetails, error) {
	cpu, err := resource.ParseQuantity("50")

	if err != nil {
		return nil, nil, fmt.Errorf("error parsing quantity: %w", err)
	}
	cpuUsed, err := resource.ParseQuantity("3")
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing quantity: %w", err)
	}

	mem, err := resource.ParseQuantity("50G")
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing quantity: %w", err)
	}
	memUsed, err := resource.ParseQuantity("3G")
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing quantity: %w", err)
	}

	storage, err := resource.ParseQuantity("1000G")
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing quantity: %w", err)
	}
	storageUsed, err := resource.ParseQuantity("60G")
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing quantity: %w", err)
	}

	return NewResourceDetails(cpu, mem, storage),
		NewResourceDetails(cpuUsed, memUsed, storageUsed), nil
}

type ResourceDetails struct {
	cpu     resource.Quantity
	mem     resource.Quantity
	storage resource.Quantity
}

func NewResourceDetails(cpu resource.Quantity, mem resource.Quantity, storage resource.Quantity) *ResourceDetails {
	return &ResourceDetails{
		cpu:     cpu,
		mem:     mem,
		storage: storage,
	}
}

func (r *ResourceDetails) Cpu() *resource.Quantity {
	return &r.cpu
}

func (r *ResourceDetails) Memory() *resource.Quantity {
	return &r.mem
}

func (r *ResourceDetails) Storage() *resource.Quantity {
	return &r.storage
}
