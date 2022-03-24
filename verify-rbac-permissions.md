```
export NS=default
export SA=default
kubectl auth can-i get endpoints --as system:serviceaccount:${NS}:${SA}
yes
kubectl auth can-i get pods --as system:serviceaccount:${NS}:${SA}
no
kubectl create -f - -o yaml << EOF
apiVersion: authorization.k8s.io/v1
kind: SubjectAccessReview
spec:
  user: system:serviceaccount:${NS}:${SA}
  resourceAttributes:
    group:
    resource: endpoints
    verb: get
    namespace: default
EOF
apiVersion: authorization.k8s.io/v1
kind: SubjectAccessReview
metadata:
  creationTimestamp: null
spec:
  resourceAttributes:
    namespace: default
    resource: endpoints
    verb: get
  user: system:serviceaccount:default:default
status:
  allowed: true
  reason: 'RBAC: allowed by ClusterRoleBinding "allow-default-sa-to-do-things" of
    ClusterRole "allow-default-sa-to-list-eps" to ServiceAccount "default/default"'
kubectl create -f - -o yaml << EOF 
apiVersion: authorization.k8s.io/v1
kind: SubjectAccessReview
spec:
  user: system:serviceaccount:${NS}:${SA}
  resourceAttributes:
    group:
    resource: pods
    verb: get
    namespace: default
EOF
apiVersion: authorization.k8s.io/v1
kind: SubjectAccessReview
metadata:
  creationTimestamp: null
spec:
  resourceAttributes:
    namespace: default
    resource: pods
    verb: get
  user: system:serviceaccount:default:default
status:
  allowed: false
```
