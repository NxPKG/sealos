module github.com/labring/sealos

go 1.20

require (
	github.com/BurntSushi/toml v1.4.0
	github.com/Masterminds/semver/v3 v3.2.1
	github.com/Masterminds/sprig/v3 v3.2.3
	github.com/containerd/containerd v1.7.18
	github.com/containers/buildah v1.37.0
	github.com/containers/common v0.60.0
	github.com/containers/image/v5 v5.32.0
	github.com/containers/ocicrypt v1.2.0
	github.com/containers/storage v1.55.0
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc
	github.com/docker/docker v27.1.1+incompatible
	github.com/docker/go-units v0.5.0
	github.com/emicklei/go-restful/v3 v3.10.1
	github.com/emirpasic/gods v1.18.1
	github.com/hashicorp/go-multierror v1.1.1
	github.com/imdario/mergo v0.3.16
	github.com/labring/image-cri-shim v0.0.0-00010101000000-000000000000
	github.com/labring/lvscare v0.0.0-00010101000000-000000000000
	github.com/labring/sreg v0.1.6
	github.com/manifoldco/promptui v0.9.0
	github.com/modood/table v0.0.0-20220527013332-8d47e76dad33
	github.com/opencontainers/go-digest v1.0.1-0.20220411205349-bde1400a84be
	github.com/opencontainers/image-spec v1.1.0
	github.com/pelletier/go-toml v1.9.5
	github.com/pkg/errors v0.9.1
	github.com/pkg/sftp v1.13.5
	github.com/schollz/progressbar/v3 v3.8.6
	github.com/sirupsen/logrus v1.9.3
	github.com/spf13/cobra v1.8.1
	github.com/spf13/pflag v1.0.5
	go.uber.org/zap v1.27.0
	golang.org/x/crypto v0.25.0
	golang.org/x/exp v0.0.0-20240613232115-7f521ea00fb8
	golang.org/x/sync v0.7.0
	golang.org/x/sys v0.22.0
	golang.org/x/term v0.22.0
	google.golang.org/grpc v1.64.1
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
	gopkg.in/yaml.v2 v2.4.0
	helm.sh/helm/v3 v3.11.2
	k8s.io/api v0.27.4
	k8s.io/apimachinery v0.27.4
	k8s.io/client-go v0.27.4
	k8s.io/kube-proxy v0.27.4
	k8s.io/kubectl v0.26.0
	k8s.io/kubelet v0.27.4
	k8s.io/kubernetes v1.27.4
	k8s.io/utils v0.0.0-20230505201702-9f6742963106
	sigs.k8s.io/controller-runtime v0.13.0
	sigs.k8s.io/kustomize/kyaml v0.14.1
	sigs.k8s.io/yaml v1.4.0
)

require (
	github.com/NYTimes/gziphandler v1.1.1 // indirect
	github.com/antlr/antlr4/runtime/Go/antlr v1.4.10 // indirect
	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd/v22 v22.5.0 // indirect
	github.com/google/cel-go v0.12.6 // indirect
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.1-0.20210315223345-82c243799c99 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.19.1 // indirect
	github.com/stoewer/go-strcase v1.2.0 // indirect
	go.etcd.io/etcd/api/v3 v3.5.7 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.7 // indirect
	go.etcd.io/etcd/client/v3 v3.5.7 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240318140521-94a12d6c2237 // indirect
	k8s.io/cloud-provider v0.0.0 // indirect
	k8s.io/controller-manager v0.27.4 // indirect
	k8s.io/cri-api v0.27.4 // indirect
	k8s.io/kms v0.27.4 // indirect
	sigs.k8s.io/apiserver-network-proxy/konnectivity-client v0.1.2 // indirect
)

