// Code generated by tutone: DO NOT EDIT
package types

// OpenInstallationCategory - Categorization of a Quickstart
type OpenInstallationCategory string

var OpenInstallationCategoryTypes = struct {
	// COMMUNITY category
	COMMUNITY OpenInstallationCategory
	// NEWRELIC category
	NEWRELIC OpenInstallationCategory
}{
	// COMMUNITY category
	COMMUNITY: "COMMUNITY",
	// NEWRELIC category
	NEWRELIC: "NEWRELIC",
}

// OpenInstallationOperatingSystem - Operating System of target environment
type OpenInstallationOperatingSystem string

var OpenInstallationOperatingSystemTypes = struct {
	// MacOS operating system
	DARWIN OpenInstallationOperatingSystem
	// Linux-based operating system
	LINUX OpenInstallationOperatingSystem
	// Windows operating system
	WINDOWS OpenInstallationOperatingSystem
}{
	// MacOS operating system
	DARWIN: "DARWIN",
	// Linux-based operating system
	LINUX: "LINUX",
	// Windows operating system
	WINDOWS: "WINDOWS",
}

// OpenInstallationPlatform - Operating System distribution
type OpenInstallationPlatform string

var OpenInstallationPlatformTypes = struct {
	// Amazon Linux operating system
	AMAZON OpenInstallationPlatform
	// CentOS operating system
	CENTOS OpenInstallationPlatform
	// Debian operating system
	DEBIAN OpenInstallationPlatform
	// RedHat Enterprise Linux operating system
	REDHAT OpenInstallationPlatform
	// SUSE operating system
	SUSE OpenInstallationPlatform
	// Ubuntu operating system
	UBUNTU OpenInstallationPlatform
}{
	// Amazon Linux operating system
	AMAZON: "AMAZON",
	// CentOS operating system
	CENTOS: "CENTOS",
	// Debian operating system
	DEBIAN: "DEBIAN",
	// RedHat Enterprise Linux operating system
	REDHAT: "REDHAT",
	// SUSE operating system
	SUSE: "SUSE",
	// Ubuntu operating system
	UBUNTU: "UBUNTU",
}

// OpenInstallationPlatformFamily - Operating System distribution family
type OpenInstallationPlatformFamily string

var OpenInstallationPlatformFamilyTypes = struct {
	// Debian distribution family
	DEBIAN OpenInstallationPlatformFamily
	// RHEL distribution family
	RHEL OpenInstallationPlatformFamily
	// openSUSE distribution family
	SUSE OpenInstallationPlatformFamily
}{
	// Debian distribution family
	DEBIAN: "DEBIAN",
	// RHEL distribution family
	RHEL: "RHEL",
	// openSUSE distribution family
	SUSE: "SUSE",
}

// OpenInstallationStability - Stability level of recipe
type OpenInstallationStability string

var OpenInstallationStabilityTypes = struct {
	// Recipe is disabled
	DISABLED OpenInstallationStability
	// Recipe is experimental
	EXPERIMENTAL OpenInstallationStability
	// Recipe is stable
	STABLE OpenInstallationStability
}{
	// Recipe is disabled
	DISABLED: "DISABLED",
	// Recipe is experimental
	EXPERIMENTAL: "EXPERIMENTAL",
	// Recipe is stable
	STABLE: "STABLE",
}

// OpenInstallationSuccessLinkType - Success link type
type OpenInstallationSuccessLinkType string

var OpenInstallationSuccessLinkTypeTypes = struct {
	// Explorer link
	EXPLORER OpenInstallationSuccessLinkType
	// Host entity link
	HOST OpenInstallationSuccessLinkType
}{
	// Explorer link
	EXPLORER: "EXPLORER",
	// Host entity link
	HOST: "HOST",
}

// OpenInstallationTargetType - Installation target type
type OpenInstallationTargetType string

