package organization_test

import (
	"github.com/cloudfoundry/cli/cf/api/quotas/quotasfakes"
	"github.com/cloudfoundry/cli/cf/commandregistry"
	"github.com/cloudfoundry/cli/cf/configuration/coreconfig"
	"github.com/cloudfoundry/cli/cf/models"
	"github.com/cloudfoundry/cli/cf/requirements"
	"github.com/cloudfoundry/cli/cf/requirements/requirementsfakes"
	testcmd "github.com/cloudfoundry/cli/testhelpers/commands"
	testconfig "github.com/cloudfoundry/cli/testhelpers/configuration"
	testterm "github.com/cloudfoundry/cli/testhelpers/terminal"

	. "github.com/cloudfoundry/cli/testhelpers/matchers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("set-quota command", func() {
	var (
		ui                  *testterm.FakeUI
		quotaRepo           *quotasfakes.FakeQuotaRepository
		requirementsFactory *requirementsfakes.FakeFactory
		configRepo          coreconfig.Repository
		deps                commandregistry.Dependency
	)

	updateCommandDependency := func(pluginCall bool) {
		deps.UI = ui
		deps.Config = configRepo
		deps.RepoLocator = deps.RepoLocator.SetQuotaRepository(quotaRepo)
		commandregistry.Commands.SetCommand(commandregistry.Commands.FindCommand("set-quota").SetDependency(deps, pluginCall))
	}

	runCommand := func(args ...string) bool {
		return testcmd.RunCLICommand("set-quota", args, requirementsFactory, updateCommandDependency, false, ui)
	}

	BeforeEach(func() {
		ui = new(testterm.FakeUI)
		quotaRepo = new(quotasfakes.FakeQuotaRepository)
		configRepo = testconfig.NewRepositoryWithDefaults()
		requirementsFactory = new(requirementsfakes.FakeFactory)
	})

	It("fails with usage when provided too many or two few args", func() {
		runCommand("org")
		Expect(ui.Outputs()).To(ContainSubstrings(
			[]string{"Incorrect Usage", "Requires", "arguments"},
		))

		runCommand("org", "quota", "extra-stuff")
		Expect(ui.Outputs()).To(ContainSubstrings(
			[]string{"Incorrect Usage", "Requires", "arguments"},
		))

	})

	It("fails requirements when not logged in", func() {
		requirementsFactory.NewLoginRequirementReturns(requirements.Failing{Message: "not logged in"})
		Expect(runCommand("my-org", "my-quota")).To(BeFalse())
	})

	Context("when logged in", func() {
		BeforeEach(func() {
			org := models.Organization{}
			org.Name = "my-org"
			org.GUID = "my-org-guid"
			orgReq := new(requirementsfakes.FakeOrganizationRequirement)
			orgReq.GetOrganizationReturns(org)
			requirementsFactory.NewOrganizationRequirementReturns(orgReq)

			requirementsFactory.NewLoginRequirementReturns(requirements.Passing{})
		})

		It("assigns a quota to an org", func() {
			quota := models.QuotaFields{Name: "my-quota", GUID: "my-quota-guid"}
			quotaRepo.FindByNameReturns(quota, nil)

			runCommand("my-org", "my-quota")

			Expect(ui.Outputs()).To(ContainSubstrings(
				[]string{"Setting quota", "my-quota", "my-org", "my-user"},
				[]string{"OK"},
			))

			Expect(quotaRepo.FindByNameArgsForCall(0)).To(Equal("my-quota"))
			orgGUID, quotaGUID := quotaRepo.AssignQuotaToOrgArgsForCall(0)
			Expect(orgGUID).To(Equal("my-org-guid"))
			Expect(quotaGUID).To(Equal("my-quota-guid"))
		})
	})
})
