SOURCE_IMAGE = os.getenv("SOURCE_IMAGE", default='your-registry.io/project/tanzu-java-web-app-source')
# LOCAL_PATH = os.getenv("LOCAL_PATH", default='.')
NAMESPACE = os.getenv("NAMESPACE", default='default')

k8s_custom_deploy(
    'hello-go-tilt-dev',
    apply_cmd="tanzu apps workload apply hello-go-tilt-dev --live-update" +
               " --local-path ./src" +
               " --source-image " + SOURCE_IMAGE +
               " --namespace " + NAMESPACE +
               " --type web" +
               " --yes >/dev/null" +
               " && kubectl get workload hello-go-tilt-dev --namespace " + NAMESPACE + " -o yaml",
    delete_cmd="tanzu apps workload delete hello-go-tilt-dev --namespace " + NAMESPACE + " --yes",
    deps=[],
    container_selector='workload',
    live_update=[
    #   sync('./target/classes', '/workspace/BOOT-INF/classes')
      sync('./web', '/app/web')
    ]
)

k8s_resource('hello-go-tilt-dev', port_forwards=["8080:8080"],
            extra_pod_selectors=[{'serving.knative.dev/service': 'hello-go-tilt-dev'}])
allow_k8s_contexts('lkimmel')