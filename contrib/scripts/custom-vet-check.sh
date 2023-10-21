#!/usr/bin/env bash
# SPDX-License-Identifier: Apache-2.0
# Copyright Authors of Cilium

# Cilium implements custom linters that can be found at
# https://github.com/khulnasoft/linters
# They performs custom static analysis checks.
"$GO" run github.com/khulnasoft/linters -timeafter.ignore inctimer -ioreadall.ignore safeio ./...
