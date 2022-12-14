package main

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/chef/automate/components/automate-cli/pkg/status"
	"github.com/chef/automate/lib/stringutils"
	ptoml "github.com/pelletier/go-toml"
	"github.com/spf13/cobra"
)

const TAINT_TERRAFORM = "for x in $(terraform state list -state=/hab/a2_deploy_workspace/terraform/terraform.tfstate | grep module); do terraform taint $x; done"

type HAModifyAndDeploy interface {
	Execute(c *cobra.Command, args []string) error
	prepare() error
	validate() error
	modifyConfig() error
	promptUserConfirmation() (bool, error)
	runDeploy() error
}

type MockNodeUtilsImpl struct {
	executeAutomateClusterCtlCommandAsyncfunc func(command string, args []string, helpDocs string) error
	getHaInfraDetailsfunc                     func() (*AutomteHAInfraDetails, *SSHConfig, error)
	genConfigfunc                             func(path string) error
	taintTerraformFunc                        func(path string) error
	isA2HARBFileExistFunc                     func() bool
	getModeFromConfigFunc                     func(path string) (string, error)
	checkIfFileExistFunc                      func(path string) bool
	pullAndUpdateConfigFunc                   func(sshUtil *SSHUtil, exceptionIps []string) (*ExistingInfraConfigToml, error)
	isManagedServicesOnFunc                   func() bool
}

func (mnu *MockNodeUtilsImpl) executeAutomateClusterCtlCommandAsync(command string, args []string, helpDocs string) error {
	return mnu.executeAutomateClusterCtlCommandAsyncfunc(command, args, helpDocs)
}
func (mnu *MockNodeUtilsImpl) getHaInfraDetails() (*AutomteHAInfraDetails, *SSHConfig, error) {
	return mnu.getHaInfraDetailsfunc()
}
func (mnu *MockNodeUtilsImpl) genConfig(path string) error {
	return mnu.genConfigfunc(path)
}
func (mnu *MockNodeUtilsImpl) taintTerraform(path string) error {
	return mnu.taintTerraformFunc(path)
}
func (mnu *MockNodeUtilsImpl) isA2HARBFileExist() bool {
	return mnu.isA2HARBFileExistFunc()
}
func (mnu *MockNodeUtilsImpl) getModeFromConfig(path string) (string, error) {
	return mnu.getModeFromConfigFunc(path)
}
func (mnu *MockNodeUtilsImpl) checkIfFileExist(path string) bool {
	return mnu.checkIfFileExistFunc(path)
}
func (mnu *MockNodeUtilsImpl) pullAndUpdateConfig(sshUtil *SSHUtil, exceptionIps []string) (*ExistingInfraConfigToml, error) {
	return mnu.pullAndUpdateConfigFunc(sshUtil, exceptionIps)
}
func (mnu *MockNodeUtilsImpl) isManagedServicesOn() bool {
	return mnu.isManagedServicesOnFunc()
}

type NodeOpUtils interface {
	executeAutomateClusterCtlCommandAsync(command string, args []string, helpDocs string) error
	getHaInfraDetails() (*AutomteHAInfraDetails, *SSHConfig, error)
	genConfig(path string) error
	taintTerraform(path string) error
	isA2HARBFileExist() bool
	getModeFromConfig(path string) (string, error)
	checkIfFileExist(path string) bool
	pullAndUpdateConfig(sshUtil *SSHUtil, exceptionIps []string) (*ExistingInfraConfigToml, error)
	isManagedServicesOn() bool
}

type NodeUtilsImpl struct{}

func NewNodeUtils() NodeOpUtils {
	return &NodeUtilsImpl{}
}

func (nu *NodeUtilsImpl) pullAndUpdateConfig(sshUtil *SSHUtil, exceptionIps []string) (*ExistingInfraConfigToml, error) {
	infra, cfg, err := nu.getHaInfraDetails()
	if err != nil {
		return nil, err
	}
	(*sshUtil).setSSHConfig(cfg)
	configPuller := NewPullConfigs(infra, *sshUtil)
	if len(exceptionIps) > 0 {
		configPuller.setExceptionIps(exceptionIps)
	}
	return configPuller.generateConfig()
}

func (nu *NodeUtilsImpl) checkIfFileExist(path string) bool {
	return checkIfFileExist(path)
}

func (nu *NodeUtilsImpl) getModeFromConfig(path string) (string, error) {
	return getModeFromConfig(path)
}

func (nu *NodeUtilsImpl) isA2HARBFileExist() bool {
	return isA2HARBFileExist()
}

func (nu *NodeUtilsImpl) taintTerraform(path string) error {
	return executeShellCommand("/bin/sh", []string{"-c", TAINT_TERRAFORM}, path)
}

func (nu *NodeUtilsImpl) readConfig(path string) (ExistingInfraConfigToml, error) {
	return readConfig(path)
}
func (nu *NodeUtilsImpl) executeAutomateClusterCtlCommandAsync(command string, args []string, helpDocs string) error {
	return executeAutomateClusterCtlCommandAsync(command, args, helpDocs)
}