var OpenInstallationTargetTypeTypes = struct {
	// APM agent installation
	APPLICATION OpenInstallationTargetType
	// Cloud provider installation
	CLOUD OpenInstallationTargetType
	// Docker container installation
	DOCKER OpenInstallationTargetType
	// Bare metal, virtual machine, or host-based installation
	HOST OpenInstallationTargetType
	// Kubernetes installation
	KUBERNETES OpenInstallationTargetType
	// Serverless installation
	SERVERLESS OpenInstallationTargetType
}{
	// APM agent installation
	APPLICATION: "APPLICATION",
	// Cloud provider installation
	CLOUD: "CLOUD",
	// Docker container installation
	DOCKER: "DOCKER",
	// Bare metal, virtual machine, or host-based installation
	HOST: "HOST",
	// Kubernetes installation
	KUBERNETES: "KUBERNETES",
	// Serverless installation
	SERVERLESS: "SERVERLESS",
}

// OpenInstallationAttributes - Custom event data attributes
type OpenInstallationAttributes struct {
	// Built-in parsing rulesets
	Logtype string `json:"logtype,omitempty" yaml:"logtype,omitempty"`
}

// OpenInstallationLogMatch - Matches partial list of the Log forwarding parameters
type OpenInstallationLogMatch struct {
	// List of custom attributes, as key-value pairs, that can be used to send additional data with the logs which you can then query.
	Attributes OpenInstallationAttributes `json:"attributes,omitempty" yaml:"attributes,omitempty"`
	// Path to the log file or files.
	File string `json:"file,omitempty" yaml:"file,omitempty"`
	// Name of the log or logs.
	Name string `json:"name" yaml:"name"`
	// Regular expression for filtering records.
	Pattern string `json:"pattern,omitempty" yaml:"pattern,omitempty"`
	// Service name (Linux Only).
	Systemd string `json:"systemd,omitempty" yaml:"systemd,omitempty"`
}

// OpenInstallationPostInstallConfiguration - Optional post-install configuration items
type OpenInstallationPostInstallConfiguration struct {
	// Message/Docs notice displayed to user after running the recipe
	Info string `json:"info,omitempty" yaml:"info,omitempty"`
}

// OpenInstallationPreInstallConfiguration - Optional pre-install configuration items
type OpenInstallationPreInstallConfiguration struct {
	// Message/Docs notice displayed to user prior to running recipe
	Info string `json:"info,omitempty" yaml:"info,omitempty"`
	// Message/Docs notice displayed to user prior to running recipe
	Prompt        string `json:"prompt,omitempty" yaml:"prompt,omitempty"`
	ExecDiscovery string `json:"execDiscovery,omitempty" yaml:"execDiscovery,omitempty"`
}

// OpenInstallationQuickstartEntityType - Entity type relation for Quickstart
type OpenInstallationQuickstartEntityType struct {
	// Domain of Entity
	Domain string `json:"domain" yaml:"domain"`
	// Type of Entity
	Type string `json:"type" yaml:"type"`
}

// OpenInstallationQuickstartsFilter - Metadata used to filter for Quickstarts
type OpenInstallationQuickstartsFilter struct {
	// Categorization of Quickstart
	Category OpenInstallationCategory `json:"category,omitempty" yaml:"category,omitempty"`
	// Entity relationship for Quickstart
	EntityType OpenInstallationQuickstartEntityType `json:"entityType,omitempty" yaml:"entityType,omitempty"`
	// Name of the Quickstart
	Name string `json:"name" yaml:"name"`
}

