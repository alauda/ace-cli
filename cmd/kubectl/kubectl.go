package kubectl

import (
	"github.com/spf13/cobra"
	"k8s.io/kubernetes/pkg/kubectl/cmd"
)

// NewKubectlCmd creates a new embedded kubectl command.
func NewKubectlCmd() *cobra.Command {
	return cmd.NewDefaultKubectlCommand()
}
