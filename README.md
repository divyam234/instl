<h1 align="center">INSTL <br> The Easiest Installer for GitHub Projects</sup></h1>

----

<p align="center">
Instl is an installation script generator for GitHub projects. <br/>

## Key Features

- 💻 Cross-Platform: Works on Windows, macOS and Linux out of the box
- 🧸 One-Click Installs: Install any GitHub project with just a single command
- 🪶 Web-Based: No need to install Instl - scripts are generated server-side
- ⚙️ Intelligent Configuration: Instl adapts to the project's structure
- 🕊️ On the fly: Installer scripts are created in real-time for each project

## Usage

Alternatively, you can create your own commands by using the following URL structure:

> [!NOTE]
> Replace `{username}` and `{reponame}` with your GitHub username and repository name.


#### Windows

```powershell
iwr https://instl.vercel.app/{username}/{reponame}/windows | iex
```

#### macOS

```bash
curl -sSL https://instl.vercel.app/{username}/{reponame}/macos | bash
```

#### Linux

```bash
curl -sSL https://instl.vercel.app/{username}/{reponame}/linux | bash
```
