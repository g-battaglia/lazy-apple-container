# User Config:

## Opening The User Config

The location of the user config will differ depending on your OS. You can open it via lazyapple by opening the application, clicking on the 'containers' panel at the top left and pressing 'o' (or pressing 'e' if your files open in vim).

Changes to the user config will only take place after closing and re-opening lazyapple

### Location:

- macOS: `~/Library/Application Support/lazyapple/config.yml` (or `~/Library/Application Support/jesseduffield/lazyapple/config.yml` if migrating from an earlier install)

## Default:

```yml
gui:
  scrollHeight: 2
  language: "auto" # one of 'auto' | 'en' | 'pl' | 'nl' | 'de' | 'tr'
  border: "rounded" # one of 'rounded' | 'single' | 'double' | 'hidden'
  theme:
    activeBorderColor:
      - green
      - bold
    inactiveBorderColor:
      - white
    selectedLineBgColor:
      - blue
    optionsTextColor:
      - blue
  returnImmediately: false
  wrapMainPanel: true
  # Side panel width as a ratio of the screen's width
  sidePanelWidth: 0.333
  # Determines whether we show the bottom line (the one containing keybinding
  # info and the status of the app).
  showBottomLine: true
  # When true, increases vertical space used by focused side panel,
  # creating an accordion effect
  expandFocusedSidePanel: false
  # Determines which screen mode will be used on startup
  screenMode: "normal" # one of 'normal' | 'half' | 'fullscreen'
  # Determines the style of the container status and container health display in the
  # containers panel. "long": full words (default), "short": one or two characters,
  # "icon": unicode emoji.
  containerStatusHealthStyle: "long"
logs:
  timestamps: false
  since: '60m' # set to '' to show all logs
  tail: '' # set to 200 to show last 200 lines of logs
oS:
  openCommand: open {{filename}}
  openLinkCommand: open {{link}}
stats:
  graphs:
    - caption: CPU (%)
      statPath: DerivedStats.CPUPercentage
      color: blue
    - caption: Memory (%)
      statPath: DerivedStats.MemoryPercentage
      color: green
```

## Color Attributes:

For color attributes you can choose an array of attributes (with max one color attribute)
The available attributes are:

- default
- black
- red
- green
- yellow
- blue
- magenta
- cyan
- white
- bold
- reverse # useful for high-contrast
- underline

## Custom Commands

You can add custom commands like so:

```yaml
customCommands:
  containers:
    - name: bash
      attach: true
      command: 'container exec {{ .Container.ID }} bash'
      serviceNames: []
```

You may use the following go templates (such as `{{ .Container.ID }}` above) in your commands:
- [`{{ .Container }}`](https://github.com/g-battaglia/lazy-apple-container/blob/apple-container/pkg/commands/container.go) and its fields. For example: `{{ .Container.ID }}`

## Replacements

You can add replacements like so:

```yaml
replacements:
  imageNamePrefixes:
    '123456789012.dkr.ecr.us-east-1.amazonaws.com': '<prod>'
    '923456789999.dkr.ecr.us-east-1.amazonaws.com': '<dev>'
```
