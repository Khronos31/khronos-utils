#!/bin/bash
# com.khronos.go | make.sh

set -e
set -u

export PROJECTROOT="$(cd $(dirname $0)&&pwd)"
readonly PROJECTROOT
cd "${PROJECTROOT}"

export BUILDROOT

export ARCH="${ARCH-$(arch)}"
if [ "${ARCH}" = arm ]; then
  ARCH=armv7
fi

export VERSION="0.0.1"

export CC=clang
export CXX=clang++
export AR=llvm-ar

export GOOS=darwin
export GOARCH
export CGO_ENABLED=1

export APPS
APPS=(clip say)

init() {
  ARCH="$1"
  BUILDROOT="${PROJECTROOT}/${ARCH}"
  if [ -e "${BUILDROOT}" ]; then
    rm -rf "${BUILDROOT}"
  fi
  mkdir -p "${BUILDROOT}"
}

build() {
  cd "${PROJECTROOT}"
  if [ "${ARCH}" = armv7 ]; then
    GOARCH=arm
  else
    GOARCH="${ARCH}"
  fi

  local destdir="${BUILDROOT}/build"
  rm -rf "${destdir}"
  mkdir -p "${destdir}"

  local bindir="${destdir}/usr/bin"
  mkdir -p "${bindir}"

  local toolname
  for toolname in "${APPS[@]}"; do
    go build -o "${bindir}/${toolname}" "./${toolname}"
  done
  
  cd "${PROJECTROOT}"
}

merge() {
  mkdir -p "${PROJECTROOT}/fat/build"
  cp -nR "${PROJECTROOT}/arm64/build/." "${PROJECTROOT}/fat/build"
  cp -nR "${PROJECTROOT}/armv7/build/." "${PROJECTROOT}/fat/build"

  cd "${PROJECTROOT}"

  find "fat/build" -type f |
  while read x; do
    if lipo -info "$x" >/dev/null 2>&1; then
      rm "$x"
      lipo -create "${x/fat/arm64}" "${x/fat/armv7}" -output "$x"
      if test -x "$x"; then
        ldid -S/usr/share/SDKs/entitlements.xml "$x"
      fi
    fi
  done
}

pack() {
  local ARCHS
  if [ "$1" = fat ]; then
    ARCHS="ARM64/ARMv7"
  elif [ "$1" = arm64 ]; then
    ARCHS=ARM64
  elif [ "$1" = armv7 ]; then
    ARCHS=ARMv7
  else
    echo "Unknown architecture." >&2
    exit 1
  fi
  BUILDROOT="${PROJECTROOT}/$1"
  cp -R "${PROJECTROOT}/deb/." "${BUILDROOT}/build"
  sed -e "/^Version:/s/%%VERSION%%/${VERSION}/" \
      -e "/^Description:/s_%%ARCHS%%_${ARCHS}_" \
      -i -- "${BUILDROOT}/build/DEBIAN/control"
  dpkg-deb -Zxz --root-owner-group --build "${BUILDROOT}/build" "${BUILDROOT}"
  # su -c "chown -R 0:0 ${BUILDROOT}/build"
  # dpkg-deb -Zxz --build "${BUILDROOT}/build" "${BUILDROOT}"
}

init arm64
build

init armv7
build

merge

pack fat
