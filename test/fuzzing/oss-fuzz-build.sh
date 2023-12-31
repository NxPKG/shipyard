#!/usr/bin/env bash

set -eu

# The oss-fuzz-build.sh is meant to only be run inside the OSS-Fuzz build environment.
# In that environment, the fuzzers are built with go build which does not include _test.go files
# in its scope and they are therefore renamed.
# When run locally by developers with go test -fuzz then the renaming is not necessary.
# Ciliums OSS-Fuzz integration can be found here: https://github.com/google/oss-fuzz/tree/master/projects/cilium


# Add a fuzz dependency because OSS-Fuzz rewrites the testing types to ones compatible with libFuzzer which is the fuzzing engine used by OSS-Fuzz:
printf "package policy\nimport _ \"github.com/AdamKorcz/go-118-fuzz-build/testing\"\n" > $SRC/cilium/pkg/policy/registerfuzzdep.go

go mod tidy && go mod vendor
mv $SRC/cilium/pkg/policy/l4_filter_test.go $SRC/cilium/pkg/policy/l4_filter_test_fuzz.go
mv $SRC/cilium/pkg/policy/policy_test.go $SRC/cilium/pkg/policy/policy_test_fuzz.go
mv $SRC/cilium/pkg/policy/rule_test.go $SRC/cilium/pkg/policy/rule_test_fuzz.go
mv $SRC/cilium/pkg/policy/selectorcache_test.go $SRC/cilium/pkg/policy/selectorcache_test_fuzz.go

compile_go_fuzzer github.com/khulnasoft/shipyard/test/fuzzing Fuzz fuzz gofuzz
compile_native_go_fuzzer github.com/khulnasoft/shipyard/pkg/monitor/format FuzzFormatEvent FuzzFormatEvent
compile_native_go_fuzzer github.com/khulnasoft/shipyard/pkg/k8s/apis/cilium.io/v2 FuzzCiliumNetworkPolicyParse FuzzCiliumNetworkPolicyParse
compile_native_go_fuzzer github.com/khulnasoft/shipyard/pkg/k8s/apis/cilium.io/v2 FuzzCiliumClusterwideNetworkPolicyParse FuzzCiliumClusterwideNetworkPolicyParse
compile_native_go_fuzzer github.com/khulnasoft/shipyard/pkg/policy FuzzResolveEgressPolicy FuzzResolveEgressPolicy
compile_native_go_fuzzer github.com/khulnasoft/shipyard/pkg/policy FuzzDenyPreferredInsert FuzzDenyPreferredInsert
compile_native_go_fuzzer github.com/khulnasoft/shipyard/pkg/policy FuzzAccumulateMapChange FuzzAccumulateMapChange
