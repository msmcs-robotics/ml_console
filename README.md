# ML_Console

> A Modular Framework for Managing Distributed Nueral Networks
----------------------------------------------------------------

## Dependencies

1. **Build**

In order to build this framework, you need go version 1.17 or later.

2. **Install**

These tools are a neccessity for scripts to execute properly.
The listed command is for Ubuntu 20.04.

```
sudo apt install -fy ssh-client sed
```

## A High Level Overview

 - Written in Go, this framework automates the creation, customization, and execution of Ansible Playbooks. 
 - This framework is incredibly modular, so there is always potential to expand its capability.

| Module Name | What it Does |
|-------------|--------------|
| Hosts | Manage Adding/Removing Nodes From Cluster |
| Shell | Manage Nodes With A More Shell-Like Experience |
| Data | Manage Transfering Data For NN Models Across Nodes |
| Model | Manage NN Models for Single-Node or Distributed Learning Accross Nodes |  

## How to work with the Modules

1. **Hosts** [see wiki]()
2. **Shell** [see wiki]()
3. **Data** [see wiki]()
4. **Model** [see wiki]()

## Adding Modules

Adding your own modules is very straightforward. See wiki [here]().
