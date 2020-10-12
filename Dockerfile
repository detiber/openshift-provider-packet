FROM registry.svc.ci.openshift.org/ocp/builder:rhel-8-golang-openshift-4.6 AS builder
WORKDIR /go/src/github.com/openshift/cluster-api-provider-packet
COPY . .
# VERSION env gets set in the openshift/release image and refers to the golang version, which interfers with our own
RUN unset VERSION \
 && make build GOPROXY=off NO_DOCKER=1 CGO_ENABLED=0

FROM registry.svc.ci.openshift.org/ocp/4.6:base
COPY --from=builder /go/src/github.com/openshift/cluster-api-provider-packet/bin/machine-controller-manager /
