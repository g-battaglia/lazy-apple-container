package i18n

// TranslationSet is a set of localised strings for a given language
type TranslationSet struct {
	NotEnoughSpace                             string
	MainTitle                                  string
	GlobalTitle                                string
	Navigate                                   string
	Menu                                       string
	MenuTitle                                  string
	Execute                                    string
	Scroll                                     string
	Close                                      string
	Quit                                       string
	ErrorTitle                                 string
	NoViewMachingNewLineFocusedSwitchStatement string
	OpenConfig                                 string
	EditConfig                                 string
	ConfirmQuit                                string
	ErrorOccurred                              string
	ConnectionFailed                           string
	UnattachableContainerError                 string
	WaitingForContainerInfo                    string
	CannotAttachStoppedContainerError          string
	CannotAccessRuntimeError                   string
	CannotKillChildError                       string

	Donate                      string
	Cancel                      string
	CustomCommandTitle          string
	BulkCommandTitle            string
	Remove                      string
	HideStopped                 string
	ForceRemove                 string
	RemoveWithVolumes           string
	MustForceToRemoveContainer  string
	Confirm                     string
	Return                      string
	FocusMain                   string
	LcFilter                    string
	StopContainer               string
	RestartingStatus            string
	StartingStatus              string
	StoppingStatus              string
	RemovingStatus              string
	RunningCustomCommandStatus  string
	RunningBulkCommandStatus    string
	Stop                        string
	Restart                     string
	Start                       string
	PreviousContext             string
	NextContext                 string
	Attach                      string
	ViewLogs                    string
	ContainersTitle             string
	StandaloneContainersTitle   string
	TopTitle                    string
	ImagesTitle                 string
	VolumesTitle                string
	NetworksTitle               string
	NoContainers                string
	NoContainer                 string
	NoImages                    string
	NoVolumes                   string
	NoNetworks                  string
	RemoveImage                 string
	RemoveVolume                string
	RemoveNetwork               string
	RemoveWithoutPrune          string
	RemoveWithoutPruneWithForce string
	RemoveWithForce             string
	PruneImages                 string
	PruneContainers             string
	PruneVolumes                string
	PruneNetworks               string
	ConfirmPruneContainers      string
	ConfirmStopContainers       string
	ConfirmRemoveContainers     string
	ConfirmPruneImages          string
	ConfirmPruneVolumes         string
	ConfirmPruneNetworks        string
	PruningStatus               string
	PressEnterToReturn          string
	DetachFromContainerShortCut string
	StopAllContainers           string
	RemoveAllContainers         string
	ViewRestartOptions          string
	ExecShell                   string
	RunCustomCommand            string
	ViewBulkCommands            string
	FilterList                  string
	OpenInBrowser               string
	SortContainersByState       string

	LogsTitle                 string
	ConfigTitle               string
	EnvTitle                  string
	StatsTitle                string
	CreditsTitle              string
	ContainerConfigTitle      string
	ContainerEnvTitle         string
	NothingToDisplay          string
	CannotDisplayEnvVariables string

	No  string
	Yes string

	LcNextScreenMode string
	LcPrevScreenMode string
	FilterPrompt     string

	FocusContainers string
	FocusImages     string
	FocusVolumes    string
	FocusNetworks   string

	Kill                string
	KillingStatus       string
	RemoveContainer     string
	ContainerNotRunning string
	NoIp                string
	NoPorts             string
	NoStats             string
}

