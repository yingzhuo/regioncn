# regioncn

```bash
docker image pull registry.cn-shanghai.aliyuncs.com/yingzhor/regioncn:1.0.0
```

## install

#### (1) k8s

```yaml
apiVersion: v1
kind: Namespace
metadata:
  name: regioncn
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: regioncn
  namespace: regioncn
spec:
  replicas: 1
  selector:
    matchLabels:
      app: regioncn
      node: regioncn
  template:
    metadata:
      name: regioncn
      namespace: regioncn
      labels:
        app: regioncn
        node: regioncn
    spec:
      containers:
        - name: regioncn-mysql
          image: "registry.cn-shanghai.aliyuncs.com/yingzhor/regioncn-mysql:1.0.0"
          imagePullPolicy: Always
          ports:
            - containerPort: 3306
        - name: region
          image: "registry.cn-shanghai.aliyuncs.com/yingzhor/regioncn:1.0.0"
          imagePullPolicy: Always
          ports:
            - containerPort: 8080
          env:
            - name: "WAIT_MYSQL"
              value: "localhost:3306"
          command:
            - "docker-entrypoint"
            - "--mysql-host=localhost"
            - "--mysql-port=3306"
            - "--type=json"
            - "--indent"
---
apiVersion: v1
kind: Service
metadata:
  name: regioncn
  namespace: regioncn
spec:
  selector:
    app: regioncn
    node: regioncn
  type: NodePort
  ports:
    - name: main
      port: 8080
      targetPort: 8080
      nodePort: 18081

```

#### (2) docker-compose

```yaml
version: "3.7"

services:
  db:
    image: registry.cn-shanghai.aliyuncs.com/yingzhor/regioncn-mysql:1.0.0
    container_name: "regioncn-mysql"
    restart: "always"

  regioncn:
    image: registry.cn-shanghai.aliyuncs.com/yingzhor/regioncn:1.0.0
    container_name: "regioncn"
    restart: "always"
    ports:
      - "18081:8080"
    links:
      - db:db
    command:
      - "docker-entrypoint"
      - "--type=protobuf"
      - "--mysql-host=db"
      - "--mysql-port=3306"
      - "--indent"
    environment:
      - "WAIT_MYSQL=db:3306"
```
