# Contri Calculator

---

![License](https://img.shields.io/badge/license-GPL%20v3-blue.svg)

It is a simple tool designed to help groups split bills and expenses fairly.

### Table of Contents

1. [How to Install](#how-to-install)
2. [How to Run](#how-to-run)
3. [Version](#version)
4. [Usage](#usage)
5. [Output](#output)
6. [Notes](#notes)

### How to Install

#### 1. **Install Go** (if not already installed)

Before using `go install`, you need to have Go installed on your system. If you don't have Go installed, you can download and install it from [Go's official website](https://golang.org/dl/).

#### 2. **Install Contri Calculator** globally using Go

Once Go is installed, you can use the following command to install the Contri Calculator globally:

```bash
go install github.com/bhpcv252/contri-calculator/cmd/contri@latest
```

#### 3. **Download the Binary (Optional)**

Alternatively, you can download the precompiled binary for your platform from the [Releases Page](https://github.com/bhpcv252/contri-calculator/releases). Choose the appropriate binary for your system and follow the instructions to install it.

### How to Run

Once installed, execute the following command to run the project:

```bash
contri
```

If you want to run the project directly from the source without installing it, you can run:

```bash
make run
```

This will:

- Build the binary.
- Run the program.

To generate a binary file without running it, you can run:

```bash
make build
```

The binary file will be available in the current directory as `contri`.

### Version

You can check the current version of Contri by running either of the following commands:

```bash
contri --version
```

or:

```bash
contri -v
```

### Usage

Upon running, you'll be prompted to enter users' information. Leave empty to finish and calculate.

#### Input:

- **Person's name** (Required)
- **Amount paid**: The amount already paid by the person, if any (Optional)
- **Budget**: How much the user can afford (Optional)

#### Example:

```
(Name Amount Budget): John 0 100
(Name Amount Budget): Jane 200
(Name Amount Budget): Dave
(Name Amount Budget): Diana 350
```

### Output

Contri will calculate how much each person should contribute or receive, based on the amount they've paid and their budget. For example:

```
John -> 100.00 (Max Affordability)
Jane -> 100.00
Dave -> 300.00
Diana -> -50.00 (To be Collected)
```

### Notes:

- Contri will attempt to divide the total expense equally among the users or based on their budgets.
- Users who paid more than their budget will have the excess to be collected from others.
