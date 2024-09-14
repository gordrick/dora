# Dora Daemon Project

## Overview

This project includes a daemon service and a GUI application interacts with MacOS to check file stats and run commands through an API.

## Table of Contents

1. [Installation Instructions](#installation-instructions)
2. [Uninstallation Instructions](#uninstallation-instructions)
3. [API Schema](#api-schema)

---

## Installation Instructions

### Daemon Installation (.pkg)

To install the daemon, follow these steps:

1. **Download the Daemon Installer**:
    - Obtain the `.pkg` installer for the daemon provided.

2. **Run the Installer**:
    - Open the terminal and navigate to the directory where the `.pkg` file is located.
    - Run the following command:
      ```bash
      sudo installer -pkg ./dora.pkg -target /
      ```

3**Start the Daemon**:
    - If not automatically started, run:
      ```bash
      sudo launchctl load /Library/LaunchDaemons/com.gordrick.dora.plist
      ```

### GUI Installation (Binary)

1. **Download the GUI Binary**:
    - Obtain the GUI binary provided.

2. **Make the Binary Executable**:
   ```bash
   chmod +x ./dora-gui

3. **Move the Binary to /Applications:**:
   ```bash
   chmod +x ./dora-gui
   
4. **Run the Binary**:
    ```bash
     /Applications/dora-gui
    ```


## Uninstallation Instructions

Run the Uninstall Script Provided:
    ```bash 
    ./scripts/uninstall
    ``` 


## API Schema

| Endpoint     | Method | Description                         | Request Body                                   | Responses                                       |
|--------------|--------|-------------------------------------|------------------------------------------------|------------------------------------------------|
| `/commands`  | POST   | Add commands to the execution queue | `{ "commands": ["<command1>", "<command2>"] }` | `200 OK`, `400 Bad Request`, `503 Service Unavailable` |
| `/logs`      | GET    | Fetch the logs from the daemon      | N/A                                            | `200 OK`, `500 Internal Server Error`           | | `200 OK` (Healthy/Degraded)                     |



