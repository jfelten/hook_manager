# Copyright 2016 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM lachlanevenson/k8s-kubectl:v1.6.0
MAINTAINER john.felten@gmail.com

ENV HELM_VERSION=v2.2.3 \
  DOCKER_API_VERSION=1.23

COPY ./hook_manager /hook_manager
COPY ./delete_hook.sh /delete_hook.sh
COPY ./install_hook.sh /install_hook.sh

RUN apk add --no-cache curl wget bash ca-certificates && update-ca-certificates\
  && chmod +x /install_hook.sh \
  && chmod +x /delete_hook.sh

# Install helm
RUN wget http://storage.googleapis.com/kubernetes-helm/helm-${HELM_VERSION}-linux-amd64.tar.gz \
  && tar xzf ./helm-${HELM_VERSION}-linux-amd64.tar.gz \
  && mv linux-amd64/helm /usr/local/bin \
  && rm -rf helm-${HELM_VERSION}-linux-amd64.tar.gz linux-amd64 

ENTRYPOINT [] 
