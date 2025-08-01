# Copyright 2021 The Kubeflow Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Makefile to generate KFP api clients from proto.

IMAGE_TAG=kfp-api-generator
# Contact chensun or HumairAK if this remote image needs an update.
REMOTE_IMAGE=ghcr.io/kubeflow/kfp-api-generator
# Assume the latest API version by default.
API_VERSION ?= v2beta1

# Keep in sync with the version used in test/release/Dockerfile.release
PREBUILT_REMOTE_IMAGE=ghcr.io/kubeflow/kfp-api-generator:1.2

CONTAINER_ENGINE ?= docker

# Generate clients using a pre-built api-generator image.
.PHONY: generate
generate: fetch-dependencies hack/generator.sh $(API_VERSION)/*.proto
	${CONTAINER_ENGINE} run --interactive --rm \
		-e API_VERSION=$(API_VERSION) \
		--user $$(id -u):$$(id -g) \
		--mount type=bind,source="$$(pwd)/../..",target=/go/src/github.com/kubeflow/pipelines \
		$(PREBUILT_REMOTE_IMAGE) /go/src/github.com/kubeflow/pipelines/backend/api/hack/generator.sh

# Fetch dependency proto
.PHONY: fetch-dependencies
fetch-dependencies: v2beta1/google/rpc/status.proto

v2beta1/google/rpc/status.proto:
	mkdir -p v2beta1/google/rpc
	wget -O v2beta1/google/rpc/status.proto https://raw.githubusercontent.com/googleapis/googleapis/047d3a8ac7f75383855df0166144f891d7af08d9/google/rpc/status.proto

# Generate clients starting by building api-generator image locally.
# Note, this should only be used for local development purposes. Once any change is made to the Dockerfile,
# we should push the new image remotely to ensure everyone is using the same tools.
.PHONY: generate-from-scratch
generate-from-scratch: .image-built hack/generator.sh $(API_VERSION)/*.proto
	${CONTAINER_ENGINE} run --interactive --rm \
		-e API_VERSION=$(API_VERSION) \
	    --user $$(id -u):$$(id -g) \
		--mount type=bind,source="$$(pwd)/../..",target=/go/src/github.com/kubeflow/pipelines \
		$(IMAGE_TAG) /go/src/github.com/kubeflow/pipelines/backend/api/hack/generator.sh

# Build a local api-generator image.
.PHONY: image
image: .image-built

# Push api-generator image remotely.
.PHONY: push
push: image
	${CONTAINER_ENGINE} tag $(IMAGE_TAG) $(REMOTE_IMAGE)
	${CONTAINER_ENGINE} push $(REMOTE_IMAGE)

# .image-built is a local token file to help Make determine the latest successful build.
.image-built: Dockerfile
	${CONTAINER_ENGINE} build ../.. -t $(IMAGE_TAG) -f Dockerfile
	touch .image-built
