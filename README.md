# ShareiT
A simple cross platform app to share files within local network!

## Builds
**Latest builds available [here](https://github.com/AnjanaMadu/ShareiT/releases/tag/latest)**

## Compile
- Download and install requirements.

`sudo apt-get -qq -y update && sudo apt-get -qq -y install git snapd && sudo snap install go`

- Clone repo

`git clone https://github.com/AnjanaMadu/ShareiT shareit && cd shareit`

- Compile

`go mod download && go build`

_If you are looking for compile app for android. Please look build [workflow file](https://github.com/AnjanaMadu/ShareiT/blob/main/.github/workflows/build.yml). You will need Android NDK r19c and fyne cmd go lib too..._