func englishSet() TranslationSet {
	return TranslationSet{
		PruningStatus:              "pruning",
		RemovingStatus:             "removing",
		RestartingStatus:           "restarting",
		StartingStatus:             "starting",
		StoppingStatus:             "stopping",
		RunningCustomCommandStatus: "running custom command",
		RunningBulkCommandStatus:   "running bulk command",

		NoViewMachingNewLineFocusedSwitchStatement: "No view matching newLineFocused switch statement",

		ErrorOccurred:                     "An error occurred! Please create an issue at https://github.com/g-battaglia/lazy-apple-container/issues",
		ConnectionFailed:                  "connection to Apple Container runtime failed. You may need to check that the 'container' CLI is installed",
		UnattachableContainerError:        "Container does not support attaching. Use exec shell ('E' key) to open an interactive shell instead",
		WaitingForContainerInfo:           "Cannot proceed until Apple Container gives us more information about the container. Please retry in a few moments.",
		CannotAttachStoppedContainerError: "You cannot attach to a stopped container, you need to start it first (which you can actually do with the 'r' key) (yes I'm too lazy to do this automatically for you) (pretty cool that I get to communicate one-on-one with you in the form of an error message though)",
		CannotAccessRuntimeError:          "Can't connect to the Apple Container runtime.\nMake sure the 'container' CLI is installed and working.",
		CannotKillChildError:              "Waited three seconds for child process to stop. There may be an orphan process that continues to run on your system.",

		Donate:  "Donate",
		Confirm: "Confirm",

		Return:                      "return",
		FocusMain:                   "focus main panel",
		LcFilter:                    "filter list",
		Navigate:                    "navigate",
		Execute:                     "execute",
		Close:                       "close",
		Quit:                        "quit",
		Menu:                        "menu",
		MenuTitle:                   "Menu",
		Scroll:                      "scroll",
		OpenConfig:                  "open lazyapple config",
		EditConfig:                  "edit lazyapple config",
		Cancel:                      "cancel",
		Remove:                      "remove",
		HideStopped:                 "hide/show stopped containers",
		ForceRemove:                 "force remove",
		RemoveWithVolumes:           "remove with volumes",
		Stop:                        "stop",
		Restart:                     "restart",
		Start:                       "start",
		PreviousContext:             "previous tab",
		NextContext:                 "next tab",
		Attach:                      "attach",
		ViewLogs:                    "view logs",
		RemoveImage:                 "remove image",
		RemoveVolume:                "remove volume",
		RemoveNetwork:               "remove network",
		RemoveWithoutPrune:          "remove without deleting untagged parents",
		RemoveWithoutPruneWithForce: "remove (forced) without deleting untagged parents",
		RemoveWithForce:             "remove (forced)",
		PruneContainers:             "prune exited containers",
		PruneVolumes:                "prune unused volumes",
		PruneNetworks:               "prune unused networks",
		PruneImages:                 "prune unused images",
		StopAllContainers:           "stop all containers",
		RemoveAllContainers:         "remove all containers (forced)",
		ViewRestartOptions:          "view restart options",
		ExecShell:                   "exec shell",
		RunCustomCommand:            "run predefined custom command",
		ViewBulkCommands:            "view bulk commands",
		FilterList:                  "filter list",
		OpenInBrowser:               "open in browser (first port is http)",
		SortContainersByState:       "sort containers by state",

		GlobalTitle:               "Global",
		MainTitle:                 "Main",
		ContainersTitle:           "Containers",
		StandaloneContainersTitle: "Standalone Containers",
		ImagesTitle:               "Images",
		VolumesTitle:              "Volumes",
		NetworksTitle:             "Networks",
		CustomCommandTitle:        "Custom Command:",
		BulkCommandTitle:          "Bulk Command:",
		ErrorTitle:                "Error",
		LogsTitle:                 "Logs",
		ConfigTitle:               "Config",
		EnvTitle:                  "Env",
		TopTitle:                  "Top",
		StatsTitle:                "Stats",
		CreditsTitle:              "About",
		ContainerConfigTitle:      "Container Config",
		ContainerEnvTitle:         "Container Env",
		NothingToDisplay:          "Nothing to display",
		CannotDisplayEnvVariables: "Something went wrong while displaying environment variables",

		NoContainers: "No containers",
		NoContainer:  "No container",
		NoImages:     "No images",
		NoVolumes:    "No volumes",
		NoNetworks:   "No networks",

		ConfirmQuit:                 "Are you sure you want to quit?",
		MustForceToRemoveContainer:  "You cannot remove a running container unless you force it. Do you want to force it?",
		NotEnoughSpace:              "Not enough space to render panels",
		ConfirmPruneImages:          "Are you sure you want to prune all unused images?",
		ConfirmPruneContainers:      "Are you sure you want to prune all stopped containers?",
		ConfirmStopContainers:       "Are you sure you want to stop all containers?",
		ConfirmRemoveContainers:     "Are you sure you want to remove all containers?",
		ConfirmPruneVolumes:         "Are you sure you want to prune all unused volumes?",
		ConfirmPruneNetworks:        "Are you sure you want to prune all unused networks?",
		StopContainer:               "Are you sure you want to stop this container?",
		PressEnterToReturn:          "Press enter to return to lazyapple (this prompt can be disabled in your config by setting `gui.returnImmediately: true`)",
		DetachFromContainerShortCut: "By default, to detach from the container press ctrl-p then ctrl-q",

		No:  "no",
		Yes: "yes",

		LcNextScreenMode: "next screen mode (normal/half/fullscreen)",
		LcPrevScreenMode: "prev screen mode",
		FilterPrompt:     "filter",

		FocusContainers: "focus containers panel",
		FocusImages:     "focus images panel",
		FocusVolumes:    "focus volumes panel",
		FocusNetworks:   "focus networks panel",

		Kill:                "kill",
		KillingStatus:       "killing",
		RemoveContainer:     "remove container",
		ContainerNotRunning: "container is not running",
		NoIp:                "n/a",
		NoPorts:             "n/a",
		NoStats:             "No stats available",
	}
}
