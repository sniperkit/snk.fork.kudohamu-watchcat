# watchcat

[![GoDoc](https://godoc.org/github.com/kudohamu/watchcat?status.svg)](https://godoc.org/github.com/kudohamu/watchcat)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

watchcat is a tool to watch activities of github's repositories.

![2017-09-20 20 37 07](https://user-images.githubusercontent.com/7100252/30647891-a2111968-9e57-11e7-9cc5-ceb9eaf7de07.png)

## Usage

full command

```sh
$ watchcat \
  --conf=https://gist.githubusercontent.com/kudohamu/XXXXXXXXX \
  --slack_webhook_url=https://hooks.slack.com/services/XXXXXXXXXXX \
  --token=xxxxx-xxxxxx-xxxxxx \
  --notifiers=std,slack \
  --interval=30m \
  w
```

### --conf (required)

path of configuration file listing repositories to watch.  
you put the **toml** file according to the format below.

```toml
[[repos]]
  owner = "golang"
  name = "go"
  targets = ["release", "commit", "issue"]
[[repos]]
  owner = "golang"
  name = "dep"
  targets = ["release"]
[[repos]]
  owner = "stretchr"
  name = "testify"
  targets = ["release"]
```

you can put to `gist`, `dropbox` and etc that be able to return content of file.  
or put file to local and specify `--conf=file://your/configuration/path`.  

now, you can specify targets below.

* **commit**
* **release**
* **issue**
* **pr**
* **tag**

### --token (recommended)

github personal access token.  
I **recommend** to use [personal access token](https://github.com/settings/tokens) to **avoid rate limiting**, if you watch a lot of repositories.

### --notifiers (optional)

you can specify notifiers for notification when watched repository is changed.  

* std - standard output
* slack - slack (incoming webhook)

### --slack_webhook_url (optional)

if you specify `slack` to notifiers, you have to set this option.

### --interval (optional)

watch interval. default is 30 minutes.