func (nu *NodeUtilsImpl) genConfig(path string) error {
	e := newExistingInfa(path)
	return e.generateConfig()
}

func (nu *NodeUtilsImpl) getHaInfraDetails() (*AutomteHAInfraDetails, *SSHConfig, error) {
	infra, err := getAutomateHAInfraDetails()
	if err != nil {
		return nil, nil, err
	}
	sshconfig := &SSHConfig{
		sshUser:    infra.Outputs.SSHUser.Value,
		sshPort:    infra.Outputs.SSHPort.Value,
		sshKeyFile: infra.Outputs.SSHKeyFile.Value,
	}
	return infra, sshconfig, nil
}

func (nu *NodeUtilsImpl) isManagedServicesOn() bool {
	return isManagedServicesOn()
}

func trimSliceSpace(slc []string) []string {
	for i := range slc {
		slc[i] = strings.TrimSpace(slc[i])
	}
	return slc
}

func modifyConfigForAddNewNode(instanceCount *string, existingPrivateIPs *[]string, newIps []string, certsIp *[]CertByIP) error {
	if len(newIps) == 0 {
		return nil
	}
	*existingPrivateIPs = append(*existingPrivateIPs, newIps...)
	inc, err := modifyInstanceCount(*instanceCount, len(newIps))
	*instanceCount = inc
	if err != nil {
		return err
	}
	if len(*certsIp) > 0 {
		for _, ip := range newIps {
			c := CertByIP{
				IP:         ip,
				PrivateKey: (*certsIp)[len(*certsIp)-1].PrivateKey,
				PublicKey:  (*certsIp)[len(*certsIp)-1].PublicKey,
			}
			*certsIp = append(*certsIp, c)
		}
	}
	return nil
}

func modifyConfigForDeleteNode(instanceCount *string, existingPrivateIPs *[]string, newIps []string, certsIp *[]CertByIP) error {
	if len(newIps) == 0 {
		return nil
	}
	*existingPrivateIPs = difference(*existingPrivateIPs, newIps)
	inc, err := modifyInstanceCount(*instanceCount, -len(newIps))
	*instanceCount = inc
	if err != nil {
		return err
	}
	if len(*certsIp) > 0 {
		for _, ip := range newIps {
			*certsIp = findAndDelete(*certsIp, ip)
		}
	}
	return nil
}

func difference(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

func findAndDelete(s []CertByIP, item string) []CertByIP {
	index := 0
	for _, i := range s {
		if i.IP != item {
			s[index] = i
			index++
		}
	}
	return s[:index]
}

func modifyInstanceCount(instanceCount string, additive int) (string, error) {
	i, err := strconv.Atoi(instanceCount)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", i+additive), nil
}

func splitIPCSV(automateIp, chefserverIp, opensearchIp, postgresIp string) (automateIpList, chefServerIpList, opensearchIpList, postgresqlIp []string) {
	if automateIp != "" {
		automateIpList = strings.Split(automateIp, ",")
		automateIpList = trimSliceSpace(automateIpList)
	}
	if chefserverIp != "" {
		chefServerIpList = strings.Split(chefserverIp, ",")
		chefServerIpList = trimSliceSpace(chefServerIpList)
	}
	if opensearchIp != "" {
		opensearchIpList = strings.Split(opensearchIp, ",")
		opensearchIpList = trimSliceSpace(opensearchIpList)
	}
	if postgresIp != "" {
		postgresqlIp = strings.Split(postgresIp, ",")
		postgresqlIp = trimSliceSpace(postgresqlIp)
	}
	return
}

func isFinalInstanceCountAllowed(current string, additive int, minAllowed int) (bool, int, error) {
	i, err := strconv.Atoi(current)
	if err != nil {
		return false, -1, err
	}
	final := i + additive
	if final < minAllowed {
		return false, final, nil
	}
	return true, final, nil
}

func checkIfPresentInPrivateIPList(existingIPArray []string, ips []string, errorPrefix string) *list.List {
	errorList := list.New()
	prefixAdder := ""
	if errorPrefix != "" {
		prefixAdder = " "
	}
	for _, ip := range ips {
		if !stringutils.SliceContains(existingIPArray, ip) {
			errorList.PushBack(fmt.Sprintf("%s%sIp %s is not present in existing list of ip addresses. Please use a different private ip.", errorPrefix, prefixAdder, ip))
		}
	}
	return errorList
}

func readConfig(path string) (ExistingInfraConfigToml, error) {
	templateBytes, err := ioutil.ReadFile(path) // nosemgrep
	if err != nil {
		return ExistingInfraConfigToml{}, status.Wrap(err, status.FileAccessError, "error in reading config toml file")
	}
	config := ExistingInfraConfigToml{}
	err = ptoml.Unmarshal(templateBytes, &config)
	if err != nil {
		return ExistingInfraConfigToml{}, status.Wrap(err, status.ConfigError, "error in unmarshalling config toml file")
	}
	return config, nil
}
