print()

def wait_for_resource_exist(namespace, type, name):
    while (name not in str(local('kubectl get {type} -n {namespace} -o name && printf "\n"'.format(type=type, namespace=namespace), quiet=True, echo_off=True))):
        local('sleep 0.1 && printf "\n"', quiet=True, echo_off=True)
        continue

# Deploy cert-manager
print('Deploying cert-manager...')
local('kubectl apply -f deploy/deps/third_party/cert-manager.yaml', quiet=True, echo_off=True)

print('Waiting on deployment/cert-manager to be ready...')
local('kubectl wait --for=condition=Available --timeout=300s -n cert-manager deployment/cert-manager', quiet=True, echo_off=True)

print('Waiting on deployment/cert-manager-cainjector to be ready...')
local('kubectl wait --for=condition=Available --timeout=300s -n cert-manager deployment/cert-manager-cainjector', quiet=True, echo_off=True)

print('Waiting on deployment/cert-manager-webhook to be ready...')
local('kubectl wait --for=condition=Available --timeout=300s -n cert-manager deployment/cert-manager-webhook', quiet=True, echo_off=True)
print()


# Deploy self-signed-issuer
print('Deploying self-signed-issuer...')
local('kubectl apply -f deploy/deps/first_party/self-signed-issuer.yaml', quiet=True, echo_off=True)
print()


# Deploy istio-operator
print('Deploying istio-operator...')
local('kubectl apply -f deploy/deps/first_party/istio-operator-ns.yaml', quiet=True, echo_off=True)
local('kubectl apply -f deploy/deps/third_party/istio-operator.yaml', quiet=True, echo_off=True)

print('Waiting on deployment/istio-operator to be ready...')
local('kubectl wait --for=condition=Available --timeout=300s -n istio-operator deployment/istio-operator', quiet=True, echo_off=True)
print()


# Deploy istio-controlplane
print('Deploying istio-controlplane...')
local('kubectl apply -f deploy/deps/first_party/istio-controlplane.yaml', quiet=True, echo_off=True)

print('Waiting on deployment/istiod to be ready...')
wait_for_resource_exist('istio-system', 'deployment', 'istiod')
local('kubectl wait --for=condition=Available --timeout=300s -n istio-system deployment/istiod', quiet=True, echo_off=True)

print('Waiting on deployment/istio-ingressgateway to be ready...')
wait_for_resource_exist('istio-system', 'deployment', 'istio-ingressgateway')
local('kubectl wait --for=condition=Available --timeout=300s -n istio-system deployment/istio-ingressgateway', quiet=True, echo_off=True)
print()


# Deploy tidb-operator
print('Deploying tidb-operator...')
local('kubectl apply -f deploy/deps/first_party/tidb-ns.yaml -n tidb-operator', quiet=True, echo_off=True)
local('kubectl apply -f deploy/deps/third_party/tidb-crd.yaml -n tidb-operator', quiet=True, echo_off=True)
local('kubectl apply -f deploy/deps/third_party/tidb-operator.yaml -n tidb-operator', quiet=True, echo_off=True)

print('Waiting on deployment/tidb-controller-manager to be ready...')
local('kubectl wait --for=condition=Available --timeout=300s -n tidb-operator deployment/tidb-controller-manager', quiet=True, echo_off=True)

print('Waiting on deployment/tidb-scheduler to be ready...')
local('kubectl wait --for=condition=Available --timeout=300s -n tidb-operator deployment/tidb-scheduler', quiet=True, echo_off=True)
print()


# Deploy tidb-cluster
print('Deploying tidb-cluster...')
local('kubectl apply -f deploy/deps/first_party/tidb-cluster.yaml -n tidb-cluster', quiet=True, echo_off=True)

print('Waiting on deployment/cluster-discovery to be ready...')
wait_for_resource_exist('tidb-cluster', 'deployment', 'cluster-discovery')
local('kubectl wait --for=condition=Available --timeout=300s -n tidb-cluster deployment/cluster-discovery', quiet=True, echo_off=True)

print('Waiting on statefulset/cluster-pd to be ready...')
wait_for_resource_exist('tidb-cluster', 'statefulset', 'cluster-pd')
local('kubectl rollout status --watch --timeout=300s -n tidb-cluster statefulset/cluster-pd', quiet=True, echo_off=True)

print('Waiting on statefulset/cluster-tikv to be ready...')
wait_for_resource_exist('tidb-cluster', 'statefulset', 'cluster-tikv')
local('kubectl rollout status --watch --timeout=300s -n tidb-cluster statefulset/cluster-tikv', quiet=True, echo_off=True)

print('Waiting on statefulset/cluster-tidb to be ready...')
wait_for_resource_exist('tidb-cluster', 'statefulset', 'cluster-tidb')
local('kubectl rollout status --watch --timeout=300s -n tidb-cluster statefulset/cluster-tidb', quiet=True, echo_off=True)
print()


# Deploy schemahero
print('Deploying schemahero...')
local('kubectl apply -f deploy/deps/first_party/schemahero-ns.yaml', quiet=True, echo_off=True)
local('kubectl apply -f deploy/deps/third_party/schemahero.yaml', quiet=True, echo_off=True)

print('Waiting on statefulset/schemahero to be ready...')
local('kubectl rollout status --watch --timeout=300s -n schemahero-system statefulset/schemahero', quiet=True, echo_off=True)
print()


# Deploy medulla
print('Deploying medulla...')
k8s_yaml(helm('deploy/chart', name='medulla', namespace='medulla-system', values='values-dev.yaml'))