require (
	dario.cat/mergo v1.0.0 // indirect
	github.com/AdaLogics/go-fuzz-headers v0.0.0-20230811130428-ced1acdcaa24 // indirect
	github.com/Azure/go-ansiterm v0.0.0-20230124172434-306776ec8161 // indirect
	github.com/MakeNowJust/heredoc v1.0.0 // indirect
	github.com/Masterminds/goutils v1.1.1 // indirect
	github.com/Microsoft/go-winio v0.6.2 // indirect
	github.com/Microsoft/hcsshim v0.12.5 // indirect
	github.com/VividCortex/ewma v1.2.0 // indirect
	github.com/acarl005/stripansi v0.0.0-20180116102854-5a71ef0e047d // indirect
	github.com/aead/serpent v0.0.0-20160714141033-fba169763ea6 // indirect
	github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/blang/semver/v4 v4.0.0 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/chzyer/readline v1.5.1 // indirect
	github.com/cilium/ebpf v0.11.0 // indirect
	github.com/container-orchestrated-devices/container-device-interface v0.5.4 // indirect
	github.com/containerd/cgroups/v3 v3.0.3 // indirect
	github.com/containerd/errdefs v0.1.0 // indirect
	github.com/containerd/log v0.1.0 // indirect
	github.com/containerd/stargz-snapshotter/estargz v0.15.1 // indirect
	github.com/containerd/typeurl/v2 v2.1.1 // indirect
	github.com/containernetworking/cni v1.2.3 // indirect
	github.com/containernetworking/plugins v1.5.1 // indirect
	github.com/containers/libtrust v0.0.0-20230121012942-c1716e8a8d01 // indirect
	github.com/containers/luksy v0.0.0-20240618143119-a8846e21c08c // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.4 // indirect
	github.com/cyberphone/json-canonicalization v0.0.0-20231217050601-ba74d44ecf5f // indirect
	github.com/cyphar/filepath-securejoin v0.3.1 // indirect
	github.com/disiqueira/gotree/v3 v3.0.2 // indirect
	github.com/distribution/distribution/v3 v3.0.0-20221208165359-362910506bc2 // indirect
	github.com/distribution/reference v0.6.0 // indirect
	github.com/docker/cli v27.1.1+incompatible // indirect
	github.com/docker/distribution v2.8.3+incompatible // indirect
	github.com/docker/docker-credential-helpers v0.8.2 // indirect
	github.com/docker/go-connections v0.5.0 // indirect
	github.com/docker/go-events v0.0.0-20190806004212-e31b211e4f1c // indirect
	github.com/docker/go-metrics v0.0.1 // indirect
	github.com/docker/libtrust v0.0.0-20160708172513-aabc10ec26b7 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/evanphx/json-patch v5.6.0+incompatible // indirect
	github.com/evanphx/json-patch/v5 v5.6.0 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/fsouza/go-dockerclient v1.11.1 // indirect
	github.com/go-errors/errors v1.4.2 // indirect
	github.com/go-jose/go-jose/v4 v4.0.2 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-logr/zapr v1.2.4 // indirect
	github.com/go-openapi/analysis v0.23.0 // indirect
	github.com/go-openapi/errors v0.22.0 // indirect
	github.com/go-openapi/jsonpointer v0.21.0 // indirect
	github.com/go-openapi/jsonreference v0.21.0 // indirect
	github.com/go-openapi/loads v0.22.0 // indirect
	github.com/go-openapi/runtime v0.28.0 // indirect
	github.com/go-openapi/spec v0.21.0 // indirect
	github.com/go-openapi/strfmt v0.23.0 // indirect
	github.com/go-openapi/swag v0.23.0 // indirect
	github.com/go-openapi/validate v0.24.0 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/godbus/dbus/v5 v5.1.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/gomodule/redigo v1.8.2 // indirect
	github.com/google/btree v1.0.1 // indirect
	github.com/google/gnostic v0.6.9 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/go-containerregistry v0.20.0 // indirect
	github.com/google/go-intervals v0.0.2 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/handlers v1.5.1 // indirect
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/gregjones/httpcache v0.0.0-20190611155906-901d90724c79 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/golang-lru v1.0.2 // indirect
	github.com/huandu/xstrings v1.4.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jinzhu/copier v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/klauspost/pgzip v1.2.6 // indirect
	github.com/kr/fs v0.1.0 // indirect
	github.com/letsencrypt/boulder v0.0.0-20240418210053-89b07f4543e0 // indirect
	github.com/liggitt/tabwriter v0.0.0-20181228230101-89fcab3d43de // indirect
	github.com/lithammer/dedent v1.1.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/mattn/go-shellwords v1.0.12 // indirect
	github.com/mattn/go-sqlite3 v1.14.22 // indirect
	github.com/miekg/pkcs11 v1.1.1 // indirect
	github.com/mistifyio/go-zfs/v3 v3.0.1 // indirect
	github.com/mitchellh/colorstring v0.0.0-20190213212951-d06e56a500db // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/moby/buildkit v0.12.5 // indirect
	github.com/moby/ipvs v1.1.0 // indirect
	github.com/moby/locker v1.0.1 // indirect
	github.com/moby/patternmatcher v0.6.0 // indirect
	github.com/moby/spdystream v0.2.0 // indirect
	github.com/moby/sys/mountinfo v0.7.2 // indirect
	github.com/moby/sys/sequential v0.5.0 // indirect
	github.com/moby/term v0.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/monochromegane/go-gitignore v0.0.0-20200626010858-205db1a8cc00 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	github.com/opencontainers/runc v1.1.13 // indirect
	github.com/opencontainers/runtime-spec v1.2.0 // indirect
	github.com/opencontainers/runtime-tools v0.9.1-0.20230914150019-408c51e934dc // indirect
	github.com/opencontainers/selinux v1.11.0 // indirect
	github.com/openshift/imagebuilder v1.2.14 // indirect
	github.com/ostreedev/ostree-go v0.0.0-20210805093236-719684c64e4f // indirect
	github.com/peterbourgon/diskv v2.0.1+incompatible // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/proglottis/gpgme v0.1.3 // indirect
	github.com/prometheus/client_golang v1.19.0 // indirect
	github.com/prometheus/client_model v0.6.0 // indirect
	github.com/prometheus/common v0.51.1 // indirect
	github.com/prometheus/procfs v0.15.1 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/seccomp/libseccomp-golang v0.10.0 // indirect
	github.com/secure-systems-lab/go-securesystemslib v0.8.0 // indirect
	github.com/shopspring/decimal v1.3.1 // indirect
	github.com/sigstore/fulcio v1.4.5 // indirect
	github.com/sigstore/rekor v1.3.6 // indirect
	github.com/sigstore/sigstore v1.8.4 // indirect
	github.com/smartystreets/goconvey v1.8.0 // indirect
	github.com/spf13/cast v1.6.0 // indirect
	github.com/stefanberger/go-pkcs11uri v0.0.0-20230803200340-78284954bff6 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	github.com/sylabs/sif/v2 v2.18.0 // indirect
	github.com/syndtr/gocapability v0.0.0-20200815063812-42c35b437635 // indirect
	github.com/tchap/go-patricia/v2 v2.3.1 // indirect
	github.com/titanous/rocacheck v0.0.0-20171023193734-afe73141d399 // indirect
	github.com/ulikunitz/xz v0.5.12 // indirect
	github.com/vbatts/tar-split v0.11.5 // indirect
	github.com/vbauerster/mpb/v8 v8.7.4 // indirect
	github.com/vishvananda/netlink v1.2.1-beta.2 // indirect
	github.com/vishvananda/netns v0.0.4 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	github.com/xlab/treeprint v1.1.0 // indirect
	go.etcd.io/bbolt v1.3.10 // indirect
	go.mongodb.org/mongo-driver v1.14.0 // indirect
	go.mozilla.org/pkcs7 v0.0.0-20210826202110-33d05740a352 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.49.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.49.0 // indirect
	go.opentelemetry.io/otel v1.24.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.22.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.21.0 // indirect
	go.opentelemetry.io/otel/metric v1.24.0 // indirect
	go.opentelemetry.io/otel/sdk v1.24.0 // indirect
	go.opentelemetry.io/otel/trace v1.24.0 // indirect
	go.opentelemetry.io/proto/otlp v1.0.0 // indirect
	go.starlark.net v0.0.0-20200306205701-8dd3e2ee1dd5 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/mod v0.18.0 // indirect
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/oauth2 v0.21.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	gomodules.xyz/jsonpatch/v2 v2.4.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240318140521-94a12d6c2237 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/apiextensions-apiserver v0.28.0 // indirect
	k8s.io/apiserver v0.27.4 // indirect
	k8s.io/cli-runtime v0.27.4 // indirect
	k8s.io/cluster-bootstrap v0.27.4 // indirect
	k8s.io/component-base v0.28.1 // indirect
	k8s.io/component-helpers v0.27.4 // indirect
	k8s.io/klog/v2 v2.120.1 // indirect
	k8s.io/kube-openapi v0.0.0-20230501164219-8b0f38b5fd1f // indirect
	oras.land/oras-go v1.2.2 // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/kustomize/api v0.13.2 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.3 // indirect
	tags.cncf.io/container-device-interface v0.8.0 // indirect
	tags.cncf.io/container-device-interface/specs-go v0.8.0 // indirect
)

