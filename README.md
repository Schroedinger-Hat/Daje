# Daje - A general purpose Ansible dotfiles installer

Configure one time and run everywhere.

## What is this

I've never liked using a big bash script to install my dotfiles, because the more you add, the more that becomes: bloated, complex to maintain, and less reproducible across multiple platforms. This is why I start write this simple and easy to mantain playbook template that works in every *nix system.

## How it works

***Create a profile, define the environment you want and let ansible do the rest.***

A **Profile** could be a generic os configuration (e.g.: [os_Archlinux](vars/os_Archlinux.yml) or [os_MacOSX](vars/os_MacOSX.yml))  or a specific profile (e.g.: [profile1](vars/profiles/profile1.yml)).

Every package and configuration specified on the Profile will be installed.

## How to use it

First of all you need to have installed [ansible](https://docs.ansible.com/ansible/latest/index.html), you could do with [pip](https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html#installing-and-upgrading-ansible-with-pip) or with [any other package manage](https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html#installing-ansible-on-specific-operating-systems).

Than create a Profile and configure it, you can find on [vars](vars/) directory some examples where is possible to take one of it as a template.

After create your own profile run:

- Generic os profile (must be in [vars](vars/) directory)

	```Bash
	ansible-playbook -K install.yml
	```

- Specific profile (must be in [vars/profiles](vars/profiles) directory)

	```Bash
	ansible-playbook -K install.yml -e 'profile=<profile_name>'
	```

## List of all variables supported
This is the complete list of every variable that is possible to configure:

| Variable     | Required  | Example   | Description |
|--------------|-----------|-----------|-------------|
| username | yes | wabri | Username where the configuration will be apply |
| group | no | staff (default: {{ username }}) | Group of username |
| install_command | yes | brew install --quiet | Package manager command to install dependencies (N.B.: This need to be a no prompt install because we will not have the access to the standard input, for more info go to [Install_command](#install_command) section) |
| update_command | no | pacman -Syyu --noconfirm | Package manager command to install dependencies (N.B.: all the rules of install_command must be apply to this variable) |
| home_dir | yes | /home/wabri | Home directory where the configuration file must be place |
| workspace_dir | no | Documents/Workspaces | Workspaces directory to create |
| default_shell | no | zsh | Default shell to apply |
| default_keyboard | no | us | Default kayboard layout to apply |
| ide_install | no | true | Is a boolean variable used to enable or not the ide installation. (N.B. [Ide vs System vs General](#ide-vs-system-vs-general)) |
| system_install | no | true | Is a boolean variable used to enable or not the system installation (N.B. [Ide vs System vs General](#ide-vs-system-vs-general)) |
| general_install | no | true | Is a boolean variable used to enable or not the general installation (N.B. [Ide vs System vs General](#ide-vs-system-vs-general)) |
| before | no | - paru <br> - homebrew | List of the tasks to apply before every other tasks. This could be helpfull for package manager installation or multiple service that need to be stop. (N.B. [List of all tasks](#list-of-all-tasks)) |
| after | no | - bluetooth | Like before, but after. (N.B. [List of all tasks](#list-of-all-tasks))|
| ide_packages | no | - git<br>- zsh<br>- vim<br>- tmux | List of packages to install for ide environment. |
| system_packages | no | - htop<br>- vagrant<br>- arandr | Like ide_packages but for system_packages. |
| general_packages | no | - firefox<br>- spotify | Like system_packages but for general_packages. |
| ide_config | no | - vim<br>- zsh<br>-ohmyzsh | List of tasks that apply configurations for the ide environment. (N.B. [List of all tasks](#list-of-all-tasks)) |
| system_config | no | - i3<br>- gtk3 | Like ide_config but for system_packages. (N.B. [List of all tasks](#list-of-all-tasks)) |
| general_config | no | - vim<br>- zsh<br>-ohmyzsh | Like system_config but for general_packages. (N.B. [List of all tasks](#list-of-all-tasks)) |

### Install_command

In order to have a polished and automated installation is necessary to set up a package manager installation command to run without any prompt for standard input. An example could be:

- For Arch distribution: `pacman -Y --noconfirm` or `paru -Y --noconfirm`
- For MacOSX: `brew install --quiet`
- For Debian distribution: `apt-get install --yes`

## Ide vs System vs General

If you're wondering why this separation from: Ide, System and General, the answer is easy to say: division of responsibilities. If I want to refresh the ide configuration or add a new package to it, I don't want to go through the system and general configuration, so I will disable the system_install and general_install and run the installation.

If you're wondering where to put a configuration task or a package installation, ask this question:

- Is for my vim or visual studio code (for example) configuration? Yes, then put on ide.
- Is for my i3 or yabai configuration? Yes, then put on system.
- Is for my firefox or spotify configuration? Yes, then put on general.

## List of all tasks

| Task     | Requirements/Description | Reference |
|----------|--------------|-------------|
| [alacritty](tasks/alacritty.yml) | Configuration file must be on .config/alacritty/alacritty.yml | [github/alacritty](https://github.com/alacritty/alacritty) |
| [bash](tasks/bash.yml) | Configuration file must be on .bashrc | [gnu/bash](https://www.gnu.org/software/bash/) |
| [bluetooth](tasks/bluetooth.yml) | These task enable systemd bluetooth service | --- |
| [dunst](tasks/dunst.yml) | Configuration file must be on .config/dunst/ | [dunst-project](https://dunst-project.org) |
| [git](tasks/git.yml) | Configuration file must be .gitconfig and .gitignore | [git](https://git-scm.com/) |
| [gtk3](tasks/gtk3.yml) | Configuration file must be .config/gtk-3.0/ | [gtk.org](https://www.gtk.org/) |
| [homebrew](tasks/homebrew.yml) | Install homebrew on MacOSX | [brew.sh](https://brew.sh/) |
| [i3](tasks/i3.yml) | Configuration file must be .config/i3/ | [i3wm.org](https://i3wm.org/) |
| [nitrogen](tasks/nitrogen.yml) | Configuration file must be .config/nitrogen/ | [github/nitrogen](https://github.com/l3ib/nitrogen/) |
| [ohmyzsh](tasks/ohmyzsh.yml) | Install oh-my-zsh | [ohmyz.sh](https://ohmyz.sh/) |
| [paru](tasks/paru.yml) | Install paru on Archlinux | [github/paru](https://github.com/morganamilo/paru) |
| [polybar](tasks/polybar.yml) | Configuration file must be .config/polybar/ | [github/polybar](https://polybar.github.io/) |
| [rofi](tasks/rofi.yml) | Configuration file must be .config/rofi/ | [davatorium/rofi](https://github.com/davatorium/rofi) |
| [rofithemes](tasks/rofithemes.yml) | Install rofi themes collection | [lr-tech/rofi-themes-collection](https://github.com/lr-tech/rofi-themes-collection) |
| [sddm-sugar-candy](tasks/sddm-sugar-candy.yml) | Install sddm sugar candy theme (works only on arch with AUR packages enabled) | [aur/sddm-sugar-candy-git](https://store.kde.org/p/1312658/) |
| [spotify](tasks/spotify.yml) | Install spotify with spicetify themes and more. Must modify task if want a different theming. (For now it works only on linux... sorry!) | [khanhas/spicetify-cli](https://github.com/khanhas/spicetify-cli) |
| [starship](tasks/starship.yml) | Install starship the minimal, blazing-fast, and infinitely customizable prompt for any shell!  | [starship.rs](starship.rs) |
| [tmux](tasks/tmux.yml) | Configuration file must be .tmux.conf  | [github/tmux](https://github.com/tmux/tmux) |
| [tpm](tasks/tpm.yml) | Install tmux plugin manager and plugins listed on .tmux.conf  | [github/tmux-plugins](https://github.com/tmux-plugins/tpm) |
| [vim-plug](tasks/vim-plug.yml) | Install vim-plug and plugins listed on .vim/plugins.vim | [github/vim-plug](https://github.com/junegunn/vim-plug) |
| [vim](tasks/vim.yml) | Configuration file must be .vimrc | [vim.org](https://www.vim.org/) |
| [zsh](tasks/zsh.yml) | Configuration file must be .zshrc | [zsh.org](https://www.zsh.org/) |

## Q&A

### Can I put everything on a single package and configuration list?

It will works, but are you sure to do it? See [Ide vs System vs General](#ide-vs-system-vs-general).

### I have created a new task and I want to add to this template, how can I do it?

Open a pull request and add that task.

## Example repository

I create these template for my dotfiles -> [wabri/dotfiles](https://github.com/Wabri/dotfiles).


