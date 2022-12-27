#!/bin/sh

# we dont want to proceed if something fails
set -e

REPO="https://github.com/alexcoder04/golored"

# check if version number is passed
if [ -z "$1" ]; then
  echo "Please pass the version number as first argument"
  exit 1
fi

VERSION="$1"
TAG_NAME="v$VERSION"

# check if version number is in right format
echo "Checking version for right format..."
if ! echo "$VERSION" | grep -qE '^[0-9]+\.[0-9]+\.[0-9]+$'; then
  echo "Your version number is not in the semver format"
  exit 1
fi

release_default(){
  echo "Creating default release"

  # check if this version already exists
  echo "Checking if $TAG_NAME already exists..."
  if git tag | grep -qE "^$TAG_NAME\$"; then
    echo "The tag $TAG_NAME already exists, please increase the version"
    exit 1
  fi

  # check whether the version number is newer
  echo "Checking if the new version is newer than latest existing..."
  LATEST_EXISTING="$(git tag --sort=-version:refname | head -n 1 | tr -d v)"
  if [ "$(echo "$LATEST_EXISTING\n$VERSION" | sort -r | head -n 1)" != "$VERSION" ]; then
    echo "Your specified version is older than the latest already existsing version"
    exit 1
  fi

  # tidy the deps
  echo "Tidying the project dependencies..."
  go mod tidy

  # check for unstaged changes
  echo "Checking for unstaged changes..."
  if [ ! -z "$(git status -s)" ]; then
    echo "You have unstaged changes, please commit them first"
    exit 1
  fi

  # create new tag
  echo "Tagging..."
  git tag "$TAG_NAME"

  # push everything, pushing the tags will trigger release generation on github
  echo "Pushing..."
  git push
  git push --tags
}

release_arch(){
  # check if we are on arch
  if ! command -v makepkg >/dev/null; then
    echo "You are not on Arch, not updating PKGBUILD"
    exit 1
  fi

  echo "Creating Arch Linux release"

  # check if the tag exists
  echo "Checking if $TAG_NAME exists..."
  if ! git tag | grep -qE "^$TAG_NAME\$"; then
    echo "The tag $TAG_NAME does not exist, please run default release first"
    exit 1
  fi

  echo "Checking if already the latest version..."
  VERSION_ARCH="$(grep -E '^pkgver=.*$' ".aur/PKGBUILD" | cut -d "=" -f2)"
  if [ "$VERSION_ARCH" = "$VERSION" ]; then
    echo "PKGBUILD already latest version"
    exit 1
  fi

  # download the .tar.gz for the tag from github and calculate the ms5sum
  echo "Calculating the md5sum..."
  cd "${TMPDIR:-/tmp}"
  wget -O "$TAG_NAME.tar.gz" "$REPO/archive/refs/tags/$TAG_NAME.tar.gz" || exit 1
  MD5SUM="$(md5sum "$TAG_NAME.tar.gz" | cut -d " " -f1)"
  cd "$OLDPWD"

  # change package version and md5sum for the PKGBUILD
  echo "Updating PKGBUILD..."
  sed -i "s/^pkgver=.*\$/pkgver=$VERSION/" ".aur/PKGBUILD"
  sed -i "s/^md5sums=('.*')\$/md5sums=('$MD5SUM')/" ".aur/PKGBUILD"

  # generate .SRCINFO
  cd ".aur"
  echo "Generating .SRCINFO"
  makepkg --printsrcinfo >.SRCINFO

  # push to AUR
  echo "Committing changes in the AUR package..."
  git add PKGBUILD .SRCINFO
  #git commit -S -m "update to $TAG_NAME"
  #echo "Pushing changes to the AUR..."
  #git push
}

case "$1" in
  default) release_default ;;
  arch) release_arch ;;
  *)
    release_default
    # wait a little, so we are sure that the tag is on github
    sleep 5
    release_arch
    ;;
esac