replace (
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc => go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.35.0
	go.opentelemetry.io/otel/metric => go.opentelemetry.io/otel/metric v0.31.0
	k8s.io/api => k8s.io/api v0.27.4
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.27.4
	k8s.io/apimachinery => k8s.io/apimachinery v0.27.4
	k8s.io/apiserver => k8s.io/apiserver v0.27.4
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.27.4
	k8s.io/client-go => k8s.io/client-go v0.27.4
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.27.4
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.27.4
	k8s.io/component-base => k8s.io/component-base v0.27.4
	k8s.io/component-helpers => k8s.io/component-helpers v0.27.4
	k8s.io/cri-api => k8s.io/cri-api v0.27.4
	k8s.io/dynamic-resource-allocation => k8s.io/dynamic-resource-allocation v0.27.4
	k8s.io/kms => k8s.io/kms v0.27.4
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.27.4
	k8s.io/kubectl => k8s.io/kubectl v0.27.4
	k8s.io/kubelet => k8s.io/kubelet v0.27.4
	k8s.io/kubernetes => k8s.io/kubernetes v1.27.4
)

replace (
	github.com/containers/common => github.com/containers/common v0.53.1-0.20230613173441-e1ea4d9a74e5
	github.com/containers/storage => github.com/containers/storage v1.46.1
	github.com/docker/docker => github.com/docker/docker v23.0.6+incompatible
	github.com/imdario/mergo => github.com/imdario/mergo v0.3.13
	github.com/vishvananda/netlink => github.com/vishvananda/netlink v1.1.0
)

replace (
	github.com/labring/image-cri-shim => ./staging/src/github.com/labring/image-cri-shim
	github.com/labring/lvscare => ./staging/src/github.com/labring/lvscare
	github.com/labring/sealos => ./
)
