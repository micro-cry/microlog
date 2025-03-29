# Contributing

Thank you for your interest in contributing to this project. Below you will find the guidelines to help you get started and ensure a smooth collaboration process.

## Automation for Unix Systems

For Unix-based systems, we provide bash automation scripts to streamline the setup and deployment process.  
- **Initial Setup:**  
  Run the script `_run/firststart.sh` once during the initial setup on your local environment. This will configure the project so that all future commits and version updates occur automatically.
  
- **Local Version Management:**  
  The local version of the project is specified in the file `_run/values/ver.txt`. Please ensure that any changes to the project are consistent with the versioning outlined in this file.

## Commit Message Format

All commits that introduce functional changes must follow this format: 

```{branch name} [{local version}] \n {detailed description of changes}```

- **Branch Name:** The name of the branch where the commit was made.
- **Local Version:** The version from `_run/values/ver.txt` should be included in square brackets.
- **Description:** A detailed explanation of the changes being introduced.

**Note:** If possible, sign your commits using GPG. This adds an extra layer of authenticity and security to your contributions.

## Code and Documentation

- **Code Changes:**  
  Ensure that your code is well-commented. Detailed inline comments that explain the logic and functionality will be highly appreciated during the review process.

- **Documentation Changes:**  
  Changes to documentation do not require the strict commit format applied to code. However, when submitting a pull request that includes documentation updates, please provide a thorough explanation of what has changed and the reasons behind these modifications.

Thank you for helping us improve the project. Your contributions make a significant difference!