// OpenInstallationRecipe - Installation instructions and definition of an instrumentation integration
type OpenInstallationRecipe struct {
	// Named list of dependencies for this recipe
	Dependencies []string `json:"dependencies" yaml:"dependencies"`
	// Description of the recipe
	Description string `json:"description" yaml:"description"`
	// Friendly name of the integration
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// The full contents of the recipe file (yaml)
	File string `json:"file" yaml:"file"`
	// The ID
	ID string `json:"id,omitempty" yaml:"id,omitempty"`
	// List of variables to prompt for input from the user
	InputVars []OpenInstallationRecipeInputVariable `json:"inputVars" yaml:"inputVars"`
	// Go-task's taskfile definiton (see https://taskfile.dev/#/usage)
	Install string `json:"install" yaml:"install"`
	// Object representing the intended install target
	InstallTargets []OpenInstallationRecipeInstallTarget `json:"installTargets" yaml:"installTargets"`
	// Tags
	Keywords []string `json:"keywords" yaml:"keywords"`
	// # Partial list of possible Log forwarding parameters
	LogMatch []OpenInstallationLogMatch `json:"logMatch" yaml:"logMatch"`
	// Short unique handle for the name of the integration
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	// Object representing optional post-install configuration items
	PostInstall OpenInstallationPostInstallConfiguration `json:"postInstall,omitempty" yaml:"postInstall,omitempty"`
	// Object representing optional pre-install configuration items
	PreInstall OpenInstallationPreInstallConfiguration `json:"preInstall,omitempty" yaml:"preInstall,omitempty"`
	// List of process definitions used to match CLI process detection
	ProcessMatch []string `json:"processMatch" yaml:"processMatch"`
	// Metadata used to recommend and install Quickstarts
	Quickstarts OpenInstallationQuickstartsFilter `json:"quickstarts,omitempty" yaml:"quickstarts,omitempty"`
	// Github repository url
	Repository string `json:"repository" yaml:"repository"`
	// Indicates stability level of recipe
	Stability OpenInstallationStability `json:"stability,omitempty" yaml:"stability,omitempty"`
	// Metadata to support generating a URL after installation success
	SuccessLinkConfig OpenInstallationSuccessLinkConfig `json:"successLinkConfig,omitempty" yaml:"successLinkConfig,omitempty"`
	// NRQL the newrelic-cli uses to validate this recipe
	// is successfully sending data to New Relic
	ValidationNRQL NRQL `json:"validationNrql,omitempty" yaml:"validationNrql,omitempty"`
}

// OpenInstallationRecipeInputVariable - Recipe input variable prompts displayed to the user prior to execution
type OpenInstallationRecipeInputVariable struct {
	// Default value of variable
	Default string `json:"default,omitempty" yaml:"default,omitempty"`
	// Name of the variable
	Name string `json:"name" yaml:"name"`
	// Message to present to the user
	Prompt string `json:"prompt,omitempty" yaml:"prompt,omitempty"`
	// Indicates a password field
	Secret bool `json:"secret,omitempty" yaml:"secret,omitempty"`
}

// OpenInstallationRecipeInstallTarget - Matrix of supported installation criteria for this recipe
type OpenInstallationRecipeInstallTarget struct {
	// OS kernel architecture
	KernelArch string `json:"kernelArch,omitempty" yaml:"kernelArch,omitempty"`
	// OS kernel version
	KernelVersion string `json:"kernelVersion,omitempty" yaml:"kernelVersion,omitempty"`
	// Operating system
	Os OpenInstallationOperatingSystem `json:"os,omitempty" yaml:"os,omitempty"`
	// Operating System distribution
	Platform OpenInstallationPlatform `json:"platform,omitempty" yaml:"platform,omitempty"`
	// Operating System distribution family
	PlatformFamily OpenInstallationPlatformFamily `json:"platformFamily,omitempty" yaml:"platformFamily,omitempty"`
	// OS distribution version
	PlatformVersion string `json:"platformVersion,omitempty" yaml:"platformVersion,omitempty"`
	// Target type
	Type OpenInstallationTargetType `json:"type,omitempty" yaml:"type,omitempty"`
}

// OpenInstallationSuccessLinkConfig - Metadata to support generating a URL after installation success
type OpenInstallationSuccessLinkConfig struct {
	// An optional filter for appending to the URL
	Filter string `json:"filter,omitempty" yaml:"filter,omitempty"`
	// The type of the link to generate
	Type OpenInstallationSuccessLinkType `json:"type" yaml:"type"`
}

// NRQL - This scalar represents a NRQL query string.
//
// See the [NRQL Docs](https://docs.newrelic.com/docs/insights/nrql-new-relic-query-language/nrql-resources/nrql-syntax-components-functions) for more information about NRQL syntax.
type NRQL string
